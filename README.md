# Go SQL driver for Rockset

This go module is a
[SQL driver](https://pkg.go.dev/database/sql/driver) 
for the
[Rockset](https://rockset.com/)
database.

It currently only support the base SQL types
* `int64` using `rockset.Int64`
* `float64`
* `bool`
* `[]byte`
* `string`
* `time.Time` using `rockset.Time`

## Usage

To use this driver you need a Rockset organization, and create an API key.

```go
	db, err := sql.Open("rockset", "rockset://")
	// check error

    rows, err := db.QueryContext(ctx, "...")
	// check error

    for rows.Next() {
        err = rows.Scan(&var1, &var2)
        // check error
    }
```
