package request

type LinkRequest struct {
	Url string `json:"url" validate:"required,url"`
}
