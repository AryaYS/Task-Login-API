package exception

type BadRequest struct {
	Error string
}

func BadRequestF(error string) BadRequest {
	return BadRequest{
		Error: error,
	}
}
