package posts

type CreatePostRequest struct {
	Title   string `json:"title" example:"My first post"`
	Content string `json:"content" example:"Hello world"`
}
