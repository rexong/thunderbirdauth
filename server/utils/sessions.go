package utils

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const sessionDuration = 24 * time.Hour

type SessionManager struct {
	Sessions map[string]time.Time
	mu       sync.RWMutex
}

// Initialise the Session Manager
func InitialiseSessionManager() *SessionManager {
	log.Println("Initialising Session Manager")
	return &SessionManager{
		Sessions: make(map[string]time.Time),
	}
}

func getExpiryTime() time.Time {
	return time.Now().Add(sessionDuration)
}

// Issue Session Token based on Unix Nano Time.
// Token is stored with an expiry time in SessionManager.
func (sm *SessionManager) IssueSessionToken() string {
	sessionToken := fmt.Sprintf("%d", time.Now().UnixNano())
	expiryTime := getExpiryTime()

	sm.mu.Lock()
	sm.Sessions[sessionToken] = expiryTime
	sm.mu.Unlock()

	return sessionToken
}

// Verify Session Token against tokens in SessionManager.
func (sm *SessionManager) VerifySessionToken(sessionToken string) bool {
	sm.mu.RLock()
	expiry, exists := sm.Sessions[sessionToken]
	sm.mu.RUnlock()

	// If non existent
	if !exists {
		log.Println("Session does not exist")
		return false
	}

	// If exist but expired
	if expiry.Before(time.Now()) {
		log.Println("Session expiried")
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
