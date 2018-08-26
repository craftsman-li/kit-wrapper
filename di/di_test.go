package di

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Database interface {
	GetValue() int
}

type databaseImpl struct {
	Started bool
}

func (d *databaseImpl) GetValue() int {
	return 42
}

func (d *databaseImpl) Open() error {
	d.Started = true
	return nil
}

func (d *databaseImpl) Close() {
}

type service struct {
	Database Database `inject:"db"`
	Started  bool
}

func (s *service) Open() error {
	s.Started = true
	return nil
}

func (s *service) Close() {
}

func TestDependencyInjection(t *testing.T) {

	Register("db", &databaseImpl{})

	Register("s", &service{})

	assert.NoError(t, Resolve())

	db, _ := GetByName("db").(*databaseImpl)
	assert.True(t, db.Started)
	s, _ := GetByName("s").(*service)
	assert.True(t, s.Started)
	assert.Equal(t, 42, s.Database.GetValue())
}
