package queries

import "errors"

type fetcherBuilder struct {
	recursive string
	query  Query
}

func createFetcherBuilder() FetcherBuilder {
	out := fetcherBuilder{
		recursive: "",
		query:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *fetcherBuilder) Create() FetcherBuilder {
	return createFetcherBuilder()
}

// WithRecursive adds a recursive query's token name to the builder
func (app *fetcherBuilder) WithRecursive(recursive string) FetcherBuilder {
	app.recursive = recursive
	return app
}

// WithQuery adds a query to the builder
func (app *fetcherBuilder) WithQuery(query Query) FetcherBuilder {
	app.query = query
	return app
}

// Now builds a new Fetcher instance
func (app *fetcherBuilder) Now() (Fetcher, error) {
	if app.recursive != "" {
		return createFetcherWithRecursive(app.recursive), nil
	}

	if app.query != nil {
		return createFetcherWithQuery(app.query), nil
	}

	return nil, errors.New("the Fetcher is invalid")
}
