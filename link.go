package ddposter

type Link struct {
	Text string `json:"text"` // 消息内容。如果太长只会部分展示 在feedCard里无此项
	Title string `json:"title"` // 消息标题
	PicUrl string `json:"picUrl"` // 图片URL
	MessageUrl string `json:"messageUrl"` // 点击消息跳转的URL
}

func (l *Link) Post(url string, at *At) error {
	return post(l, "link", url, at)
}
