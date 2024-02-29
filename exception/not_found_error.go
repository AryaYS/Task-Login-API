package exception

type NotFoundError struct {
	Error string
}

func NotFoundErrorF(error string) NotFoundError {
	return NotFoundError{
		Error: error,
	}
}
