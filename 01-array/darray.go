package darray

import "errors"

// ErrIndexOutOfRange 下标越界时返回的错误
var ErrIndexOutOfRange = errors.New("index out of range")

// DynamicArray 动态数组：底层连续存储，容量不足时自动扩容
type DynamicArray struct {
	data []int // 底层切片，cap 即容量
	size int   // 当前元素个数
}

// New 创建动态数组，initialCap 为初始容量（可为 0）
func New(initialCap int) *DynamicArray {
	if initialCap < 0 {
		panic("initialCap 必须大于等于 0")
	}

	return &DynamicArray{
		data: make([]int, initialCap),
		size: 0,
	}
}

// Len 返回当前元素个数
func (a *DynamicArray) Len() int {
	return a.size
}

// Cap 返回当前底层容量
func (a *DynamicArray) Cap() int {
	return len(a.data)
}

// Get 返回下标 i 的元素；越界返回 ErrIndexOutOfRange
func (a *DynamicArray) Get(i int) (int, error) {
	if i < 0 || i >= a.size {
		return 0, ErrIndexOutOfRange
	}

	return a.data[i], nil
}

// Set 把下标 i 的元素改为 v；越界返回 ErrIndexOutOfRange
func (a *DynamicArray) Set(i, v int) error {
	if i < 0 || i >= a.size {
		return ErrIndexOutOfRange
	}

	a.data[i] = v
	return nil
}

// Append 尾部追加 v；容量不足时自动扩容
func (a *DynamicArray) Append(v int) {
	l := len(a.data)
	if a.size == l {
		olddata := a.data

		if l == 0 {
			l = 2
		} else {
			l *= 2
		}

		a.data = make([]int, l)
		copy(a.data, olddata)
	}

	a.data[a.size] = v
	a.size++
}

// Insert 在下标 i 处插入 v，i 及其后的元素后移；i 可为 size（等价尾部追加），越界返回错误
func (a *DynamicArray) Insert(i, v int) error {
	if i < 0 || i > a.size {
		return ErrIndexOutOfRange
	}

	if i == a.size {
		a.Append(v)
		return nil
	}

	newdatalen := len(a.data)
	if len(a.data) == a.size {
		newdatalen *= 2
	}
	newdata := make([]int, newdatalen)

	for j := 0; j < a.size; j++ {
		if j == i {
			newdata[j] = v
		}

		if j < i {
			newdata[j] = a.data[j]
		} else {
			newdata[j+1] = a.data[j]
		}
	}

	a.data = newdata
	a.size++
	return nil
}

// Remove 删除下标 i 的元素并返回其值，后续元素前移；越界返回 ErrIndexOutOfRange
func (a *DynamicArray) Remove(i int) (int, error) {
	if i < 0 || i >= a.size {
		return 0, ErrIndexOutOfRange
	}

	val := 0
	for j := 0; j < a.size; j++ {
		if j == i {
			val = a.data[j]
			continue
		}

		if j > i {
			a.data[j-1] = a.data[j]
		}
	}

	a.size--
	return val, nil
}
