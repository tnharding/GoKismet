package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

func downloadFile(filename string) {
	url := "http://192.168.82.1:4501/logfiles/" + filename
	fmt.Println("Downloading file", url)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	resp.Write(file)
}

// Retrieve each individual file
func retrievelogFiles(hrefs []string) {
	// Iterate over all of the Token's attributes until we find an "href"
	downloadFile(hrefs[0])
	//	for _, href := range hrefs {
	//		downloadFile(href)
	//	}

}

// Helper function to pull the href attribute from a Token
func getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return
}

func getAvailableLogFiles() (hrefs []string, err error) {
	hrefs = make([]string, 0)

	resp, err := http.Get("http://192.168.82.1:4501/logfiles/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Received link data")

		z := html.NewTokenizer(resp.Body)

		for {
			tt := z.Next()

			switch {
			case tt == html.ErrorToken:
				return hrefs, nil

			case tt == html.StartTagToken:
				t := z.Token()

				isAnchor := t.Data == "a"
				if !isAnchor {
					continue
				}

				ok, url := getHref(t)
				if !ok {
					continue
				}

				//fmt.Println("found link", url)
				hrefs = append(hrefs, url)
			}
		}
	}

	return nil, err
}

func main() {

	//	for {
	conn, err := net.Dial("tcp", "192.168.82.1:4501")
	if err != nil { //error connecting to kismet server
		fmt.Println("Error connecting to kismet server.")
		time.Sleep(time.Millisecond * 60000) //We will try to connect once a minute
		//continue
	}
	//We have connected to the server.
	//Do something
	fmt.Println("We have connected to the kismet server.")
	hrefs, err := getAvailableLogFiles()
	if len(hrefs) == 0 || err != nil {
		fmt.Println("No log files fount")
	} else {
		retrievelogFiles(hrefs)
	}
	conn.Close() //close the connection
	//time.Sleep(time.Millisecond * 60000) //We will try to connect once a minute
}

//}
