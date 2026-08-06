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
// 思路：size 是唯一权威，越界判断只看 i 与 size，与底层容量无关
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

// grow 倍增扩容并搬运旧数据。
// 思路：扩容是 Append 与 Insert 的公共需求，抽出来消除重复；倍增策略保证
// 摊还复杂度 O(1)。调用约定：仅在容量正好满时调用，因此翻倍一次必然够用
// （2*cap >= cap+1），无需循环；容量为 0 时先给基础容量 2。
// 不变量：扩容前后元素内容与顺序完全不变，size 不受影响。
func (a *DynamicArray) grow() {
	old := a.data
	capacity := len(old)
	if capacity == 0 {
		capacity = 2
	} else {
		capacity *= 2
	}

	a.data = make([]int, capacity)
	copy(a.data, old)
}

// Append 尾部追加 v；容量不足时自动扩容。
// 思路：写入位置必须是自增前的 size，顺序反了会整体错位一位（经典坑）。
// 不变量：追加后新元素恰在 data[size-1]，其余元素不动。
func (a *DynamicArray) Append(v int) {
	if a.size == len(a.data) {
		a.grow()
	}

	a.data[a.size] = v
	a.size++
}

// Insert 在下标 i 处插入 v，i 及其后的元素后移；i 可为 size（等价尾部追加），越界返回错误。
// 思路：合法区间 [0, size]，超出会破坏"无空洞"不变量。先保证容量，再原地
// 从后往前搬移腾位——方向必须是后往前，反了会覆盖尚未搬走的元素。
// i == size 时搬移循环不执行，天然退化为尾部追加，无需特判。
func (a *DynamicArray) Insert(i, v int) error {
	if i < 0 || i > a.size {
		return ErrIndexOutOfRange
	}

	if a.size == len(a.data) {
		a.grow()
	}

	for j := a.size; j > i; j-- {
		a.data[j] = a.data[j-1]
	}

	a.data[i] = v
	a.size++
	return nil
}

// Remove 删除下标 i 的元素并返回其值，后续元素前移；越界返回 ErrIndexOutOfRange。
// 思路：删除只减不增，物理空间永远够，不需要扩容也不需要新数组——原地从
// 前往后覆盖即可（方向与前移方向一致，被读的位置总是未被写的）。
// 注意：j 上限是 size-2，最后一个元素无需搬移。
func (a *DynamicArray) Remove(i int) (int, error) {
	if i < 0 || i >= a.size {
		return 0, ErrIndexOutOfRange
	}

	v := a.data[i]
	for j := i; j < a.size-1; j++ {
		a.data[j] = a.data[j+1]
	}

	a.size--
	return v, nil
}
