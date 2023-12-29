package dream

import (
	"context"
	"sync"
	"time"
)

// Store is a concurrent-safe, in-memory key-value store.
type Store struct {
	mu sync.RWMutex
	kv map[string]data

	cancel context.CancelFunc
}

type data struct {
	value any
	putAt time.Time
}

// Delete removes the key and its associated value. If the key doesn't
// exist, Delete is a no-op.
func (s *Store) Delete(k string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.kv, k)
}

// Exists checks if a key exists in the in-memory store.
func (s *Store) Exists(k string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.kv[k]
	return ok
}

// Put adds the given key-value pair to the database. If the key already
// exists, it is updated.
func (s *Store) Put(k string, v any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.kv[k] = data{
		value: v,
		putAt: time.Now(),
	}
}

// Get returns the value associated with the given key. If the key
// doesn't exist, it returns nil.
func (s *Store) Get(k string) any {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if v, ok := s.kv[k]; ok {
		return v.value
	}
	return nil
}

// GetString returns the value associated with the given key as a
// string. If the value is not a string or the key doesn't exist, it
// returns an empty string.
func (s *Store) GetString(k string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return ""
	}
	if x, ok := v.value.(string); ok {
		return x
	}
	return ""
}

// GetBool returns the value associated with the given key as a boolean.
// If the value is not a bool or the key doesn't exist, it returns
// false.
func (s *Store) GetBool(k string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return false
	}
	if x, ok := v.value.(bool); ok {
		return x
	}
	return false
}

// GetInt returns the value associated with the given key as an int. If
// the value is not an int or the key doesn't exist, it returns 0.
func (s *Store) GetInt(k string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return 0
	}
	if x, ok := v.value.(int); ok {
		return x
	}
	return 0
}

// GetInt8 returns the value associated with the given key as an int8.
// If the value is not an int8 or the key doesn't exist, it returns 0.
func (s *Store) GetInt8(k string) int8 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return 0
	}
	if x, ok := v.value.(int8); ok {
		return x
	}
	return 0
}

// GetInt16 returns the value associated with the given key as an int16.
// If the value is not an int16 or the key doesn't exist, it returns 0.
func (s *Store) GetInt16(k string) int16 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return 0
	}
	if x, ok := v.value.(int16); ok {
		return x
	}
	return 0
}

// GetInt32 returns the value associated with the given key as an
// int32. If the value is not an int32 or the key doesn't exist, it
// returns 0.
func (s *Store) GetInt32(k string) int32 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return 0
	}
	if x, ok := v.value.(int32); ok {
		return x
	}
	return 0
}

// GetInt64 returns the value associated with the given key as an
// int64. If the value is not an int64 or the key doesn't exist, it
// returns 0.
func (s *Store) GetInt64(k string) int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return 0
	}
	if x, ok := v.value.(int64); ok {
		return x
	}
	return 0
}

// GetUint returns the value associated with the given key as an
// uint. If the value is not an uint or the key doesn't exist, it
// returns 0.
func (s *Store) GetUint(k string) uint {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return 0
	}
	if x, ok := v.value.(uint); ok {
		return x
	}
	return 0
}

// GetUint8 returns the value associated with the given key as an
// uint8. If the value is not an uint8 or the key doesn't exist, it
// returns 0.
func (s *Store) GetUint8(k string) uint8 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return 0
	}
	if x, ok := v.value.(uint8); ok {
		return x
	}
	return 0
}

// GetUint16 returns the value associated with the given key as an
// uint16. If the value is not an uint16 or the key doesn't exist, it
// returns 0.
func (s *Store) GetUint16(k string) uint16 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return 0
	}
	if x, ok := v.value.(uint16); ok {
		return x
	}
	return 0
}

// GetUint32 returns the value associated with the given key as an
// uint32. If the value is not an uint32 or the key doesn't exist, it
// returns 0.
func (s *Store) GetUint32(k string) uint32 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return 0
	}
	if x, ok := v.value.(uint32); ok {
		return x
	}
	return 0
}

// GetUint64 returns the value associated with the given key as an
// uint64. If the value is not an uint64 or the key doesn't exist, it
// returns 0.
func (s *Store) GetUint64(k string) uint64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return 0
	}
	if x, ok := v.value.(uint64); ok {
		return x
	}
	return 0
}

// GetFloat32 returns the value associated with the given key as a
// float32. If the value is not a float32 or the key doesn't exist, it
// returns 0.
func (s *Store) GetFloat32(k string) float32 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return 0
	}
	if x, ok := v.value.(float32); ok {
		return x
	}
	return 0
}

// GetFloat64 returns the value associated with the given key as a
// float64. If the value is not a float64 or the key doesn't exist, it
// returns 0.
func (s *Store) GetFloat64(k string) float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.kv[k]
	if !ok {
		return 0
	}
	if x, ok := v.value.(float64); ok {
		return x
	}
	return 0
}

// StopCleanUp stops the clean up loop.
func (s *Store) StopCleanUp() {
	s.cancel()
}

func (s *Store) cleanUp(ctx context.Context, interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
		}

		s.mu.Lock()
		for k, v := range s.kv {
			if time.Since(v.putAt) > interval {
				delete(s.kv, k)
			}
		}
		s.mu.Unlock()
	}
}
