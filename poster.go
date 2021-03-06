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

type DingBody interface {
	TypeString() string
}

type dDPoster struct {
	MsgType string         `json:"msgtype"`
	At *At                 `json:"at,omitempty"`
	FeedCard FeedCard     `json:"feedCard,omitempty"`
	ActionCard ActionCard `json:"actionCard,omitempty"`
	Link Link             `json:"link,omitempty"`
	Markdown Markdown     `json:"markdown,omitempty"`
	Text Text             `json:"text,omitempty"`
}

func Send(body DingBody, url string, at *At) error {
	if len(url) == 0 {
		return fmt.Errorf("发送地址不能为空")
	}
	m := new(dDPoster)
	m.At = at
	if t, ok := body.(Text); ok {
		m.MsgType = t.TypeString()
		m.Text = t
	} else if t, ok := body.(Link); ok {
		m.MsgType = t.TypeString()
		m.Link = t
	} else if t, ok := body.(Markdown); ok {
		m.MsgType = t.TypeString()
		m.Markdown = t
	} else if t, ok := body.(FeedCard); ok {
		m.MsgType = t.TypeString()
		m.FeedCard = t
	} else if t, ok := body.(ActionCard); ok {
		m.MsgType = t.TypeString()
		m.ActionCard = t
	}
	if len(m.MsgType) == 0 {
		return fmt.Errorf("内容不能为空")
	}
	bc, err := json.Marshal(m)
	if err != nil {
		return err
	}
	fmt.Println(string(bc))
	resp, err := http.Post(url, "application/json;charset=utf-8", bytes.NewReader(bc))
	if err != nil {
		return err
	}
	out, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(out))
	return nil
}
