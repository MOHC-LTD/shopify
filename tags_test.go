package shopify

import (
	"testing"
)

// TestTags_Split_One tests that one tag gets split correctly
func TestTags_Split_One(t *testing.T) {
	expectedTagNum := 1
	tags := Tags("tag1")

	splittedTags := tags.Split()

	if len(splittedTags) != expectedTagNum {
		t.Errorf("Was expecting: %v got %v", expectedTagNum, len(splittedTags))
	}
}

// TestTags_Split_Multiple tests that multiple tags get split correctly
func TestTags_Split_Multiple(t *testing.T) {
	expectedTagNum := 4
	tags := Tags("tag1 test, tag2 , tag3,tag4")

	splittedTags := tags.Split()

	if len(splittedTags) != expectedTagNum {
		t.Errorf("Was expecting: %v got %v", expectedTagNum, len(splittedTags))
	}

	if splittedTags[0] != "tag1 test" {
		t.Errorf("Was expecting: %v got %v", "tag1", splittedTags[0])
	}

	if splittedTags[1] != "tag2" {
		t.Errorf("Was expecting: %v got %v", "tag2", splittedTags[1])
	}

	if splittedTags[2] != "tag3" {
		t.Errorf("Was expecting: %v got %v", "tag3", splittedTags[2])
	}

	if splittedTags[3] != "tag4" {
		t.Errorf("Was expecting: %v got %v", "tag4", splittedTags[3])
	}
}

// TestTags_Split_Empty tests that no split happens for no tags
func TestTags_Split_Empty(t *testing.T) {
	expectedTagNum := 0
	tags := Tags("")

	splittedTags := tags.Split()

	if len(splittedTags) != expectedTagNum {
		t.Errorf("Was expecting: %v got %v", expectedTagNum, len(splittedTags))
	}
}

// TestGetTagValue_Golden tests you can get the separated value of a tag via a delimiter
func TestGetTagValue_Golden(t *testing.T) {
	tagToFind := "tag2"
	expectedTagValue := "info2"
	testTags := []string{"tag1:info1", "tag2:info2", "tag3:info3"}

	tagValue := GetTagValue(testTags, tagToFind, ":")

	if tagValue != expectedTagValue {
		t.Errorf("Was expecting: %v got %v", expectedTagValue, tagValue)
	}
}

// TestGetTagValue_NotFound tests you get an empty string when the tag value was not found
func TestGetTagValue_NotFound(t *testing.T) {
	tagToFind := "missingTag"
	expectedTagValue := ""
	testTags := []string{"tag1:info1", "tag2:info2", "tag3:info3"}

	tagValue := GetTagValue(testTags, tagToFind, ":")

	if tagValue != expectedTagValue {
		t.Errorf("Was expecting: %v got %v", expectedTagValue, tagValue)
	}
}

// TestTags_Add_NonExisting tests that a new tag gets added when no tags already exist
func TestTags_Add_NonExisting(t *testing.T) {
	expectedTags := "tag1"
	tags := Tags("")
	newTags := tags.Add("tag1")

	if string(newTags) != expectedTags {
		t.Errorf("Was expecting: %v got %v", expectedTags, newTags)
	}
}

// TestTags_Add_WhenOneExisting tests that a new tag gets added to the end when one already exists
func TestTags_Add_WhenOneExisting(t *testing.T) {
	expectedTags := "tag1,tag2"
	tags := Tags("tag1")
	newTags := tags.Add("tag2")

	if string(newTags) != expectedTags {
		t.Errorf("Was expecting: %v got %v", expectedTags, newTags)
	}
}

// TestTags_Add_WhenMultipleExisting tests that a new tag gets added to the end when multiple already exist
func TestTags_Add_WhenMultipleExisting(t *testing.T) {
	expectedTags := "tag1,tag2,tag3"
	tags := Tags("tag1 ,tag2")
	newTags := tags.Add("tag3")

	if string(newTags) != expectedTags {
		t.Errorf("Was expecting: %v got %v", expectedTags, newTags)
	}
}

// TestTags_RemoveByKey_LowerCase_Start tests that it removes a tag from the start of the tag string
func TestTags_RemoveByKey_LowerCase_Start(t *testing.T) {
	tagToRemove := "tag1"
	tags := Tags("tag1, tag2 , tag3,tag4")
	expectedResult := "tag2,tag3,tag4"

	newTags := tags.RemoveByKey(tagToRemove)

	if string(newTags) != expectedResult {
		t.Errorf("Was expecting: %v got %v", expectedResult, newTags)
	}
}

// TestTags_RemoveByKey_LowerCase_Middle tests that it removes a tag from somewhere in the middle of the tag string
func TestTags_RemoveByKey_LowerCase_Middle(t *testing.T) {
	tagToRemove := "tag2"
	tags := Tags("tag1, tag2 , tag3,tag4")
	expectedResult := "tag1,tag3,tag4"

	newTags := tags.RemoveByKey(tagToRemove)

	if string(newTags) != expectedResult {
		t.Errorf("Was expecting: %v got %v", expectedResult, newTags)
	}
}

// TestTags_RemoveByKey_LowerCase_End tests that it removes a tag from the end of the tag string
func TestTags_RemoveByKey_LowerCase_End(t *testing.T) {
	tagToRemove := "tag4"
	tags := Tags("tag1, tag2 , tag3,tag4")
	expectedResult := "tag1,tag2,tag3"

	newTags := tags.RemoveByKey(tagToRemove)

	if string(newTags) != expectedResult {
		t.Errorf("Was expecting: %v got %v", expectedResult, newTags)
	}
}

// TestTags_RemoveByKey_Uppercase tests that it removes a tag from the tag string when the requested tag key is uppercase
func TestTags_RemoveByKey_Uppercase(t *testing.T) {
	tagToRemove := "TAG2"
	tags := Tags("tag1, tag2 , tag3,tag4")
	expectedResult := "tag1,tag3,tag4"

	newTags := tags.RemoveByKey(tagToRemove)

	if string(newTags) != expectedResult {
		t.Errorf("Was expecting: %v got %v", expectedResult, newTags)
	}
}

// TestTags_RemoveByKey_Uppercase_Tag tests that it removes a tag from the tag string when the tag to remove is uppercase
func TestTags_RemoveByKey_Uppercase_Tag(t *testing.T) {
	tagToRemove := "tag2"
	tags := Tags("tag1, TAG2 , tag3,tag4")
	expectedResult := "tag1,tag3,tag4"

	newTags := tags.RemoveByKey(tagToRemove)

	if string(newTags) != expectedResult {
		t.Errorf("Was expecting: %v got %v", expectedResult, newTags)
	}
}
