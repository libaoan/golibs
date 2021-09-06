package main

import (
	"fmt"
	"sort"
)

// 学生成绩结构体
type StuScore struct {
	name  string // 姓名
	score int    // 成绩
}

type StuScores []StuScore

//Len()
func (s StuScores) Len() int {
	return len(s)
}

//Less(): 成绩将有低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

//Swap()
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}

	// 打印未排序的 stus 数据
	fmt.Println("Default:\n\t", stus)
	//StuScores 已经实现了 sort.Interface 接口 , 所以可以调用 Sort 函数进行排序
	sort.Sort(stus)
	// 判断是否已经排好顺序，将会打印 true
	fmt.Println("IS Sorted?\n\t", sort.IsSorted(stus))
	// 打印排序后的 stus 数据
	fmt.Println("Sorted:\n\t", stus)

	// reverse操作, Reverse是对stus对象的封装，改变了Less接口的实现
	sort.Sort(sort.Reverse(sort.Reverse(sort.Reverse(stus))))
	fmt.Println("Reverse Sorted:\n\t", stus)

	// Search
	sort.Sort(stus)
	fmt.Println("Search Greate Than 90:\n\t", stus[sort.Search(len(stus), func(i int) bool { return stus[i].score >= 91 })])

	s := []int{5, 2, 6, 3, 1, 4} // 未排序的切片数据
	sort.Ints(s)
	fmt.Println(s) // 将会输出[1 2 3 4 5 6]

	s = []int{5, 2, 6, 3, 1, 4} // 未排序的切片数据
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s) // 将会输出[6 5 4 3 2 1]
	// 注意，SearchInts() 的使用条件为：**切片 a 已经升序排序** 以下是一个错误使用的例子：
	fmt.Println(sort.SearchInts(s, 2)) // 将会输出 0 而不是 1
	// 正确的调用
	sort.Ints(s)
	fmt.Println(s)
	fmt.Println(sort.SearchInts(s, 2))

}
