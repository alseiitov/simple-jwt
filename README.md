# simple-jwt
A very simple JWT implementation written in Go.

 
### Install
```sh
$ go get github.com/alseiitov/simple-jwt
```


### Usage
```go
package main

import (
	"fmt"

	sjwt "github.com/alseiitov/simple-jwt"
)

var secret = "53cr3tk3y"

func main() {
	// Create new token
	header := make(map[string]interface{})
	header["alg"] = "HS256"
	header["typ"] = "JWT"

	payload := make(map[string]interface{})
	payload["user"] = "alseiitov"

	jwt := sjwt.New(header, payload)
	token, _ := jwt.Sign(secret)

	// Verify token
	jwt2, _ := sjwt.Parse(token)
	if err := jwt2.Verify(token, secret); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK!")
		fmt.Println(jwt.Payload["user"])
	}
}

```