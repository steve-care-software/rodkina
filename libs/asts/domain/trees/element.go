package trees

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
)

type element struct {
	contents Contents
	grammar  grammars.Container
}

func createElement(
	contents Contents,
) Element {
	return createElementInternally(contents, nil)
}

func createElementWithGrammar(
	contents Contents,
	grammar grammars.Container,
) Element {
	return createElementInternally(contents, grammar)
}

func createElementInternally(
	contents Contents,
	grammar grammars.Container,
) Element {
	out := element{
		grammar:  grammar,
		contents: contents,
	}

	return &out
}

// Fetch fetches a tree or value by name
func (obj *element) Fetch(name string, elementIndex uint) (Tree, Element, error) {
	if obj.HasGrammar() {
		if obj.grammar.Name() == name {
			return nil, obj, nil
		}
	}

	list := obj.contents.List()
	for _, oneContent := range list {
		if !oneContent.IsTree() {
			continue
		}

		tree, element, err := oneContent.Tree().Fetch(name, elementIndex)
		if err != nil {
			continue
		}

		if tree != nil {
			return tree, nil, nil
		}

		if element != nil {
			return nil, element, nil
		}
	}

	str := fmt.Sprintf("there is no Tree or Element associated to the given name: %s", name)
	return nil, nil, errors.New(str)
}

// Bytes returns the element's bytes
func (obj *element) Bytes(includeChannels bool) []byte {
	output := []byte{}
	list := obj.contents.List()
	for _, oneContent := range list {
		if oneContent.IsValue() {
			value := oneContent.Value()
			if includeChannels && value.HasPrefix() {
				output = append(output, value.Prefix().Bytes(includeChannels)...)
			}

			output = append(output, value.Content())
			continue
		}

		output = append(output, oneContent.Tree().Bytes(includeChannels)...)
	}

	return output
}

// IsSuccessful returns true if successful, false otherwise
func (obj *element) IsSuccessful() bool {
	if !obj.HasGrammar() {
		return true
	}

	amount := obj.Amount()
	if obj.grammar.IsElement() {
		cardinality := obj.grammar.Element().Cardinality()
		min := cardinality.Min()
		if amount < min {
			return false
		}

		if cardinality.HasMax() {
			pMax := cardinality.Max()
			if amount > *pMax {
				return false
			}
		}

		return true
	}

	requestedAmount := uint(0)
	composeList := obj.grammar.Compose().List()
	for _, oneCompose := range composeList {
		requestedAmount += oneCompose.Occurences()
	}

	if amount < requestedAmount {
		return false
	}

	return true
}

// Contents returns the contents
func (obj *element) Contents() Contents {
	return obj.contents
}

// Grammar returns the grammar
func (obj *element) Grammar() grammars.Container {
	return obj.grammar
}

// HasGrammar returns true if there is a grammar, false otherwise
func (obj *element) HasGrammar() bool {
	return obj.grammar != nil
}

// Amount returns the amount
func (obj *element) Amount() uint {
	return uint(len(obj.contents.List()))
}
