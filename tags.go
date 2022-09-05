package shopify

import "strings"

// Tags are tags attached to the order, formatted as a string of comma-separated values. Tags are additional short descriptors, commonly used for filtering and searching. Each individual tag is limited to 40 characters in length.
type Tags string

// NewTags combines a list of tags into a single shopify tag string
func NewTags(tags []string) Tags {
	return Tags(strings.Join(tags, ", "))
}

// Split parses the tag string and returns a list of the individual tags
func (tags Tags) Split() []string {
	splitStrings := strings.Split(string(tags), ", ")
	if len(splitStrings) == 1 && splitStrings[0] == "" {
		return []string{}
	}

	return splitStrings
}
