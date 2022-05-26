package animegirls_books_test

import (
	"testing"

	animegirls "github.com/nicolito128/animegirls-holding-programming-books"
)

func TestRequest(t *testing.T) {
	t.Run("Should return a valid response", func(t *testing.T) {
		body, err := animegirls.Request("Go")
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if len(body) == 0 {
			t.Error("Body is empty")
		}
	})

	t.Run("Should return an error", func(t *testing.T) {
		_, err := animegirls.Request("")
		if err == nil {
			t.Error("Expected an error")
		}
	})
}

func TestGetImages(t *testing.T) {
	t.Run("Should return a valid slice", func(t *testing.T) {
		images, err := animegirls.GetImages("Go")
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if len(images) == 0 {
			t.Error("Images is empty")
		}
	})

	t.Run("Should return an error", func(t *testing.T) {
		_, err := animegirls.GetImages("")
		if err == nil {
			t.Error("Expected an error")
		}
	})
}

func TestGetRandomImage(t *testing.T) {
	t.Run("Should return a valid image", func(t *testing.T) {
		image, err := animegirls.GetRandomImage("Go")
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if len(image) == 0 {
			t.Error("Image is empty")
		}
	})

	t.Run("Should return an error", func(t *testing.T) {
		_, err := animegirls.GetRandomImage("")
		if err == nil {
			t.Error("Expected an error")
		}
	})
}
