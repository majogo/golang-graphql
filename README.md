## Description

A GraphQL server in golang that contains music data from different artists. The music is order in three main objects `artist` ->  `disc` -> `track` -> `artist`. Queries for requesting all object from a type are given just as requesting specific objects. Mutation for create/deleting artists are added. Some example data are given.

The example was shown at the 2. Golang Meetup in Karlsruhe on the 24.10.2018.

## How to use
Exec "go run ./server/server.go" to start GraphQL server. Connect to http://localhost:8081/ for GraphQL playground.
The dependencies can be install by `glide` with `glide install` and `glide update`.

## Modify schema
The following steps are necessary to modify the schema:
* modify the `schema.graphql`
* code generation by `go:generate gorunpkg github.com/99designs/gqlgen`
* depending on the modification, adding and implementing new resolver(s)

