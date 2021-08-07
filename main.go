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

        salesTotal := s.GetTotal()
        expensesTotal := e.GetTotal()
        salesFirstTwoDays := s.GetTotalWithinRange(0, 2)
        expensesFirstTwoDays := e.GetTotalWithinRange(0, 2)

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
