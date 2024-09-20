package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

type Item struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	PubDate     time.Time `xml:"pubDate"`
}

type Channel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	PubDate     time.Time `xml:"pubDate"`
	Items       []Item    `xml:"item"`
}

// create the function to generate the rss feeds
func generateRSSFeed() ([]byte, error) {
	items := []Item{
		{
			Title:       "Article 1",
			Link:        "https://example.com/article1",
			Description: "This is the description of Article 1.",
			PubDate:     time.Now(),
		},
		{
			Title:       "Article 2",
			Link:        "https://example.com/article2",
			Description: "This is the description of Article 2.",
			PubDate:     time.Now().Add(-1 * time.Hour), // One hour ago
		},
		{
			Title:       "Artricle 2",
			Link:        "https://example.com/article3",
			Description: "This is the description for Article 3.",
			PubDate:     time.Now().Add(-2 * time.Hour), // Two hours ago
		},
		{
			Title:       "Artricle 4",
			Link:        "https://example.com/article4",
			Description: "This is the description for Article 4.",
			PubDate:     time.Now().Add(-3 * time.Hour), // Three hours ago
		},
	}

	// RSS feed data
	feed := Channel{
		Title:       "Sample RSS Feed",
		Link:        "https://example.com",
		Description: "This is a sample RSS feed generated using Golang.",
		PubDate:     time.Now(),
		Items:       items,
	}

	// Marshal the data to XML
	xmlData, err := xml.MarshalIndent(feed, "", "   ")
	if err != nil {
		return nil, err
	}

	// Add XML header to the feed
	rssFeed := []byte(xml.Header + string(xmlData))

	return rssFeed, nil

}

func main() {
	// generate the rss feeds
	rssFeed, err := generateRSSFeed()
	if err != nil {
		fmt.Println("Error generating rss feeds:", err)
		return
	}

	// Write feed to a file
	file, err := os.Create("feed.xml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	_, err = file.Write(rssFeed)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("RSS feed generated and saved to feed.xml")

}
