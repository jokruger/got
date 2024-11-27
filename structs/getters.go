package structs

import (
	. "github.com/jokruger/got/ifs"
)

// GetID returns the ID of the given IDProvider
func GetID[I any, T IDProvider[I]](s T) I {
	return s.GetID()
}

// GetName returns the name of the given NameProvider
func GetName[N NameProvider](s N) string {
	return s.GetName()
}
