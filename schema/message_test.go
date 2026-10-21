package schema_test

import (
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suprsend/qlbridge/schema"
)

func TestMessage(t *testing.T) {
	k := schema.NewKeyUint(uint64(7))
	assert.Equal(t, driver.Value(uint64(7)), k.Key())
}
