package dto

type GenericResponse[T any] struct {
	Status   string `json:"status"`
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Response *T     `json:"response"`
}

func OK[T any](response *T) *GenericResponse[T] {
	return &GenericResponse[T]{
		Status:   "Success",
		Success:  true,
		Message:  "Success",
		Response: response,
	}
}

func OKWithMessage[T any](response *T, message string) *GenericResponse[T] {
	return &GenericResponse[T]{
		Status:   "Success",
		Success:  true,
		Message:  message,
		Response: response,
	}
}

func Error[T any](response *T) *GenericResponse[T] {
	return &GenericResponse[T]{
		Status:   "Error",
		Success:  false,
		Message:  "Failed",
		Response: response,
	}
}

func ErrorWithMessage[T any](response *T, message string) *GenericResponse[T] {
	return &GenericResponse[T]{
		Status:   "Error",
		Success:  false,
		Message:  message,
		Response: response,
	}
}
