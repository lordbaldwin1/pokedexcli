package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))
	cache.Add("https://example1.com", []byte("testdata"))
	cache.Add("https://example2.com", []byte("testdata"))
	cache.Add("https://example3.com", []byte("testdata"))
	cache.Add("https://example4.com", []byte("testdata"))
	cache.Add("https://example5.com", []byte("testdata"))
	cache.Add("https://example6.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	_, ok = cache.Get("https://example1.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	_, ok = cache.Get("https://example2.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	_, ok = cache.Get("https://example3.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	_, ok = cache.Get("https://example4.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	_, ok = cache.Get("https://example5.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	_, ok = cache.Get("https://example6.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
	_, ok = cache.Get("https://example1.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
	_, ok = cache.Get("https://example2.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
	_, ok = cache.Get("https://example3.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
	_, ok = cache.Get("https://example4.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
	_, ok = cache.Get("https://example5.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
	_, ok = cache.Get("https://example6.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
