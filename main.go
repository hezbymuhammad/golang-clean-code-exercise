package main

import (
        "fmt"

        "github.com/hezbymuhammad/golang-clean-code-exercise/models"
)

func main() {
        salesData := []int{1000, 22000, 33000, 44000, 55000, 66000, 77000}
        expensesData := []int{15000, 25000, 35000, 45000, 55000, 65000, 75000}

        s := models.NewSales(salesData)
        e := models.NewExpenses(expensesData)

        salesTotal := 0
        expensesTotal := 0
        salesFirstTwoDays := 0
        expensesFirstTwoDays := 0

        // TODO: create separate method for calculating sales and expense
        // i.e. salesTotal := s.GetTotal()
        // i.e. expensesTotal := e.GetTotal()
        for i := 0; i <= 6; i++ {
                salesTotal += s.Get(i)
                expensesTotal += e.Get(i)
        }

        // TODO: create separate method for calculating sales and expense
        // i.e. salesFirstTwoDays := s.GetTotalWithinRange(0, 2)
        // i.e. expensesFirstTwoDays := e.GetTotalWithinRange(0, 2)
        for i := 0; i <= 1; i++ {
                salesFirstTwoDays += s.Get(i)
                expensesFirstTwoDays += e.Get(i)
        }

        fmt.Println("=======##=======")
        fmt.Printf("Sales in the first 2 days is: %v\n", salesFirstTwoDays)
        fmt.Printf("Expenses in the first 2 days is: %v\n", expensesFirstTwoDays)
        fmt.Printf("Total profit in the first 2 days is: %v\n", salesFirstTwoDays - expensesFirstTwoDays)
        fmt.Println("")
        fmt.Printf("Sales total is: %v\n", salesTotal)
        fmt.Printf("Expenses total is: %v\n", expensesTotal)
        fmt.Printf("Total profit is: %v\n", salesTotal - expensesTotal)
        fmt.Println("=======##=======")
}
