package session

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func AssertGetInstance(t *testing.T) *SessionManager {
	sm := GetInstance()

	assert.True(t, (sm != nil), "Got a valid instance")

	return sm
}

func createSession(t *testing.T, u string) *UserSession {
	sess := &UserSession{}
	Create(sess, u)

	assert.True(t, len(sess.ID) > 0)

	return sess
}

func expiryTest(ch chan bool, sessID string, sleepInterval int64) {
	time.Sleep(time.Duration(sleepInterval) * time.Second)
	ch <- GetInstance().hasExpired(sessID)
}

func TestSessionManagerGetInstanceSuccess(t *testing.T) {
	AssertGetInstance(t)
}

func TestSessionManagerGetInstanceMultipleInstance(t *testing.T) {
	sm := AssertGetInstance(t)
	sm2 := AssertGetInstance(t)

	assert.Same(t, sm, sm2, "Same instance")
}

func TestSessionManagerInsertSession(t *testing.T) {
	u := "foo"
	sm := AssertGetInstance(t)
	sess := createSession(t, u)

	inserted := sm.insert(sess)

	assert.Truef(t, inserted, "%s session inserted", u)

}

func TestSessionManagerFindSession(t *testing.T) {
	u := "foo"
	sm := AssertGetInstance(t)
	sess := createSession(t, u)

	inserted := sm.insert(sess)

	assert.Truef(t, inserted, "'%s' session inserted", u)
	ss := sm.find(sess.ID)

	assert.Same(t, sess, ss, "inserted session for '%s' found", u)
}

func TestSessionManagerExpiredSession(t *testing.T) {
	u := "foo"
	sm := AssertGetInstance(t)
	sess := createSession(t, u)

	inserted := sm.insert(sess)

	assert.Truef(t, inserted, "'%s' session inserted", u)

	ch := make(chan bool)
	go expiryTest(ch, sess.ID, (EXPIRE_INTERVAL + 1))
	assert.Truef(t, <-ch, "Session '%s' has expired", sess.ID)
}

func TestSessionManagerNotExpiredSession(t *testing.T) {
	u := "foo"
	sm := AssertGetInstance(t)
	sess := createSession(t, u)

	inserted := sm.insert(sess)

	assert.Truef(t, inserted, "'%s' session inserted", u)

	expired := make(chan bool)
	go expiryTest(expired, sess.ID, 5)
	assert.Falsef(t, <-expired, "Session '%s' has not expired", sess.ID)
}
