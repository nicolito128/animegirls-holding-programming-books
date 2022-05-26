# Anime Girls Holding Programming Books
This module provides a simple functionality to obtain images of anime-girls holding well-known programming books. The module is based on the eponymous project [`cat-milk/Anime-Girls-Holding-Programming-Books`](https://github.com/cat-milk/Anime-Girls-Holding-Programming-Books). All image rights and credits go to their respective authors, this repository is intended only to provide a simple functionality to use the images provided.

The module does not handle, and does not plan to handle, +18 (adult content) images.

## How to use
Import the module:
```go
import (
    animegirls "github.com/nicolito128/animegirls-holding-programming-books"
)
```

Now you can use it with the functions GetImages or GetRandomImage. For example:
```go
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
```

Accepted languages are only those available in the original `cat-milk` repository. If you want more images then contribute, it's open source.

## All
* `Using GetImages`:
```go
im, err := animegirls_books.GetImages("ocaml")
if err != nil {
	panic(err)
}

// im is a []string
fmt.Println(im...)
```

* `Using GetRandomImage`:
```go
im, err := animegirls_books.GetRandomImage("go")
if err != nil {
	panic(err)
}

fmt.Println(im)
```

* `Using IsLanguage`:
```go
lang, err := animegirls_books.IsLanguage("ruby")
if err != nil {
	fmt.Println("Not language")
} else {
	fmt.Println(lang)
}
```

* `Using Request`:
```go
body, err := animegirls_books.Request("dart")
if err != nil {
	panic(err)
}

// Body of: https://github.com/cat-milk/Anime-Girls-Holding-Programming-Books/tree/master/<language>
fmt.Println(string(body))
```

## Interest links
* [Waffer discord bot](https://github.com/nicolito128/waffer)
* [Anime-Girls-Holding-Programming-Books](https://github.com/cat-milk/Anime-Girls-Holding-Programming-Books)