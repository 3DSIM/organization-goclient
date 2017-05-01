# organization-goclient
Client for interacting with the organization api

## Background Info
We use https://goswagger.io to generate our Go APIs and clients.  This allows
us to build our APIs in a "design first" manner.

First we create a `swagger.yaml` file that defines the API.  Then we run a command
to generate the server code.

Additionally, this allows us to automatically generate client code.  The code in this
directory was all generated using the `go-swagger` tools.


## Organization
* `organization` - the client package that adds convenience methods for common operations
* `genclient` - the generated client code
* `models` - the generated models

## Regenerating code
First install the swagger generator.  Currently we are using version 0.10.0 of https://github.com/go-swagger/go-swagger.

For mac users:
* brew tap go-swagger/go-swagger
* brew install go-swagger

For windows users:
* See https://github.com/go-swagger/go-swagger for options

The code generator needs a specification file.  The specification for the organization API is stored in `github.com/3dsim/organization-api/swagger.yaml`.  Assuming that project
is cloned as a sibling project, the command to run to generate new client code is:
```
swagger generate client -A OrganizationAPI -f ../organization-api/swagger.yaml --client-package genclient
```

* Generate fakes using counterfeiter
```
go get github.com/maxbrunsfeld/counterfeiter
```
From inside package folder
```
go generate
```

## Using the client
TODO

## Client to API version compatibility

| Organization API | Organization Client |
| ------------- | ------------- |
| TODO  | TODO |
