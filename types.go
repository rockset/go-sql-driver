package rockset

import (
	"errors"
	"fmt"
	"time"
)

var ErrUnsupportedConversion = errors.New("unsupported conversion")

// Time is a time.Time that knows how to convert itself from a RFC3339 string.
type Time struct{ time.Time }

// Scan converts the argument a into Time
func (t *Time) Scan(a any) error {
	switch tp := a.(type) {
	case string:
		t0, err := time.Parse(time.RFC3339Nano, a.(string))
		if err != nil {
			return err
		}
		*t = Time{t0}
		return nil
	default:
		return fmt.Errorf("%w: can't convert %T to %T", ErrUnsupportedConversion, tp, *t)
	}
}

// Int64 is used to handle int values, as the [Rockset go client] uses json.Unmarshal into an
// interface value, which makes all numbers become float64.
//
// To unmarshal JSON into an interface value, Unmarshal stores one of these in the interface value:
//   - bool, for JSON booleans
//   - float64, for JSON numbers
//   - string, for JSON strings
//   - []interface{}, for JSON arrays
//   - map[string]interface{}, for JSON objects
//   - nil for JSON null
//
// [Rockset go client]: https://github.com/rockset/rockset-go-client
type Int64 int64

// Scan converts the argument a into Int64
func (i *Int64) Scan(a any) error {
	switch t := a.(type) {
	case float64:
		*i = Int64(int64(a.(float64)))
		return nil
	default:
		return fmt.Errorf("%w: can't convert %T to %T", ErrUnsupportedConversion, t, *i)
	}
}
