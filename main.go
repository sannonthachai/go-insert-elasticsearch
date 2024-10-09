package main

import (
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200/",
		},
		Username: "elastic",
		Password: "1234",
	}

	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	createIndexWithMapping(es)

	docs := generateData()

	for _, bod := range docs {
		res, err := es.Index(
			"test-1234",            // Index name
			strings.NewReader(bod), // Document to index
			es.Index.WithDocumentID(generateRandomString(20)),
			es.Index.WithRefresh("true"),
			es.Index.WithPretty(),
			es.Index.WithTimeout(100),
		)
		if err != nil {
			log.Fatalf("Error indexing document: %s", err)
		}
		log.Println(res)
	}
}

func createIndexWithMapping(client *elasticsearch.Client) {
	index := "test-1234"
	mapping := `
    {
      "mappings": {
        "properties": {
          "bpmnProcessId": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "flowNodes": {
            "properties": {
              "id": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "name": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              }
            }
          },
          "id": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "key": {
            "type": "float"
          },
          "name": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "partitionId": {
            "type": "long"
          },
          "startedByForm": {
            "type": "boolean"
          },
          "version": {
            "type": "long"
          }
        }
      },
      "settings": {
        "routing": {
          "allocation": {
            "include": {
              "_tier_preference": "data_content"
            }
          }
        },
        "number_of_shards": "1",
        "number_of_replicas": "1"
      },
    }`

	res, err := client.Indices.Create(
		index,
		client.Indices.Create.WithBody(strings.NewReader(mapping)),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
