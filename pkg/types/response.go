package types

type Meta struct {
	CurrentPage      int `json:"current_page"`
	CountDataPerPage int `json:"count_data_per_page"`
	TotalData        int `json:"total_data"`
	TotalPage        int `json:"total_page"`
}

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Meta    *Meta  `json:"meta,omitempty"`
	Data    T      `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}
