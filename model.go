package ddposter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type At struct {
	AtMobiles []string `json:"atMobiles"` // 被@人的手机号（在content里添加@人的手机号）
	IsAtAll bool `json:"isAtAll"` // 是否@所有人
}

type CommonModel struct {
	MsgType string `json:"msgType"`
	At *At `json:"at,omitempty"`
	FeedCard *FeedCard `json:"feedCard,omitempty"`
	ActionCard *ActionCard `json:"actionCard,omitempty"`
	Link *Link `json:"link,omitempty"`
	Markdown *Markdown `json:"markdown,omitempty"`
	Text *Text `json:"text,omitempty"`
}

func (r *CommonModel) configType() bool {
	empty := false
	if r.Text != nil {
		r.MsgType = "text"
	} else if r.ActionCard != nil {
		r.MsgType = "actionCard"
	} else if r.Link != nil {
		r.MsgType = "link"
	} else if r.FeedCard != nil {
		r.MsgType = "feedCard"
	} else if r.Markdown != nil {
		r.MsgType = "markdown"
	} else {
		empty = true
	}
	return empty
}

func (r *CommonModel)Post(dingURL string) error {
	if len(dingURL) == 0 {
		return fmt.Errorf("发送地址不能为空")
	}

	empty := r.configType()
	if len(r.MsgType) == 0 {
		return fmt.Errorf("消息类型不能为空")
	}
	if empty {
		return fmt.Errorf("消息类型对应的内容不能为空")
	}
	bc, err := json.Marshal(r)
	if err != nil {
		return err
	}

	resp, err := http.Post(dingURL, "application/json;charset=utf-8", bytes.NewReader(bc))
	if err != nil {
		return err
	}
	out, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(out))
	return nil
}