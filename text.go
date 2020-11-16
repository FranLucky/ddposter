package ddposter

type Text struct {
	Content string `json:"content"` // 消息内容
}

func (t *Text) Post(url string, at *At) error {
	return post(t, "text", url, at)
}
