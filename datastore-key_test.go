package data

import (
	"testing"

	"cloud.google.com/go/datastore"
)

func TestEncodeDecode(t *testing.T) {
	key1 := datastore.NameKey("Kind1", "name1", nil)
	encoded1 := DatastoreKeyToString(key1)
	decoded1 := DatastoreKeyFromString(encoded1)
	if decoded1.Kind != key1.Kind {
		t.Error()
	}
	if decoded1.Name != key1.Name {
		t.Error()
	}
	key2 := datastore.IDKey("Kind2", 2, key1)
	encoded2 := DatastoreKeyToString(key2)
	decoded2 := DatastoreKeyFromString(encoded2)
	if decoded2.Kind != key2.Kind {
		t.Error()
	}
	if decoded2.Name != key2.Name {
		t.Error()
	}
	if decoded2.Parent.Kind != key1.Kind {
		t.Error()
	}
	if decoded2.Parent.Name != key1.Name {
		t.Error()
	}
	var key3 datastore.Key
	encoded3 := DatastoreKeyToString(&key3)
	decoded3 := DatastoreKeyFromString(encoded3)
	if decoded3 != nil {
		t.Error(decoded3)
	}
}

func TestInvalidKeys(t *testing.T) {
	key1 := datastore.NameKey("Kind1", "name1", nil)
	encoded1 := DatastoreKeyToString(key1)
	decoded1 := DatastoreKeyFromString(encoded1 + "garbage")
	if decoded1 != nil {
		t.Error(decoded1)
	}
	decoded2 := DatastoreKeyFromString("garbage")
	if decoded2 != nil {
		t.Error(decoded2)
	}
	decoded3 := DatastoreKeyFromString("")
	if decoded3 != nil {
		t.Error(decoded3)
	}
}
