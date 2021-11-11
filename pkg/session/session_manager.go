/*
	Used for Session Management and handle expiry and idle timeout logic.
*/
package session

import (
	"sync"
	"time"
)

const (
	IDLE_TIME_OUT = 10
)

var lock = &sync.Mutex{}

var session_store = make(map[string]*UserSession)

type SessionManager struct {
}

var manager *SessionManager

/* Make sure that there is a single instance of SessionManager */
func GetInstance() *SessionManager {
	if manager != nil {
		return manager
	}

	lock.Lock()
	defer lock.Unlock()
	manager = &SessionManager{}
	return manager
}

/* Inserts a new authenticated session for a given user */
func (m *SessionManager) insert(s *UserSession) bool {
	lock.Lock()
	defer lock.Unlock()
	session_store[s.ID] = s

	return true
}

/* Find an unexpired user session given a Session ID */
func (m *SessionManager) find(sessionID string) *UserSession {
	if session, ok := session_store[sessionID]; ok && !m.hasExpired(sessionID) {
		return session
	}

	return nil
}

/* Delete a session from session store */
func (m *SessionManager) remove(sessionID string) {
	if _, ok := session_store[sessionID]; ok {
		lock.Lock()
		defer lock.Unlock()
		delete(session_store, sessionID)
	}
}

/* Verify if the session has expired or crossed the idle timeout */
func (m *SessionManager) hasExpired(sessionID string) bool {
	if ss, ok := session_store[sessionID]; ok {
		if ss != nil {
			now := time.Now().Unix()

			if (now >= ss.expireTime) || (now >= ss.lastAccessed+int64(IDLE_TIME_OUT)) {
				m.remove(sessionID)
				return true
			}
		}
	}

	return false
}
