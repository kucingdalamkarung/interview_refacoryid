package model

type User []struct {
	ID       int        `json:"id"`
	Username string     `json:"username"`
	Profile  Profile    `json:"profile"`
	Articles []Articles `json:"articles:"`
}
type Profile struct {
	FullName string   `json:"full_name"`
	Birthday string   `json:"birthday"`
	Phones   []string `json:"phones"`
}
type Articles struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	PublishedAt string `json:"published_at"`
}
