package resources

type Non200Error struct {
	Title      string `json:"title"`
	Detail     string `json:"detail"`
	Type       string `json:"type"`
	Status     string `json:"-"`
	StatusCode int    `json:"-"`
}

type PartialError struct {
	ResourceType string `json:"resource_type"`
	Field        string `json:"field"`
	Parameter    string `json:"parameter"`
	ResourceId   string `json:"resource_id"`
	Title        string `json:"title"`
	Section      string `json:"section"`
	Detail       string `json:"detail"`
	Value        string `json:"value"`
	Type         string `json:"type"`
}
