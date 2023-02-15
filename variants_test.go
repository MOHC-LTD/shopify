package shopify

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

// TestVariants_Exists tests that a variant can be checked to see if it exists by its ID
func TestVariants_Exists(t *testing.T) {
	variants := Variants{}
	variants = append(variants, Variant{ID: 1})

	if !variants.Exists(1) {
		t.Errorf("Was expecting: %v got %v", true, !variants.Exists(1))
	}

	if variants.Exists(2) {
		t.Errorf("Was expecting: %v got %v", false, variants.Exists(2))
	}
}

// TestVariants_Get tests that a variant can be got by its ID
func TestVariants_Get(t *testing.T) {
	variants := Variants{}
	expectedVariant := Variant{ID: 1}
	variants = append(variants, expectedVariant)

	gotVariant := variants.Get(1)

	if !cmp.Equal(gotVariant, expectedVariant) {
		t.Errorf("Was expecting: %v got %v", expectedVariant, gotVariant)
	}
}
