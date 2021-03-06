package main

import "fmt"

// 简单定义
func sum0(a int, b int) int {
	return a + b
}

// 相邻同类型可以省略前边参数的类型，只声明最后一个
func sum1(a, b int) int {
	return a + b
}

// 我们甚至还可以给返回值命名，这个时候需要通过赋值的方式来更新结果，而且 return 可以不用带返回值
func sum2(a, b int) (res int) {
	res = a + b
	return
}

// 可变参数
func sum3(init int, vals ...int) int {
	sum := init
	fmt.Println(vals, len(vals))
	for _, val := range vals {
		sum += val
	}
	return sum
}

// 返回多个值
func sum4(init int, vals ...int) (int, int) {
	sum := init
	fmt.Println(vals, len(vals))
	for _, val := range vals {
		sum += val
	}
	return sum, len(vals)
}

func testSum() {
	fmt.Println(sum0(1, 2))
	fmt.Println(sum1(1, 2))
	fmt.Println(sum2(1, 2))
	fmt.Println(sum3(0, 1, 2, 3))
	fmt.Println(sum3(0, []int{1, 2, 3}...)) // 传入一个slice
	fmt.Println(sum4(0, 1, 2, 3))
}

func changeStr(s string) {
	s = "hehe"
	fmt.Println(s)
}

func test_changeStr() {
	name := "lao wang"
	changeStr(name)
	fmt.Println(name) // 打印出来还是 "lao wang"，没有修改成功
}
func changeMap(m map[string]string) {
	m["王八"] = "绿豆"
}

func changeString(s *string) {
	*s = "new lao wang"
}

func testAnonymousFunc() {
	func(s string) {
		fmt.Println(s)
	}("hehe")
}

func testFuncType() {
	myPrint := func(s string) { fmt.Println(s) }
	myPrint("hello go")
}

func testMapFunc() {
	funcMap := map[string]func(int, int) int{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
	}

	fmt.Println(funcMap["add"](3, 2))
	fmt.Println(funcMap["sub"](3, 2))
}

func Double(n int) int {
	return n * 2
}
func Apply(n int, f func(int) int) int {
	return f(n) // f 的类型是 "func(int) int"
}
func funcAsParam() {
	fmt.Println(Apply(10, Double))
}

func FilterIntSlice(intVals []int, predicate func(i int) bool) []int {
	res := make([]int, 0)
	for _, val := range intVals {
		if predicate(val) {
			res = append(res, val)
		}
	}
	return res
}

// 闭包示例
func testClosure() {
	suffix := ".go"
	addSuffix := func(name string) string {
		return name + suffix // 这里使用到了 suffix 这个变量，所以 addSuffix 就是一个闭包
	}
	fmt.Println(addSuffix("hello_world"))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// s := "lao wang"
	// changeString(&s)
	// fmt.Println(s)

	// m := map[string]string{"name": "lao wang"}
	// changeMap(m)
	// fmt.Println(m) // map[name:lao wang 王八:绿豆]
	// testAnonymousFunc()

	// testFuncType()
	// testMapFunc()
	// funcAsParam()

	// ints := []int{1, 2, 3, 4, 5}
	// isOdd := func(i int) bool { return i%2 != 0 } // 是奇数
	// fmt.Println(FilterIntSlice(ints, isOdd))      // [1 3 5]

	// testClosure()
	fmt.Println(fib(10))
}
