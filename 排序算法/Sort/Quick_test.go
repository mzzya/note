package Sort

import (
	"math/rand"
	"testing"
	"time"
)

var SourceAry []int

type TestAry struct {
	Used bool
	Arry []int
}

var TestArys []TestAry

func TestMain(m *testing.M) {
	aryLen := 10000 * 100
	SourceAry = make([]int, aryLen)
	SourceAry = GetIntAry(aryLen)
	TestArys = make([]TestAry, 100)
	for i := 0; i < len(TestArys); i++ {
		var item = TestAry{
			Used: false,
		}
		copy(item.Arry, SourceAry)
		TestArys = append(TestArys, item)
	}
	m.Run()
}

func GetIntAry(aryLen int) []int {
	rand.Seed(time.Now().UnixNano())
	ArrayA := make([]int, aryLen)
	for i := 0; i < aryLen; i++ {
		ArrayA[i] = rand.Intn(aryLen)
	}
	return ArrayA
}

func TestSubSync(t *testing.T) {
	for j := 0; j < 2; j++ {
		t.Run("A:1", func(t *testing.T) {
			t.Parallel()
			for i := 0; i < 5; i++ {
				t.Log("a\n")
				time.Sleep(time.Second)
			}
		})
	}
}

func TestQuickStandard(t *testing.T) {
	dataAry := make([]int, len(SourceAry))
	copy(dataAry, SourceAry)
	QuickStandard(dataAry)
	// t.Logf("SourceAry:%+v\n", SourceAry)
	// t.Logf("dataAry:%+v\n", dataAry)
	for i := 0; i < len(SourceAry)-1; i++ {
		if dataAry[i] > dataAry[i+1] {
			t.Fail()
		}
	}
}

func TestQuickSort(t *testing.T) {
	dataAry := make([]int, len(SourceAry))
	copy(dataAry, SourceAry)
	QuickSort(dataAry)
	// t.Logf("SourceAry:%+v\n", SourceAry)
	// t.Logf("dataAry:%+v\n", dataAry)
	for i := 0; i < len(SourceAry)-1; i++ {
		if dataAry[i] > dataAry[i+1] {
			t.Fail()
		}
	}
}

func TestQuickSort_Chan(t *testing.T) {
	dataAry := make([]int, len(SourceAry))
	copy(dataAry, SourceAry)
	QuickSort_Chan(dataAry)
	// t.Logf("SourceAry:%+v\n", SourceAry)
	// t.Logf("dataAry:%+v\n", dataAry)
	for i := 0; i < len(SourceAry)-1; i++ {
		if dataAry[i] > dataAry[i+1] {
			t.Fail()
		}
	}
}

func TestQuickSort_Sync(t *testing.T) {
	dataAry := make([]int, len(SourceAry))
	copy(dataAry, SourceAry)
	QuickSort_Sync(dataAry)
	// t.Logf("SourceAry:%+v\n", SourceAry)
	// t.Logf("dataAry:%+v\n", dataAry)
	for i := 0; i < len(SourceAry)-1; i++ {
		if dataAry[i] > dataAry[i+1] {
			t.Fail()
		}
	}
}

// func TestQuickSort2(t *testing.T) {
// 	dataAry := make([]int, len(SourceAry))
// 	copy(dataAry, SourceAry)
// 	QuickSort(dataAry)
// 	// t.Logf("SourceAry:%+v\n", SourceAry)
// 	// t.Logf("dataAry:%+v\n", dataAry)
// 	for i := 0; i < len(SourceAry)-1; i++ {
// 		if dataAry[i] > dataAry[i+1] {
// 			t.Fail()
// 		}
// 	}
// 	//t.Parallel()
// }
func BenchmarkQuickStandard(b *testing.B) {
	dataAry := make([]int, len(SourceAry))
	copy(dataAry, SourceAry)
	b.ReportAllocs()
	QuickStandard(dataAry)
}

func BenchmarkQuickSort(b *testing.B) {
	dataAry := make([]int, len(SourceAry))
	copy(dataAry, SourceAry)
	b.ReportAllocs()
	QuickSort(dataAry)
}

func BenchmarkQuickSort_Chan(b *testing.B) {
	dataAry := make([]int, len(SourceAry))
	copy(dataAry, SourceAry)
	b.ReportAllocs()
	QuickSort_Chan(dataAry)
}

func BenchmarkQuickSort_Sync(b *testing.B) {
	dataAry := make([]int, len(SourceAry))
	copy(dataAry, SourceAry)
	b.ReportAllocs()
	QuickSort_Sync(dataAry)
}

// func BenchmarkQuickSort2(b *testing.B) {
// 	dataAry := make([]int, len(SourceAry))
// 	copy(dataAry, SourceAry)
// 	b.ReportAllocs()
// 	QuickSort2(dataAry)
// }

// func BenchmarkQuickSort2_Chan(b *testing.B) {
// 	dataAry := make([]int, len(SourceAry))
// 	copy(dataAry, SourceAry)
// 	b.ReportAllocs()
// 	QuickSort2_Chan(dataAry)
// 	//b.Errorf("%+v", dataAry)
// }
