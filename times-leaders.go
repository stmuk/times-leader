package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://m.apps.thetimes.co.uk/tto/feeds/deviceApp/article3015612.ece", nil)
	req.Header.Add("X-Requested-With", "uk.co.thetimes")
	req.Header.Add("X-NewsInternational-Times-Token", "XXXXXXXXXXXXXXX")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	type Item struct {
		Story   string `xml:"story"`
		Title   string `xml:"title"`
		PubDate string `xml:"pubDate"`
	}

	type Rss struct {
		Item []Item `xml:"channel>item"`
	}

	r := Rss{}

	xml.Unmarshal(body, &r)

	fmt.Print("<html><head><meta charset=\"utf-8\" /></head><body>")
	for _, item := range r.Item {
		fmt.Printf("%s<br/><br>\n", item.PubDate)
		fmt.Printf("%s", item.Title)
		fmt.Printf("%s<br/><br>\n", item.Story)
	}
	fmt.Print("</body></html>")

}
