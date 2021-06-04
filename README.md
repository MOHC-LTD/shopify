# shopify

Package `shopify` provides Shopify types and interfaces.

See the [Shopify developer documentation](https://shopify.dev).

## Contents

- [Github authentication](#github-authentication)
- [Installation](#installation)
- [Implementations](#implementations)
- [How to contribute](#how-to-contribute)

## Github authentication

As this is a private Github repository, you will need to set up private go modules.

First set the `GOPRIVATE` go environment variable.

```sh
go env -w GOPRIVATE=github.com/MOHC-LTD
```

Generate a [Github personal access token](https://github.com/settings/tokens), and set up
global Github authentication on your machine

```sh
git config --global url."https://${username}:${access_token}@github.com".insteadOf "https://github.com"
```

## Docker

To build applications that consuming this module using docker, you will need to allow the docker container to authenticate with Github.

Do this by adding the following lines to your Dockerfile.

```sh
ARG authToken

RUN go env -w GOPRIVATE=github.com/MOHC-LTD

RUN apk add git

RUN git config --global url."https://golang:$authToken@github.com".insteadOf "https://github.com"
```

Then, when building your container, set the docker argument `authToken` to the value of your Github access token.

## Installation

Install the module using

```sh
go get -u github.com/MOHC-LTD/shopify
```

## Implementations

- [mockshopify](https://github.com/MOHC-LTD/mockshopify) is a [gomock](https://github.com/golang/mock) implementation of the [shopify](https://github.com/MOHC-LTD/shopify) package.
- [httpshopify](https://github.com/MOHC-LTD/httpshopify) is a http REST implementation of the [shopify](https://github.com/MOHC-LTD/shopify) package. It communicates with Shopify through the [admin API](https://shopify.dev/docs/admin-api/rest/reference).

## How to contribute

Something missing or not working as expected? See our [contribution guide](./CONTRIBUTING.md).
