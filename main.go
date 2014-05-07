// hello project main.go
package main

import (
	"flag"
	"github.com/ed0wolf/gojistatic"
)

func main() {
	flag.Parse()

	gojistatic.Start()
}
