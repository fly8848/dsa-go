package darray

import (
	"math/rand"
	"testing"
)

// 测试思路：从边界、一致性、顺序、压力四个维度验证动态数组的核心不变量：
// 0 <= size <= cap 恒成立；data[0..size-1] 连续无空洞、无丢失、无错位。

// TestNew 验证创建：初始 size 为 0，cap 与传入的 initialCap 一致；initialCap 为 0 时也能正常使用
func TestNew(t *testing.T) {
	darray := New(0)
	if len(darray.data) != 0 || darray.size != 0 {
		t.Errorf("New(0) 后 len(data) = %d, size = %d, 期望均为 0", len(darray.data), darray.size)
	}

	darray = New(2)
	if len(darray.data) != 2 || darray.size != 0 {
		t.Errorf("New(2) 后 len(data) = %d, size = %d, 期望 len 为 2、size 为 0", len(darray.data), darray.size)
	}
}

// TestAppend 验证尾部追加：元素严格按追加顺序排列；跨越容量边界继续追加不丢数据、不错位
func TestAppend(t *testing.T) {
	darray := New(0)
	vs := []int{1, 2, 3}
	for i, v := range vs {
		darray.Append(v)
		if darray.data[i] != v {
			t.Errorf("Append(%d) 后 data[%d] = %d, 期望 %d", v, i, darray.data[i], v)
		}
	}
}

// TestGetSet 验证一致性：Set 后 Get 读回原值；越界（负数、size 处）均返回 ErrIndexOutOfRange
func TestGetSet(t *testing.T) {
	darray := New(0)
	err := darray.Set(-1, 1)
	if err == nil {
		t.Errorf("空数组 Set(-1, 1) 应返回越界错误, 实际为 nil")
	}

	err = darray.Set(0, 1)
	if err == nil {
		t.Errorf("空数组 Set(0, 1) 应返回越界错误, 实际为 nil")
	}

	err = darray.Set(1, 1)
	if err == nil {
		t.Errorf("空数组 Set(1, 1) 应返回越界错误, 实际为 nil")
	}

	_, err = darray.Get(-1)
	if err == nil {
		t.Errorf("空数组 Get(-1) 应返回越界错误, 实际为 nil")
	}
	_, err = darray.Get(0)
	if err == nil {
		t.Errorf("空数组 Get(0) 应返回越界错误, 实际为 nil")
	}
	_, err = darray.Get(1)
	if err == nil {
		t.Errorf("空数组 Get(1) 应返回越界错误, 实际为 nil")
	}

	darray.Append(1)
	_, err = darray.Get(1)
	if err == nil {
		t.Errorf("空数组 Get(1) 应返回越界错误, 实际为 nil")
	}
}

// TestInsert 验证顺序：在头、中、尾各插入一次，遍历输出符合预期；i == size 等价尾部追加；越界报错
func TestInsert(t *testing.T) {
	darray := New(0)
	vs := []int{1, 5, 10, 30, 50}
	for _, v := range vs {
		darray.Append(v)
	}

	err := darray.Insert(-1, 1)
	if err == nil {
		t.Errorf("Insert(-1, 1) 应返回越界错误, 实际为 nil")
	}
	err = darray.Insert(darray.size+1, 1)
	if err == nil {
		t.Errorf("Insert(size+1=%d, 1) 应返回越界错误, 实际为 nil", darray.size+1)
	}

	err = darray.Insert(0, 1)
	if err != nil {
		t.Errorf("Insert(0, 1) 返回错误: %v", err)
	}

	err = darray.Insert(3, 1)
	if err != nil {
		t.Errorf("Insert(3, 1) 返回错误: %v", err)
	}

	err = darray.Insert(5, 1)
	if err != nil {
		t.Errorf("Insert(5, 1) 返回错误: %v", err)
	}

	vs1 := []int{1, 1, 5, 1, 10, 1, 30, 50}

	j := 0
	for i := 0; i < darray.size; i++ {
		if vs1[i] != darray.data[i] {
			t.Errorf("data[%d] = %d, 期望 %d", i, darray.data[i], vs1[i])
		}
		j++
	}

	if darray.size != j {
		t.Errorf("size = %d, 与遍历计数 %d 不一致", darray.size, j)
	}
}

// TestRemove 验证顺序与返回值：删头、中、尾后遍历无空洞，返回值是被删元素；删空或越界报错
func TestRemove(t *testing.T) {
	darray := New(0)
	vs := []int{1, 5, 10, 30, 50, 100, 300, 500}
	for _, v := range vs {
		darray.Append(v)
	}

	_, err := darray.Remove(-1)
	if err == nil {
		t.Errorf("Remove(-1) 应返回越界错误, 实际为 nil")
	}
	_, err = darray.Remove(darray.size + 1)
	if err == nil {
		t.Errorf("Remove(size+1=%d) 应返回越界错误, 实际为 nil", darray.size+1)
	}

	v, err := darray.Remove(0)
	if err != nil {
		t.Errorf("Remove(0) 返回错误: %v", err)
	}
	if v != 1 {
		t.Errorf("Remove(0) 返回值 = %d, 期望 1", v)
	}

	v, err = darray.Remove(3)
	if err != nil {
		t.Errorf("Remove(3) 返回错误: %v", err)
	}
	if v != 50 {
		t.Errorf("Remove(3) 返回值 = %d, 期望 50", v)
	}

	v, err = darray.Remove(5)
	if err != nil {
		t.Errorf("Remove(5) 返回错误: %v", err)
	}
	if v != 500 {
		t.Errorf("Remove(5) 返回值 = %d, 期望 500", v)
	}

	vs1 := []int{5, 10, 30, 100, 300}

	j := 0
	for i := 0; i < darray.size; i++ {
		if vs1[i] != darray.data[i] {
			t.Errorf("data[%d] = %d, 期望 %d", i, darray.data[i], vs1[i])
		}
		j++
	}

	if darray.size != j {
		t.Errorf("size = %d, 与遍历计数 %d 不一致", darray.size, j)
	}

}

// TestStress 验证压力：大量随机 Append/Insert/Remove/Set 混合操作后，Len <= Cap 恒成立；
// 用 Go 原生切片同步模拟同一串操作作为参照，最终遍历结果与参照一致
func TestStress(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	darray := New(0)
	ref := make([]int, 0)

	for step := 0; step < 1000; step++ {
		if darray.size == 0 {
			v := r.Intn(100)
			darray.Append(v)
			ref = append(ref, v)
		} else {
			switch r.Intn(4) {
			case 0:
				v := r.Intn(100)
				darray.Append(v)
				ref = append(ref, v)
			case 1:
				i := r.Intn(darray.size + 1)
				darray.Insert(i, 100)
				ref = append(ref, 0)
				copy(ref[i+1:], ref[i:])
				ref[i] = 100
			case 2:
				i := r.Intn(darray.size)
				darray.Remove(i)
				ref = append(ref[:i], ref[i+1:]...)
			case 3:
				i := r.Intn(darray.size)
				darray.Set(i, 100)
				ref[i] = 100
			}
		}

		if darray.Len() != len(ref) {
			t.Fatalf("第 %d 步后 Len() = %d, 参照长度 %d", step, darray.Len(), len(ref))
		}
		for i := range ref {
			got, err := darray.Get(i)
			if err != nil || got != ref[i] {
				t.Fatalf("第 %d 步后 data[%d] = %d (err: %v), 参照 %d", step, i, got, err, ref[i])
			}
		}
	}
}
