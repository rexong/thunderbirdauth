package http

import (
	"fmt"
	"sync"
	"time"
)

const sessionDuration = 24 * time.Hour

type SessionManager struct {
	Sessions map[string]time.Time
	mu       sync.RWMutex
}

func NewSessionManager() SessionManager {
	return SessionManager{
		Sessions: make(map[string]time.Time),
	}
}

func getExpiryTime() time.Time {
	return time.Now().Add(sessionDuration)
}

func (sm *SessionManager) IssueSessionToken() string {
	sessionToken := fmt.Sprintf("%d", time.Now().UnixNano())
	expiryTime := getExpiryTime()

	sm.mu.Lock()
	sm.Sessions[sessionToken] = expiryTime
	sm.mu.Unlock()

	return sessionToken
}

func (sm *SessionManager) VerifySessionToken(sessionToken string) bool {
	sm.mu.RLock()
	expiry, exists := sm.Sessions[sessionToken]
	sm.mu.RUnlock()
	if !exists {
		return false
	}
	if expiry.Before(time.Now()) {
		sm.mu.Lock()
		delete(sm.Sessions, sessionToken)
		sm.mu.Unlock()
		return false
	}
	newExpiry := getExpiryTime()
	sm.mu.Lock()
	sm.Sessions[sessionToken] = newExpiry
	sm.mu.Unlock()
	return true
}

func (sm *SessionManager) GetSessionExpiryByToken(sessionToken string) time.Time {
	return sm.Sessions[sessionToken]
}
