package models

type PhotoInfo struct {
	ID        uint64 `json:"id,omitempty"`
	ImageURL  string `json:"imageURL" db:"image_url"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
}
