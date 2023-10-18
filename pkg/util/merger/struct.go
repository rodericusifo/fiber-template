package merger

import (
	"github.com/imdario/mergo"
)

func MergeStructs[T comparable](withOverride bool, structs ...T) T {
	dst := new(T)

	if !withOverride {
		for _, s := range structs {
			mergo.Merge(dst, s)
		}

		return *dst
	}

	for _, s := range structs {
		mergo.Merge(dst, s, mergo.WithOverride)
	}

	return *dst
}
