package dtos

type RegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
