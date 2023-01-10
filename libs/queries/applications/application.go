package queries

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/rodkina/libs/asts/domain/grammars"
	"github.com/steve-care-software/rodkina/libs/asts/domain/trees"
	"github.com/steve-care-software/rodkina/libs/queries/domain/queries"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Matches returns true if the query matches the grammar, false otherwise
func (app *application) Matches(grammar grammars.Grammar, query queries.Query) (bool, error) {
	return true, nil
}

// Execute executes a query on a data tree
func (app *application) Execute(query queries.Query, treeIns trees.Tree) (interface{}, bool, []byte, error) {
	ins, isValid, err := app.queryFetch(query, treeIns, nil)
	if err != nil {
		return nil, false, nil, err
	}

	remaining := []byte{}
	if treeIns.HasRemaining() {
		remaining = treeIns.Remaining()
	}

	return ins, isValid, remaining, nil
}

func (app *application) queryFetch(query queries.Query, tree trees.Tree, previous map[string]queries.Query) (interface{}, bool, error) {
	contentsList, err := app.queryContents(query, tree)
	if err != nil {
		return nil, false, err
	}

	if previous == nil {
		previous = map[string]queries.Query{}
	}

	previous[query.Token().Name()] = query
	return app.queryInstance(query, contentsList, previous)
}

func (app *application) queryContents(query queries.Query, tree trees.Tree) ([]trees.Content, error) {
	queryToken := query.Token()
	treeName := tree.Grammar().Name()
	if queryToken.Name() != treeName {
		str := fmt.Sprintf("the contents cannot be retrieved because the tree (token: %s) do not match the query (token: %s)", treeName, queryToken.Name())
		return nil, errors.New(str)
	}

	block := tree.Block()
	if !block.HasSuccessful() {
		str := fmt.Sprintf("the contents cannot be retrieved because the tree (token: %s) contains no successful line", treeName)
		return nil, errors.New(str)
	}

	cpt := uint(0)
	queryElement := queryToken.Element()
	elements := tree.Block().Successful().Elements().List()
	for _, oneElement := range elements {
		contents := oneElement.Contents()
		if !oneElement.HasGrammar() {
			if queryElement.Name() != queryToken.ReverseName() {
				continue
			}

			if queryElement.Index() == cpt {
				return app.tokenRefine(queryToken, contents)
			}

			cpt++
		}

		if oneElement.Grammar().Name() != queryElement.Name() {
			continue
		}

		if queryElement.Index() == cpt {
			return app.tokenRefine(queryToken, contents)
		}

		cpt++
	}

	return []trees.Content{}, nil
}

func (app *application) queryInstance(query queries.Query, contentList []trees.Content, previous map[string]queries.Query) (interface{}, bool, error) {
	fn := query.Fn()
	inside := query.Inside()
	insList, err := app.insideInstances(inside, contentList, previous)
	if err != nil {
		return nil, false, err
	}

	if len(insList) <= 0 {
		return nil, false, nil
	}

	if fn.IsSingle() {
		var param interface{}
		if len(insList) > 0 {
			param = insList[0]
		}

		singleFn := fn.Single()
		return singleFn(param)
	}

	multiFn := fn.Multi()
	return multiFn(insList)
}

func (app *application) tokenRefine(token queries.Token, contents trees.Contents) ([]trees.Content, error) {
	if !token.HasContent() {
		return contents.List(), nil
	}

	pIndex := token.Content()
	list := contents.List()
	listLength := uint(len(list))
	if listLength <= *pIndex {
		str := fmt.Sprintf("the contents cannot be refined because the token query requires a content (index: %d) but the list (length: %d) is too small", *pIndex, listLength)
		return nil, errors.New(str)
	}

	return []trees.Content{
		list[*pIndex],
	}, nil
}

func (app *application) insideInstances(inside queries.Inside, contentList []trees.Content, previous map[string]queries.Query) ([]interface{}, error) {
	if inside.IsFn() {
		fn := inside.Fn()
		if fn.IsSingle() {
			if len(contentList) < 0 {
				str := fmt.Sprintf("%d content instances were expected in the Content list, %d were provided", 1, len(contentList))
				return nil, errors.New(str)
			}

			singleFn := fn.Single()
			return singleFn(contentList[0])
		}

		multiFn := fn.Multi()
		return multiFn(contentList)
	}

	output := []interface{}{}
	fetchers := inside.Fetchers()
	fetchersList := fetchers.List()
	for _, oneContent := range contentList {
		row := []interface{}{}
		for _, oneFetcher := range fetchersList {
			if oneContent.IsValue() {
				continue
			}

			tree := oneContent.Tree()
			if oneFetcher.IsQuery() {
				query := oneFetcher.Query()
				ins, isValid, err := app.queryFetch(query, tree, previous)
				if err != nil {
					return nil, err
				}

				if !isValid {
					continue
				}

				row = append(row, ins)
				continue
			}

			recursive := oneFetcher.Recursive()
			if query, ok := previous[recursive]; ok {
				ins, isValid, err := app.queryFetch(query, tree, previous)
				if err != nil {
					return nil, err
				}

				if !isValid {
					continue
				}

				row = append(row, ins)
				continue
			}

			str := fmt.Sprintf("the recursive Query's Token (name: %s) could not be found in the previous iterations", recursive)
			return nil, errors.New(str)

		}

		output = append(output, row...)
	}

	return output, nil
}
