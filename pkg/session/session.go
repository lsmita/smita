/*
	Package session provides persistent session for Trisolian domains.

	Note: Currently it is using a mock UserDB cache to get authenticated.
*/
package session

import (
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/trisolaria/connectulum/pkg/users"
)

const (
	EXPIRE_INTERVAL = 15
)

type UserSession struct {
	ID           string // session ID - an unique UUID
	userName     string // Username to which this session is associated with. Might be replaced by accessToken
	expireTime   int64  // To make session life finite
	lastAccessed int64  // To keep track of idleTimeout
}

/* Populate the user session with relevant fields on successful authentication */
func Create(s *UserSession, user string) {
	s.userName = user
	now := time.Now().Unix()
	s.expireTime = now + int64(EXPIRE_INTERVAL)
	s.lastAccessed = now
	uid, err := uuid.NewV4()

	if err == nil {
		s.ID = uid.String()
	}
}

/* Authenticates using user login credentials and creates a session object for
** future use by clients.
 */
func (s *UserSession) Authenticate(username, password string) bool {
	if users.Valid(username, password) {
		Create(s, username)
		GetInstance().insert(s)
		return true
	}

	return false
}
