package quotehub

import (
	"errors"
	"math/rand"
)

// RandomQuote returns a random quote from the library.
func RandomQuote() (Quote, error) {
	if len(quotes) == 0 {
		return Quote{}, errors.New("no quotes available")
	}
	return quotes[random.Intn(len(quotes))], nil
}

// GetQuote returns a random quote from the specified category.
func GetQuote(category string) (Quote, error) {
	var filtered []Quote
	for _, q := range quotes {
		if q.Category == category {
			filtered = append(filtered, q)
		}
	}
	if len(filtered) == 0 {
		return Quote{}, errors.New("no quotes found in the specified category")
	}
	return filtered[rand.Intn(len(filtered))], nil
}

// ListCategories returns all unique categories in the library.
func ListCategories() []string {
	categorySet := make(map[string]struct{})
	for _, q := range quotes {
		categorySet[q.Category] = struct{}{}
	}

	categories := make([]string, 0, len(categorySet))
	for category := range categorySet {
		categories = append(categories, category)
	}
	return categories
}
