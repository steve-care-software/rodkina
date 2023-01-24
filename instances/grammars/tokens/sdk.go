package tokens

import (
	"github.com/steve-care-software/rodkina/instances/grammars/components"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
)

const byteLength = 256

// NewToken creates a new token instance
func NewToken() Token {
	component := components.NewComponent()
	return createToken(
		component,
	)
}

// Token represents the token component
type Token interface {
	VariableName() grammars.Token
	Sha512Hex() grammars.Token
	AnyHexCharacter() grammars.Token
	AnyLetter() grammars.Token
	AnyByte() grammars.Token
	AnyNumber() grammars.Token
	AToFUpperCaseLetters() grammars.Token
	AToFLowerCaseLetters() grammars.Token
	UpperCaseLetters() grammars.Token
	LowerCaseLetters() grammars.Token
}
