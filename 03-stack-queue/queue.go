package stackqueue

import "errors"

// ErrQueueEmpty 空队列上执行 Dequeue/Peek 时返回的错误
var ErrQueueEmpty = errors.New("queue is empty")

// ErrQueueFull 满队列上执行 Enqueue 时返回的错误
var ErrQueueFull = errors.New("queue is full")

// Queue 队列：先进先出（FIFO），环形数组实现。
// 为什么环形：若用普通数组，出队把后面的元素整体前移是 O(n)；
// 环形让 front/rear 两个下标沿数组"绕圈"，出队只动指针不搬元素，O(1)。
//
// 关键设计：用 size 字段直接记录元素个数，用 front 指向队首元素、
// rear 指向下一个写入位置。移动指针时必须对容量取模，绕到末尾跳回开头。
//
// 纠错点：
//  1. front/rear 前进要取模，忘了就数组越界；
//  2. "空"和"满"两种状态都要有出口——空时 Dequeue/Peek 返回
//     ErrQueueEmpty，满时 Enqueue 返回 ErrQueueFull，不能静默吞掉；
//  3. front 与 rear 相撞并不总是"满"，空时两者也相等，
//     区分靠 size 而不是下标比较。
type Queue struct {
	data  []int // 环形存储区，长度即容量
	front int   // 队首元素下标
	rear  int   // 下一个写入位置下标
	size  int   // 当前元素个数
}

// NewQueue 创建容量为 capacity 的空队列。capacity <= 0 时视为非法参数。
func NewQueue(capacity int) *Queue {
	panic("TODO")
}

// Enqueue 入队：x 写入 rear 位置，rear 前进一步（取模）。
// 满队列返回 ErrQueueFull 且队列保持不变。
// 自查点：入队后 size 加一；连续入队到满为止，一个不丢。
func (q *Queue) Enqueue(x int) error {
	panic("TODO")
}

// Dequeue 出队：返回 front 位置的元素，front 前进一步（取模）。
// 空队列返回 ErrQueueEmpty。
// 自查点：出队顺序与入队顺序完全一致；size 减一。
func (q *Queue) Dequeue() (int, error) {
	panic("TODO")
}

// Peek 看队首：返回 front 位置的元素但不移除。空队列返回 ErrQueueEmpty。
func (q *Queue) Peek() (int, error) {
	panic("TODO")
}

// IsEmpty 判断队列是否为空（size == 0）
func (q *Queue) IsEmpty() bool {
	panic("TODO")
}

// Len 返回当前元素个数
func (q *Queue) Len() int {
	panic("TODO")
}

// Cap 返回队列容量。测试用它验证"满"状态。
func (q *Queue) Cap() int {
	panic("TODO")
}
