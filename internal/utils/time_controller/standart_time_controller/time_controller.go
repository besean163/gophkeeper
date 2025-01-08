package standarttimecontroller

import "time"

type TimeController struct{}

func NewTimeController() TimeController {
	return TimeController{}
}

func (s TimeController) Now() int64 {
	return time.Now().Unix()
}
