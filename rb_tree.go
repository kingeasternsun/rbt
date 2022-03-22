package rbt

import "github.com/sirupsen/logrus"

type RedBlackTree struct {
	Root *RedBlackNode
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (t *RedBlackTree) binaryInsert(key Key, v interface{}) *RedBlackNode {

	node := NewRedBlackNode(key, v)
	cur := t.Root
	if cur == nil {
		t.Root = node
		return node
	}

	var parent *RedBlackNode // 记录父亲节点
	var cmpResult CmpResult  // 记录每次比较结果
	for cur != nil {
		parent = cur
		cmpResult = key.Cmp(cur.Key)
		switch cmpResult {
		case CMP_EQ:
			cur = cur.Right
		case CMP_MORE:
			cur = cur.Right
		case CMP_LESS:
			cur = cur.Left
		}
	}

	if cmpResult == CMP_LESS {
		parent.SetLeftChild(node)
	} else {
		parent.SetRightChild(node)
	}

	return node
}

func (t *RedBlackTree) Insert(key Key, v interface{}) {
	p := t.binaryInsert(key, v)
	var left, right *RedBlackNode

	for {
		if p.Parent == nil {
			p.SetColor(BLACK)
			return
		}

		if p.Parent.Color == BLACK {
			return
		}

		p, left, right = t.Rotate(p, p.Parent, p.Parent.Parent)

		p.SetColor(RED)
		left.SetColor(BLACK)
		right.SetColor(BLACK)
	}

}

// 旋转 https://www.bilibili.com/video/BV1Sr4y117nP 0:32:07 讲解  0:51:53代码
//
func (t *RedBlackTree) Rotate(n, p, g *RedBlackNode) (root *RedBlackNode, left *RedBlackNode, right *RedBlackNode) {
	if n == p.Left && p == g.Left {
		return t.RotateLeftSlash_c(n, p, g)
	}

	if n == p.Right && p == g.Right {
		return t.RotateRightSlash_c(n, p, g)
	}

	if n == p.Right && p == g.Left {
		return t.RotateLeftCurve_c(n, p, g)
	}

	return t.RotateRightCurve_c(n, p, g)
}

/*
    G         P
  P    ->   N   G
N

pull up P node directly
直接提升P节点
*/
func (t *RedBlackTree) RotateLeftSlash_a(n, p, g *RedBlackNode) (root *RedBlackNode, left *RedBlackNode, right *RedBlackNode) {
	gp := g.Parent
	pRight := p.Right

	gIsLeft := false

	if gp != nil && g == gp.Left {
		gIsLeft = true
	}

	// finish pull up node p
	p.SetRightChild(g)
	g.SetLeftChild(pRight)

	if gp == nil {
		p.Parent = nil
		t.Root = p
	} else if gIsLeft {
		gp.SetLeftChild(p)
	} else {
		gp.SetRightChild(p)
	}
	return p, n, g
}

/*
    G         P        P
  P    ->   G    ->  N   G
N         N

swap key、value and color of G P firstly, so we need not consider the parent of g
首先交换G、P两个节点的key、value、color, 就不用考虑g节点的父节点了
*/
func (t *RedBlackTree) RotateLeftSlash_b(n, p, g *RedBlackNode) (root *RedBlackNode, left *RedBlackNode, right *RedBlackNode) {

	pRight := p.Right
	gRight := g.Right // important
	// swap p g
	SwapContentOfRedBlackNode(g, p)
	g, p = p, g

	p.SetLeftChild(n)
	p.SetRightChild(g)
	g.SetLeftChild(pRight)
	g.SetRightChild(gRight) // important, because the g is not the original p

	return p, n, g
}

/*
    G         G        P
  P    ->   N   P ->  N  G
N

 we let N as P'Left child, G as P' right child , then swap content of G and P
 先把N变成P的左节点，G变成P的右节点，然后交换G和P里面的key、value、color
*/
func (t *RedBlackTree) RotateLeftSlash_c(n, p, g *RedBlackNode) (root *RedBlackNode, left *RedBlackNode, right *RedBlackNode) {

	pRight := p.Right
	gRight := g.Right // important

	g.SetLeftChild(n)
	g.SetRightChild(p)

	p.SetLeftChild(pRight)
	p.SetRightChild(gRight)

	SwapContentOfRedBlackNode(g, p)
	g, p = p, g

	// now g has the content of the p
	return p, n, g
}

/*
G               P
  P      ->  G    N
     N

pull up P node directly
直接提升P节点
*/
func (t *RedBlackTree) RotateRightSlash_a(n, p, g *RedBlackNode) (root *RedBlackNode, left *RedBlackNode, right *RedBlackNode) {
	gp := g.Parent
	pLeft := p.Left

	gIsLeft := false

	if gp != nil && g == gp.Left {
		gIsLeft = true
	}

	// finish pull up node p
	p.SetLeftChild(g)
	g.SetRightChild(pLeft)

	if gp == nil {
		p.Parent = nil
		t.Root = p
	} else if gIsLeft {
		gp.SetLeftChild(p)
	} else {
		gp.SetRightChild(p)
	}
	return p, g, n
}

/*
G         P             P
  P    ->   G     ->  G    N
     N         N

swap key、value and color of G P firstly, so we need not consider the parent of g
首先交换G、P两个节点的key、value、color, 就不用考虑g节点的父节点了
*/
func (t *RedBlackTree) RotateRightSlash_b(n, p, g *RedBlackNode) (root *RedBlackNode, left *RedBlackNode, right *RedBlackNode) {

	pLeft := p.Left
	gLeft := g.Left // important
	// swap p g
	SwapContentOfRedBlackNode(g, p)
	g, p = p, g

	p.SetLeftChild(g)
	p.SetRightChild(n)
	g.SetLeftChild(gLeft) // important, because the g is not the original p
	g.SetRightChild(pLeft)

	return p, g, n
}

/*
G             G        P
  P    ->   P   N ->  G  N
     N

 we let P as G'Left child, N as G' right child , then swap content of G and P
 先把P变成G的左节点，N变成G的右节点，然后交换G和P里面的key、value、color
*/
func (t *RedBlackTree) RotateRightSlash_c(n, p, g *RedBlackNode) (root *RedBlackNode, left *RedBlackNode, right *RedBlackNode) {

	gLeft := g.Left
	pLeft := p.Left

	g.SetLeftChild(p)
	g.SetRightChild(n)
	p.SetLeftChild(gLeft)
	p.SetRightChild(pLeft)

	SwapContentOfRedBlackNode(g, p)
	g, p = p, g

	return p, g, n
}

/*
	G             N
  P      ->    P     G
	N

pull up N node directly, should consider the parent of node G
直接提升N节点,需要考虑G的父节点
*/
func (t *RedBlackTree) RotateLeftCurve_a(n, p, g *RedBlackNode) (root *RedBlackNode, left *RedBlackNode, right *RedBlackNode) {
	gp := g.Parent
	gIsLeft := false
	nLeft := n.Left
	nRight := n.Right

	if gp != nil && g == gp.Left {
		gIsLeft = true
	}

	n.SetLeftChild(p)
	n.SetRightChild(g)
	p.SetRightChild(nLeft)
	g.SetLeftChild(nRight)

	if gp == nil {
		t.Root = n
		n.Parent = nil
	} else if gIsLeft {
		gp.SetLeftChild(n)
	} else {
		gp.SetRightChild(n)
	}

	return n, p, g
}

/*
	G             N			N
  P      ->    P      ->  P   G
	N			  G

first swap content of node N and G, then set G as right child of node N
首先交换节点N和G的key、value、color， 然后把节点G变为节点N的右节点

*/

func (t *RedBlackTree) RotateLeftCurve_b(n, p, g *RedBlackNode) (root *RedBlackNode, left *RedBlackNode, right *RedBlackNode) {
	nLeft := n.Left
	nRight := n.Right
	gRight := g.Right

	SwapContentOfRedBlackNode(g, n)
	g, n = n, g // by doing this , the code would be more readable

	n.SetRightChild(g)
	p.SetRightChild(nLeft)
	g.SetLeftChild(nRight)
	g.SetRightChild(gRight)

	return n, p, g
}

/*
	G             G			 N
  P      ->    P     N ->  P   G
	N

first set N as right child of node G, then swap key,value,color of node G and G
首先把节点N变为节点G的右节点，然后交换节点N和G的key、value、color
*/

func (t *RedBlackTree) RotateLeftCurve_c(n, p, g *RedBlackNode) (root *RedBlackNode, left *RedBlackNode, right *RedBlackNode) {
	gRight := g.Right
	nLeft := n.Left
	nRight := n.Right

	g.SetRightChild(n)
	p.SetRightChild(nLeft)
	n.SetLeftChild(nRight)
	n.SetRightChild(gRight)

	SwapContentOfRedBlackNode(g, n)
	g, n = n, g

	return n, p, g
}

/*
	G             N
       P ->    G     P
	N

pull up N node directly, should consider the parent of node G
直接提升N节点,需要考虑G的父节点
*/
func (t *RedBlackTree) RotateRightCurve_a(n, p, g *RedBlackNode) (root *RedBlackNode, left *RedBlackNode, right *RedBlackNode) {
	gp := g.Parent
	gIsLeft := false
	nLeft := n.Left
	nRight := n.Right

	if gp != nil && g == gp.Left {
		gIsLeft = true
	}

	n.SetLeftChild(g)
	n.SetRightChild(p)
	g.SetRightChild(nLeft)
	p.SetLeftChild(nRight)

	if gp == nil {
		t.Root = n
		n.Parent = nil
	} else if gIsLeft {
		gp.SetLeftChild(n)
	} else {
		gp.SetRightChild(n)
	}

	return n, g, p
}

/*
	G           N				 N
       P ->        P     ->   G     P
	N           G

first swap content of node N and G, then set G as right child of node N
首先交换节点N和G的key、value、color， 然后把节点G变为节点N的左节点
*/
func (t *RedBlackTree) RotateRightCurve_b(n, p, g *RedBlackNode) (root *RedBlackNode, left *RedBlackNode, right *RedBlackNode) {
	nLeft := n.Left
	nRight := n.Right
	gLeft := g.Left

	SwapContentOfRedBlackNode(g, n)
	g, n = n, g // by doing this , the code would be more readable

	n.SetLeftChild(g)
	g.SetLeftChild(gLeft)
	g.SetRightChild(nLeft)
	p.SetLeftChild(nRight)

	return n, g, p
}

/*
	G           G				 N
       P ->  N     P     ->   G     P
	N

first set N as left child of node G, then swap key,value,color of node G and G
首先把节点N变为节点G的左节点，然后交换节点N和G的key、value、color
*/
func (t *RedBlackTree) RotateRightCurve_c(n, p, g *RedBlackNode) (root *RedBlackNode, left *RedBlackNode, right *RedBlackNode) {
	nLeft := n.Left
	nRight := n.Right
	gLeft := g.Left

	g.SetLeftChild(n)
	n.SetLeftChild(gLeft)
	n.SetRightChild(nLeft)
	p.SetLeftChild(nRight)

	SwapContentOfRedBlackNode(g, n)
	g, n = n, g

	return n, g, p
}

func (t *RedBlackTree) Find(v Key) *RedBlackNode {
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
func (t *RedBlackTree) FindMostLeft(cur *RedBlackNode) *RedBlackNode {
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

// 查找节点 cur 右子树中最右边的节点，也就是最大节点
func (t *RedBlackTree) FindMostRight(cur *RedBlackNode) *RedBlackNode {
	for cur.Right != nil {
		cur = cur.Right
	}
	return cur
}

func (t *RedBlackTree) Check() (height int, ret bool) {

	h, ret := Check(t.Root)
	logrus.Infof(" %d %v", h, ret)
	return h, ret
}

func Check(root *RedBlackNode) (height int, ret bool) {
	if root == nil {
		return 1, true
	}

	leftHeight, leftRet := Check(root.Left)
	if !leftRet {
		return leftHeight, false
	}

	rightHeight, rightRet := Check(root.Right)
	if !rightRet {
		return rightHeight, false
	}

	if root.Color == RED {
		if root.Left != nil && root.Left.Color != BLACK {
			return leftHeight, false
		}
		if root.Right != nil && root.Right.Color != BLACK {
			return leftHeight, false
		}
	}

	if leftHeight != rightHeight {
		logrus.Errorf("Check(%d) %d", leftHeight, rightHeight)
		return leftHeight, false
	}

	if root.GetColor() == RED {
		return leftHeight, true
	}
	return leftHeight + 1, true
}
