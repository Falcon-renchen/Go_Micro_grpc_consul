package product

import "strconv"

type Product struct {
	ProductID int `json:"product_id"`
	ProductName string `json:"product_name"`
}

func NewProduct(id int, pname string) *Product {
	return &Product{
		ProductID:   id,
		ProductName: pname,
	}
}
func NewProdList(n int) []*Product {
	ret := make([]*Product, 0)

	for i:=0; i<n; i++ {
		ret = append(ret, NewProduct(100+i,"pname"+strconv.Itoa(100+i)))
	}
	return ret
}