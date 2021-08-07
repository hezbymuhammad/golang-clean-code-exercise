package models_test

import (
        "testing"

        "github.com/stretchr/testify/suite"

        "github.com/hezbymuhammad/golang-clean-code-exercise/models"
)

type TransactionsTestSuite struct {
        suite.Suite
        transactions models.Transactions
}

func TestTransactions(t *testing.T) {
        suite.Run(t, new(TransactionsTestSuite))
}

func (tr *TransactionsTestSuite) SetupTest() {
        tr.transactions = models.NewTransactions([]int{1,2,3})
}

func (tr *TransactionsTestSuite) TestSuccessGet() {
        tr.Assert().Equal(tr.transactions.Get(0), 1)
}

func (tr *TransactionsTestSuite) TestNotFoundGet() {
        tr.Assert().Equal(tr.transactions.Get(5), 0)
}

func (tr *TransactionsTestSuite) TestSuccessGetTotal() {
        tr.Assert().Equal(tr.transactions.GetTotal(), 6)
}

func (tr *TransactionsTestSuite) TestSuccessGetTotalWithinRange() {
        tr.Assert().Equal(tr.transactions.GetTotalWithinRange(0, 1), 3)
}
