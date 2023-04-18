package rockset

import (
	"context"
	"database/sql/driver"

	rs "github.com/rockset/rockset-go-client"
)

type Conn struct {
	rs *rs.RockClient
}

func (c *Conn) Close() error {
	return nil
}

// Begin is not supported.
func (c *Conn) Begin() (driver.Tx, error) {
	return nil, ErrNotImplemented
}

// Prepare is not supported.
func (c *Conn) Prepare(_ string) (driver.Stmt, error) {
	return nil, ErrNotImplemented
}

// Rollback is not supported.
func (c *Conn) Rollback() error {
	return ErrNotImplemented
}

// QueryContext executes a query.
func (c *Conn) QueryContext(ctx context.Context, query string, args []driver.Value) (driver.Rows, error) {
	if len(args) > 0 {
		// TODO detailed error message
		return nil, ErrNotImplemented
	}

	qr, err := c.rs.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	return &results{qr: qr}, nil
}

// Query is used to execute an SQL query.
// Deprecated: use QueryContext instead
func (c *Conn) Query(query string, args []driver.Value) (driver.Rows, error) {
	return c.QueryContext(context.Background(), query, args)
}

func (c *Conn) Ping(ctx context.Context) error {
	return c.rs.Ping(ctx)
}
