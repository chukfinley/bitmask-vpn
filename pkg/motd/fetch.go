package motd

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"0xacab.org/leap/bitmask-vpn/pkg/config"
)

func FetchLatest() []Message {
	empty := []Message{}
	if os.Getenv("SKIP_MOTD") == "1" {
		return empty
	}
	url := ""
	switch config.Provider {
	case "riseup.net":
		url = "https://downloads.leap.se/motd/riseup/motd.json"
	default:
		return empty
	}
	log.Println("Fetching MOTD for", config.Provider)
	b, err := fetchURL(url)
	if err != nil {
		log.Println("WARN Error fetching json from", url)
		return empty
	}
	allMsg, err := getFromJSON(b)
	if err != nil {
		log.Println("WARN Error parsing json from", url)
		return empty
	}
	valid := empty[:]
	if allMsg.Length() != 0 {
		log.Printf("There are %d pending messages\n", allMsg.Length())
	}
	for _, msg := range allMsg.Messages {
		if msg.IsValid() {
			valid = append(valid, msg)
		}
	}
	return valid
}

func fetchURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}