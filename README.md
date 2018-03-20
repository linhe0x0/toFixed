# toFixed

[![Build Status](https://travis-ci.org/sqrthree/toFixed.svg?branch=master)](https://travis-ci.org/sqrthree/toFixed)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](http://makeapullrequest.com)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/sqrthree/toFixed/blob/master/LICENSE)

> Formats a number using fixed-point notation.

## Installing

Use [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install and update:

```shell
go get -u github.com/sqrthree/toFixed
```

## Usage

Package toFixed will export a function `ToFixed` that formats a number using fixed-point notation.

## Examples

```go
ToFixed(1.2345678, 2) // => 1.23 float64 
ToFixed(1.2345678, 5) // => 1.23457 float64
```

Go to [./examples](https://github.com/sqrthree/toFixed/blob/master/examples/hello-world.go) to see details.

## Note

Function `ToFixed` will return a float64 value, if you want to get a string value, you can use `fmt.Sprintf()`.

```go
package main

import (
	"fmt"
)

func main() {
	var f float64 = 18.923487203
  
	fmt.Println(fmt.Sprintf("%.5f", f))
}
```

## Tests

```shell
go test -v
```

---

> [sqrtthree.com](http://sqrtthree.com/) &nbsp;&middot;&nbsp;
> GitHub [@sqrthree](https://github.com/sqrthree) &nbsp;&middot;&nbsp;
> Twitter [@sqrtthree](https://twitter.com/sqrtthree)
