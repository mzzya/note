package Sort

import (
	"sort"
	"sync"
)

//快速排序算法
// func main() {
// 	//QuickSort(Algorithm.ArrayA)
// 	Algorithm.ArrayA = []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
// 	QuickSort2(Algorithm.ArrayA)
// 	Algorithm.ShowArrayA(Algorithm.ArrayA)
// }

// 第二种写法
func QuickSort2(values []int) {
	if len(values) <= 1 {
		return
	}
	key, i := values[0], 1
	left, right := 0, len(values)-1
	for left < right {
		//log.Println(values)
		if values[i] > key {
			values[i], values[right] = values[right], values[i]
			right--
		} else {
			values[i], values[left] = values[left], values[i]
			left++
			i++
		}
	}
	// log.Println("========================================")
	// log.Println(values)
	values[left] = key
	QuickSort2(values[:left])
	QuickSort2(values[left+1:])
}

func QuickSort2_Chan(values []int) {
	if len(values) <= 1 {
		return
	}
	key, i := values[0], 1
	left, right := 0, len(values)-1
	for left < right {
		//log.Println(values)
		if values[i] > key {
			values[i], values[right] = values[right], values[i]
			right--
		} else {
			values[i], values[left] = values[left], values[i]
			left++
			i++
		}
	}
	// log.Println("========================================")
	// log.Println(values)
	values[left] = key
	signal := make(chan int, 2)
	go func() {
		QuickSort2_Chan(values[:left])
		signal <- 1
	}()
	go func() {
		QuickSort2_Chan(values[left+1:])
		signal <- 1
	}()

	<-signal
	<-signal
	//close(signal)
}

func QuickStandard(values []int) {
	sort.Ints(values)
}

func QuickSort(values []int) {
	if len(values) <= 1 {
		return
	}
	//KEY为第一个值
	key := values[0]
	//左右移动位置
	left, right := 0, len(values)-1
	for left != right {
		//循环从右侧找比key小数据
		for left < right && values[right] >= key {
			right--
		}
		//right位比key小则将rigth位的值放在左侧边位
		//
		values[left] = values[right]
		//循环从左侧开始找比key大数据 然后将该数据放在之前右侧的空档为,
		for left < right && values[left] <= key {
			left++
		}
		values[right] = values[left]
	}
	values[left] = key
	QuickSort(values[:left])
	QuickSort(values[left+1:])
}

func QuickSort_Chan(values []int) {
	if len(values) <= 1 {
		return
	}
	//KEY为第一个值
	key := values[0]
	//左右移动位置
	left, right := 0, len(values)-1
	for left < right {
		//循环从右侧找比key小数据
		for left < right && values[right] >= key {
			right--
		}
		//right位比key小则将rigth位的值放在左侧边位
		values[left] = values[right]
		//循环从左侧开始找比key大数据 然后将该数据放在之前右侧的空档为
		for left < right && values[left] <= key {
			left++
		}
		values[right] = values[left]
	}
	values[left] = key
	signal := make(chan int, 2)
	go func() {
		QuickSort(values[:left])
		signal <- 1
	}()
	go func() {
		QuickSort(values[left+1:])
		signal <- 1
	}()
	<-signal
	<-signal
	close(signal)
}

func QuickSort_Sync(values []int) {
	if len(values) <= 1 {
		return
	}
	//KEY为第一个值
	key := values[0]
	//左右移动位置
	left, right := 0, len(values)-1
	for left < right {
		//循环从右侧找比key小数据
		for left < right && values[right] >= key {
			right--
		}
		//right位比key小则将rigth位的值放在左侧边位
		values[left] = values[right]
		//循环从左侧开始找比key大数据 然后将该数据放在之前右侧的空档为
		for left < right && values[left] <= key {
			left++
		}
		values[right] = values[left]
	}
	values[left] = key

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		QuickSort_Sync(values[:left])
		wg.Done()
	}()
	go func() {
		QuickSort_Sync(values[left+1:])
		wg.Done()
	}()
	wg.Wait()
	//close(signal)
}
