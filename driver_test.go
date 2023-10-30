package rockset_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	rockset "github.com/rockset/go-sql-driver"
)

const (
	createQuery = `INSERT INTO commons."go-sql-driver"
SELECT
    1 AS f_int,
    2.2 AS f_float,
    'foo' AS f_string,
    true AS f_bool,
    CURRENT_TIMESTAMP AS f_timestamp,
    1680632196642 AS t_unix`

	selectQuery = `SELECT
    f_int,
    f_float,
    f_string,
    f_bool,
    f_timestamp,
    t_unix
FROM
    commons."go-sql-driver"
`
)

func TestDriver(t *testing.T) {
	SkipUnlessEnv(t, "ROCKSET_APIKEY")

	ctx := context.TODO()
	ds := fmt.Sprintf("rockset://%s@%s", os.Getenv("ROCKSET_APIKEY"), os.Getenv("ROCKSET_APISERVER"))
	db, err := sql.Open("rockset", ds)
	require.NoError(t, err)

	rows, err := db.QueryContext(ctx, selectQuery)
	require.NoError(t, err)

	defer func() {
		err = rows.Close()
		assert.NoError(t, err)
	}()

	for rows.Next() {
		var fInt64 int64
		var fFloat64 float64
		var fString string
		var fBool bool
		var fTime rockset.Time
		// fUnix is a UNIX timestamp - it special as it becomes 1.680632196642e12, and needs help
		// to be converted into an int64
		// TODO see if it is possible to use a NamedValueChecker on the driver to handle this
		var fUnix rockset.Int64
		err := rows.Scan(&fInt64, &fFloat64, &fString, &fBool, &fTime, &fUnix)
		require.NoError(t, err)
		assert.Equal(t, int64(1), fInt64)
		assert.Equal(t, 2.2, fFloat64)
		assert.Equal(t, "foo", fString)
		assert.True(t, fBool)
		assert.Equal(t, int64(1681741419), fTime.Unix())
		assert.Equal(t, rockset.Int64(1680632196642), fUnix)
	}
}

func TestParseConnStr(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		key    string
		server string
		err    error
	}{
		{"correct", "rockset://apikey@api.usw2a1.rockset.com", "apikey", "api.usw2a1.rockset.com", nil},
		{"from env", "rockset://", "", "", nil},
		{"wrong prefix", "foobar://apikey@api.usw2a1.rockset.com", "", "", rockset.ErrConnectionString},
		{"no server", "rockset://apikey", "", "", rockset.ErrConnectionString},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			key, server, err := rockset.ParseConnStr(tst.s)
			if tst.err != nil {
				assert.ErrorIs(t, err, tst.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tst.key, key)
				assert.Equal(t, tst.server, server)
			}
		})
	}
}
