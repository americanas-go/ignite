package ignite

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	separator = ","
)

// represents the values expected at ignite options struct field tag.
type IgniteOptionTag struct {
	// option config key.
	// Defined as `config:"example"`
	// If not defined it will be inferred by camelCasing the field name.
	Config string
	// default value if none option value is provided.
	// Defined as `default:"123"`.
	// Multiple values use commas: `default:"value1,value2"`.
	// Maps should be declared as: `"key1:value2,key2:value2"`
	// If not present the default value will be the field type zero.
	Default any
	// describes what the option field represents.
	// Defined as `desc:"example of description"`.
	// If not present it will be the empty string.
	Description string
	// option config path without options root, because it may change.
	// It is generated according to the tags added to the options struct.
	Path string
}

// tag as string: <path>: <default_value> // <description>
func (t *IgniteOptionTag) String() string {
	return fmt.Sprintf("%v:\t%v\t// %v", t.Path, t.Default, t.Description)
}

// returns all tags from ignite options.
func GetTags(o IgniteOptions) []*IgniteOptionTag {
	t := reflect.TypeOf(o)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	tags := []*IgniteOptionTag{}
	for i := 0; i < t.NumField(); i++ {
		tags = append(tags, getTags(t.Field(i), "")...)
	}
	return tags
}

func getTags(f reflect.StructField, p string) []*IgniteOptionTag {
	tags := []*IgniteOptionTag{}
	config := getTagValue(f, "config")
	if config == "" {
		//if config not present, it uses field camelcased name
		config = lowerFirst(f.Name)
	}
	path := fmt.Sprintf("%s.%s", p, config)
	if f.Type.Kind() == reflect.Struct {
		for i := 0; i < f.Type.NumField(); i++ {
			tags = append(tags, getTags(f.Type.Field(i), path)...)
		}
	} else {
		tag := &IgniteOptionTag{
			Config:      config,
			Default:     getDefault(f),
			Description: getTagValue(f, "desc"),
			Path:        path,
		}
		tags = append(tags, tag)
	}
	return tags
}

func getDefault(f reflect.StructField) any {
	d := getTagValue(f, "default")
	switch f.Type.Kind() {
	case reflect.Map:
		m := map[string]string{}
		if d != "" {
			keyVals := strings.Split(d, separator)
			for _, keyVal := range keyVals {
				kv := strings.SplitN(keyVal, ":", 2)
				m[kv[0]] = kv[1]
			}
		}
		return m
	case reflect.Slice:
		return strings.Split(d, separator)
	}
	return d
}

func getTagValue(f reflect.StructField, key string) string {
	if value, ok := f.Tag.Lookup(key); ok {
		return value
	}
	return ""
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}
