package datadog

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/americanas-go/log"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
)

type spanKey struct {
	ConnectionID string
	RequestID    int64
}

type monitor struct {
	sync.Mutex
	spans     map[spanKey]ddtrace.Span
	startOpts []ddtrace.StartSpanOption
}

func (m *monitor) Started(ctx context.Context, evt *event.CommandStartedEvent) {
	hostname, port := peerInfo(evt)
	b, _ := bson.MarshalExtJSON(evt.Command, false, false)

	var resource string
	coll := collName(evt)
	if coll != "" {
		resource = "mongo." + coll + "." + evt.CommandName
	} else {
		resource = "mongo." + evt.CommandName
	}

	opts := []ddtrace.StartSpanOption{
		tracer.SpanType(ext.SpanTypeMongoDB),
		tracer.ResourceName(resource),
		tracer.Tag(ext.DBInstance, evt.DatabaseName),
		tracer.Tag(ext.DBStatement, string(b)),
		tracer.Tag(ext.DBType, "mongo"),
		tracer.Tag(ext.PeerHostname, hostname),
		tracer.Tag(ext.PeerPort, port),
	}

	if parent, ok := tracer.SpanFromContext(ctx); ok {
		opts = append(opts, tracer.ChildOf(parent.Context()))
	}

	opts = append(opts, m.startOpts...)
	span, _ := tracer.StartSpanFromContext(ctx, "mongodb.query", opts...)
	key := spanKey{
		ConnectionID: evt.ConnectionID,
		RequestID:    evt.RequestID,
	}
	m.Lock()
	m.spans[key] = span
	m.Unlock()
}

func collName(e *event.CommandStartedEvent) string {
	coll := e.Command.Lookup(e.CommandName)
	collName, _ := coll.StringValueOK()
	return collName
}

func (m *monitor) Succeeded(ctx context.Context, evt *event.CommandSucceededEvent) {
	m.Finished(&evt.CommandFinishedEvent, nil)
}

func (m *monitor) Failed(ctx context.Context, evt *event.CommandFailedEvent) {
	m.Finished(&evt.CommandFinishedEvent, fmt.Errorf("%s", evt.Failure))
}

func (m *monitor) Finished(evt *event.CommandFinishedEvent, err error) {
	key := spanKey{
		ConnectionID: evt.ConnectionID,
		RequestID:    evt.RequestID,
	}
	m.Lock()
	span, ok := m.spans[key]
	if ok {
		delete(m.spans, key)
	}
	m.Unlock()
	if !ok {
		return
	}
	span.Finish(tracer.WithError(err))
}

// NewMonitor creates a new mongodb event CommandMonitor.
func NewMonitor(opts ...ddtrace.StartSpanOption) *event.CommandMonitor {
	log.Debug("contrib/go.mongodb.org/mongo-driver/mongo: Creating Monitor")
	m := &monitor{
		spans:     make(map[spanKey]ddtrace.Span),
		startOpts: opts,
	}
	return &event.CommandMonitor{
		Started:   m.Started,
		Succeeded: m.Succeeded,
		Failed:    m.Failed,
	}
}

func peerInfo(evt *event.CommandStartedEvent) (hostname, port string) {
	hostname = evt.ConnectionID
	port = "27017"
	if idx := strings.IndexByte(hostname, '['); idx >= 0 {
		hostname = hostname[:idx]
	}
	if idx := strings.IndexByte(hostname, ':'); idx >= 0 {
		port = hostname[idx+1:]
		hostname = hostname[:idx]
	}
	return hostname, port
}
