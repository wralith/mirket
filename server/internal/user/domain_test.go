package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := NewUser(NewUserOpts{
		Username:       "test",
		Email:          "test@test.com",
		HashedPassword: []byte("test"),
	})

	assert.Len(t, user.ID, 11) // Weird thing to put here?
	assert.Nil(t, user.DeletedAt)
	assert.Equal(t, user.CreatedAt, user.UpdatedAt)
}
