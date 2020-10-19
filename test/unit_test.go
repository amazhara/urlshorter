package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jminds/internal/urlshort"
)

const (
	url         = "http://localhost:8080/"
	contentType = "application/json"
)

var (
	fakeRequests = urlshort.Shorts{
		{
			"https://www.google.com/",
			"/g",
		},
		{
			"https://www.facebook.com/",
			"/f",
		},
	}
)

// TestAdd test /add API endpoint
func TestAdd(t *testing.T) {
	body, err := json.Marshal(&fakeRequests)

	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Post(url+"add", contentType, bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal(fmt.Errorf("Wrong status code = %v", res.StatusCode))
	}
}

func TestRedirect(t *testing.T) {
	res1, err := http.Get(url + fakeRequests[0].Short)

	if err != nil {
		t.Fatal(err)
	}

	res2, err := http.Get(url + fakeRequests[1].Short)

	if err != nil {
		t.Fatal(err)
	}

	if res1.StatusCode != http.StatusOK {
		t.Fatal(fmt.Errorf("Wrong status code = %v", res1.StatusCode))
	}

	if res2.StatusCode != http.StatusOK {
		t.Fatal(fmt.Errorf("Wrong status code = %v", res2.StatusCode))
	}
}
