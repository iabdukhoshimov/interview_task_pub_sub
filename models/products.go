package models

type ProductMessage struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Image     Image  `json:"image"`
	Thumbnail Image  `json:"thumbnail"`
}

type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
