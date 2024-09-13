package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strconv"

	"example.com/greetings"
	"example.com/greetings/subpackagename"
)

// 定义接口，泛类型的接口实现，意义大不大样子
type graphics interface {
	getArea() float64
}

/* 定义结构体 */
type Circle struct {
	radius float64
	area   float64
	name   string
}

// 定义结构体的方法  -- 也可用接口
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

// 返回闭包匿名函数
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 函数变量声明
var add = func(a, b int) int {
	return a + b
}

// 函数作为参数
func fc(operation func(int, int) int, x, y int) int {
	return operation(x, y)
}

// 函数接受一个数组作为参数，不修改原数组
func modifyArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

// 函数接受一个数组的指针作为参数，修改原数组
func modifyArrayWithPointer(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}
}

// 切片信息
func sprintSlice(x []int) string {
	return fmt.Sprintf("len=%d cap=%d ", len(x), cap(x))
}

func init() {
	log.Println("init:$USER not set")

}

func main() {
	var c1 = Circle{}
	c1.radius = 10.01
	c1.area = c1.getArea()
	c1.name = "圆"
	fmt.Println("圆的面积 = ", c1.area)

	numX := getSequence()
	fmt.Println("计数器numX", numX())
	fmt.Println("计数器numX", numX())

	fmt.Println("add = ", add(1, 2))

	fmt.Println("fc add = ", fc(add, 10, 10))

	var graphics1 = graphics(Circle{radius: 13})
	fmt.Println("graphics1 = ", graphics1.getArea())

	// 数组声明
	var numbers1 = [5]int{}
	var numbers2 = [5]int{1, 2, 3, 4, 5}
	var numbers3 = []int{1, 2, 3, 4, 5} // 切片，保持切片声明使用，不要存在引用关系
	var numbers4 = numbers2[2:4]        // 数组的切片，不要使用该方式，会存在引用关系。
	var numbers5 = numbers3[2:4]
	fmt.Println("numbers1 = ", numbers1, len(numbers1))
	fmt.Println("numbers2 = ", numbers2, len(numbers2))
	fmt.Println("numbers3 = ", numbers3, sprintSlice(numbers3))
	fmt.Println("numbers4 = ", numbers4, sprintSlice(numbers4))
	fmt.Println("numbers5 = ", numbers5, sprintSlice(numbers5))
	// 切片是索引，指向原数组，修改会影响原数组
	numbers4 = append(numbers4, 6)              // 影响原数据
	numbers5 = append(numbers5, 6, 7, 8, 9, 10) // 发生自动扩容，拷贝数据，不影响原数据

	fmt.Println("numbers2 = ", numbers2, len(numbers2))
	fmt.Println("numbers3 = ", numbers3, sprintSlice(numbers3))
	fmt.Println("numbers4 = ", numbers4, sprintSlice(numbers4))
	fmt.Println("numbers5 = ", numbers5, sprintSlice(numbers5))

	// 指针

	// 为指针变量赋值
	var fp = &c1.radius
	fmt.Println("fp = ", fp, *fp)
	var cp = &c1
	fmt.Println("cp = ", cp, *cp)

	myArray := []int{1, 2, 3, 4, 5}
	// 函数接受一个数组作为参数，不修改原数组
	modifyArray(myArray)
	fmt.Println("Array after modifyArray:", myArray)
	// 函数接受一个数组的指针作为参数，修改原数组
	modifyArrayWithPointer(&myArray)
	fmt.Println("Array after modifyArray:", myArray)

	// 多维数组
	var a2 = [2][3]int{}
	fmt.Println("a2 = ", a2, len(a2), len(a2[0]))
	var s2 = [][]int{}
	s2 = append(s2, []int{1, 2, 3})
	fmt.Println("s2 = ", s2, len(s2), len(s2[0]))

	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
	for key, value := range myArray {
		fmt.Println(key, ":", value, " ")
	}

	// map
	m := map[string]int{"apple": 1, "banana": 2, "orange": 3}
	// 判断某个键值是否存在
	v, ok := m["apple"]
	if ok {
		fmt.Println("apple = ", v)
	}

	// 类型转换
	radiusI := int(c1.radius)
	fmt.Println("radiusI = ", radiusI)

	numI, _ := strconv.Atoi("10")
	numF, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Println("num = ", numI, numF)

	str1 := strconv.Itoa(10)
	fmt.Println("str1 = ", str1)
	str2 := strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Println("str2 = ", str2)

	// greetings
	fmt.Println(greetings.Hello("world xx"))
	fmt.Println(subpackagename.Hello("world xx"))

Loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break Loop
			}
			fmt.Println("i = ", i, "j = ", j)
		}
	}

	c := '#' // 字符
	switch c {
	default: // default 语句用于没有匹配的情况，默认最后执行
		fmt.Println("c is not a special character")
	case ' ', '?', '&', '=', '#', '+', '%':
		fmt.Println("c is a special character")
	}

	fmt.Printf("%q %q %q\n", "x'x", "abc", "abc	")

	fmt.Println(min(1, 2, 3, 4, 5))
	var addr = *(flag.String("addr", ":1718", "http service address")) // name不能重复定义 1718端口，用于生成二维码
	flag.Parse()
	fmt.Println("addr = ", addr)
	fmt.Println(flag.Args())

	// json字符串解析
	jsonString := `{"name": "John Doe", "age": 30, "email": "john.doe@example.com", "location": {"city": "New York", "country": "USA"}}`
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonString), &result)
	fmt.Println(result)
	result["name"] = "Tom"
	result["location"].(map[string]interface{})["city"] = "Shanghai"
	fmt.Println(result)
}
