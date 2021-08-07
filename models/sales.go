package models

type Sales struct {
        data map[int]int
}

func NewSales(d []int) Sales {
        data := make(map[int]int)

        for i, v := range d {
                data[i] = v
        }

        return Sales{data: data}
}

func (s *Sales) Get(idx int) int {
        return s.data[idx]
}
