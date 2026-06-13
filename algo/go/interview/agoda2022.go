package i

import (
	"errors"
	"fmt"
)

// [Question]: Good Tuple count, 2 out of 3 elements in a row are the same.
func Agoda_2022_OA1(seq []int) {

}

// [Question]: Print sorted score bucket count
func Agoda_2022_OA2() {

}

// [Question]: Text operations with cursor
func Agoda_2022_OA3() {

}

// [Question]: Reverse a struct Collection with Reverse way
//
//	input: {6, 4, 1, 5, 7, 8, 9}  -> ouput: {9, 8, 7, 5, 1, 4, 6}
//	input: {1, 8, 6, 4, 3, 2, 5}  -> ouput: {5, 2, 3, 4, 6, 8, 1}
func Agoda_CodeSignal_1() {
	input1 := []int{6, 4, 1, 5, 7, 8, 9}
	input2 := []int{1, 8, 6, 4, 3, 2, 5}

	c1 := Collection{data: input1}
	c2 := Collection{data: input2}

	reversed_c1 := c1.Reverse()
	reversed_c2 := c2.Reverse()

	fmt.Println(reversed_c1)
	fmt.Println(reversed_c2)
}

type Collection struct {
	data []int
}

func (c Collection) First() (error, int) {
	if len(c.data) == 0 {
		return errors.New("No elements available"), 0
	}
	return nil, c.data[0]
}

func (c Collection) Last() (error, int) {
	if len(c.data) == 0 {
		return errors.New("No elements available"), 0
	}
	return nil, c.data[len(c.data)-1]
}

func (c Collection) IsEmpty() bool {
	return len(c.data) == 0
}

func (c Collection) DropFirst() Collection {
	return Collection{data: c.data[1:]}
}

func (c Collection) DropLast() Collection {
	dsize := len(c.data)
	return Collection{data: c.data[:dsize-1]}
}

func (c Collection) Append(input int) *Collection {
	return &Collection{data: append(c.data, input)}
}

func (c Collection) Reverse() *Collection {
	if c.IsEmpty() {
		return &Collection{data: make([]int, 0)}
	}

	remained := c.DropFirst()

	subReversed := remained.Reverse()

	err, first := c.First()

	if err == nil {
		subReversed = subReversed.Append(first)
	}

	return subReversed
}

func (c Collection) Reverse_loop() Collection {
	dsize := len(c.data)
	copied := make([]int, len(c.data))

	for i := 0; i < dsize; i++ {
		copied[i] = c.data[dsize-i-1]
	}

	return Collection{data: copied}
}

// [Question]: two sum question with optimization
// [Time Complexity]: AppendData: O(1) -> O(n), checkSum: O(n) -> O(1)
var sumSlice []int
var sumMap map[int]bool = make(map[int]bool)
var sumDataMap map[int]bool = make(map[int]bool)

func Agoda_CodeSignal_3() {
	AppendData(2)
	AppendData(3)
	AppendData(6)
	AppendData(2)
	AppendData(7)
	AppendData(9)

	fmt.Println(checkSum(5))
	fmt.Println(checkSum(10))
	fmt.Println(checkSum(0))
	fmt.Println(checkSum(15))
	fmt.Println(checkSum(4))
	fmt.Println(checkSum(13))
}

func AppendData(input int) {
	// [Solution 1]:
	// sumSlice = append(sumSlice, input)

	// [Solution 2]:
	for i := 0; i < len(sumSlice); i++ {
		sumMap[input+sumSlice[i]] = true
	}
	sumSlice = append(sumSlice, input)

	// [Solution 3]: Assume max two sum is 19, O(19) => O(1)
	for i := 0; i < 19; i++ {
		if sumDataMap[i-input] {
			sumMap[i] = true
		}
	}
	sumDataMap[input] = true
}

func checkSum(sum int) bool {
	// [Solution 1]:
	// sumMap := make(map[int]bool)
	// for i := 0; i < len(datas); i++ {
	// 	v := sumSlice[i]
	// 	if sumMap[sum-v] {
	// 		return true
	// 	}
	// 	sumMap[v] = true
	// }
	// return false
	// [Solution 2/3]:
	return sumMap[sum]
}
