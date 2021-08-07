package models

type transactions struct {
        data map[int]int
}

// TODO: implement this
// i.e. go test is green
func NewTransactions(d []int) Transactions {
        data := make(map[int]int)

        for i, v := range d {
                data[i] = v
        }

        return &transactions{data: data}
}

// TODO: implement this
// i.e. go test is green
func (t *transactions) Get(idx int) int {
        return t.data[idx]
}

func (t *transactions) GetTotal() int {
        total := 0

        for _, v := range t.data {
                total += v
        }

        return total
}

func (t *transactions) GetTotalWithinRange(i, j int) int {
        total := 0

        for ; i <= j && i < len(t.data); i++ {
                total += t.data[i]
        }

        return total
}
