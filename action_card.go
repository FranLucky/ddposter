package ddposter

type Btn struct {
	Title string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type ActionCard struct{
	Title string `json:"title"` // 首屏会话透出的展示内容
	Text string `json:"text"` // markdown格式的消息
	BtnOrientation string `json:"btnOrientation"` // 0-按钮竖直排列，1-按钮横向排列
	SingleTitle string `json:"singleTitle"` // 单个按钮的标题。(设置此项和singleURL后btns无效)
	SingleURL string `json:"singleURL"` // 点击singleTitle按钮触发的URL
	Btns []Btn `json:"btns"` // 按钮
}

