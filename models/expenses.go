package models

type expenses struct {
        data map[int]int
}

// TODO: instantiate from transaction model
func NewExpenses(d []int) Transactions {
        data := make(map[int]int)

        for i, v := range d {
                data[i] = v
        }

        // i.e. return NewTransactions(d)
        return &expenses{data: data}
}

func (e *expenses) Get(idx int) int {
        return e.data[idx]
}
