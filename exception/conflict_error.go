package exception

type Unauthorized struct {
	Error string
}

func NewUnauthorizedf(error string) Unauthorized {
	return Unauthorized{
		Error: error,
	}
}
