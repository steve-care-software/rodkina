package grammars

type suite struct {
	isValid bool
	content Compose
}

func createSuiteWithValid(
	valid Compose,
) Suite {
	return createSuiteInternally(true, valid)
}

func createSuiteWithInvalid(
	invalid Compose,
) Suite {
	return createSuiteInternally(false, invalid)
}

func createSuiteInternally(
	isValid bool,
	content Compose,
) Suite {
	out := suite{
		isValid: isValid,
		content: content,
	}

	return &out
}

// IsValid returns true if valid, false otherwise
func (obj *suite) IsValid() bool {
	return obj.isValid
}

// Content returns the the content
func (obj *suite) Content() Compose {
	return obj.content
}
