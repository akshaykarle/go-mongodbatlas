# go-mongodbatlas [![Build Status](https://travis-ci.org/akshaykarle/go-mongodbatlas.svg?branch=master)](https://travis-ci.org/akshaykarle/go-mongodbatlas) [![GoDoc](https://godoc.org/github.com/akshaykarle/go-mongodbatlas/mongodb?status.png)](https://godoc.org/github.com/akshaykarle/go-mongodbatlas/mongodb) [![codecov](https://codecov.io/gh/akshaykarle/go-mongodbatlas/branch/master/graph/badge.svg)](https://codecov.io/gh/akshaykarle/go-mongodbatlas)
A Go client library for the [MongoDB Atlas API](https://docs.atlas.mongodb.com/api/).

## Getting started
MongoDB Atlas uses the [Digest Access Authentication](https://tools.ietf.org/html/rfc2069) and doesn't support Basic Auth. Follow the examples in [examples](/examples) directory. We create a httpClient using [go-http-digest-auth-client](https://github.com/xinsnake/go-http-digest-auth-client) and pass it over to mongodbatlas. To run a mongo atlas cluster example, just run:
```
go run examples/clusters.go <username> <mongodb-atlas-api-key> <group-id>
```

## Development & Contributing
### Installing dependencies
```
go get github.com/golang/lint/golint
go get github.com/akshaykarle/go-http-digest-auth-client
go get -v -t ./mongodbatlas
```

### Testing
Run `./test.sh`

## TODO
* Add support for Monitoring & Logging
