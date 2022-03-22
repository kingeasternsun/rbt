package rbt

import (
	"github.com/smartystreets/assertions"
)

// https://www.bilibili.com/video/BV1oq4y1d7jv?spm_id_from=333.999.0.0  0:35:0 双黑节点

func (t *RedBlackTree) Remove(v Key) (ret bool) {
	ret = true
	cur := t.Find(v)
	if cur == nil {
		return false
	}

	for {
		// case 1
		if cur.Left == nil && cur.Right == nil {
			t.removeCaseOne(cur)
			return
		}

		// case 2
		if (cur.Left == nil && cur.Right != nil) || (cur.Left != nil && cur.Right == nil) {
			t.removeCaseTwo(cur)
			return
		}

		// case 3
		cur = t.removeCaseThree(cur)
	}

}

func (t *RedBlackTree) removeCaseOne(cur *RedBlackNode) {

	parent := cur.Parent
	// case 1.1  just remove
	if cur.GetColor() == RED {

		if cur == parent.Left {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		cur.Clear()
		return
	}

	// case 1.2  double black
	t.doubleBlackFix(cur)

	if cur.Parent == nil {
		t.Root = nil
	} else if cur == parent.Left {
		parent.Left = nil
	} else {
		parent.Right = nil
	}

	cur.Clear()
}

// case 2 : one of the children is null, the other must be red,  and the cur must be black
func (t *RedBlackTree) removeCaseTwo(cur *RedBlackNode) {

	assertions.ShouldEqual(cur.GetColor(), BLACK)
	candicate := cur.Left
	if cur.Right != nil {
		candicate = cur.Right
	}

	assertions.ShouldEqual(candicate.GetColor(), RED)

	// just give key,value of candicate to cur , leave color unchanged , and delete candicate
	cur.AcceptKeyValue(candicate)
	cur.Left = nil
	cur.Right = nil

	candicate.Clear()

}

// case 3: all of the children are not null
func (t *RedBlackTree) removeCaseThree(cur *RedBlackNode) *RedBlackNode {
	rightNode := cur.Right
	//found the leftMost of rightNode, then give key value of leftMost to cur, finally delete leftMost

	leftMost := t.FindMostLeft(rightNode)
	cur.AcceptKeyValue(leftMost)
	return leftMost // to case 1 or case 2
}

// double black,  we delete node in the end, so we need not to consider the case that the double black is null
// https://www.bilibili.com/video/BV1oq4y1d7jv 0:09:42
// attention: the nil node is considered to BLACK too   https://www.bilibili.com/video/BV19L411G72Y 1:20:30
// We first do double black fix then delete the node , by this method we avoid the situation that the double black is null
func (t *RedBlackTree) doubleBlackFix(db *RedBlackNode) {
	for db.Parent != nil {
		p, s, n, f := t.getPSNF(db)
		assertions.ShouldNotBeNil(p)
		assertions.ShouldNotBeNil(s)

		// first consider the case that n and f all black
		if n.IsBlack() && f.IsBlack() {
			// case : 0111
			if p.GetColor() == RED {
				assertions.ShouldEqual(s.GetColor(), BLACK)

				// db give one BLACK to p , and set color of s to RED
				p.SetColor(BLACK)
				s.SetColor(RED)
				return
			}

			// case : 1111
			if p.IsBlack() && s.IsBlack() {
				// db give on BLACK to p , and set color of s to RED,   p become double BLACK
				s.SetColor(RED)
				db = p
				continue
			}

			/* case : 1011
			  P				S
			 / \		   / \
			DB  S 	->    P   F
			   /  \      / \
			  N    F    DB  N
			*/
			if p.IsBlack() && s.IsRED() {

				// this is a programming trick, we exchange the color of p with s firstly
				// because after Rotate(), the pointer of p,s may be changed to different pointer
				p.SetColor(RED)
				s.SetColor(BLACK)

				t.Rotate(f, s, p)
				continue
			}

		}

		/* case : y101
		  P				N
		 / \		   / \
		DB  S 	->    P   S
		   /  \      /     \
		  N    F    DB      F
		*/
		if n.IsRED() {
			pColor := p.GetColor()
			root, left, right := t.Rotate(n, s, p)
			root.SetColor(pColor)
			left.SetColor(BLACK)
			right.SetColor(BLACK)
			return
		}

		/* case : y110
		  P				S
		 / \		   / \
		DB  S 	->    P   F
		   /  \      / \
		  N    F    DB  N
		*/
		pColor := p.GetColor()
		root, left, right := t.Rotate(f, s, p)
		root.SetColor(pColor)
		left.SetColor(BLACK)
		right.SetColor(BLACK)
		return

	}

}

// https://www.bilibili.com/video/BV19L411G72Y 1:25:16  n or f may be null
func (t *RedBlackTree) getPSNF(db *RedBlackNode) (p, s, n, f *RedBlackNode) {
	p = db.Parent
	if db == p.Left {
		s = p.Right
		n = s.Left  // the node near to db
		f = s.Right // the node far from db
	} else {
		s = p.Left
		n = s.Right // the node near to db
		f = s.Left  // the node far from db
	}
	return
}
