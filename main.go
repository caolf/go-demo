package main

import (
	"fmt"
	"go-demo/hello"
)

func main() {

	//TestPackage()
	//TestArrayType()

	//TestStructType()

	//TestSlice()
	//TestSliceParams()

	//容易产生的问题
	// 如何解决 ？？？
	TestSliceBug()
}

// TestPackage 包管理
func TestPackage() {
	hello.PrintHello()
	hello.Test()
	fmt.Printf("使用其他package的变量：%v\n", hello.AA)
}

// TestArrayType 数组类型
func TestArrayType() {
	a := [3]int{1, 2}
	fmt.Println("数组初始值:", a)
	fmt.Println("数组初始值:", len(a), cap(a))

	a[1] = 124
	fmt.Printf("数组赋值前：%p , %v\n", &a, a)
	c := a
	fmt.Printf("数组赋值c=a: %p\n", &c)

	fmt.Printf("数组a值传递的地址:%p\n", &a)
	TestArrayWhenIsParam(a)       // 不同于其他语言 不建议使用
	TestArrayWhenIsParamPoint(&a) // 与其他语言一样
	fmt.Printf("数组a值传递的地址:%p, %v\n", &a, a)
}

func TestArrayWhenIsParamPoint(i *[3]int) {
	fmt.Printf("参数数组指针时：%p, %v\n", i, *i)
	(*i)[0]++
	fmt.Printf("参数数组指针时：%p, %v\n", i, *i)
}

// TestArrayWhenIsParam 数组作为参数时
func TestArrayWhenIsParam(a [3]int) {
	fmt.Printf("参数是数组时:%p, %v,  %T\n", &a, a, a)
	a[1] = 33
	fmt.Printf("参数是数组时：%p , %v\n", &a, a)
}

// TestStructType 结构体
func TestStructType() {
	b := [...]struct {
		name string
		age  int
	}{
		{"user1", 1},
		{"user2", 2},
	}

	fmt.Printf("结构体：%v, %T, %p \n", b, b, &b)
	for index, v := range b {
		fmt.Println(index, v)
	}

	a := b[0]
	fmt.Printf("结构体参数调用前：%v, %T, %p \n", a, a, &a)
	TestPrintStruct(a)
	fmt.Printf("结构体参数调用后：%v, %T, %p \n", a, a, &a)

}

func TestPrintStruct(b struct {
	name string
	age  int
}) {
	fmt.Printf("结构体参数：%v, %T, %p \n", b, b, &b)
	b.age = 100
	fmt.Printf("结构体参数：%v, %T, %p \n", b, b, &b)
}

func TestSliceParams() {
	s1 := []int{12, 34}
	fmt.Printf("切片作为参数前： %T, %p, %v \n", s1, &s1, s1)
	s1 = s1
	fmt.Printf("切片作为参数前： %T, %p, %v \n", s1, &s1, s1)
	TestSliceWhenIsParamsPoint(&s1)
	//传数组指针
	TestSliceWhenIsParams(s1)
	fmt.Printf("切片作为参数后： %T, %p, %v \n", s1, &s1, s1)

	//两个切片地址，但是指的值是一样
	arrayB := s1[:]
	s1[0]++
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)
	fmt.Printf("s1 : %p , %v\n", &s1, s1)

}

func TestSliceWhenIsParams(s1 []int) {
	fmt.Printf("切片作为参数： %T, %p, %v \n", s1, &s1, s1)
	s1[1] += 100
	fmt.Printf("切片作为参数： %T, %p, %v \n", s1, &s1, s1)
	s2 := append(s1, 200)
	fmt.Printf("切片作为参数： %T, %p, %v \n", s2, &s2, s2)

}

func TestSliceWhenIsParamsPoint(s1 *[]int) {
	fmt.Printf("切片作为指针： %T, %p, %v \n", s1, s1, &s1)
	(*s1)[1] += 100
	fmt.Printf("切片作为指针： %T, %p, %v \n", s1, s1, &s1)
}

func TestSlice() {
	//1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("切片的初始值：是空")
	} else {
		fmt.Println("切片的初始值：不是空")
	}
	// 2.:= 声明
	s2 := []int{}
	// 3.make()
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	// 4.初始化赋值
	var s4 []int = make([]int, 2, 3)
	fmt.Println(s4, len(s4), cap(s4))
	s5 := []int{1, 2, 3}
	fmt.Println(s5, len(s5), cap(s5))
	// 5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包
	s6 = arr[1:4]
	fmt.Println(s6)
}

func TestSliceStuff(value []string) {
	fmt.Printf("value=%v, %T\n", value, value)

	value2 := value[:]
	value2 = append(value2, "b")
	fmt.Printf("value=%v, %p, value2=%v, %p,\n", value, &value, value2, &value2)

	value2[0] = "z"
	fmt.Printf("value=%v, value2=%v\n", value, value2)
}

func TestSliceBug() {
	slice1 := []string{"a"} // 长度 1, 容量 1

	TestSliceStuff(slice1)
	// Output:
	// value=[a] -- ok
	// value=[a], value2=[a b] -- ok: value 未改变, value2 被更新
	// value=[a], value2=[z b] -- ok: value 未改变, value2 被更新

	slice10 := make([]string, 1, 10) // 长度 1, 容量 10
	slice10[0] = "a"

	TestSliceStuff(slice10)
	// Output:
	// value=[a] -- ok
	// value=[a], value2=[a b] -- ok: value 未改变, value2 被更新
	// value=[z], value2=[z b] -- WTF?!? value 改变了???
}
