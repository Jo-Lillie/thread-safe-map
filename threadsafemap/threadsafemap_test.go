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

// test delete and exists function
