package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var channelIDs = []string{
	//"UCJFZiqLMntJufDCHc6bQixg",
	//"UCp6993wxpyDPHUpavwDFqgg",
	//"UCDqI2jOz0weumE8s7paEk6g",
	//"UC-hM6YJuNYVAmUWxeIr9FeA",
	//"UC5CwaMl1eIgY8h02uZw7u8A",
	//"UCdn5BQ06XqgXoAxIhbqw5Rg",
	//"UCQ0UDLQCjY0rmuxCDE38FGg",
	//"UC1CfXB_kRs3C-zaeTG3oGyg",
	//"UCFTLzh12_nrtzqBPsTCqenA",
	//"UCD8HOxPs4Xvsm8H0ZxXGiBw",
	//"UC1suqwovbL1kzsoaZgFZLKg",
	//"UCp3tgHXw_HI0QMk1K8qh3gQ",
	//"UCXTpFs_3PqI41qX2d9tL2Rw",
	//"UC1opHUrw8rvnsadT-iGp7Cg",
	//"UCvzGlP9oQwU--Y0r9id_jnA",
	//"UC7fk0CB07ly8oSl0aqKkqFg",
	//"UC1DCedRgGHBdm81E1llLhOQ",
	//"UCl_gCybOJRIgOXw6Qb4qJzQ",
	//"UCvInZx9h3jC2JzsIzoOebWg",
	//"UCCzUftO8KOVkV4wQG1vkUvg",
	//"UCdyqAaZDKHXg4Ahi7VENThQ",
	//"UCZlDXzGoo7d44bwdNObFacg",
	//"UCS9uQI-jC3DE0L4IpXyvr6w",
	//"UCqm3BQLlJfvkTsX_hvm0UmA",
	//"UC1uv2Oq6kNxgATlCiez59hw",
	//"UCa9Y57gfeY0Zro_noHRVrnw",
	//"UCFKOVgVbGmX65RxO3EtH3iw",
	//"UCAWSyEs_Io8MtpY3m-zqILA",
	//"UCUKD-uaobj9jiqB-VXt71mA",
	//"UCK9V2B22uJYu3N7eR_BT9QA",
	//"UCp-5t9SrOQwXMU7iIjQfARg",
	//"UCvaTdHTWBGv3MKj3KVqJVCw",
	//"UChAnqc_AY5_I3Px5dig3X1Q",
	"UCOyYb1c43VlX9rc_lT6NKQw",
	"UCP0BspO_AMEe3aQqqpo89Dg",
	"UCAoy6rzhSf4ydcYjJw3WoVg",
	"UCYz_5n-uDuChHtLo7My1HnQ",
	"UC727SQYUvx5pDDGQpTICNWg",
	"UChgTyjG-pdNvxxhdsXfHQ5Q",
	"UCyl1z3jo3XHR1riLFKG5UAg",
	"UCL_qhgtOy0dy1Agp8vkySQg",
	"UCoSrY_IQQVpmIRZ9Xf-y93g",
	"UCMwGHR0BTZuLsmjY_NT5Pwg",
	"UCHsx4Hqa-1ORjQTh9TYDhww",
}

func main() {
	callbackURL := os.Args[1]
	for _, id := range channelIDs {
		log.Println(id)
		data := url.Values{}
		data.Set("hub.callback", callbackURL)
		data.Set("hub.topic", "https://www.youtube.com/xml/feeds/videos.xml?channel_id="+id)
		data.Set("hub.verify", "sync")
		data.Set("hub.mode", "subscribe")
		res, err := http.Post("https://pubsubhubbub.appspot.com/", "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.Status)
	}
}
