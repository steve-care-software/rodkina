package queries

type fetcher struct {
	recursive string
	query  Query
}

func createFetcherWithRecursive(
	recursive string,
) Fetcher {
	return createFetcherInternally(recursive, nil)
}

func createFetcherWithQuery(
	query Query,
) Fetcher {
	return createFetcherInternally("", query)
}

func createFetcherInternally(
	recursive string,
	query Query,
) Fetcher {
	out := fetcher{
		recursive: recursive,
		query:  query,
	}

	return &out
}

// IsRecursive returns true if recursive, false otherwise
func (obj *fetcher) IsRecursive() bool {
	return obj.recursive != ""
}

// Recursive returns the recursive query's token name
func (obj *fetcher) Recursive() string {
	return obj.recursive
}

// IsQuery returns true if query, false otherwise
func (obj *fetcher) IsQuery() bool {
	return obj.query != nil
}

// Query returns the query if any
func (obj *fetcher) Query() Query {
	return obj.query
}
