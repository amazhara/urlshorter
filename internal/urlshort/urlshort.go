package urlshort

import (
	"fmt"

	"github.com/pkg/errors"
)

// Shorts hold URL and it`s short
type Shorts []struct {
	URL   string `json:"url" bson:"url"`
	Short string `json:"short" bson:"short"`
}

// toMap convert shorts to map[url]short
func (s Shorts) toMap() map[string]string {
	shorts := make(map[string]string, len(s))
	for _, short := range s {
		shorts[short.Short] = short.URL
	}

	return shorts
}

// Validation
func (s Shorts) validate() error {
	storedShorts, err := GetShorts()

	if err != nil {
		return err
	}

	for _, short := range s {
		// Check that all fields are filled
		if short.URL == "" || short.Short == "" {
			return errors.New("Url and Short fields should be not empty")
		}

		// Add / to string
		if short.Short[0] != '/' {
			short.Short = "/" + short.Short
		}

		// Check if short exist in db
		if _, ok := storedShorts[short.Short]; ok {
			return fmt.Errorf("%s already exists", short.Short)
		}
	}

	return nil
}

// GetShorts return map[url]shortURL
func GetShorts() (map[string]string, error) {
	shorts, err := findShorts()

	if err != nil {
		return nil, errors.Wrap(err, "When reading from mongo")
	}

	return shorts.toMap(), nil
}

// SaveShorts save new short URLs
func SaveShorts(s Shorts) error {
	if err := s.validate(); err != nil {
		return errors.Wrap(err, "Validation error")
	}

	if err := saveShorts(s); err != nil {
		return err
	}

	return nil
}
