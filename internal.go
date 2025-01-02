package quotehub

import (
	"math/rand"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	// quotes stores all the loaded quotes.
	quotes []Quote

	random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

/*
Private internal functions
*/

// loadQuotes parses the embedded YAML content and loads quotes.
func loadQuotes() error {
	err := yaml.Unmarshal([]byte(quotesYAML), &quotes)
	if err != nil {
		return err
	}
	return nil
}

// init loads the quotes.
func init() {
	// Load quotes when the package is initialized
	err := loadQuotes()
	if err != nil {
		panic("Failed to load quotes: " + err.Error())
	}
}
