package components

type GroupFocusCursor struct {
	Group int
	Index int
}

func (fc *GroupFocusCursor) Move(group, index int) {
	fc.Group = group
	fc.Index = index
}
