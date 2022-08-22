package shopify

import (
	"strconv"

	"github.com/jaswdr/faker"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Tests that the correct type comes back when getting a metafield by its key
var _ = Describe("Checking Metafield types are correct when using the GetByKey method", func() {
	var fake faker.Faker

	BeforeEach(func() {
		fake = faker.New()
	})

	Describe("Given I have a single metafield", func() {
		Context("When I ask to get that metafield by its key when its of type boolean", func() {
			It("Then the value is returned in the correct type", func() {
				// Expects
				booleanString := fake.Boolean().BoolString("true", "false")
				expectedValue, err := strconv.ParseBool(booleanString)
				Expect(err).To(Not(HaveOccurred()))

				metafield := Metafield{}
				fake.Struct().Fill(&metafield)
				metafield.Type = BooleanMetaFieldType
				metafield.Value = booleanString
				metafields := Metafields{metafield}

				retValue, err := metafields.GetByKey(metafield.Key)

				Expect(err).To(Not(HaveOccurred()))
				Expect(retValue.(bool)).To(Equal(expectedValue))
			})
		})

		Context("When I ask to get that metafield by its key when its of type string", func() {
			It("Then the value is returned in the correct type", func() {
				expectedValue := fake.Color().Hex()

				metafield := Metafield{}
				fake.Struct().Fill(&metafield)
				metafield.Type = ColorMetaFieldType
				metafield.Value = expectedValue
				metafields := Metafields{metafield}

				retValue, err := metafields.GetByKey(metafield.Key)

				Expect(err).To(Not(HaveOccurred()))
				Expect(retValue.(string)).To(Equal(expectedValue))
			})
		})

		Context("When I ask to get that metafield by its key when its of type integer", func() {
			It("Then the value is returned in the correct type", func() {
				// expected
				expectedValue := fake.Int64()

				metafield := Metafield{}
				fake.Struct().Fill(&metafield)
				metafield.Type = NumberIntegerMetaFieldType
				metafield.Value = strconv.FormatInt(expectedValue, 10)
				metafields := Metafields{metafield}

				retValue, err := metafields.GetByKey(metafield.Key)

				Expect(err).To(Not(HaveOccurred()))
				Expect(retValue.(int64)).To(Equal(expectedValue))
			})
		})

		Context("When I ask to get that metafield by an invalid key", func() {
			It("Then the correct error message is returned", func() {
				metafield := Metafield{}
				fake.Struct().Fill(&metafield)
				metafield.Key = ""
				metafields := Metafields{metafield}

				_, err := metafields.GetByKey("123")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
