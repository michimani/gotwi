package resources

type Non200Error struct {
	Title      string `json:"title"`
	Detail     string `json:"detail"`
	Type       string `json:"type"`
	Status     string `json:"-"`
	StatusCode int    `json:"-"`
}
