package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"shodan/shodan"

	"github.com/joho/godotenv"
)

func main() {
	var query string
	flag.StringVar(&query, "q", "", "query for search host")
	flag.Parse()
	if query == "" {
		log.Fatalln("query (--q) is required=")
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	if apiKey == "" {
		log.Fatalln("api key is missing")
	}
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", info)
	host, err := s.HostSearch(query)
	if err != nil {
		log.Fatalln(err)
	}
	for _, h := range host.Matches {
		fmt.Println(h.IPString, h.Port)
	}
}
