package rockset

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"

	rs "github.com/rockset/rockset-go-client"
)

var (
	ErrNotImplemented   = errors.New("not implemented")
	ErrConnectionString = errors.New("incorrect connection string")
)

const driverName = "rockset"

func init() {
	sql.Register(driverName, &Driver{})
}

type Driver struct{}

// Open connects to Rockset using a connection string of the format
//   - rockset://
//   - rockset://apikey@apiserver
//
// The former connection string falls back to using the environment variables
//   - ROCKSET_APIKEY
//   - ROCKSET_APISERVER
func (d *Driver) Open(name string) (driver.Conn, error) {
	var options []rs.RockOption
	key, server, err := ParseConnStr(name)
	if err != nil {
		return nil, err
	}

	if key != "" && server != "" {
		options = append(options, rs.WithAPIServer(server), rs.WithAPIKey(key))
	}

	client, err := rs.NewClient(options...)
	if err != nil {
		return nil, err
	}

	return &Conn{client}, nil
}

var connStrRe = regexp.MustCompile(`^rockset://((\w+)@([\w\\.]+))?$`)

// ParseConnStr extracts apikey and apiserver from a connection string like rockset://apikey@apiserver
func ParseConnStr(s string) (key string, server string, err error) {
	m := connStrRe.FindStringSubmatch(s)
	if m == nil {
		return "", "", fmt.Errorf("no match: %w", ErrConnectionString)
	}
	switch l := len(m); l {
	case 1:
		return "", "", nil
	case 4:
		return m[2], m[3], nil
	default:
		return "", "", fmt.Errorf("wrong match count %d: %w", l, ErrConnectionString)
	}
}
