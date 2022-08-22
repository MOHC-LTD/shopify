package shopify

import (
	"strconv"
	"testing"

	"github.com/jaswdr/faker"
)

/**
Given I have a single metafield
When I ask to get that metafield by its key when its of type boolean
Then the value is returned is the correct type
*/
func TestMetafields_GetByKeyBoolean(t *testing.T) {
	fake := faker.New()
	// Expects
	booleanString := fake.Boolean().BoolString("true", "false")
	expectedValue, _ := strconv.ParseBool(booleanString)

	metafield := Metafield{}
	fake.Struct().Fill(&metafield)
	metafield.Type = BooleanMetaFieldType
	metafield.Value = booleanString
	metafields := Metafields{metafield}

	retValue, err := metafields.GetByKey(metafield.Key)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}

	if retValue.(bool) != expectedValue {
		t.Errorf("Was expecting: %v got %v", retValue.(bool), expectedValue)
	}
}

/**
Given I have a single metafield
When I ask to get that metafield by its key when its of type string
Then the value is returned is the correct type
*/
func TestMetafields_GetByKeyString(t *testing.T) {
	fake := faker.New()
	expectedValue := fake.Color().Hex()

	metafield := Metafield{}
	fake.Struct().Fill(&metafield)
	metafield.Type = ColorMetaFieldType
	metafield.Value = expectedValue
	metafields := Metafields{metafield}

	retValue, err := metafields.GetByKey(metafield.Key)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}

	if retValue.(string) != expectedValue {
		t.Errorf("Was expecting: %v got %v", retValue.(bool), expectedValue)
	}
}

/**
Given I have a single metafield
When I ask to get that metafield by its key when its of type integer
Then the value is returned is the correct type
*/
func TestMetafields_GetByKeyInteger(t *testing.T) {
	fake := faker.New()
	expectedValue := fake.Int64()

	metafield := Metafield{}
	fake.Struct().Fill(&metafield)
	metafield.Type = NumberIntegerMetaFieldType
	metafield.Value = strconv.FormatInt(expectedValue, 10)
	metafields := Metafields{metafield}

	retValue, err := metafields.GetByKey(metafield.Key)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}

	if retValue.(int64) != expectedValue {
		t.Errorf("Was expecting: %v got %v", retValue.(bool), expectedValue)
	}
}

/**
Given I have a single metafield
When I ask to get that metafield by an invalid key
Then the value is returned is the correct type
*/
func TestMetafields_GetByKeyInvalid(t *testing.T) {
	fake := faker.New()
	metafield := Metafield{}
	fake.Struct().Fill(&metafield)
	metafield.Key = ""
	metafields := Metafields{metafield}

	_, err := metafields.GetByKey("123")
	if err == nil {
		t.Errorf("An error was thrown: %v", err)
	}
}
