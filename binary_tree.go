package rbt

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) Insert(key Comparable, v interface{}) (ret bool) {
	ret = true
	node := NewBinaryTreeNode(key, v)
	cur := t.Root
	if cur == nil {
		t.Root = node
		return
	}

	var parent *BinaryTreeNode // 记录父亲节点
	var cmpResult CmpResult    // 记录每次比较结果
	for cur != nil {
		parent = cur
		cmpResult = key.Cmp(cur.Key)
		switch cmpResult {
		case CMP_EQ:
			return false
		case CMP_LESS:
			cur = cur.Right
		case CMP_MORE:
			cur = cur.Left
		}
	}

	if cmpResult == CMP_MORE {
		parent.SetLeftChild(node)
	} else {
		parent.SetRightChild(node)
	}

	return
}

// learn from https://www.bilibili.com/video/BV17P4y1h74 0:54:22
func (t *BinaryTree) Remove(v Comparable) (ret bool) {
	ret = true
	cur := t.Find(v)
	if cur == nil {
		return false
	}

	// case 1: left and right are null
	if cur.Left == nil && cur.Right == nil {
		isRoot := cur.Parent == nil
		if isRoot {
			t.Root = nil
		} else {
			parent := cur.Parent
			// cur is the left child of parent
			if cur == parent.Left {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}

		cur.Clear()
		return

	}

	// case 2: one of the children is null , and the other is not null
	if (cur.Left == nil && cur.Right != nil) || (cur.Left != nil && cur.Right == nil) {
		t.RemoveCaseTwo_b(cur)
		return
	}

	// case 3:
	t.RemoveCaseThree(cur)

	return
}

// 第二种场景，其中一个子节点是空，另外一个非空
// 第一种方法，直接用 candidateNode 执行节点替换 cur, 需要考虑 cur 是否是根节点 以及 cur是左节点还是右节点
func (t *BinaryTree) RemoveCaseTwo_a(cur *BinaryTreeNode) {
	isRoot := cur.Parent == nil
	var candidateNode *BinaryTreeNode
	if cur.Left != nil {
		candidateNode = cur.Left
	}
	if cur.Right != nil {
		candidateNode = cur.Right
	}

	if isRoot {
		t.Root = candidateNode
		candidateNode.Parent = nil
	} else {
		parent := cur.Parent
		candidateNode.Parent = parent
		if cur == parent.Left {
			parent.Left = candidateNode
		} else {
			parent.Right = candidateNode
		}
	}

	cur.Clear()
}

// 第二种场景，其中一个子节点是空，另外一个非空
// 第二种方法, 把 candidateNode 的Key、Value、Left、Right 复制给 cur
//  cur的左右节点分别指向 candidateNode的左右节点，candidateNode的左右节点的父亲节点指向 cur
//  最后只需要删除 candidateNode 就可以
func (t *BinaryTree) RemoveCaseTwo_b(cur *BinaryTreeNode) {

	var candidateNode *BinaryTreeNode
	if cur.Left != nil {
		candidateNode = cur.Left
	}
	if cur.Right != nil {
		candidateNode = cur.Right
	}

	cur.CloneValue(candidateNode)
	cur.SetLeftChild(candidateNode.Left)
	cur.SetRightChild(candidateNode.Right)

	cur.Clear()
}

// 第三种场景，两个节点都是非空节点
//  right := cur.Right
//  记录
// 3.1 if right.Left is null，just give right 's key and value to cur, let right.Right as cur's Right child
//     and then just clean right
// 3.1 如果 right.Left is null, 只需要把 right的 key、Value复制给 cur， 然后让 right.Right作为 cur 的右子节点
//     最后只需要清除 right 就可以了
// 3.2 if right.Left is not null, get the most left node of right , we name it leftMost. just give leftMost 's key and value to cur,
//     and let leftMost's Right node as leftMost'Parent 's Left node, and then clean leftMost
// 3.2 如果 right.Left 非空，获取right出发最左边的节点，命名为leftMost. 然后把leftMost节点的key、value复制给cur，
//     然后把leftMost的右子节点作为leftMost父节点的左节点，最后清除leftMost
func (t *BinaryTree) RemoveCaseThree(cur *BinaryTreeNode) {

	right := cur.Right
	// case: 3.1
	if right.Left == nil {
		cur.CloneValue(right)
		cur.SetRightChild(right.Right)
		right.Clear()
		return
	}

	// case:3.2
	leftMost := t.FindMostLeft(right)
	cur.CloneValue(leftMost)
	parent := leftMost.Parent
	parent.SetLeftChild(leftMost.Right)
	leftMost.Clear()

}

func (t *BinaryTree) Find(v Comparable) *BinaryTreeNode {
	cur := t.Root

	for cur != nil {
		switch v.Cmp(cur.Key) {
		case CMP_EQ:
			return cur
		case CMP_LESS:
			cur = cur.Right
		case CMP_MORE:
			cur = cur.Left
		}
	}
	return nil
}

// 查找节点 cur 左子树中最左边的节点，也就是最小节点
func (t *BinaryTree) FindMostLeft(cur *BinaryTreeNode) *BinaryTreeNode {
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

// 查找节点 cur 右子树中最右边的节点，也就是最大节点
func (t *BinaryTree) FindMostRight(cur *BinaryTreeNode) *BinaryTreeNode {
	for cur.Right != nil {
		cur = cur.Right
	}
	return cur
}
