package schema_test

import (
	"database/sql/driver"
	"testing"

	"github.com/suprsend/qlbridge/datasource"

	"github.com/stretchr/testify/assert"
	"github.com/suprsend/qlbridge/datasource/memdb"
	"github.com/suprsend/qlbridge/schema"
)

func TestApplySchema(t *testing.T) {
	a := schema.NewApplyer(func(s *schema.Schema) schema.Source {
		sdb := datasource.NewSchemaDb(s)
		s.InfoSchema.DS = sdb
		return sdb
	})
	reg := schema.NewRegistry(a)
	a.Init(reg)

	inrow := []driver.Value{122, "bob", "bob@email.com"}
	db, err := memdb.NewMemDbData("users", [][]driver.Value{inrow}, []string{"user_id", "name", "email"})
	assert.Equal(t, nil, err)

	s := schema.NewSchema("hello")
	s.DS = db
	err = a.AddOrUpdateOnSchema(s, s)
	assert.Equal(t, nil, err)

	err = a.AddOrUpdateOnSchema(s, "not_real")
	assert.NotEqual(t, nil, err)

	a.Drop(s, s)

	err = a.Drop(s, "fake")
	assert.NotEqual(t, nil, err)
}
