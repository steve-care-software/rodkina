package grammars

type compose struct {
	name string
	list []ComposeElement
}

func createCompose(
	name string,
	list []ComposeElement,
) Compose {
	out := compose{
		name: name,
		list: list,
	}

	return &out
}

// Name returns the name
func (obj *compose) Name() string {
	return obj.name
}

// List returns the list of elements
func (obj *compose) List() []ComposeElement {
	return obj.list
}
