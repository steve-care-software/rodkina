package applications

import "github.com/steve-care-software/rodkina/libs/asts/domain/grammars"

type stack struct {
	token grammars.Token
	lines map[int][]byte
}
