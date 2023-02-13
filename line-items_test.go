package shopify

import "testing"

// TestLineItem_GetPropertyByName tests that you can get an line item order property by its name
func TestLineItem_GetPropertyByName(t *testing.T) {
	lineItem := LineItem{}
	property := struct {
		Name  string
		Value string
	}{Name: "TestProp", Value: "TestValue"}
	lineItem.Properties = append(lineItem.Properties, property)

	prop := lineItem.GetPropertyByName("TestProp")
	if prop.Value != "TestValue" {
		t.Errorf("Was expecting: %v got %v", "TestProp", prop.Value)
	}
}
