package requestid

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RequestID represents a requestID plugin for resty client.
type RequestID struct {
	options *Options
}

// NewRequestIDWithConfigPath returns a new requestID plugin with options from config path.
func NewRequestIDWithConfigPath(path string) (*RequestID, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewRequestIDWithOptions(o), nil
}

// NewRequestIDWithOptions returns a new requestID plugin with options.
func NewRequestIDWithOptions(options *Options) *RequestID {
	return &RequestID{options: options}
}

// Register registers a new requestID plugin on resty client.
func Register(ctx context.Context, client *resty.Client) error {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	plugin := NewRequestIDWithOptions(o)
	return plugin.Register(ctx, client)
}

// Register registers this requestID plugin on resty client.
func (i *RequestID) Register(ctx context.Context, client *resty.Client) error {

	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling requestID middleware in resty")

	client.OnBeforeRequest(requestId)

	logger.Debug("requestID middleware successfully enabled in resty")

	return nil
}

func requestId(client *resty.Client, request *resty.Request) error {

	ctx := request.Context()

	idValue, ok := ctx.Value("requestId").(string)
	if !ok {
		id, err := uuid.NewUUID()
		if err != nil {
			return err
		}
		idValue = id.String()
	}

	request.SetHeader("X-Request-ID", idValue)

	return nil
}
