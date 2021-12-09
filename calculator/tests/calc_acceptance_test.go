package calc_test

import (
	"fmt"

	"github.com/cucumber/godog"
	calc "github.com/devdinu/go-play/calculator"
)

type CalcSuite struct {
	Suite *godog.ScenarioContext
	calc  *calc.Calculator
}

func (cs CalcSuite) calculatorIsCleared() error {
	cs.calc.Clear()
	return nil
}

func (cs CalcSuite) iAdd(arg1 int) error {
	cs.calc.Add(arg1)
	return nil
}

func (cs CalcSuite) iMultiplyBy(arg1 int) error {
	cs.calc.MultiplyBy(arg1)
	return nil
}

func (cs CalcSuite) iPress(arg1 int) error {
	cs.calc.Press(arg1)
	return nil
}

func (cs CalcSuite) iSubtract(arg1 int) error {
	cs.calc.Sub(arg1)
	return nil
}

func (cs CalcSuite) theResultShouldBe(expectedResult int) error {
	result := cs.calc.Result()
	if result == expectedResult {
		return nil
	}
	return fmt.Errorf("%d doesn't match expectation: %d", result, expectedResult)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	suite := &CalcSuite{
		calc:  &calc.Calculator{},
		Suite: ctx,
	}

	ctx.Step(`^calculator is cleared$`, suite.calculatorIsCleared)
	ctx.Step(`^i add (\d+)$`, suite.iAdd)
	ctx.Step(`^i multiply by (\d+)$`, suite.iMultiplyBy)
	ctx.Step(`^i press (\d+)$`, suite.iPress)
	ctx.Step(`^i subtract (\d+)$`, suite.iSubtract)
	ctx.Step(`^the result should be (\d+)$`, suite.theResultShouldBe)
}
