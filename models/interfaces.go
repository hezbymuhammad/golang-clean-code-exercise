package models

type Transactions interface {
        Get(idx int) int
        GetTotalWithinRange(i, j int) int
        GetTotal() int
}
