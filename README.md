# GO-React starter

## Getting started

### Setup - Without Docker

- Download and install [golang](https://golang.org)

- Download and install [postgres](https://www.postgresql.org/download/)

  - [Setup Postgres](https://www.codementor.io/engineerapart/getting-started-with-postgresql-on-mac-osx-are8jcopb): Setting up postgres on a mac

  - [Setup postgres- windows](https://www.robinwieruch.de/postgres-sql-windows-setup/): Setting up postgres on windows

### Configure project - Without Docker

Use the queries in the [server/db/.psql](./server/db/.postgres) file to setup the database.

Enter the DB creds in the [server/config/](./server/config/config.go) file

Navigate to the server directory

```bash
> cd server
> go run main.go
```

This will start the go server.

To start the react app navigate to the client directory

```bash
> cd client
> yarn install
> yarn start
```

## Backend

### Endpoints

- /ping [GET]

- /register [POST]

```js
       {
         name String,
         email String,
         password String
       }
```

- /login [POST]

```js
       {
         email String,
         password String
       }
```

- /session [GET] (Auth required)

- /customers [GET] (Auth required)

- /customers/:id [GET] (Auth required)

- /customers/:id/paymentmethods [GET] (Auth required)

- /customers/:id/paymentmethods/count [GET] (Auth required)

- /customers [POST] - Add or Update (Auth required)

```js
    {
      customerId: String,
      registrationTime: number,
      email: String,
      emailVerifiedTime: number,
      name: String,
      familyName: String,
      givenName: String,
      telephone: String,
      telephoneVerifiedTime: number,
      telephoneCountry: String,
      location: {
        country: String,
        postalCode: String,
        latitude: double,
        longitude: double,
        addresseeName: String,
        street1: String,
        street2: String,
        neighbourhood: String,
        zone: String,
        city: String,
        region: String,
        poBoxNumber: String,
      },
    }
```

- /customers/:id [DELETE] (Auth required)

- /payments [GET] (Auth required)

- /payments/:id [GET] (Auth required)

- /payments [POST] (Auth required)

```js
    {
      paymentMethodId: String,
      customerId: String,
      methodType: String,
      instrumentId: String,
      cardBin: String,
      cardLastFour: String,
      expiryMonth: number,
      expiryYear: number,
      eWallet: String,
      nameOnCard: String,
      billingAddress: {
        country: String,
        postalCode: String,
        latitude: double,
        longitude: double,
        addresseeName: String,
        street1: String,
        street2: String,
        neighbourhood: String,
        zone: String,
        city: String,
        region: String,
        poBoxNumber: String,
      },
      successfulRegistration: bool,
      registrationTime: number,
      lastVerified: number,
    }
```

- /payments/:id [DELETE] (Auth required)

### Routes

- /ping
- /login
- /register
- /session
- /customers
- /customers/:id
- /customers/:id/paymentmethods
- /customers/:id/paymentmethods/count
- /customers
- /customers/:id
- /payments
- /payments/:id
- /payments
- /payments/:id

## Frontend

### Routes

- / (Redirect login...)
- /login (Login User)
- /register (Register new User)
- /customers (Customers - static list )

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
This will start the go server.

To start the react app navigate to the client directory
