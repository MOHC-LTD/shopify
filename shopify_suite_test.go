package shopify_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestShopify(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopify Suite")
}
