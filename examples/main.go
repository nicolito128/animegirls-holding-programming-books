package main

import (
	"fmt"

	animegirls "github.com/nicolito128/animegirls-holding-programming-books"
)

func main() {
	im, err := animegirls.GetRandomImage("go")
	if err != nil {
		panic(err)
	}

	fmt.Println(im)
}
