package references

type reference struct {
	contentKeys ContentKeys
	commits     Commits
}

func createReference(
	contentKeys ContentKeys,
	commits Commits,
) Reference {
	out := reference{
		contentKeys: contentKeys,
		commits:     commits,
	}

	return &out
}

// ContentKeys returns the contentKeys
func (obj *reference) ContentKeys() ContentKeys {
	return obj.contentKeys
}

// Commits returns the commits, if any
func (obj *reference) Commits() Commits {
	return obj.commits
}
