package linkedlist

import "errors"

// ErrIndexOutOfRange 下标越界时返回的错误
var ErrIndexOutOfRange = errors.New("index out of range")

// ErrNotFound 未找到目标元素时返回的错误
var ErrNotFound = errors.New("value not found")

// Node 链表节点：一个值 + 指向下一个节点的指针
type Node struct {
	Val  int
	Next *Node
}

// LinkedList 单向链表：哨兵节点 + 长度。
// 为什么用哨兵（sentinel）：真实头节点没有"前驱"，导致 Insert/Delete 都要为
// i==0 特判。哨兵是一个不存值的虚拟节点，永远站在真实链表的"下标 -1"处，
// 让"前驱"这个角色在头部也有了着落——从哨兵出发走 i 步，恰好走到下标 i-1
// 的节点（i==0 时走 0 步就是哨兵本身），头/中/尾三种情况共用同一条逻辑，
// 空表/头插/头删的特判全部消失。空表时链表只剩哨兵，sentinel.Next == nil，
// 空表也因此不用特判。
type LinkedList struct {
	sentinel *Node // 哨兵节点，sentinel.Next 才是第一个真实节点
	size     int   // 真实节点个数
}

// New 创建空链表（只有哨兵，无真实节点）
func New() *LinkedList {
	return &LinkedList{
		sentinel: &Node{},
	}
}

// Len 返回节点个数
func (l *LinkedList) Len() int {
	return l.size
}

// Insert 在下标 i 处插入 v，0 <= i <= Len()，i == Len() 等价尾部追加；越界返回错误。
// 思路：插入位置由"前驱节点"唯一确定。从哨兵出发走 i 步，落点就是前驱
// （哨兵 = 下标 -1，走 i 步 = 下标 i-1）；i==0 时走 0 步落在哨兵上，头插
// 无需特判，i==size 时落在最后一个节点上，尾插也无需特判。落点处"新节点
// 先接好后继、前驱再指向新节点"，顺序反了会丢链。
// 不变量：插入后从头遍历仍能完整走到尾，新节点恰好位于前驱与后继之间。
func (l *LinkedList) Insert(i, v int) error {
	if i < 0 || i > l.size {
		return ErrIndexOutOfRange
	}

	prev := l.sentinel
	for ; i > 0; i-- { // 用 i 本身当步数，走一步减一，省一个计数器变量
		prev = prev.Next
	}

	prev.Next = &Node{Val: v, Next: prev.Next}
	l.size++
	return nil
}

// Delete 删除下标 i 的节点并返回其值；越界返回错误。
// 思路：与 Insert 完全对称——从哨兵走 i 步到前驱，先保存被删值，再让前驱
// 的 Next 跳过被删节点。i==0 时前驱是哨兵，头删同样无需特判。
// 纠错点：保存值必须在跳过之前，顺序反了旧节点就丢了；i >= size 时被删
// 节点不存在，必须拦截（i == size 是合法插入位，但不是合法删除位）。
func (l *LinkedList) Delete(i int) (int, error) {
	if i < 0 || i >= l.size {
		return 0, ErrIndexOutOfRange
	}

	prev := l.sentinel
	for ; i > 0; i-- {
		prev = prev.Next
	}

	del := prev.Next
	val := del.Val
	prev.Next = del.Next
	l.size--
	return val, nil
}

// Search 返回第一个值为 v 的节点的下标；未找到返回错误。
// 思路：从第一个真实节点（哨兵之后）开始沿 Next 走，走到 nil 为止——
// 不会死循环的保证就是链的最后一个节点 Next 必为 nil。下标计数器与指针
// 同步前进，命中时计数器恰好就是下标（与 Traverse 的下标天然对齐）。
func (l *LinkedList) Search(v int) (int, error) {
	i := 0
	curr := l.sentinel.Next
	for curr != nil {
		if curr.Val == v {
			return i, nil
		}
		curr = curr.Next
		i++
	}
	return 0, ErrNotFound
}

// Traverse 按顺序返回所有元素。
// 为什么预分配容量：make([]int, 0, l.size) 按已知长度一次分配底层数组，
// 避免 append 反复扩容搬移（摊还上是常数，但常数也小一点）；空表时返回
// 非 nil 空切片，调用方无需判空。
func (l *LinkedList) Traverse() []int {
	res := make([]int, 0, l.size)

	curr := l.sentinel.Next
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}

	return res
}
