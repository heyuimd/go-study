package main

import (
	"fmt"
	"strings"
	"time"
)

func lenAndUpper(name string) (int, string) {
	defer fmt.Println("done")
	return len(name), strings.ToUpper(name)
}

func main() {
	names := []string{"a"}
	names = append(names, "b")
	fmt.Println(names)

	me := map[string]string{"1": "a", "2": "b"}
	me["3"] = "c"

	for k, v := range me {
		fmt.Println(k, v)
	}

	_, upperName := lenAndUpper("hello")
	fmt.Println(upperName)

	sayMe("a", "b", "c")

	fmt.Println(canDrink(10))

	// make()로 가변 길이 배열 만들기
	a := make([]int, 5)
	fmt.Printf("a := make([]int, 5)의\t")
	printSlice(a)

	b := make([]int, 0, 5)
	fmt.Printf("b := make([]int, 0, 5)의\t")
	printSlice(b)

	c := b[:2]
	fmt.Printf("c := b[:2]의\t\t")
	printSlice(c)

	d := c[2:5]
	fmt.Printf("d := c[2:5]의\t\t")
	printSlice(d)

	// 한번에 여러개 원소를 추가할 수 있습니다.
	d = append(d, 1, 2, 3)
	fmt.Printf("d = append(d, 1,2,3)후\t")
	printSlice(d)

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			i, ":",
			pos(i),
			neg(-2*i),
		)
	}

	fmt.Println("① empty interface에 대해")
	var i interface{}
	describe(i)

	fmt.Println("① i = 42에 대해")
	i = 42
	describe(i)
	fmt.Println("① i = \"hello\"에 대해")
	i = "hello"
	describe(i)

	go say("② 다른 루틴")
	say("① 이 루틴")

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("100\t밀리초 지났어요.")
		case <-boom:
			fmt.Println("500\t밀리초 지났어요. 종료합니다!")
			return
		default:
			fmt.Println("50\t밀리초 지났어요.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func describe(i interface{}) {
	fmt.Printf("인터페이스 i의 (값, 타입) : (%v, %T)\n", i, i)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func canDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	} else {
		return true
	}
}

func sayMe(me ...string) {
	for idx, name := range me {
		fmt.Println(idx, name)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
