package animegirls_books

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var Languages = []string{"Al", "APL", "ASM", "Ada", "Agda", "Algorithms", "Architecture", "Beef", "C#", "C++", "C", "CSS", "Cobol", "Compilers", "D", "Dart", "Delphi", "Design Patterns", "Editors", "Elixir", "Elm", "F#", "FORTH", "Fortran", "GDScript", "Go", "Haskell", "HoTT", "HolyC", "Idris", "Java", "Javascript", "Kotlin", "Lisp", "Lua", "Math", "Memes", "Mixed", "MongoDB", "Nim", "OCaml", "Objective-C", "Other", "PHP", "Perl", "Personification", "Prolog", "Python", "Quantum Computing", "R", "Racket", "RayTracing", "ReCT", "Regex", "Ruby", "Rust", "SICP", "SQL", "Scala", "Shell", "Smalltalk", "Solidity", "Swift", "Systems", "Typescript", "Uncategorized", "Unity", "Unreal", "V", "VHDL", "Verilog", "WebGL"}

var rawLink = "https://raw.githubusercontent.com/cat-milk/Anime-Girls-Holding-Programming-Books/master/"

func GetRandomImage(l string) (string, error) {
	lang, err := IsLanguage(l)
	if err != nil {
		return "", err
	}

	images, err := GetImages(lang)
	if err != nil {
		return "", err
	}

	// Initializing the random seed
	rand.Seed(time.Now().UnixNano())

	rbIndex := rand.Intn(len(images))
	rbImage := images[rbIndex]
	return rbImage, nil
}

func GetImages(l string) ([]string, error) {
	lang, err := IsLanguage(l)
	if err != nil {
		return []string{}, err
	}

	body, err := Request(lang)
	if err != nil {
		return []string{}, err
	}

	// The regular expression to find all <fileName>.jpg images
	regJpg, err := regexp.Compile(`\w+?(?:(-)\w|\w)+\w+[a-zA-Z]\.jpg`)
	if err != nil {
		return []string{}, err
	}

	// The regular expression to find all <fileName>.png images
	regPng, err := regexp.Compile(`\w+?(?:(-)\w|\w)+\w+[a-zA-Z]\.png`)
	if err != nil {
		return []string{}, err
	}

	// Remove duplicate strings
	jpgIm := regJpg.FindAllString(string(body), -1)
	jpgIm = removeDuplicateStr(jpgIm)

	pngIm := regPng.FindAllString(string(body), -1)
	pngIm = removeDuplicateStr(pngIm)

	// Combine the two slices
	im := append(jpgIm, pngIm...)

	// Making each item a raw link
	for i, item := range im {
		im[i] = concatRawLink(lang, item)
	}

	return im, nil
}

func Request(l string) ([]byte, error) {
	lang, err := IsLanguage(l)
	if err != nil {
		return []byte{}, err
	}

	link := fmt.Sprintf("https://github.com/cat-milk/Anime-Girls-Holding-Programming-Books/tree/master/%s", lang)

	res, err := http.Get(link)
	if err != nil {
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func IsLanguage(str string) (string, error) {
	var ok bool
	var l string
	err := errors.New("Not a language")

	for _, lang := range Languages {
		if strings.ToLower(str) == strings.ToLower(lang) {
			ok = true
			l = lang
			break
		}
	}

	if !ok {
		return "", err
	}

	return l, nil
}

func concatRawLink(lang string, im string) string {
	return fmt.Sprintf("%s%s/%s", rawLink, lang, im)
}

func removeDuplicateStr(slice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}

	for _, item := range slice {
		if item == "fluidicon.png" || item == "favicon.png" {
			continue
		}

		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	return list
}
