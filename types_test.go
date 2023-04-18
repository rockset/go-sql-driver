package rockset_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	rockset "github.com/rockset/go-sql-driver"
)

func TestTime(t *testing.T) {
	tests := []struct {
		name string
		a    any
		time time.Time
		err  bool
	}{
		{name: "incorrect format", a: "foo", err: true},
		{name: "wrong type", a: 100, err: true},
		{"correct", "2023-04-17T14:23:39.800340Z", time.Unix(1681741419, 0), false},
	}

	var t0 rockset.Time

	for _, tst := range tests {
		err := t0.Scan(tst.a)
		if tst.err {
			assert.Error(t, err)
		} else {
			assert.Equal(t, tst.time.Unix(), t0.Unix())
		}
	}
}

func TestInt64(t *testing.T) {
	tests := []struct {
		a   any
		i   rockset.Int64
		err bool
	}{
		{a: "foo", err: true},
		{a: [4]byte{}, err: true},
		{123.0, 123, false},
	}

	var i rockset.Int64

	for _, tst := range tests {
		err := i.Scan(tst.a)
		if tst.err {
			assert.Error(t, err)
		} else {
			assert.Equal(t, tst.i, i)
		}
	}
}
