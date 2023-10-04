package dto

type Entry struct {
	UUID        string
	Title       string `json:"title"`
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
	Encrypted   *bool  `json:"is_encrypted"`
	CreateDate  string
}
