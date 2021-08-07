package models

type transactions struct {
        data map[int]int
}

func NewTransactions(d []int) Transactions {
        data := make(map[int]int)

        return &transactions{data: data}
}

func (t *transactions) Get(idx int) int {
        return 0
}
