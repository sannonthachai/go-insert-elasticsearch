package main

import (
	"github.com/elastic/go-elasticsearch/v8"
)

func deleteIndex(client *elasticsearch.Client) {
	listIndex := []string{
		"jaeger-span-2024-10-01",
		"jaeger-span-2024-10-02",
		"jaeger-span-2024-10-03",
		"jaeger-span-2024-10-04",
		"jaeger-span-2024-10-05",
		"jaeger-span-2024-10-06",
		"jaeger-service-2024-10-01",
		"jaeger-service-2024-10-02",
		"jaeger-service-2024-10-03",
		"jaeger-service-2024-10-04",
		"jaeger-service-2024-10-05",
		"jaeger-service-2024-10-06",
	}
	client.Indices.Delete(listIndex)
}
