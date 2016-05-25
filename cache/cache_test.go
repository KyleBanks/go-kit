package cache

import (
	"strconv"
	"testing"
	"time"
)

var (
	ValidHost = "localhost:6379"

	TestKey   = "Test:" + strconv.Itoa(int(time.Now().Unix()))
	TestValue = "Value:" + strconv.Itoa(int(time.Now().Unix()))
)

func TestNew(t *testing.T) {
	cache := New(ValidHost)
	if cache == nil {
		t.Error("Nil cache returned!")
	}
}

func TestCache_PutString(t *testing.T) {
	cache := New(ValidHost)

	if ok, err := cache.PutString(TestKey, TestValue); err != nil {
		t.Error(err)
	} else if ok != "OK" {
		t.Error("Unexpected cache response:", ok)
	}
}

func TestCache_GetString(t *testing.T) {
	cache := New(ValidHost)

	if res, err := cache.GetString(TestKey); err != nil {
		t.Error(err)
	} else if res != TestValue {
		t.Error("Unexpected result:", res)
	}
}

func TestCache_Delete(t *testing.T) {
	cache := New(ValidHost)

	// Delete should not return an error
	if err := cache.Delete(TestKey); err != nil {
		t.Error(err)
	}

	// Calling again, on a deleted key, should still not fail
	if err := cache.Delete(TestKey); err != nil {
		t.Error(err)
	}
}

func TestCache_GetString_DeletedKey(t *testing.T) {
	cache := New(ValidHost)

	if _, err := cache.GetString(TestKey); err == nil {
		t.Error("Expected error getting deleted key!")
	}
}
