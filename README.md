# shopify

Package `shopify` provides Shopify types and interfaces.

See the [Shopify developer documentation](https://shopify.dev).

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Implementations](#implementations)
- [How to contribute](#how-to-contribute)

## Installation

Install the module using

```sh
go get -u github.com/MOHC-LTD/shopify
```

## Usage

Use this package to create functionality backed by Shopify.

```go
type ExampleService struct {
    shop shopify.Shop
}

func NewExampleService(shop shopify.Shop) ExampleService {
    return ExampleService{
        shop,
    }
}

func (service ExampleService) CloseOrder(orderID int64) {
    err := service.shop.Orders().Close(orderID)

    if err != nil {
        return err
    }

    return nil
}
```

Then either test the functionality using the [mockshopify](https://github.com/MOHC-LTD/mockshopify) package, or inject the [httpshopify](https://github.com/MOHC-LTD/httpshopify)
to create a real connection to Shopify.

## Implementations

- [mockshopify](https://github.com/MOHC-LTD/mockshopify) is a [gomock](https://github.com/golang/mock) implementation of the [shopify](https://github.com/MOHC-LTD/shopify) package.
- [httpshopify](https://github.com/MOHC-LTD/httpshopify) is a http REST implementation of the [shopify](https://github.com/MOHC-LTD/shopify) package. It communicates with Shopify through the [admin API](https://shopify.dev/docs/admin-api/rest/reference).

## How to contribute

Something missing or not working as expected? See our [contribution guide](./CONTRIBUTING.md).
