package cachev1

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestCache_GetAfterExpiry_standard(t *testing.T) {
	c := New[string, int](1 * time.Second)
	c.Set("answer", 42)

	// Jump past the TTL.
	time.Sleep(1100 * time.Millisecond)

	_, ok := c.Get("answer")
	if ok {
		t.Fatal("expected key to be expired after TTL")
	}
}

func TestCache_GetAfterExpiry_synctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := New[string, int](1 * time.Second)
		c.Set("answer", 42)

		// Jump past the TTL.
		time.Sleep(1100 * time.Millisecond)

		_, ok := c.Get("answer")
		if ok {
			t.Fatal("expected key to be expired after TTL")
		}
	})
}
