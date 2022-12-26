package data

import "cloud.google.com/go/datastore"

// to be used with github.com/tinylib/msgp / msg:shim

func DatastoreKeyToString(k *datastore.Key) string {
	return k.Encode()
}

func DatastoreKeyFromString(s string) *datastore.Key {
	if k, err := datastore.DecodeKey(s); err == nil {
		return k
	}
	return nil
}
