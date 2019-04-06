# Bitly

[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/retgits/bitly)

If you'd like to use [Bitly](https://bitly.com) to shorten, brand, share, or retrieve data from links using [Go](https://golang.org), this is the right module you're looking for. This lib is a wrapper on top of the [v4 Bitly API](https://dev.bitly.com/v4_documentation.html).

## Why build this?

I think [Bitly](https://bitly.com) is an awesome service, but to interact with it using the API, and marshaling and unmarshalling the data every time was a bit of a mess. Because there wasn't a good Go module available yet, I decided to build my own.

## Usage

To use the Bitly module, you'll need to create a [Generic Access Token](https://bitly.is/accesstoken). That access token can be used to create the _bitly client_.

```go
import (
	"fmt"

	"github.com/retgits/bitly/client"
)

func main() {
	bitly := client.NewClient().WithAccessToken("<myAccessToken>")
	fmt.Println(bitly.AccessToken)
}
```

Depending on which type of resource you want to access, you'll need to import one of the services

```go
import (
    "github.com/retgits/bitly/client/bitlinks" // If you want to use the bitlinks resource
    "github.com/retgits/bitly/client/bsds" // If you want to use the bsds resource
    "github.com/retgits/bitly/client/groups" // If you want to use the groups resource
    "github.com/retgits/bitly/client/organizations" // If you want to use the organizations
    "github.com/retgits/bitly/client/users" // If you want to use the users resource
)
```

For example, getting the details of the user would be

```go
package main

import (
	"fmt"

	"github.com/retgits/bitly/client"
	"github.com/retgits/bitly/client/users"
)

func main() {
	// Generate a new Bitly client
	bitly := client.NewClient().WithAccessToken("<myAccessToken>")
	// Create a new users service
	usersSvc := users.New(bitly)
	// Get the details of your user
	me, err := usersSvc.RetrieveUser()
	if err != nil {
		fmt.Println(err.Error())
	}
    fmt.Printf("User details: %+v\n", me)
    // Result
    // User details: {Created:2018-08-05T22:08:53+0000 Modified:2018-08-05T22:29:42+0000 Login:myOrganization IsActive:true Is2FaEnabled:false Name:John Doe Emails:[{Email:user@example.org IsPrimary:true IsVerified:true}] IsSsoUser:false DefaultGroupGUID:myGuid}
}
```

## Contributing

If something is missing, or if you'd like to suggest new features feel free to [create an issue](https://github.com/retgits/bitly/issues/new) or a [PR](https://github.com/retgits/bitly/compare)! The code is structured as

```text
├── LICENSE
├── README.md
├── client
│   ├── bitlinks       <-- Bitlinks service
│   │   ├── api.go     <-- The types and helper methods for the service
│   │   └── service.go <-- The methods that can be used with this module
│   ├── bsds           <-- BSDs service
│   │   ├── api.go
│   │   └── service.go
│   ├── groups         <-- Groups service
│   │   ├── api.go
│   │   └── service.go
│   ├── http.go
│   ├── organizations  <-- Organizations service
│   │   ├── api.go
│   │   └── service.go
│   └── users          <-- Users service
│       ├── api.go
│       └── service.go
└── go.mod
```

## License

See the [LICENSE](./LICENSE) file in the repository

## Acknowledgements

A most sincere thanks to the team of [Bitly](https://bitly.com), for building such an awesome service that I enjoy every day!

_This package is not endorsed by Bitly_