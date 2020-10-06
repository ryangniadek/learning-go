//Echo2 prints its command line arguments and the index of those arguments each on a new line
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("%d %s\n", index, arg)
	}
}
