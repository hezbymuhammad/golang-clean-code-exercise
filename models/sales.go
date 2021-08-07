package models

type sales struct {
        data map[int]int
}

func NewSales(d []int) Transactions {
        data := make(map[int]int)

        for i, v := range d {
                data[i] = v
        }

        return &sales{data: data}
}

func (s *sales) Get(idx int) int {
        return s.data[idx]
}
