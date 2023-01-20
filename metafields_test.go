package shopify

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/jaswdr/faker"
)

/*
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

	returnedMetafield, err := metafields.GetByKey(metafield.Key, metafield.Namespace)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}

	if returnedMetafield.Value.(bool) != expectedValue {
		t.Errorf("Was expecting: %v got %v", returnedMetafield.Value.(bool), expectedValue)
	}
}

/*
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

	returnedMetafield, err := metafields.GetByKey(metafield.Key, metafield.Namespace)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}

	if returnedMetafield.Value.(string) != expectedValue {
		t.Errorf("Was expecting: %v got %v", returnedMetafield.Value.(bool), expectedValue)
	}
}

/*
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

	returnedMetafield, err := metafields.GetByKey(metafield.Key, metafield.Namespace)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}

	if returnedMetafield.Value.(int64) != expectedValue {
		t.Errorf("Was expecting: %v got %v", returnedMetafield.Value.(bool), expectedValue)
	}
}

/*
GIVEN I have a single metafield
WHEN I ask to get that metafield by its key when its of type list.single_line_text_field
THEN the value returned is of the correct type and what is expected
*/
func TestMetafields_GetByKeyListSingleLineTextField(t *testing.T) {
	// Expects
	fake := faker.New()
	rawMetafieldValue := "[\"aa\",\"bb\"]" // encoded as json array
	expectedValue := []string{"aa", "bb"}

	metafield := Metafield{}
	fake.Struct().Fill(&metafield)
	metafield.Type = ListSingleLineTextFieldMetaFieldType
	metafield.Value = rawMetafieldValue
	metafields := Metafields{metafield}

	returnedMetafield, err := metafields.GetByKey(metafield.Key, metafield.Namespace)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}

	if !reflect.DeepEqual(returnedMetafield.Value.([]string), expectedValue) {
		t.Errorf("Was expecting: %v got %v", returnedMetafield.Value.([]string), expectedValue)
	}
}

/*
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

	_, err := metafields.GetByKey("123", "custom")
	if err == nil {
		t.Errorf("An error was thrown: %v", err)
	}
}

/*
GIVEN I have an array of metafields with one metafield that has "custom" namespace
WHEN I ask to get metafields by this namespace
THEN the length of the metafields array returned must be what is expected
*/
func TestMetafields_GetByNamespaceOneMetafield(t *testing.T) {
	// Expects
	fake := faker.New()
	metafieldsNamespace := "custom"
	expectedReturnedAmount := 1

	metafield := Metafield{}
	fake.Struct().Fill(&metafield)
	metafield.Namespace = metafieldsNamespace
	metafields := Metafields{metafield}

	returnedMetafields, err := metafields.GetByNamespace("custom")
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}

	if !reflect.DeepEqual(len(returnedMetafields), expectedReturnedAmount) {
		t.Errorf("Was expecting: %v got %v", len(returnedMetafields), expectedReturnedAmount)
	}

	if !reflect.DeepEqual(metafields, returnedMetafields) {
		t.Errorf("Was expecting: %v got %v", metafields, returnedMetafields)
	}
}

/*
GIVEN I have a list of metafields with different namespace than "custom"
WHEN I ask to get metafields with "custom" namespace
THEN the returned value must be an empty array of Metafields
*/
func TestMetafields_GetByNamespaceNoMatch(t *testing.T) {
	// Expects
	fake := faker.New()
	metafieldsNamespace := "namespace"
	expectedValue := Metafields{}

	metafield := Metafield{}
	fake.Struct().Fill(&metafield)
	metafield.Namespace = metafieldsNamespace
	metafields := Metafields{metafield}

	returnedMetafields, err := metafields.GetByNamespace("custom")
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}

	if !reflect.DeepEqual(returnedMetafields, expectedValue) {
		t.Errorf("Was expecting: %v got %v", returnedMetafields, expectedValue)
	}
}
