package model

type Response struct {
	// Message é como o Go lê.
	// `json:"message"` é como o site (JSON) vai ler.
	Message string `json:"message"`
}
