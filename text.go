package ddposter

type Text struct {
	Content string `json:"content"` // 消息内容
}

func (t Text) TypeString() string {
	return "text"
}
