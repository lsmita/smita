package session

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	userName string
	password string
}

func AuthenticateUser(u User) (bool, *UserSession) {
	uSession := &UserSession{}

	return uSession.Authenticate(u.userName, u.password), uSession
}

func assertTrue(t *testing.T, valid bool, sess *UserSession, u User) {
	assert.Truef(t, valid, "%s is authenticated", u.userName)
	assert.NotEmpty(t, sess.ID)
	assert.NotZero(t, sess.expireTime)
	assert.NotZero(t, sess.lastAccessed)
	assert.Equal(t, sess.userName, u.userName)
}

func assertFalse(t *testing.T, valid bool, sess *UserSession, u User) {
	assert.Falsef(t, valid, "%s is not authenticated", u.userName)
	assert.Empty(t, sess.ID)
	assert.Zero(t, sess.expireTime)
	assert.Zero(t, sess.lastAccessed)
	assert.Empty(t, sess.userName)
}

func TestAuthenticateUserValid(t *testing.T) {
	u := User{"foo", "bar"}
	valid, uSession := AuthenticateUser(u)
	assertTrue(t, valid, uSession, u)
}

func TestAuthenticateUserNotFoundInvalid(t *testing.T) {
	u := User{"smita", "bar"}
	valid, uSession := AuthenticateUser(u)
	assertFalse(t, valid, uSession, u)
}

func TestAuthenticateIncorrectPasswordInvalid(t *testing.T) {
	u := User{"good", "bar"}
	valid, uSession := AuthenticateUser(u)
	assertFalse(t, valid, uSession, u)
}
