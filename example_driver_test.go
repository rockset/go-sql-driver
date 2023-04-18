package rockset_test

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/rockset/go-sql-driver"
)

// Example usage of the Rockset SQL driver
func ExampleDriver() {
	ctx := context.TODO()

	// connect using environment variables ROCKSET_APIKEY and ROCKSET_APISERVER
	db, err := sql.Open("rockset", "rockset://")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.QueryContext(ctx, `SELECT
    kind,
    count(kind) as total
FROM
    commons._events
group by
    kind
order by
    total DESC
`)
	if err != nil {
		log.Fatal(err)
	}

	var kind string
	var total int64
	for rows.Next() {
		err = rows.Scan(&kind, &total)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Output:
}
