package response

import (
	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func ResponseSuccess[T any](message string, data T, meta *pkg_types.Meta) pkg_types.Response[T] {
	return pkg_types.Response[T]{
		Success: true,
		Message: message,
		Meta:    meta,
		Data:    data,
	}
}
