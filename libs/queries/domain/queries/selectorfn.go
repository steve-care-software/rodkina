package queries

type queryFn struct {
	single SingleQueryFn
	multi  MultiQueryFn
}

func createQueryFnWithSingle(
	single SingleQueryFn,
) QueryFn {
	return createQueryFnInternally(single, nil)
}

func createQueryFnWithMulti(
	multi MultiQueryFn,
) QueryFn {
	return createQueryFnInternally(nil, multi)
}

func createQueryFnInternally(
	single SingleQueryFn,
	multi MultiQueryFn,
) QueryFn {
	out := queryFn{
		single: single,
		multi:  multi,
	}

	return &out
}

// IsSingle returns true if single, false otherwise
func (obj *queryFn) IsSingle() bool {
	return obj.single != nil
}

// Single returns the single query func, if any
func (obj *queryFn) Single() SingleQueryFn {
	return obj.single
}

// IsMulti returns true if multi, false otherwise
func (obj *queryFn) IsMulti() bool {
	return obj.multi != nil
}

// Multi returns the single multi func, if any
func (obj *queryFn) Multi() MultiQueryFn {
	return obj.multi
}
