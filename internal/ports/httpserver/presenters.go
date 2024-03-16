package httpserver

type UserResponse struct {
	Url string `json:"url" example:"http://example.com"`
	Err string `json:"error" example:"some error message"`
}
