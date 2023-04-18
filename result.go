package rockset

import (
	"io"

	"database/sql/driver"

	"github.com/rockset/rockset-go-client/openapi"
)

type results struct {
	qr          openapi.QueryResponse
	resultIndex int
}

func (r *results) Columns() []string {
	columns := make([]string, len(r.qr.ColumnFields), len(r.qr.ColumnFields))
	for i, f := range r.qr.ColumnFields {
		columns[i] = f.Name
	}

	return columns
}

func (r *results) Close() error {
	return nil
}

func (r *results) Next(dest []driver.Value) error {
	if r.resultIndex >= len(r.qr.Results) {
		return io.EOF
	}

	for i, col := range r.qr.ColumnFields {
		res := r.qr.Results[r.resultIndex]
		dest[i] = res[col.Name]
	}
	r.resultIndex++

	return nil
}
