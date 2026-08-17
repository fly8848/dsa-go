package stackqueue

import "errors"

// ErrStackEmpty 空栈上执行 Pop/Peek 时返回的错误
var ErrStackEmpty = errors.New("stack is empty")

// Stack 栈：后进先出（LIFO）。
// 为什么用切片：栈只在"同一端"进出，切片尾部天然是栈顶，
// 追加与收缩都是 O(1)，不需要下标移动，是栈的最简宿主。
// 纠错点：Pop/Peek 前必须判空，空栈返回 ErrStackEmpty，
// 而不是返回零值让调用方误以为"取出 0"成功。
type Stack struct {
	data []int // 切片尾部即栈顶
}

// NewStack 创建空栈
func NewStack() *Stack {
	panic("TODO")
}

// Push 压栈：x 成为新栈顶。任何情况下都成功。
func (s *Stack) Push(x int) {
	panic("TODO")
}

// Pop 弹栈：返回栈顶并移除。空栈返回 ErrStackEmpty。
// 自查点：弹出后 Len 减一；连续 Pop 的顺序恰好是 Push 的逆序。
func (s *Stack) Pop() (int, error) {
	panic("TODO")
}

// Peek 看栈顶：返回栈顶但不移除。空栈返回 ErrStackEmpty。
// 自查点：Peek 连续调用多次，结果不变，Len 不变。
func (s *Stack) Peek() (int, error) {
	panic("TODO")
}

// IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	panic("TODO")
}

// Len 返回栈内元素个数
func (s *Stack) Len() int {
	panic("TODO")
}
