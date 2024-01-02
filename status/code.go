package status

type Codes struct {
	SUCCESS uint
	FAILED  uint
	Message map[uint]string
}

var ApiCode = &Codes{
	SUCCESS: 200,
	FAILED:  0,
}

func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS: "成功",
		ApiCode.FAILED:  "失败",
	}
}

func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
