package main

import "fmt"

func main() {
	fmt.Println("hello go")
	var i int32 = 32
	fmt.Println(i)
	j := "sada"
	c := []byte(j)
	c[0] = 'h'
	d := string(c)
	fmt.Println(d)
	fmt.Println(j)

	str := "hello"
	str1 := "c" + str[1:]
	fmt.Println(str1)

	arra := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arra)
	slice1 := arra[:4]
	fmt.Println(slice1)
	slice2 := slice1[0:2]
	fmt.Println(slice2)
	slice1[0] = 5
	slice1 = append(slice1, 60)
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	var map1 map[string]int
	map1 = make(map[string]int)
	map1["Go"] = 1
	fmt.Println(map1)

	map2 := make(map[string]int)
	map2["python"] = 23
	fmt.Println(map2)

	map3 := map[string]int{"Go": 1, "python": 2, "C++": 3}
	fmt.Println(map3)

	mapreturn1, ok := map3["Go"]
	fmt.Println(mapreturn1, ok)

	for v, k := range map3 {
		fmt.Println(v, k)
	}

	r1 := Rectangle{3, 4}
	fmt.Println(r1.area(), r1.width)

}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	r.width = 6
	return r.width * r.height
}
