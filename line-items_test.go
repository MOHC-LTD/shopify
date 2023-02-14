package shopify

import "testing"

// TestLineItem_GetPropertyByName tests that you can get a line item order property by its name
func TestLineItem_GetPropertyByName(t *testing.T) {
	lineItem := LineItem{}
	property := Property{Name: "TestProp", Value: "TestValue"}
	lineItem.Properties = append(lineItem.Properties, property)

	expectedValue := "TestValue"

	prop := lineItem.GetPropertyByName("TestProp")
	if prop.Value != expectedValue {
		t.Errorf("Was expecting: %v got %v", expectedValue, prop.Value)
	}
}
