package main

import (
	// "bufio"
	"cmp"
	"fmt"
	"math"
	"math/rand"
	"os"
	"slices"
	"strings"
	"unsafe"
)

type Person struct {
	Name string
	Age int
	Salary float64
}

func main() {
	res1 := cmp.Compare(1, 2)
	res2 := cmp.Less(1, 2)
	// res3 := cmp.Than(2, 3)
	fmt.Printf("res1: %v, res2: %v\n", res1, res2)

	os.Stdout.WriteString("Hello world!\n")

	fmt.Printf("%#b\n", 12)
	fmt.Printf("%#x\n", 12)

	num := 23
	p := &num
	fmt.Printf("p: %p\n", p)

	fmt.Printf("%08d\n", 20)

	fmt.Printf("%08b\n", 20)

	// var buf [1024]byte
	// n, _ := os.Stdin.Read(buf[:])
	// os.Stdout.Write(buf[:n])

	// var a, b int
	// fmt.Scanln(&a, &b)
	// fmt.Printf("%d + %d = %d\n", a, b, a + b)

	// n := 3
	// s := make([]int, n)
	// for i := range n {
	// 	fmt.Scan(&s[i])
	// }
	// fmt.Println(s)

	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	if line == "exit" {
	// 		break
	// 	}
	// 	fmt.Println("scan", line)
	// }

	no := 55
	switch {
	case no >= 100:
		fmt.Println("优秀！")
	case no >= 60:
		fmt.Println("不错！")
	case no >= 50:
		fallthrough
	case no >= 40:
		fmt.Println("很差")
	default:
		fmt.Println("特别差")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i <= j {
				fmt.Printf("%d * %d = %d ", i, j, i * j)
			}
		}
		fmt.Println()
	}

	sequence := "hello world"
	for index, value := range sequence {
		fmt.Println("index = ", index, " value = ", value)
		fmt.Printf("index = %d, value = %q\n", index, value)
	}

	for value := range 5 {
		fmt.Printf("value = %+v\n", value)
	}

	nums := [...]int{1, 2, 3, 4}
	for index, item := range nums {
		fmt.Printf("index=%d, item=%d\n", index, item)
	}
	fmt.Printf("nums[2] is %d\n", nums[2])
	fmt.Printf("数组的长度是：%d\n", len(nums))

	arr := [5]int{1, 2, 3, 4, 5}
	// slice := arr[:]
	// slice[0] = 0
	// fmt.Printf("array: %v\n", arr)
	// fmt.Printf("slice: %v\n", slice)

	slice1 := slices.Clone(arr[:])
	slice1[0] = 0
	fmt.Printf("array: %v\n", arr)
	fmt.Printf("slice1: %v\n", slice1)

	nums1 := make([]int, 0)
	nums1 = append(nums1, 1, 2, 3, 4)
	fmt.Println(len(nums1), cap(nums1))

	fmt.Println("=====")

	str := `原生字符串由反引号表示，不支持转义，支持多行书写
	\t原生字符串里面所有的字符都会原封不动的输出，包括换行和缩进。`
	fmt.Println("str content: ", str)

	fmt.Println(str[12])
	fmt.Println(str[10:15])

	str1 := "this is a string"
	str2 := "这是一个字符串"
	fmt.Println("str1: ", len(str1), " str2: ", len(str2))

	var dst, src string
	src = "this is a string"
	desBytes := make([]byte, len(src))
	copy(desBytes, src)
	dst = string(desBytes)
	fmt.Println("src: ", src, " dst: ", dst)

	dst1 := strings.Clone(src)
	fmt.Println("dst: ", dst1)

	str3 := "hello, "
	str4 := "world!"
	str5 := str3 + str4
	fmt.Println(str5)

	stringBuilder := strings.Builder{}
	stringBuilder.WriteString("this is a string")
	stringBuilder.WriteString(", that is a int")
	fmt.Println(stringBuilder.String())

	str6 := "Hello, 世界"
	for _, r := range str6 {
		fmt.Printf("%d, %x, %s\n", r, r, string(r))
	}

	runes := []rune(str6)
	for i := 0; i < len(runes); i++ {
		fmt.Println(string(runes[i]))
	}

	mp := map[int]string{
		0: "a",
		1: "b",
	}
	// mp = make(map[int]string, 2)
	mp[0] = "c"
	for k, v := range mp {
		fmt.Printf("k = %d, v = %v\n", k, v)
	}

	delete(mp, 1)
	for k, v := range mp {
		fmt.Printf("k = %d, v = %v\n", k, v)
	}

	fmt.Println("=====")

	fmt.Println(math.NaN())

	clear(mp)
	fmt.Println("length of mp: ", len(mp))

	fmt.Println("=====")

	set := make(map[int]struct{}, 3)
	for i := 0; i < 3; i++ {
		set[rand.Intn(100)] = struct{}{}
	}
	fmt.Println(set)

	fmt.Println("=====")
	num1 := 2
	p1 := &num1
	rawNum := *p1
	fmt.Println("rawNum: ", rawNum)

	var numPtr *int
	numPtr = new(int)
	fmt.Println(numPtr)

	num2 := 10
	numPtr = &num2
	fmt.Println("numPtr: ", *numPtr)

	fmt.Println("=====")

	res3 := func(a, b int) int {
		return a + b
	} (1, 2)
	fmt.Println("res3: ", res3)

	people := []Person{
		{Name: "Alice", Age: 25, Salary: 5000.0},
	    {Name: "Charlie", Age: 28, Salary: 5500.0},
	    {Name: "Bob", Age: 30, Salary: 6000.0},
	}
	
	slices.SortFunc(people, func(p1, p2 Person) int {
		if p1.Name > p2.Name {
			return 1
		} else if p1.Name == p2.Name {
			return 0
		} else {
			return -1
		}
	})

	fmt.Println(people)

	do()

	type Empty struct {}
	fmt.Println(unsafe.Sizeof(Empty{}))

	p2 := &people[1]
	fmt.Println(p2.Name, p2.Age, p2.Salary)
}

func do() {
	defer func() {
		fmt.Println("defer 1")
	}()
	fmt.Println("defer 2")
}
