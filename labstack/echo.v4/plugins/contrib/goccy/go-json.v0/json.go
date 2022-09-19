package json

import (
	"context"
	"fmt"
	"net/http"

	e "github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	"github.com/goccy/go-json"
	"github.com/labstack/echo/v4"
)

// Register registers a new goccy/go-json plugin for echo.
func Register(ctx context.Context, server *e.Server) error {
	o, err := NewOptions()
	if err != nil {
		return err
	}
	h := NewJsonWithOptions(o)
	return h.Register(ctx, server)
}

// Json represents a goccy/go-json plugin for echo.
type Json struct {
	options *Options
}

// NewJsonWithOptions returns a new goccy/go-json plugin with options.
func NewJsonWithOptions(options *Options) *Json {
	return &Json{options: options}
}

// NewJsonWithConfigPath returns a new goccy/go-json plugin with options from config path.
func NewJsonWithConfigPath(path string) (*Json, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewJsonWithOptions(o), nil
}

// NewJson returns a new goccy/go-json plugin with default options.
func NewJson() *Json {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewJsonWithOptions(o)
}

// Register registers this goccy/go-json plugin for echo.
func (i *Json) Register(ctx context.Context, server *e.Server) error {

	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling goccy/go-json encoder in echo")

	instance := server.Instance()
	instance.JSONSerializer = &DefaultJSONSerializer{}

	return nil
}

// DefaultJSONSerializer implements JSON encoding using encoding/json.
type DefaultJSONSerializer struct{}

// Serialize converts an interface into a json and writes it to the response.
// You can optionally use the indent parameter to produce pretty JSONs.
func (d DefaultJSONSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := json.NewEncoder(c.Response())

	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

// Deserialize reads a JSON from a request body and converts it into an interface.
func (d DefaultJSONSerializer) Deserialize(c echo.Context, i interface{}) error {

	err := json.NewDecoder(c.Request().Body).Decode(i)

	if ute, ok := err.(*json.UnmarshalTypeError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset)).SetInternal(err)
	} else if se, ok := err.(*json.SyntaxError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())).SetInternal(err)
	}
	return err
}
