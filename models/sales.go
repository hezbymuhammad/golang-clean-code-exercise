package models

type sales struct {
        data map[int]int
}

// TODO: instantiate from transaction model
func NewSales(d []int) Transactions {
        data := make(map[int]int)

        for i, v := range d {
                data[i] = v
        }

        // i.e. return NewTransactions(d)
        return &sales{data: data}
}

func (s *sales) Get(idx int) int {
        return s.data[idx]
}
