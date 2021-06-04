# shopify

Package `shopify` provides Shopify types and interfaces.

See the [Shopify developer documentation](https://shopify.dev).

## Contents

- [Github authentication](#github-authentication)
- [Installation](#installation)
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

## Installation

Install the module using

```sh
go get -u github.com/MOHC-LTD/shopify
```

## How to contribute

Something missing or not working as expected? See our [contribution guide](./CONTRIBUTING.md).
