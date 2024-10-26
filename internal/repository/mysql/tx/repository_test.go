package tx

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
	"gorm.io/gorm"
)

func TestNew(t *testing.T) {
	mockDB := &gorm.DB{}

	want := &Repository{
		db: mockDB,
	}
	got := New(mockDB)
	assert.Equal(t, want, got)
}
