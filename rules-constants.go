package shopify

const (
	// RuleRelationNumberGreaterThan is number relation where the column value is greater than the condition.
	RuleRelationNumberGreaterThan = "greater_than"
	// RuleRelationNumberLessThan is the number relation where the column value is less than the condition.
	RuleRelationNumberLessThan = "less_than"
	// RuleRelationNumberEquals is the number relation where the column value is equal to the condition.
	RuleRelationNumberEquals = "equals"
	// RuleRelationNumberNotEquals is the number relation where the column value is not equal to the condition.
	RuleRelationNumberNotEquals = "not_equals"
	// RuleRelationtTextEquals is the text relation that checks if the column value is equal to the condition value.
	RuleRelationtTextEquals = "equals"
	// RuleRelationTextNotEquals is the text relation that checks if the column value is not equal to the condition value.
	RuleRelationTextNotEquals = "not_equals"
	// RuleRelationTextStartWith is the text relation that checks if the column value starts with the condition value.
	RuleRelationTextStartWith = "starts_with"
	// RuleRelationTextEndWith is the text relation that checks if the column value ends with the condition value.
	RuleRelationTextEndWith = "ends_with"
	// RuleRelationTextContains is the text relation that checks if the column value contains the condition value.
	RuleRelationTextContains = "contains"
	// RuleRelationTextNotContains is the text relation that checks if the column value does not contain the condition value.
	RuleRelationTextNotcontains = "not_contains"
)
