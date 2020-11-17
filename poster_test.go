package ddposter

import "testing"

const dingURL = "https://oapi.dingtalk.com/robot/send?access_token=f57f56d736238e2981d79ff6c7862dff593d8133ce9938de9269ecc285be3208"
const imgURL = "https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2396361575,51762536&fm=26&gp=0.jpg"
const BaiduURL = "https://www.baidu.com"
func TestSend(t *testing.T) {
	type args struct {
		body DingBody
		url  string
		at   *At
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"text", args{body: &Text{Content: "["}, url: dingURL, at: nil}, false},
		{"link", args{body: &Link{Text: "[", Title: "aa", PicUrl: imgURL, MessageUrl: BaiduURL}, url: dingURL, at: nil}, false},
		{"markdown", args{body: &Markdown{
			Title: "[",
			Text:  "### 标题",
		}, url: dingURL, at: nil}, false},
		{"wholeCard", args{body: &ActionCard{
			Title:          "[",
			Text:           "whole",
			SingleTitle:    "前往",
			SingleURL:      BaiduURL,
		}, url: dingURL, at: nil}, false},
		{"buttonActionCard", args{body: &ActionCard{
			Title:          "[",
			Text:           "buttons",
			Btns:           []Btn{
				{Title: "a", ActionURL: BaiduURL},
				{Title: "b", ActionURL: BaiduURL},
			},
		}, url: dingURL, at: nil}, false},
		{"feedcard", args{body: &FeedCard{Links: []Link{ {Text:"aaa", Title: "[", PicUrl: imgURL, MessageUrl: BaiduURL} }}, url: dingURL, at: nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Send(tt.args.body, tt.args.url, tt.args.at); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
