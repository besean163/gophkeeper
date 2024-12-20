package components

type GroupFocusCursor struct {
	Group int
	Index int
}

func NewGroupFocusCursor(group, index int) *GroupFocusCursor {
	return &GroupFocusCursor{
		Group: group,
		Index: index,
	}
}

func (fc *GroupFocusCursor) Move(group, index int) {
	fc.Group = group
	fc.Index = index
}
