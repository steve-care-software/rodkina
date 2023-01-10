package grammars

import (
	"testing"

	grammar_applications "github.com/steve-care-software/rodkina/libs/asts/applications"
)

func TestGrammar_coverage_Success(t *testing.T) {
	grammarApp := grammar_applications.NewApplication()
	ins := NewGrammar()
	coverages, err := grammarApp.Coverages(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if coverages.ContainsError() {
		list := coverages.List()
		for _, oneCoverage := range list {
			executions := oneCoverage.Executions()
			token := oneCoverage.Token()
			executionsList := executions.List()
			for idx, oneExecution := range executionsList {
				expectation := oneExecution.Expectation()
				result := oneExecution.Result()
				if expectation.IsValid() && result.IsError() {
					t.Errorf("the token (name: %s) execution (index: %d) was expected to be valid, but contains an error: %s", token.Name(), idx, result.Error())
					continue
				}

				if !expectation.IsValid() && result.IsTree() {
					t.Errorf("the token (name: %s) execution (index: %d) was expected to be invalid, found: %s", token.Name(), idx, result.Tree().Bytes(true))
					continue
				}
			}
		}
	}
}

func TestGrammar_withScript_Success(t *testing.T) {
	grammarApp := grammar_applications.NewApplication()
	ins := NewGrammar()
	script := `
		// this is the root entry point:
		@myValue;

		// those are the channels:
		-myChannel;
		-mySecondChannel [prev];
		-myThirdChannel [:next];
		-myFourthChanel [prev:next];

		// this is a value token:
		myValue: myCompose
			---
			valid: myValidCompose;
			invalid : myInvalidCompose
					& mySecondInvalidCompose
					;
		;

		// this is an external token:
		myExternalToken: cf113f0af255e83f32351a3c32c05fc824e46119f93fb00bfece497421cd4e790b0d682a7bb54d3136c87fdd9222f2ed6a36c904958b0a797b98a22d9d94601c
		---
			valid	: firstValidCompose
					& secondValidCompose
					;
		;
	`

	treeIns, err := grammarApp.Execute(ins, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if treeIns.HasRemaining() {
		t.Errorf("the tree was expected to NOT contain remaining data")
		return
	}

}
