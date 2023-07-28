package request

type LinkStoreRequest struct {
	Url string `json:"url" validate:"required,url"`
}
