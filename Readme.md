# Online Shopping Checkout

## Development

### Run CI Build

    make ci

### Run Tests

    make test

### Build Artefact

    make build

### Start

    make run

## Assumptions

This is a list of assumptions made during development for simplicity of implementation, and where requirements are vague:

1. The cart id will be provided by the client.  It will be some unique identifer tied to the user.

1. It is the responsibility of the client to collate items for checkout.

1. The inventory is static.  Orders are not made and stock is not depleted.

1. No panics will occur.

1. In the real world, the inventory would be backed by a real data store.

1. Prices don't change while service is running - this way we don't need to fetch price all the time.  Can just inject into promotions.  A real microservice might need to fetch prices.

1. Google *3 for the price of 2* offer applies for multiple quantities of 3.

1. Users won't try to order more than is in stock.

1. There is no requirement to add the free Raspberry Pi to the cart when a Macbook Pro is ordered.  Just because it is free doesn't necessarily mean the user wants one.

1. Discount percentages are always in whole percentages and are never less than 1%.

1. Unit prices of products are always 1c or higher.

1. Unit prices of products are always in whole cents.

1. Small quantities added to cart.

1. Prices are those typical for consumer goods and not very high.


# GraphQL

Please find the graphql in the `schema.graphql` file in the project root.