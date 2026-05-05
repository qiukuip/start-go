package main

import (
	// "bufio"
	"cmp"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"reflect"
	"slices"
	"strings"
	"time"
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

	fmt.Println("=====")

	person1 := Person{
		"Dove",
		30,
		2000.00,
	}
	fmt.Println(person1)

	person2 := NewPerson("张三", 10, 1500.00)
	fmt.Println(*person2)

	person3 := NewPerson1(
		WithName("李四"),
		WithAge(23),
		WithSalary(1200.00),
	)
	fmt.Println(*person3)

	fmt.Println("=====")

	var intSlice IntSlice
	intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(intSlice.Get(1))
	intSlice.Set(1, 10)
	fmt.Println(intSlice)
	fmt.Println(intSlice.Len())

	company := Company{
		CraneA{},
	}
	company.Build()
	fmt.Println()
	company.crane = CraneB{}
	company.Build()

	fmt.Println("=====")

	file, err := os.OpenFile("../files/test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else if err != nil {
		fmt.Println("打开文件出错")
	} else {
		fmt.Println("打开文件成功")
	}
	fmt.Println(file.Name())
	
	defer file.Close()

	fmt.Println("=====")

	fileInfo, err := os.Stat("../files/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", fileInfo)
	
	fmt.Println("=====")

	file1, err := os.OpenFile("../files/test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("访问文件异常")
	} else {
		fmt.Println("打开文件成功", file1.Name())
		bytes, err := ReadFile(file)
		if err != nil {
			fmt.Println("文件读取异常")
		} else {
			fmt.Println(string(bytes))
		}
		file1.Close()
	}

	fmt.Println("=====")

	bytes, err := os.ReadFile("../files/test.txt")
	if err != nil {
		fmt.Println("打开文件失败")
	} else {
		fmt.Println(string(bytes))
	}

	fmt.Println("=====")

	file2, err := os.OpenFile("../files/test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件出错")
	}
	bytes2, err := io.ReadAll(file2)
	if err != nil {
		fmt.Println("打开文件出错")
	} else {
		fmt.Println(string(bytes2))
	}
	file2.Close()

	fmt.Println("=====")

	file3, err := os.OpenFile("../files/test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件出错")
	} else {
		offset, err := file3.WriteString("你好，世界！")
		if err != nil {
			fmt.Println("写入文件出错")
		} else {
			fmt.Println("offset: ", offset)
		}
		file3.Close()
	}

	bytes3, err := os.ReadFile("../files/test.txt")
	if err != nil {
		fmt.Println("打开文件出错")
	} else {
		fmt.Println(string(bytes3))
	}

	fmt.Println("=====")

	dir, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, entry := range dir {
			fmt.Println(entry.Name())
		}
	}

	fmt.Println("=====")

	n := 5
	for f := range Fibonacci(n) {
		fmt.Println(f)
	}

	// panic("执行完了喵")
	// fatal("出错了喵")
	
	err1 := errors.New("这是一个错误")
	err2 := fmt.Errorf("这是格式化参数的错误%d\n", 1)
	fmt.Println(err1)
	fmt.Println(err2)

	wrapErr := fmt.Errorf("错误，%w\n", err1)
	fmt.Println(wrapErr.Error())

	fmt.Println("====")

	initDatabase("1", 8000)

	fmt.Println("=====")

	str7 := "hello world!"
	reflectType := reflect.TypeOf(str7)
	fmt.Println(reflectType)

	var eface any
	eface = map[string]int{}
	rType := reflect.TypeOf(eface)
	fmt.Println(rType.Key().Kind())
	fmt.Println(rType.Elem().Kind())

	fmt.Println("====")

	fmt.Println(reflect.TypeOf("hello world!").Comparable())
	fmt.Println(reflect.TypeOf(1024).Comparable())
	fmt.Println(reflect.TypeOf([]int{}).Comparable())
	fmt.Println(reflect.TypeOf(struct{}{}).Comparable())

	rInface := reflect.TypeOf(new(MyInterface)).Elem()
	fmt.Println(reflect.TypeOf(new(MyStruct)).Elem().Implements(rInface))
	fmt.Println(reflect.TypeOf(new(OtherStruct)).Elem().Implements(rInface))

	fmt.Println(reflect.TypeOf(new(MyStruct)).Elem().ConvertibleTo(rInface))
	fmt.Println(reflect.TypeOf(new(OtherStruct)).Elem().ConvertibleTo(rInface))

	num3 := 12345
	rValue := reflect.ValueOf(num3)
	fmt.Println(rValue)
	fmt.Println(rValue.Type())

	num4 := new(int)
	*num4 = 12345
	rValue1 := reflect.ValueOf(num4).Elem()
	fmt.Println(rValue1)
	fmt.Println(rValue1.Interface())

	fmt.Println("=====")

	fmt.Println("start")
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
		time.Sleep(time.Millisecond)
	}
	time.Sleep(time.Millisecond)
	fmt.Println("end")

	fmt.Println("=====")

	intCh := make(chan int)
	defer close(intCh)
	go func() {
		intCh <- 12345
	}()
	n1 := <-intCh

	fmt.Println(n1)

	fmt.Println("=====")

	ch := make(chan int, 5)
	chW := make(chan struct{})
	chR := make(chan struct{})
	
	defer func() {
		close(ch)
		close(chW)
		close(chR)
	}()

	// 往管道中写入数据
	go func() {
		for i := range 10 {
			ch <- i
			fmt.Println("write to ch <- ", i)
			fmt.Println("len ch: ", len(ch), ", cap ch: ", cap(ch))
		}
		chW <- struct{}{}
	}()

	// 从管道中读取数据
	go func() {
		for range 10 {
			time.Sleep(time.Millisecond)
			fmt.Println("read ", <-ch)
		}
		chR <- struct{}{}
	}()

	fmt.Println("write finished", <-chW)
	fmt.Println("read finished", <-chR)

	fmt.Println("=====")

	ch1 := make(chan struct{})
	defer close(ch1)
	go func() {
		fmt.Println(2)
		ch1 <- struct{}{}
		fmt.Println("write finished")
	}()
	<-ch1
	fmt.Println("read finished")
}

func do() {
	defer func() {
		fmt.Println("defer 1")
	}()
	fmt.Println("defer 2")
}

func NewPerson(name string, age int, salary float64) *Person {
	return &Person{Name: name, Age: age, Salary: salary}
}

type PersonOptions func(p *Person)

func WithName(name string) PersonOptions {
	return func(p *Person) {
		p.Name = name
	}
}

func WithAge(age int) PersonOptions {
	return func(p *Person) {
		p.Age = age
	}
}

func WithSalary(salary float64) PersonOptions {
	return func(p *Person) {
		p.Salary = salary
	}
}

func NewPerson1(options ...PersonOptions) *Person {
	p := &Person{}
	for _, option := range options {
		option(p)
	}
	return p
}

type IntSlice []int

func (i IntSlice) Get(index int) int {
	return i[index]
}

func (i IntSlice) Set(index, value int) {
	i[index] = value
}

func (i IntSlice) Len() int {
	return len(i)
}

type MyInt int

func (myInt MyInt) Set(value int) {
	myInt = MyInt(value)
}

type Crane interface {
	JackUp() string
	Hoist() string
}

type CraneA struct {
	work int
}

func (c CraneA) Work() {
	fmt.Println("使用技术A")
}

func (c CraneA) JackUp() string {
	c.Work()
	return "CraneA jack up!"
}

func (c CraneA) Hoist() string {
	c.Work()
	return "CraneA hoist"
}

type CraneB struct {
	boot string
}

func (b CraneB) Work() {
	fmt.Println("使用技术B")
}

func (b CraneB) JackUp() string {
	b.Work()
	return "CraneB jack up"
}

func (b CraneB) Hoist() string {
	b.Work()
	return "CraneB hoist"
}

type Company struct {
	crane Crane
}

func (c *Company) Build() {
	fmt.Println(c.crane.JackUp())
	fmt.Println(c.crane.Hoist())
	fmt.Println("建设完成")
}

func ReadFile(file *os.File) ([]byte, error) {
	buffer := make([]byte, 0, 512)
	for {
		if len(buffer) == cap(buffer) {
			// 扩容
			buffer = append(buffer, 0)[:len(buffer)]
		}
		offset, err := file.Read(buffer[len(buffer):cap(buffer)])
		buffer = buffer[:len(buffer) + offset]
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
			}
			return buffer, err
		}
	}
}

func Fibonacci(n int) func(yield func(int) bool) {
	a, b, c := 0, 1, 1
	return func(yield func(int) bool) {
		for range n {
			if !yield(a) {
				return
			}
			a, b := b, c
			c = a + b
		}
	}
}

type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

func initDatabase(host string, port int) {
	defer fmt.Println("A")
	defer fmt.Println("B")
	if len(host) == 0 || port == 0 {
		panic("无效的数据库连接参数")
	} else {
		fmt.Println("数据库已连接")
		defer fmt.Println("C")
	}
}

func dangerOp() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	panic("dangerOp panic")
	defer fmt.Println("3")
}

type MyInterface interface {
	My() string
}

type MyStruct struct{
}

func (m MyStruct) My() string {
	return "my"
}

type OtherStruct struct {
}

func (o OtherStruct) String() string {
	return "other"
}
