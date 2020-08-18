package elasticsearch

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

var Es *elasticsearch.Client

func init() {
	var (
		r map[string]interface{}
	)

	Es, _ = elasticsearch.NewDefaultClient()
	res, err := Es.Info()
	if err != nil {
		log.Fatalf("ElasticSearch Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("ElasticSearch Error: %s", res.String())
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("ElasticSearch Error parsing the response body: %s", err)
	}
	log.Printf("ElasticSearch Client: %s", elasticsearch.Version)
	log.Printf("ElasticSearch Server: %s", r["version"].(map[string]interface{})["number"])
}
