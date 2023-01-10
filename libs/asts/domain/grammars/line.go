package grammars

type line struct {
	containers []Container
}

func createLine(
	containers []Container,
) Line {
	out := line{
		containers: containers,
	}

	return &out
}

// Containers returns the containers
func (obj *line) Containers() []Container {
	return obj.containers
}
