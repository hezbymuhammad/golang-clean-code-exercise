package models_test

import (
        "testing"

        "github.com/stretchr/testify/suite"

        "github.com/hezbymuhammad/golang-clean-code-exercise/models"
)

type SalesTestSuite struct {
        suite.Suite
        sales models.Transactions
}

func TestSales(t *testing.T) {
        suite.Run(t, new(SalesTestSuite))
}

func (s *SalesTestSuite) SetupTest() {
        s.sales = models.NewSales([]int{1,2,3})
}

func (s *SalesTestSuite) TestSuccessGet() {
        s.Assert().Equal(s.sales.Get(0), 1)
}

func (s *SalesTestSuite) TestNotFoundGet() {
        s.Assert().Equal(s.sales.Get(5), 0)
}
