package json

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	e "github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/encoder"
	"github.com/labstack/echo/v4"
)

// Register registers a new sonic plugin for fiber.
func Register(ctx context.Context, server *e.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewSonicWithOptions(o)
	return h.Register(ctx, server)
}

// Sonic represents a sonic plugin for fiber.
type Sonic struct {
	options *Options
}

// NewSonicWithOptions returns a new sonic plugin with options.
func NewSonicWithOptions(options *Options) *Sonic {
	return &Sonic{options: options}
}

// NewSonicWithConfigPath returns a new sonic plugin with options from config path.
func NewSonicWithConfigPath(path string) (*Sonic, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewSonicWithOptions(o), nil
}

// NewSonic returns a new sonic plugin with default options.
func NewSonic() *Sonic {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewSonicWithOptions(o)
}

// Register registers this sonic plugin for fiber.
func (i *Sonic) Register(ctx context.Context, server *e.Server) error {

	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling bytedance/sonic encoder in echo")

	e := server.Instance()
	e.JSONSerializer = &DefaultJSONSerializer{}

	return nil
}

// DefaultJSONSerializer implements JSON encoding using encoding/json.
type DefaultJSONSerializer struct{}

// Serialize converts an interface into a json and writes it to the response.
// You can optionally use the indent parameter to produce pretty JSONs.
func (d DefaultJSONSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := encoder.NewStreamEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

// Deserialize reads a JSON from a request body and converts it into an interface.
func (d DefaultJSONSerializer) Deserialize(c echo.Context, i interface{}) error {

	err := decoder.NewStreamDecoder(c.Request().Body).Decode(i)
	if ute, ok := err.(*json.UnmarshalTypeError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset)).SetInternal(err)
	} else if se, ok := err.(*json.SyntaxError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())).SetInternal(err)
	}
	return err
}
