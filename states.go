package wasabi

import (
	"encoding/json"
	"fmt"
)

type State[T any] struct {
	model
	value T
}

func NewState[T any](value T) *State[T] {
	return &State[T]{
		model: newModel(),
		value: value,
	}
}

func copy(src interface{}, dst interface{}) {
	b, _ := json.Marshal(src)
	json.Unmarshal(b, dst)
}

func (s *State[T]) Set(value T) {
	s.value = value
	s.Notify()
}

func (s *State[T]) Get() T {
	var v T
	copy(s.value, &v)
	return v
}

func (s *State[T]) String() string {
	return fmt.Sprintf("%#v", s.value)
}
