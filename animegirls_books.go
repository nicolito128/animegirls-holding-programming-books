package animegirls_books

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var rawLink = "https://raw.githubusercontent.com/cat-milk/Anime-Girls-Holding-Programming-Books/master/"

func GetRandomImage(lang string) (string, error) {
	lang = toTitleCase(lang)
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

func GetImages(lang string) ([]string, error) {
	lang = toTitleCase(lang)
	body, err := Request(lang)
	if err != nil {
		return []string{}, err
	}

	// The regular expression to find all <fileName>.jpg images
	regJpg, err := regexp.Compile(`\w+[a-zA-Z]\.jpg`)
	if err != nil {
		return []string{}, err
	}

	// The regular expression to find all <fileName>.png images
	regPng, err := regexp.Compile(`\w+[a-zA-Z]\.png`)
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

func Request(lang string) ([]byte, error) {
	lang = toTitleCase(lang)
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

func toTitleCase(str string) string {
	title := strings.ToUpper(string(str[0])) + strings.ToLower(str[1:])
	return title
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
