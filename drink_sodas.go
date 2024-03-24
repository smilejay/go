package main

import "fmt"

// 题目：商店为回收汽水瓶，规定3个空瓶换一瓶汽水。
// 一个人买了10瓶汽水喝完之后又拿空瓶去换汽水，请问他一共可以喝多少瓶汽水（包括他已喝的10瓶汽水）？
// 注：可以找商店借一个空瓶

// 喝汽水的函数
func drinkSodas(bottles int) int {
	sodasDrank := bottles
	// 用空瓶换汽水
	for bottles >= 3 {
		bottles = bottles - 3
		sodasDrank++
		bottles += 1
	}
	// 可借一个空瓶换汽水，故只要剩余空瓶为2瓶就可以再喝一瓶汽水
	// 从上面一个循环出来后，bottles只能为0、1、2这三种情况
	if bottles == 2 {
		bottles--
		sodasDrank++
	}
	return sodasDrank
}

func main() {
	// 初始的汽水数量
	initialSodas := 10
	// 计算总共喝的汽水数量
	totalSodasDrank := drinkSodas(initialSodas)
	fmt.Println("他一共可以喝", totalSodasDrank, "瓶汽水")
}
