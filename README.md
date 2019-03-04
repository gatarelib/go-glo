# go-glo

[![GoDoc](https://godoc.org/github.com/jackmcguire1/go-twitch-ext?status.svg)](https://godoc.org/github.com/jackmcguire1/go-glo)

[![Build Status](https://travis-ci.org/jackmcguire1/go-glo.svg?branch=master)](https://travis-ci.org/jackmcguire1/go-glo)
[![Go Report Card](https://goreportcard.com/badge/github.com/jackmcguire1/go-glo)](https://goreportcard.com/report/github.com/jackmcguire1/go-glo)


![GIT KRAKEN](https://cdn.worldvectorlogo.com/logos/gitkraken.svg)

[git]:      https://git-scm.com/
[golang]:   https://golang.org/
[releases]: https://github.com/jackmcguire1/Flexion-Coding-Challenge/releases/
[modules]:  https://github.com/golang/go/wiki/Modules

>A library to help interact with GitKraken [Glo Boards API](https://support.gitkraken.com/developers/api/)
<br>

## Supported Endpoints & Features

**API Endpoints:**
>This package supports the following v1 [Glo Boards API endpoints](https://gloapi.gitkraken.com/v1/docs/)

**Boards**

- [x] Get Boards
- [x] Get Boards by ID

**Columns**
- [x] Create column
- [x] Edit column
- [x] Delete column

**Cards**
- [x] Create Card
- [x] Edit Card
- [x] Delete Card
- [x] Get Cards
- [x] Get Card By ID
- [x] Get Cards By Column ID

**Attachments**
- [x] Create Attachment
- [x] Get Attachments

**Comments**
- [x] Create Comment
- [x] Edit Comment
- [x] Delete Comment
- [x] Get Comments By Card ID

**User**
- [x] Get User

## Installing
`go get github.com/jackmcguire1/go-glo`

## Example
```Go
package main

import (
	"log"
	"os"

	"github.com/jackmcguire1/go-glo"
)

var token string

func init() {
	token = os.Getenv("TOKEN")
}

func main() {
	client := glo.NewClient(token)
	user, err := client.GetUser()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user.ID, user.Name, user.Username, user.Email)
}
```

## Development

To develop `go-glo` or interact with its source code in any meaningful way, be
sure you have the following installed:

### Prerequisites

- [Git][git]
- [Go 1.11][golang]+

You will need to activate [Modules][modules] for your version of Go, generally
by invoking `go` with the support `GO111MODULE=on` environment variable set.

## FAQ
Please refer to [Git Kraken Documentation](https://support.gitkraken.com/developers/overview/) for any
further reading.
## License

[MIT]: https://opensource.org/licenses/MIT

The source code for go-glo is released under the [MIT License][MIT].

## Donations
All donations are appreciated!

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](http://paypal.me/crazyjack12)