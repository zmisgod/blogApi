package models

import "time"

//WpsPostMedia 文章媒体信息
type WpsPostMedia struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	MediaType int       `json:"media_type"`
	MediaURL  string    `json:"media_url"`
	IsForeign int       `json:"is_foreign"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
