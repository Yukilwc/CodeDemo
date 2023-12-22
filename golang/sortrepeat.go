package main

// 聚合数组中的同类
import "fmt"

// 优先聚合国家，然后是城市，然后是街区，不做任何另外排序

type Element struct {
	Country string
	City    string
	Street  string
}
type CompareFunc = func(a, b Element) bool

func main() {
	var list = []Element{
		{
			Country: "CHINA",
			City:    "QINGDAO",
			Street:  "Street-a",
		},
		{
			Country: "JAPAN",
			City:    "City1",
			Street:  "Street-b",
		},
		{
			Country: "CHINA",
			City:    "SHANGHAI",
			Street:  "Street-a",
		},
		{
			Country: "JAPAN",
			City:    "City1",
			Street:  "Street-c",
		},
		{
			Country: "CHINA",
			City:    "QINGDAO",
			Street:  "Street-d",
		},
		{
			Country: "CHINA",
			City:    "QINGDAO",
			Street:  "Street-a",
		},
	}
	newList := aggregate(&list, func(a, b Element) bool {
		return a.Street == b.Street
	})
	newList = aggregate(newList, func(a, b Element) bool {
		return a.City == b.City
	})
	newList = aggregate(newList, func(a, b Element) bool {
		return a.Country == b.Country
	})

	for _, v := range *newList {
		fmt.Printf("\n元素%+v\n", v)
	}
}

func find(twoDimensionalSlice *[][]Element, v *Element, compare CompareFunc) (int, bool) {
	index := 0
	ok := false
	for i, l := range *twoDimensionalSlice {
		if compare(l[0], *v) {
			index = i
			ok = true
			break
		}
	}
	return index, ok
}
func aggregate(list *[]Element, compare CompareFunc) *[]Element {
	twoDimensionalSlice := make([][]Element, 0)
	for _, v := range *list {
		findIndex, ok := find(&twoDimensionalSlice, &v, compare)
		if !ok {
			newSpan := make([]Element, 0)
			newSpan = append(newSpan, v)
			twoDimensionalSlice = append(twoDimensionalSlice, newSpan)
		} else {
			twoDimensionalSlice[findIndex] = append(twoDimensionalSlice[findIndex], v)
		}
	}
	res := make([]Element, 0)
	for _, v := range twoDimensionalSlice {
		fmt.Printf("\n 二维数组: %+v\n", v)
		res = append(res, v...)
	}
	return &res
}
