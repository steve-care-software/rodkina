package tokens

import (
	"fmt"

	"github.com/steve-care-software/rodkina/instances/grammars/components"
	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
)

type token struct {
	component components.Component
}

func createToken(
	component components.Component,
) Token {
	out := token{
		component: component,
	}

	return &out
}

// VariableName returns the variable name token
func (app *token) VariableName() grammars.Token {
	return app.component.Token().FromBlock(
		"variableName",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.LowerCaseLetters(), app.component.Cardinality().Once()),
				app.component.Element().FromToken(app.AnyLetter(), app.component.Cardinality().Cardinality(0, nil)),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			"m":          true,
			"myVariable": true,
			"MyVariable": false,
			"0Variable":  false,
		}),
	)
}

// Sha512Hex returns the sha512 token
func (app *token) Sha512Hex() grammars.Token {
	amount := uint(128)
	return app.component.Token().FromBlock(
		"sha512Hex",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.AnyHexCharacter(), app.component.Cardinality().Cardinality(amount, &amount)),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			"cf113f0af255e83f32351a3c32c05fc824e46119f93fb00bfece497421cd4e790b0d682a7bb54d3136c87fdd9222f2ed6a36c904958b0a797b98a22d9d94601c": true,
			"cf113f0af255e83f32351a3c32c05fc824e46119f93fb00bfece497421cd4e790b0d682a7bb54d3136c87fdd9222f2ed6a36c904958b0a797b98a22d9d9460":   false,
			"gf113f0af255e83f32351a3c32c05fc824e46119f93fb00bfece497421cd4e790b0d682a7bb54d3136c87fdd9222f2ed6a36c904958b0a797b98a22d9d94601c": false,
		}),
	)
}

// AnyHexCharacter returns the [a-fA-Z0-9] token
func (app *token) AnyHexCharacter() grammars.Token {
	return app.component.Token().FromBlock(
		"anyHexChar",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.AToFLowerCaseLetters(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.AToFUpperCaseLetters(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.AnyNumber(), app.component.Cardinality().Once()),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			"a": true,
			"b": true,
			"c": true,
			"d": true,
			"e": true,
			"f": true,
			"g": false,
			"h": false,
			"i": false,
			"j": false,
			"k": false,
			"l": false,
			"m": false,
			"n": false,
			"o": false,
			"p": false,
			"q": false,
			"r": false,
			"s": false,
			"t": false,
			"u": false,
			"v": false,
			"w": false,
			"x": false,
			"y": false,
			"z": false,
			"A": true,
			"B": true,
			"C": true,
			"D": true,
			"E": true,
			"F": true,
			"G": false,
			"H": false,
			"I": false,
			"J": false,
			"K": false,
			"L": false,
			"M": false,
			"N": false,
			"O": false,
			"P": false,
			"Q": false,
			"R": false,
			"S": false,
			"T": false,
			"U": false,
			"V": false,
			"W": false,
			"X": false,
			"Y": false,
			"Z": false,
			"0": true,
			"1": true,
			"2": true,
			"3": true,
			"4": true,
			"5": true,
			"6": true,
			"7": true,
			"8": true,
			"9": true,
		}),
	)
}

// AnyLetter returns the [a-zA-Z] token
func (app *token) AnyLetter() grammars.Token {
	return app.component.Token().FromBlock(
		"anyLetter",
		app.component.Block().FromLines([]grammars.Line{
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.UpperCaseLetters(), app.component.Cardinality().Once()),
			}),
			app.component.Line().FromElements([]grammars.Element{
				app.component.Element().FromToken(app.LowerCaseLetters(), app.component.Cardinality().Once()),
			}),
		}),
		app.component.Suite().Suites(map[string]bool{
			"a": true,
			"b": true,
			"c": true,
			"d": true,
			"e": true,
			"f": true,
			"g": true,
			"h": true,
			"i": true,
			"j": true,
			"k": true,
			"l": true,
			"m": true,
			"n": true,
			"o": true,
			"p": true,
			"q": true,
			"r": true,
			"s": true,
			"t": true,
			"u": true,
			"v": true,
			"w": true,
			"x": true,
			"y": true,
			"z": true,
			"A": true,
			"B": true,
			"C": true,
			"D": true,
			"E": true,
			"F": true,
			"G": true,
			"H": true,
			"I": true,
			"J": true,
			"K": true,
			"L": true,
			"M": true,
			"N": true,
			"O": true,
			"P": true,
			"Q": true,
			"R": true,
			"S": true,
			"T": true,
			"U": true,
			"V": true,
			"W": true,
			"X": true,
			"Y": true,
			"Z": true,
			"0": false,
		}),
	)
}

// AnyByte returns the any byte token
func (app *token) AnyByte() grammars.Token {
	validData := map[string]bool{}
	tokenName := "anyByte"
	elementsList := []grammars.Element{}
	for i := 0; i < byteLength; i++ {
		element := app.component.Element().FromValue(byte(i))
		elementsList = append(elementsList, element)
		validData[fmt.Sprintf("%d", i)] = true
	}

	return app.component.Token().AnyElement(tokenName, elementsList, app.component.Suite().Suites(validData))
}

// AnyNumber returns the [0-9] token
func (app *token) AnyNumber() grammars.Token {
	characters := "0123456789"
	return app.component.Token().AnyCharacter("anyNumber", characters)
}

// AToFUpperCaseLetters returns the [A-F] token
func (app *token) AToFUpperCaseLetters() grammars.Token {
	characters := "ABCDEF"
	return app.component.Token().AnyCharacter("aToFUpperCaseLetters", characters)
}

// AToFLowerCaseLetters returns the [a-f] token
func (app *token) AToFLowerCaseLetters() grammars.Token {
	characters := "abcdef"
	return app.component.Token().AnyCharacter("aToFLowerCaseLetter", characters)
}

// UpperCaseLetters returns the uppercase letter token
func (app *token) UpperCaseLetters() grammars.Token {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return app.component.Token().AnyCharacter("uppercaseLetter", characters)
}

// LowerCaseLetters returns the lowercase letters token
func (app *token) LowerCaseLetters() grammars.Token {
	characters := "abcdefghijklmnopqrstuvwxyz"
	return app.component.Token().AnyCharacter("lowerCaseLetter", characters)
}
