package bootstrap

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func mockOpen(driverName, dataSourceName string) (*sql.DB, error) {
	return nil, errors.New("mock error")
}
func TestNewPostgresClient(t *testing.T) {
	client := NewPostgresClient("user=username dbname=testdb sslmode=disable", 10, 5)
	assert.NotNil(t, client)
	assert.NotNil(t, client.Db)
	assert.IsType(t, &sql.DB{}, client.Db)
}
