package models

type expenses struct {
        data map[int]int
}

func NewExpenses(d []int) Transactions {
        data := make(map[int]int)

        for i, v := range d {
                data[i] = v
        }

        return &expenses{data: data}
}

func (e *expenses) Get(idx int) int {
        return e.data[idx]
}
