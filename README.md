# Elasticsearch for Gophers

This project contains an end-to-end example that showcases different features from the official [Go Client for Elasticsearch](https://github.com/elastic/go-elasticsearch) that you can use as a reference about how to get started with Elasticsearch in your Go projects. It is not intended to provide the full spectrum of what the client is capable of — but it certainly puts you on the right track.

You can run this code with an Elasticsearch instance running locally, to which you can leverage the [Docker Compose code](./docker-compose.yml) available in the project. Alternatively, you can also run this code with an Elasticsearch instance from Elastic Cloud that can be easily created using the [Terraform code](./elastic-cloud.tf) also available in the project.

## Examples available in this project

The data model from this project is about a collection of movies available in the file [movies.json](./movies.json). This file will be [loaded](logic/movies.go) in memory and made available within the context, which the other functions will look up and work with.

```bash
docker compose -f run-with-collector.yaml up -d
```

# License

This project is licensed under the [Apache 2.0 License](./LICENSE).