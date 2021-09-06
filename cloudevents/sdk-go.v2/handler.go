package cloudevents

import (
	"context"

	v2 "github.com/cloudevents/sdk-go/v2"
)

// Handler is a function for callback on receipt and handle of a cloudevent.
type Handler func(ctx context.Context, in v2.Event) (*v2.Event, error)
