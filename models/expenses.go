package models

type Expenses struct {
        data map[int]int
}

func NewExpenses(d []int) Expenses {
        data := make(map[int]int)

        for i, v := range d {
                data[i] = v
        }

        return Expenses{data: data}
}

func (e *Expenses) Get(idx int) int {
        return e.data[idx]
}
