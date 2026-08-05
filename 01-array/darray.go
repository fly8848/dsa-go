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
func New(initialCap int) *DynamicArray { panic("TODO") }

// Len 返回当前元素个数
func (a *DynamicArray) Len() int { panic("TODO") }

// Cap 返回当前底层容量
func (a *DynamicArray) Cap() int { panic("TODO") }

// Get 返回下标 i 的元素；越界返回 ErrIndexOutOfRange
func (a *DynamicArray) Get(i int) (int, error) { panic("TODO") }

// Set 把下标 i 的元素改为 v；越界返回 ErrIndexOutOfRange
func (a *DynamicArray) Set(i, v int) error { panic("TODO") }

// Append 尾部追加 v；容量不足时自动扩容
func (a *DynamicArray) Append(v int) { panic("TODO") }

// Insert 在下标 i 处插入 v，i 及其后的元素后移；i 可为 size（等价尾部追加），越界返回错误
func (a *DynamicArray) Insert(i, v int) error { panic("TODO") }

// Remove 删除下标 i 的元素并返回其值，后续元素前移；越界返回 ErrIndexOutOfRange
func (a *DynamicArray) Remove(i int) (int, error) { panic("TODO") }
