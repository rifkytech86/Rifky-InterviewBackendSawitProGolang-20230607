package bootstrap

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPostgresClient(t *testing.T) {
	client := NewPostgresClient("user=username dbname=testdb sslmode=disable", 10, 5)
	assert.NotNil(t, client)
	assert.NotNil(t, client.Db)
	assert.IsType(t, &sql.DB{}, client.Db)
}
