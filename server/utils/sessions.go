package utils

import (
	"fmt"
	"sync"
	"time"
)

const sessionDuration = 24 * time.Hour

type SessionManager struct {
	sessions map[string]time.Time
	mu       sync.RWMutex
}

// Initialise the Session Manager
func InitialiseSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]time.Time),
	}
}

func getExpiryTime() time.Time {
	return time.Now().Add(sessionDuration)
}

// Issue Session Token based on Unix Nano Time.
// Token is stored with an expiry time in SessionManager.
func (sm *SessionManager) issueSessionToken() string {
	sessionToken := fmt.Sprintf("%d", time.Now().UnixNano())
	expiryTime := getExpiryTime()

	sm.mu.Lock()
	sm.sessions[sessionToken] = expiryTime
	sm.mu.Unlock()

	return sessionToken
}

// Verify Session Token against tokens in SessionManager.
func (sm *SessionManager) verifySessionToken(sessionToken string) bool {
	sm.mu.RLock()
	expiry, exists := sm.sessions[sessionToken]
	sm.mu.RUnlock()

	// If non existent
	if !exists {
		return false
	}

	// If exist but expired
	if !expiry.Before(time.Now()) {
		sm.mu.Lock()
		delete(sm.sessions, sessionToken)
		sm.mu.Unlock()
		return false
	}

	newExpiry := getExpiryTime()
	sm.mu.Lock()
	sm.sessions[sessionToken] = newExpiry
	sm.mu.Unlock()
	return true
}
