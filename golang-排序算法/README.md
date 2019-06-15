# 算法

- [排序算法](#排序算法)
    - [快速排序](#快速排序)
    - [冒泡排序](#冒泡排序)
	- [直接插入排序](#直接插入排序)
	- [选择排序](#选择排序)
	- [归并排序](#归并排序)

## 排序算法

### 快速排序
```
func Quick(array []int) {
	if len(array) <= 1 {
		return
	}
	//将第一个值作为比较项
	key := array[0]
	//定义左侧和右侧边间
	i, j := 0, len(array)-1
	//当循环内还存在数据时
	for i < j {
		//从右侧开始找大于key的值，存在则从右侧缩小范围
		for i < j && array[j] >= key {
			j--
		}
		//从右侧找到小于key的值时，将值给换到最左侧
		array[i] = array[j]
		//从左侧开始找小于key的值，存在则从左侧缩小范围
		for i < j && array[i] <= key {
			i++
		}
		array[j] = array[i]
	}
	//循环退出时必然存在i=j 所以下方 i 或 j 无所谓
	array[i] = key
	Quick(array[:i])
	Quick(array[i+1:])
}
```
### 冒泡排序
```
//Bubble 冒泡排序
func Bubble(array []int) {
	for i := 0; i < len(array)-1; i++ {
		//内循环从下一个数开始，如果比第一个数小则交换位置
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
		//这样一圈下来，最小的数一定在最左侧
	}
}
```
### 直接插入排序
```
//Insert 插入排序
func Insert(array []int) {
	//234 75788
	for i := 1; i < len(array); i++ {
		//如果当前的数大于已排序好的最后一个数则跳过
		if array[i] > array[i-1] {
			continue
		}
		//简版判断 当前值比前一个值小则交换两者数据
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j-1], array[j] = array[j], array[j-1]
		}

		// //保存当前值
		// var temp = array[i]
		// var j int
		// //从已经排好的队尾取值与当前数字比较，如果取出的值大于当前值，则将队尾值向后移动
		// for j = i - 1; j >= 0 && array[j] > temp; j-- {
		// 	array[j+1] = array[j]
		// }
		// //最坏j=-1,排好的队列全部右移
		// array[j+1] = temp
	}
}
```
### 选择排序
```
//Select 选择排序
func Select(array []int) {
	for i := 0; i < len(array)-1; i++ {
		//假设下标未0位最小项
		var minIndex = i
		//排除已排序好的位置
		for j := i + 1; j < len(array); j++ {
			//在剩余数字中寻找最小值 即为已排序完成的最大值
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
}
```
### 归并排序
	
```
//merge 归并排序
func merge(array []int) []int {
	var aryLen = len(array)
	if aryLen <= 1 {
		return array
	}
	//首先拆成最小交换单元数据
	var partCount = aryLen / 2
	left := merge(array[0:partCount])
	right := merge(array[partCount:])
	return mergeCore(left, right)
}

//mergeCore 归并排序合并方法
func mergeCore(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	//循环结束时 必定有一个集合为空 另一个集合有1+以上值最坏全部值，
	//切改值必定大于已排序完成的集合
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
```