package quotehub

// Quote represents the structure of a quote.
type Quote struct {
	Text     string   `json:"text" yaml:"text"`
	Author   string   `json:"author" yaml:"author"`
	Category string   `json:"category" yaml:"category"`
	Tags     []string `json:"tags" yaml:"tags"`
	Year     int      `json:"year" yaml:"year"`
}
