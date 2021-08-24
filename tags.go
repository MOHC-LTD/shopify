package shopify

import "strings"

// SplitTags parses a shopify tag string into a list of the individual tags
func SplitTags(tags string) []string {
	return strings.Split(tags, ", ")
}

// ToTags combines a list of tags into a single shopify tag string
func ToTags(tags []string) string {
	return strings.Join(tags, ", ")
}
