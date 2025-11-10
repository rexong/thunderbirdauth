package ldapserver

import (
	"fmt"
	"log"
	"time"

	badger "github.com/dgraph-io/badger/v4"
)

type Store struct {
	DB *badger.DB
}

func connectStore(path string) (*Store, error) {
	log.Printf("LDAP Store: Creating New Store at %s...", path)
	opts := badger.DefaultOptions(path)
	opts.Logger = nil

	db, err := badger.Open(opts)
	if err != nil {
		log.Printf("LDAP Store: Failed to open Badger DB")
		return nil, fmt.Errorf("LDAP Store: failed to open Badger DB: %w", err)
	}
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			log.Printf("LDAP Store: Running Garbage Collection on DB")
			db.RunValueLogGC(0.5)
			log.Printf("LDAP Store: Garbage Collection Complete")
		}
	}()

	store := Store{DB: db}
	seed(store)
	log.Printf("LDAP Store: Store Connected")
	return &store, nil
}

func (s *Store) Get(dn string) ([]byte, error) {
	log.Printf("LDAP Store: Retrieving %s", dn)
	var value []byte
	key := []byte(dn)
	err := s.DB.View(getEntryByKeyTransaction(key, &value))
	if err == badger.ErrKeyNotFound {
		log.Printf("LDAP Store: %s Not Found", dn)
		return nil, nil
	}

	log.Printf("LDAP Store: %s Found", dn)
	return value, err
}

func (s *Store) Set(entry *LdapEntry) error {
	log.Printf("LDAP Store: Setting %s entry", entry.DN)
	key := []byte(entry.DN)
	entryData, err := entry.marshal()
	if err != nil {
		log.Printf("LDAP Store: Unable to Set %s into Store", entry.DN)
		return nil
	}
	err = s.DB.Update(setEntryTransaction(key, entryData))
	if err == nil {
		log.Printf("LDAP Store: Set %s into Store", entry.DN)
	}
	return err
}

func (s *Store) View(targetUID string) (*LdapEntry, error) {
	var foundEntry *LdapEntry
	err := s.DB.View(findEntryByUIDTransaction(targetUID, &foundEntry))
	return foundEntry, err
}

func seed(s Store) {
	log.Println("Seeding mock data...")
	adminUser, user1 := createDefaultEntry()
	if err := s.Set(adminUser); err != nil {
		log.Fatalf("Failed to save admin user: %v", err)
	}
	log.Printf("Admin account created with DN: %s", config.AdminDN)

	if err := s.Set(user1); err != nil {
		log.Fatalf("Failed to save mock user: %v", err)
	}
	log.Printf("Regular user account created: %s", user1.DN)
}
