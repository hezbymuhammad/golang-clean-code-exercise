package models_test

import (
        "testing"

        "github.com/stretchr/testify/suite"

        "github.com/hezbymuhammad/golang-clean-code-exercise/models"
)

type ExpenseTestSuite struct {
        suite.Suite
        expenses models.Transactions
}

func TestExpenses(t *testing.T) {
        suite.Run(t, new(ExpenseTestSuite))
}

func (e *ExpenseTestSuite) SetupTest() {
        e.expenses = models.NewExpenses([]int{1,2,3})
}

func (e *ExpenseTestSuite) TestSuccessGet() {
        e.Assert().Equal(e.expenses.Get(0), 1)
}

func (e *ExpenseTestSuite) TestNotFoundGet() {
        e.Assert().Equal(e.expenses.Get(5), 0)
}
