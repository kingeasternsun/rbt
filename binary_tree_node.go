package rbt

type CmpResult int

const (
	CMP_EQ   CmpResult = 0
	CMP_LESS CmpResult = -1
	CMP_MORE CmpResult = 1
)

type Comparable interface {
	Cmp(c Comparable) CmpResult
}

type BinaryTreeNode struct {
	Key         Comparable
	Value       interface{}
	Parent      *BinaryTreeNode
	Left, Right *BinaryTreeNode
}

func NewBinaryTreeNode(key Comparable, value interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{Key: key, Value: value}
}

func (n *BinaryTreeNode) SetLeftChild(left *BinaryTreeNode) {
	n.Left = left
	if left.Value != nil {
		left.Parent = n
	}
}

func (n *BinaryTreeNode) SetRightChild(right *BinaryTreeNode) {
	n.Right = right
	if right != nil {
		right.Parent = n
	}

}

func (n *BinaryTreeNode) Clear() {
	n.Left = nil
	n.Right = nil
	n.Parent = nil
	n.Value = nil
	n.Key = nil
}

func (n *BinaryTreeNode) CloneValue(v *BinaryTreeNode) {
	n.Key = v.Key
	n.Value = v.Value
}
