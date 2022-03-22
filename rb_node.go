/*
红黑树的定义与性质 https://www.bilibili.com/video/BV1nL41147PY, 其中 0:56:42 开始讲解和2-3-4树的关系
*/
package rbt

type Color int

const (
	RED   Color = 1
	BLACK Color = 2
)

type RedBlackNode struct {
	Key         Key
	Value       interface{}
	Color       Color
	Parent      *RedBlackNode
	Left, Right *RedBlackNode
}

func NewRedBlackNode(key Key, value interface{}) *RedBlackNode {
	return &RedBlackNode{Key: key, Value: value, Color: RED}
}

func (n *RedBlackNode) SetLeftChild(left *RedBlackNode) {
	n.Left = left
	if left != nil {
		left.Parent = n
	}
}

func (n *RedBlackNode) SetRightChild(right *RedBlackNode) {
	n.Right = right
	if right != nil {
		right.Parent = n
	}
}

func (n *RedBlackNode) Clear() {
	n.Left = nil
	n.Right = nil
	n.Parent = nil
	n.Value = nil
	// n.Key = nil
}

func (n *RedBlackNode) CloneValue(v *RedBlackNode) {
	n.Key = v.Key
	n.Value = v.Value
}

func (n *RedBlackNode) SetColor(c Color) {
	n.Color = c
}

func (n *RedBlackNode) GetColor() Color {
	return n.Color
}

func (n *RedBlackNode) IsBlack() bool {
	return n == nil || n.Color == BLACK
}

func (n *RedBlackNode) IsRED() bool {
	return !n.IsBlack()
}

func (n *RedBlackNode) AcceptKeyValue(other *RedBlackNode) {
	n.Key = other.Key
	n.Value = other.Value
}

func SwapContentOfRedBlackNode(a, b *RedBlackNode) {
	a.Key, b.Key = b.Key, a.Key
	a.Value, b.Value = b.Value, a.Value
	a.Color, b.Color = b.Color, a.Color
}
