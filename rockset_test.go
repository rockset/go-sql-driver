package rockset_test

import (
	"os"
	"testing"
)

func SkipUnlessEnv(t *testing.T, name string) {
	if _, found := os.LookupEnv(name); !found {
		t.Skipf("skipping as environment %s is not set", name)
	}
}
