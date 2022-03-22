package rbt

type Key struct {
	Key int
}

func (k Key) Cmp(c Key) CmpResult {
	if k.Key == c.Key {
		return CMP_EQ
	}

	if k.Key > c.Key {
		return CMP_MORE
	}
	return CMP_LESS
}
