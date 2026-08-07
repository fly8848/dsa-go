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

// LinkedList 单向链表：只保留头指针与长度。
// 设计思路：节点不要求连续内存，靠 Next 串成链；代价是按下标访问必须从头
// 走，访问第 i 个节点是 O(i)——与数组"随机访问 O(1)、中间插入 O(n)"正好互补。
// 头指针有两种方案：直接指向第一个节点，或指向哨兵（dummy）节点。想清楚
// 空表、删头节点时各自怎么处理，选一种并保持一致。
type LinkedList struct {
	head *Node
	size int
}

// New 创建空链表
func New() *LinkedList {
	return &LinkedList{
		head: nil,
		size: 0,
	}
}

// Len 返回节点个数
func (l *LinkedList) Len() int {
	return l.size
}

// Insert 在下标 i 处插入 v，0 <= i <= Len()，i == Len() 等价尾部追加；越界返回错误。
// 思路：插入位置由"前驱节点"唯一确定——找到前驱，让新节点先接好后继，
// 再让前驱指向新节点；接的顺序错了会丢链。
// 不变量：插入后从头遍历仍能完整走到尾，新节点恰好位于前驱与后继之间。
// 自查：在头、中、尾各插一次，遍历输出顺序符合预期。
func (l *LinkedList) Insert(i, v int) error {
	if i < 0 || i > l.size {
		return ErrIndexOutOfRange
	}

	if i == 0 {
		// 头
		l.head = &Node{
			Val:  v,
			Next: l.head,
		}
	} else {
		// 中 尾
		j := 0
		curr := l.head
		for {
			if i-1 == j {
				curr.Next = &Node{
					Val:  v,
					Next: curr.Next,
				}
				break
			}

			curr = curr.Next
			j++
		}
	}

	l.size++
	return nil
}

// Delete 删除下标 i 的节点并返回其值；越界返回错误。
// 思路：删除同样只关心前驱——让前驱的 Next 跳过被删节点即可。摘除后
// 被删节点从链表不可达，其 Next 无需置空（GC 会自行回收）；只有未来 API
// 对外暴露节点时才需要置空防御。
// 纠错点：空表删除、删除唯一节点后链表应为空、删头节点时前驱是谁？
// 自查：删头、删尾、删中间各一次，Len 与遍历结果同步正确。
func (l *LinkedList) Delete(i int) (int, error) {
	if i < 0 || i >= l.size || l.head == nil {
		return 0, ErrIndexOutOfRange
	}

	if i == 0 {
		val := l.head.Val
		l.head = l.head.Next
		l.size--
		return val, nil
	}

	j := 0
	curr := l.head
	pre := l.head
	for {
		if i == j {
			pre.Next = curr.Next
			l.size--
			return curr.Val, nil
		}

		pre = curr
		curr = curr.Next
		j++
	}
}

// Search 返回第一个值为 v 的节点的下标；未找到返回错误。
// 思路：从头沿 Next 走，走到 nil 为止——不会死循环的保证就是链的最后
// 一个节点 Next 必为 nil，这就是"链必须有终点"。
func (l *LinkedList) Search(v int) (int, error) {
	i := 0
	curr := l.head
	for {
		if curr == nil {
			return 0, ErrNotFound
		}

		if curr.Val == v {
			return i, nil
		} else {
			curr = curr.Next
		}

		i++
	}
}

// Traverse 按顺序返回所有元素。
// 自查：与 Insert 顺序一一对应；空表返回空切片。
func (l *LinkedList) Traverse() []int {
	res := make([]int, 0)

	curr := l.head
	for {
		if curr == nil {
			break
		}
		res = append(res, curr.Val)
		curr = curr.Next
	}

	return res
}
