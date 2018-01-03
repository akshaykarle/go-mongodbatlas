# mongdb-atlas-go
A Go client library for the [MongoDB Atlas API](https://docs.atlas.mongodb.com/api/).

## Getting started
MongoDB Atlas uses the [Digest Access Authentication](https://tools.ietf.org/html/rfc2069) and doesn't support Basic Auth. Follow the example in [example.go](example.go). We create a httpClient using [go-http-digest-auth-client](https://github.com/xinsnake/go-http-digest-auth-client) and pass it over to mongodb. To run the example, just run:
```
go run example.go username mongodb-atlas-api-key group-id
```

## TODO
* Add support for Clusters
* Add support for Projects
* Add support for Setting up VPC peering
* Add support for Monitoring & Logging
