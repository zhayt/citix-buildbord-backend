package models

type News struct {
	TotalResults int        `json:"totalResults,omitempty"`
	Articles     []*Article `json:"articles,omitempty"`
}

type Article struct {
	Author      string `json:"author,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"urlToImage,omitempty"`
	ArticleURL  string `json:"url,omitempty"`
	PublishedAt string `json:"publishedAt,omitempty"`
}
