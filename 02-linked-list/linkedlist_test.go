package linkedlist

import (
	"math/rand"
	"slices"
	"testing"
)

// 测试思路：从边界、一致性、顺序、压力四个维度验证链表的核心不变量：
// ① size 与"从头可遍历的节点数"恒等；② 任意操作后 Traverse 与操作序列一致；
// ③ 链有终点（无环无断链）；④ 越界操作不改变链表。
// 测试通过公开 API 操作，用 Traverse + Len 验证结果（它们就是链表结构的镜像，
// 不需要直接摸内部字段）。

// TestNew 验证创建：空链表 size 为 0、sentinel.Next 为 nil（无真实节点）。
// 越界/未找到等行为归各自方法的测试（TestDelete/TestSearch），不在此重复。

func TestNew(t *testing.T) {
	l := New()

	if l.size != 0 {
		t.Errorf("New() 后 size = %d, 期望 0", l.size)
	}
	if l.sentinel.Next != nil {
		t.Errorf("New() 后 sentinel.Next 应为 nil, 实际指向节点")
	}
}

// TestInsert 验证顺序与边界：在头、中、尾各插入一次，遍历输出符合预期；
// i == Len() 等价尾部追加；i < 0 或 i > Len() 报 ErrIndexOutOfRange，
// 且报错后链表内容不变（用 Traverse 验证）
func TestInsert(t *testing.T) {
	l := New()

	size := 4
	for i := 0; i < size; i++ {
		l.Insert(i, i)
	}

	err := l.Insert(-1, -1)
	if err == nil {
		t.Errorf("Insert(-1, -1) 应返回 ErrIndexOutOfRange, 实际为 nil")
	}
	err = l.Insert(size+1, -1)
	if err == nil {
		t.Errorf("Insert(size+1=%d, -1) 应返回 ErrIndexOutOfRange, 实际为 nil", size+1)
	}

	traverseRes := l.Traverse()
	res := []int{0, 1, 2, 3}
	if !slices.Equal(traverseRes, res) {
		t.Errorf("Traverse() = %v, 期望 %v", traverseRes, res)
	}

	l.Insert(0, 5)
	l.Insert(3, 6)
	l.Insert(6, 7)

	traverseRes = l.Traverse()
	res = []int{5, 0, 1, 6, 2, 3, 7}

	if !slices.Equal(traverseRes, res) {
		t.Errorf("Traverse() = %v, 期望 %v", traverseRes, res)
	}
}

// TestDelete 验证顺序与返回值：删头、删中间、删尾，返回值分别正确、遍历无空洞；
// 删除唯一节点后链表为空（Len 0、Traverse 空）；i < 0 或 i >= Len() 报错，
// 且报错后链表内容不变
func TestDelete(t *testing.T) {
	l := New()

	size := 4
	for i := 0; i < size; i++ {
		l.Insert(i, i)
	}

	_, err := l.Delete(-1)
	if err == nil {
		t.Errorf("Delete(-1) 应返回 ErrIndexOutOfRange, 实际为 nil")
	}
	_, err = l.Delete(size + 1)
	if err == nil {
		t.Errorf("Delete(size+1=%d) 应返回 ErrIndexOutOfRange, 实际为 nil", size+1)
	}

	v, _ := l.Delete(0)
	if v != 0 {
		t.Errorf("Delete(0) 返回值 = %d, 期望 0", v)
	}
	v, _ = l.Delete(1)
	if v != 2 {
		t.Errorf("Delete(1) 返回值 = %d, 期望 2", v)
	}
	v, _ = l.Delete(1)
	if v != 3 {
		t.Errorf("Delete(1) 返回值 = %d, 期望 3", v)
	}

	traverseRes := l.Traverse()
	res := []int{1}

	if !slices.Equal(traverseRes, res) {
		t.Errorf("Traverse() = %v, 期望 %v", traverseRes, res)
	}

	v, _ = l.Delete(0)
	if v != 1 {
		t.Errorf("Delete(0)（删唯一节点）返回值 = %d, 期望 1", v)
	}
	traverseRes = l.Traverse()
	res = make([]int, 0)

	if !slices.Equal(traverseRes, res) {
		t.Errorf("Traverse() = %v, 期望 %v", traverseRes, res)
	}
}

// TestSearch 验证查找语义：命中返回第一个匹配的下标（重复值取最小下标）；
// 未命中返回 ErrNotFound；空表搜索同样返回 ErrNotFound
func TestSearch(t *testing.T) {
	l := New()
	_, err := l.Search(-1)
	if err == nil {
		t.Errorf("Search(-1) 应返回 ErrNotFound, 实际为 nil")
	}

	size := 5
	for i := 0; i < size; i++ {
		l.Insert(i, i)
	}

	_, err = l.Search(-1)
	if err == nil {
		t.Errorf("Search(-1) 应返回 ErrNotFound, 实际为 nil")
	}
	_, err = l.Search(6)
	if err == nil {
		t.Errorf("Search(6) 应返回 ErrNotFound, 实际为 nil")
	}

	l.Insert(5, 4)
	v, err := l.Search(4)
	if err != nil {
		t.Errorf("Search(4) 返回错误: %v", err)
	}
	if v != 4 {
		t.Errorf("Search(4) 返回下标 = %d, 期望 4", v)
	}
}

// TestTraverse 验证顺序：与插入顺序一一对应；空表返回非 nil 空切片（可安全 len/遍历）
func TestTraverse(t *testing.T) {
	l := New()

	traverseRes := l.Traverse()
	if traverseRes == nil {
		t.Errorf("空表 Traverse() 返回 nil, 期望非 nil 空切片")
	}
	if !slices.Equal(traverseRes, []int{}) {
		t.Errorf("空表 Traverse() = %v, 期望空", traverseRes)
	}

	size := 5
	for i := 0; i < size; i++ {
		l.Insert(i, i+1)
	}

	traverseRes = l.Traverse()
	if !slices.Equal(traverseRes, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Traverse() = %v, 期望 [1 2 3 4 5]", traverseRes)
	}
}

// TestStress 验证压力：固定种子的随机混合 Insert/Delete（下标取合法范围），
// 用 Go 原生切片同步模拟同一串操作作影子模型，每步断言：
// Len() 与影子长度一致，Traverse() 与影子完全相等
func TestStress(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	l := New()
	ref := make([]int, 0)

	for step := 0; step < 1000; step++ {
		if l.Len() == 0 {
			// 空表只能插，不能删
			v := r.Intn(100)
			l.Insert(0, v)
			ref = append(ref, v)
		} else {
			switch r.Intn(2) {
			case 0: // 随机位置插入
				i := r.Intn(l.Len() + 1)
				v := r.Intn(100)
				if err := l.Insert(i, v); err != nil {
					t.Fatalf("第 %d 步 Insert(%d, %d) 返回错误: %v", step, i, v, err)
				}
				ref = append(ref, 0)
				copy(ref[i+1:], ref[i:])
				ref[i] = v
			case 1: // 随机位置删除
				i := r.Intn(l.Len())
				v, err := l.Delete(i)
				if err != nil {
					t.Fatalf("第 %d 步 Delete(%d) 返回错误: %v", step, i, err)
				}
				if v != ref[i] {
					t.Fatalf("第 %d 步 Delete(%d) 返回值 = %d, 参照 %d", step, i, v, ref[i])
				}
				ref = append(ref[:i], ref[i+1:]...)
			}
		}

		if l.Len() != len(ref) {
			t.Fatalf("第 %d 步后 Len() = %d, 参照长度 %d", step, l.Len(), len(ref))
		}
		got := l.Traverse()
		if !slices.Equal(got, ref) {
			t.Fatalf("第 %d 步后 Traverse() = %v, 参照 %v", step, got, ref)
		}
	}
}
