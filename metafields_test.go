package shopify

import (
	"strconv"
	"testing"

	"github.com/jaswdr/faker"
)

/**
GIVEN I have a single metafield
WHEN I ask to get that metafield by its key when its of type boolean
THEN the value returned is of the correct type and what is expected
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

	returnedMetafield, err := metafields.GetByKey(metafield.Key)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}

	if returnedMetafield.Value.(bool) != expectedValue {
		t.Errorf("Was expecting: %v got %v", returnedMetafield.Value.(bool), expectedValue)
	}
}

/**
GIVEN I have a single metafield
WHEN I ask to get that metafield by its key when its of type string
THEN the value returned is of the correct type and what is expected
*/
func TestMetafields_GetByKeyString(t *testing.T) {
	fake := faker.New()
	expectedValue := fake.Color().Hex()

	metafield := Metafield{}
	fake.Struct().Fill(&metafield)
	metafield.Type = ColorMetaFieldType
	metafield.Value = expectedValue
	metafields := Metafields{metafield}

	returnedMetafield, err := metafields.GetByKey(metafield.Key)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}

	if returnedMetafield.Value.(string) != expectedValue {
		t.Errorf("Was expecting: %v got %v", returnedMetafield.Value.(bool), expectedValue)
	}
}

/**
GIVEN I have a single metafield
WHEN I ask to get that metafield by its key when its of type integer
THEN the value returned is of the correct type and what is expected
*/
func TestMetafields_GetByKeyInteger(t *testing.T) {
	fake := faker.New()
	expectedValue := fake.Int64()

	metafield := Metafield{}
	fake.Struct().Fill(&metafield)
	metafield.Type = NumberIntegerMetaFieldType
	metafield.Value = strconv.FormatInt(expectedValue, 10)
	metafields := Metafields{metafield}

	returnedMetafield, err := metafields.GetByKey(metafield.Key)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}

	if returnedMetafield.Value.(int64) != expectedValue {
		t.Errorf("Was expecting: %v got %v", returnedMetafield.Value.(bool), expectedValue)
	}
}

/**
GIVEN I have a single metafield
WHEN I ask to get that metafield by an invalid key
THEN the value returned is of the correct type and what is expected
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
