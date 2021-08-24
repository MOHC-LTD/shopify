package shopify

import "strings"

// Tags is a string that contains any number of short descriptors commonly used for filtering and searching, separated by commas.
type Tags string

// NewTags combines a list of tags into a single shopify tag string
func NewTags(tags []string) Tags {
	return Tags(strings.Join(tags, ", "))
}

// Split parses the tag string and returns a list of the individual tags
func (tags Tags) Split() []string {
	return strings.Split(string(tags), ", ")
}
