package main

import "fmt"

type Products interface {
	GetProductList() []string
}

type AdvancedProducts interface {
	GetListAllProducts() []string
	UpdatePrice(name string) int
}

type Product struct{}

func (p Product) GetListAllProducts() []string {
	return []string{"banana", "apple", "cucumber"}
}

func (p Product) UpdatePrice(name string) int {
	prices := map[string]int{
		"banana":   50,
		"apple":    30,
		"cucumber": 20,
	}
	return prices[name]
}

type ProductAdapter struct {
	AdvancedProducts
}

func (pa ProductAdapter) GetProductList() []string {
	return pa.GetListAllProducts()
}

func main() {
	advancedProduct := Product{}

	adapter := ProductAdapter{AdvancedProducts: advancedProduct}

	fmt.Println(adapter.GetProductList())
	fmt.Println(adapter.UpdatePrice("banana"))
}

// Паттерн "Адаптер" бывает полезен, когда у нас есть старые и новые методы, которые между собой не совместимы, но нужно как-то использовать
// Плюсы: Делает работу более гибким и расширяемым, четкое разделение ответственности (handler -> service -> repository)
// Минусы: Увеличение количества классов, Увеличение сложности(в случае добавления много адаптеров, работа становится сложной)
// Реальные примеру: драйвера бд psql и другие, конвертация формата данных
