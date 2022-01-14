package extra

import (
	"context"

	iresty "github.com/americanas-go/ignite/lib/resty.v2/resty"
	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Adds a header X-Request-ID with an UUID to each request made by resty client.
func RequestID(ctx context.Context, w *iresty.Wrapper) error {
	options := w.Options.Plugins.RequestID
	if !options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling requestID middleware in resty")

	w.Instance.OnBeforeRequest(requestId)

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
