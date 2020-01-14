package threadsafemap_test

import (
	"testing"
	"github.com/puppetlabs/thread-safe-map/threadsafemap"
)

func TestRead(t *testing.T) {

	safeMap := threadsafemap.New(map[string]interface{}{"test": "this is a test"})

	got, _ := safeMap.Read("test")
	want := "this is a test"

	if got != want {
		t.Errorf("got %q want %q given, %q", got , want, "test")
	}
}