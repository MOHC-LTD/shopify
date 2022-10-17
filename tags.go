package shopify

import (
	"fmt"
	"strings"
)

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

// GetTagValue returns the value of a tag based on the key provided. The separator character is used to separate the key from value key:value
func GetTagValue(tags []string, key string, separator string) string {
	for _, tag := range tags {
		if strings.Contains(tag, key) {
			return strings.Split(tag, separator)[1]
		}
	}

	return ""
}

// Add adds a new tag at the beginning of the comma-separated tag list
func (tags Tags) Add(newTag string) Tags {
	return Tags(fmt.Sprintf("%v, %v", newTag, tags))
}

// RemoveFromListByKey removes the first item found from the tag list that contains the specified key
func (tags Tags) RemoveFromListByKey(key string) Tags {
	tagsList := tags.Split()

	for i, tag := range tagsList {
		if strings.Contains(tag, key) {
			tagsList = append(tagsList[:i], tagsList[i+1:]...)
			break
		}
	}

	return NewTags(tagsList)
}
