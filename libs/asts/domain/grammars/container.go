package grammars

type container struct {
	element Element
	compose Compose
}

func createContainerWithElement(
	element Element,
) Container {
	return createContainerInternally(element, nil)
}

func createContainerWithCompose(
	compose Compose,
) Container {
	return createContainerInternally(nil, compose)
}

func createContainerInternally(
	element Element,
	compose Compose,
) Container {
	out := container{
		element: element,
		compose: compose,
	}

	return &out
}

// Name returns the name
func (obj *container) Name() string {
	if obj.IsElement() {
		return obj.element.Name()
	}

	return obj.compose.Name()
}

// IsElement returns true if there is an element, false otherwise
func (obj *container) IsElement() bool {
	return obj.element != nil
}

// Element returns the element, if any
func (obj *container) Element() Element {
	return obj.element
}

// IsCompose returns true if there is a compose, false otherwise
func (obj *container) IsCompose() bool {
	return obj.compose != nil
}

// Compose returns the compose, if any
func (obj *container) Compose() Compose {
	return obj.compose
}
