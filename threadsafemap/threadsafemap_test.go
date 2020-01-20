package threadsafemap_test

import (
	"testing"
	"github.com/puppetlabs/thread-safe-map/threadsafemap"
)

func TestRead(t *testing.T) {
	safeMap := threadsafemap.New(map[string]interface{}{"test": "this is a test"})

	t.Run("known key", func(t *testing.T) {
		got, _ := safeMap.Read("test")
		want := "this is a test"

		if got != want {
			t.Errorf("got %q want %q", got , want)
		}
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := safeMap.Read("unknown")

		if err == nil {
			t.Error(err)
		}
	})
}

func TestWrite(t *testing.T) {
	safeMap := threadsafemap.New(map[string]interface{}{})

	safeMap.Write("yes", "it worked")
	got, _ := safeMap.Read("yes")
	want := "it worked"

	if got != want {
		t.Error("key-value pair not found")
	}
}

func TestExists(t *testing.T) {
	safeMap := threadsafemap.New(map[string]interface{}{"test": "this is a test"})
	
	t.Run("existing key", func(t *testing.T) { 
		_, exists := safeMap.Exists("test")
		if !exists {
			t.Errorf("key doesn't exist")
		}
	})

	t.Run("non-existing key", func(t *testing.T) {
		_, exists := safeMap.Exists("purple")
		if exists {
			t.Errorf("key exists when it should not")
		} 
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	safeMap := threadsafemap.New(map[string]interface{}{key: "test words"})
	safeMap.Delete(key)

	_, err := safeMap.Read(key)
	if err == nil {
		t.Errorf("Expected %q to be deleted", key)
	}
}
