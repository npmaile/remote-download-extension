package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/downloader", func(_ http.ResponseWriter, r *http.Request) {
		var reqURL string
		json.NewDecoder(r.Body).Decode(&reqURL)
		err := downloadFile(reqURL)
		if err != nil {
			fmt.Println(err.Error())
		}
	})
	http.ListenAndServe(":5432", nil)
}

func downloadFile(rawURL string) (err error) {
	urlParsed, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	urlParts := strings.Split(urlParsed.Path, "/")
	os.MkdirAll(urlParsed.Host+"/"+strings.Join(urlParts[:len(urlParts)-1], "/"), os.ModeDir)
	// Create the file
	out, err := os.Create(urlParsed.Host + "/" + urlParsed.Path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(rawURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
