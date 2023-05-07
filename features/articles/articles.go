package articles

type Article struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Tags        []string `json:"tags"`
	Cover       string   `json:"cover"`
	Categories  []string `json:"categories"`
}
