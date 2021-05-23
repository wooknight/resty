package data

import (
	"testing"
)

func TestPut(t *testing.T) {
	key, val := "bun", "pav"
	Put(key, val)
	var val1 string
	var ok error
	if val1, ok = Get(key); ok != nil {
		t.Errorf("Could not find value for key %s", key)
	}
	if val1 != val {
		t.Errorf("Check failed .Expected - %s\tGot-%s\n", val, val1)
	}
}

func TestDelete(t *testing.T) {
	key, val := "bun", "pav"
	var val1 string
	var ok error
	Put(key, val)
	if val1, ok = Get(key); ok != nil {
		t.Errorf("Could not find value %s for key %s", val1 , key)
	}
	Delete(key)
	if val1, ok = Get(key); ok == nil {
		t.Errorf("Could not delete value for key %s", key)
	}
}
