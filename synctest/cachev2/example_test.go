package cachev2_test

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/jonbodner/testtalk/synctest/cachev2"
)

func TestCache_Clean(t *testing.T) {
	c := cachev2.New[string, string](2*time.Second, 1*time.Second)
	c.Set("key", "hello")
	val, ok := c.Get("key")
	if !ok {
		t.Fatal("expected key to be present")
	}
	if val != "hello" {
		t.Fatalf("got %q, want %q", val, "hello")
	}
	time.Sleep(4 * time.Second)
	val, ok = c.Get("key")
	if ok {
		t.Fatal("expected key to not be present")
	}
	if c.Stats.RemovedByGet != 0 {
		t.Error("expected removed by Get to be 0")
	}
	if c.Stats.RemovedBySweep != 1 {
		t.Error("expected removed by Sweep to be 1")
	}
}

func TestCache_Clean_Synctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := cachev2.New[string, string](2*time.Second, 1*time.Second)
		c.Set("key", "hello")
		val, ok := c.Get("key")
		if !ok {
			t.Fatal("expected key to be present")
		}
		if val != "hello" {
			t.Fatalf("got %q, want %q", val, "hello")
		}
		time.Sleep(4 * time.Second)
		val, ok = c.Get("key")
		if ok {
			t.Fatal("expected key to not be present")
		}
		if c.Stats.RemovedByGet != 0 {
			t.Error("expected removed by Get to be 0")
		}
		if c.Stats.RemovedBySweep != 1 {
			t.Error("expected removed by Sweep to be 1")
		}
	})
}

func TestCache_Clean_Synctest_Fixed(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := cachev2.New[string, string](2*time.Second, 1*time.Second)
		t.Cleanup(c.Done)
		c.Set("key", "hello")
		val, ok := c.Get("key")
		if !ok {
			t.Fatal("expected key to be present")
		}
		if val != "hello" {
			t.Fatalf("got %q, want %q", val, "hello")
		}
		time.Sleep(4 * time.Second)
		synctest.Wait()
		val, ok = c.Get("key")
		if ok {
			t.Fatal("expected key to not be present")
		}
		if c.Stats.RemovedByGet != 0 {
			t.Error("expected removed by Get to be 0")
		}
		if c.Stats.RemovedBySweep != 1 {
			t.Error("expected removed by Sweep to be 1")
		}
	})
}
