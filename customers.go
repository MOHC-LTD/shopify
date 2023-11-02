package shopify

import (
	"fmt"
	"net/url"
	"time"
)

// Customer is a shopify customer
type Customer struct {
	// ID is the ID of the customer.
	ID int64
	// Email is the unique email address of the customer.
	Email string
	// Phone is the unique phone number (E.164 format) for this customer.
	Phone string
	// FirstName is the customer's first name.
	FirstName string
	// LastName is the customer's last name.
	LastName string
	// Tags are the tags linked to the customer.
	Tags Tags
	// CreatedAt is the date and time when the customer was created.
	CreatedAt time.Time
	// UpdatedAt is the date and time (ISO 8601 format) when the customer information was last updated.
	UpdatedAt time.Time
	// Addresses are the addresses linked to the customer.
	Addresses CustomerAddresses
	// Metafields are any meta field data linked to the customer
	Metafields Metafields
}

// Customers is a collection of customers
type Customers []Customer

// CustomerSearchQuery contains all the supported search queries
type CustomerSearchQuery struct {
	AcceptsMarketing        string
	ActivationDate          string
	Address1                string
	Address2                string
	City                    string
	Company                 string
	Country                 string
	CustomerDate            string
	CustomerFirstName       string
	CustomerID              string
	CustomerLastName        string
	CustomerTag             string
	Email                   string
	EmailMarketingState     string
	FirstName               string
	FirstOrderDate          string
	ID                      string
	LastAbandonedOrderDate  string
	LastName                string
	MultipassIdentifier     string
	OrdersCount             string
	OrderDate               string
	Phone                   string
	Province                string
	ShopID                  string
	State                   string
	Tag                     string
	TotalSpent              string
	UpdatedAt               string
	VerifiedEmail           string
	ProductSubscriberStatus string
}

// String converts the search query into a string
func (c *CustomerSearchQuery) String() string {
	params := url.Values{}

	queryData := map[string]string{
		"accepts_marketing":         c.AcceptsMarketing,
		"activation_date":           c.ActivationDate,
		"address1":                  c.Address1,
		"address2":                  c.Address2,
		"city":                      c.City,
		"company":                   c.Company,
		"country":                   c.Country,
		"customer_date":             c.CustomerDate,
		"customer_first_name":       c.CustomerFirstName,
		"customer_id":               c.CustomerID,
		"customer_last_name":        c.CustomerLastName,
		"customer_tag":              c.CustomerTag,
		"email":                     c.Email,
		"email_marketing_state":     c.EmailMarketingState,
		"first_name":                c.FirstName,
		"first_order_date":          c.FirstOrderDate,
		"id":                        c.ID,
		"last_abandoned_order_date": c.LastAbandonedOrderDate,
		"last_name":                 c.LastName,
		"multipass_identifier":      c.MultipassIdentifier,
		"orders_count":              c.OrdersCount,
		"order_date":                c.OrderDate,
		"phone":                     c.Phone,
		"province":                  c.Province,
		"shop_id":                   c.ShopID,
		"state":                     c.State,
		"tag":                       c.Tag,
		"total_spent":               c.TotalSpent,
		"updated_at":                c.UpdatedAt,
		"verified_email":            c.VerifiedEmail,
		"product_subscriber_status": c.ProductSubscriberStatus,
	}

	for queryKey, queryValue := range queryData {
		if queryValue != "" {
			params.Add(fmt.Sprintf("%v:", queryKey), queryValue)
		}
	}

	return params.Encode()
}

// CustomerRepository maintains the customers of a shop.
type CustomerRepository interface {
	// Update updates a single customer
	Update(customer Customer) (Customer, error)
	// Get gets customer with the provided id
	Get(id int64) (Customer, error)
	// GetByQuery gets the customers matching the passed query
	GetByQuery(fields []string, query CustomerSearchQuery) (Customers, error)
}

// CustomerRepository manages customers
type CustomerRepository interface {
	// Orders retrieves a list of orders belonging to a customer
	Orders(id int64, query OrderQuery) (Orders, error)
}
