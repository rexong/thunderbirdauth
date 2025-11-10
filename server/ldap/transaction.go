package ldapserver

import (
	badger "github.com/dgraph-io/badger/v4"
	"log"
)

func findEntryByUIDTransaction(targetUid string, foundEntry **LdapEntry) func(txn *badger.Txn) error {
	return func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			var entryData []byte
			item := it.Item()
			key := item.Key()
			err := item.Value(func(val []byte) error {
				entryData = make([]byte, len(val))
				copy(entryData, val)
				return nil
			})
			if err != nil {
				log.Printf("Error retrieving value for %s: %v", string(key), err)
				continue
			}
			entry, err := unmarshalEntry(entryData)
			if err != nil {
				log.Printf("Error unmarshalling entry %s: %v", string(key), err)
				continue
			}
			if entry.UID == targetUid {
				*foundEntry = entry
				return nil
			}
		}
		return nil
	}
}

func getEntryByKeyTransaction(key []byte, value *[]byte) func(txn *badger.Txn) error {
	return func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		*value, err = item.ValueCopy(nil)
		return err
	}
}

func setEntryTransaction(key, entryData []byte) func(txn *badger.Txn) error {
	return func(txn *badger.Txn) error {
		return txn.Set(key, entryData)
	}
}
