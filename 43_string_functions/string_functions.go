package main

import (
	"fmt"
	"strings"
)

var log = fmt.Println

func main() {
	log("Join:", strings.Join([]string{"a", "b"}, "-"))
	log("Count:", strings.Count("helllo", "l"))
	log("Repeat:", strings.Repeat("a", 3))
	log("Replace:", strings.Replace("foooooooo", "o", "0", 3))
}
