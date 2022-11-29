package testing

import (
	dps "github.com/xiiiew/dingtalk-push-golang-sdk"
	"strconv"
	"sync"
	"testing"
	"time"
)

// 测试发送text
func TestSendText(t *testing.T) {
	endpoint := "http://127.0.0.1:8080"
	secret := "SEC277e2bf7f71ce33ed3ef83047ccf5e042f6a9ac4689367a4658734e99c22e385"
	accessToken := "7c58eafd7663d1b784062a58069538f8a9ed3daba75333cb558d5290fbf2c1a5"
	dpsClient, err := dps.NewDpsClient(endpoint, secret, accessToken)
	if err != nil {
		t.Error(err)
	}
	msg := &dps.TextConfig{
		Text: dps.TextFieldText{
			"sdk测试",
		},
		At: dps.TextFieldAt{
			AtMobiles: nil,
			AtUserIds: nil,
			IsAtAll:   false,
		},
	}
	err = dpsClient.Send(msg)
	if err != nil {
		t.Error(err)
	}
}

// 测试发送link
func TestSendLink(t *testing.T) {
	endpoint := "http://127.0.0.1:8080"
	secret := "SEC277e2bf7f71ce33ed3ef83047ccf5e042f6a9ac4689367a4658734e99c22e385"
	accessToken := "7c58eafd7663d1b784062a58069538f8a9ed3daba75333cb558d5290fbf2c1a5"
	dpsClient, err := dps.NewDpsClient(endpoint, secret, accessToken)
	if err != nil {
		t.Error(err)
	}
	msg := &dps.LinkConfig{
		Link: dps.LinkFieldLink{
			Text:       "text",
			Title:      "title",
			PicUrl:     "www.baidu.com",
			MessageUrl: "www.baidu.com",
		},
	}

	err = dpsClient.Send(msg)
	if err != nil {
		t.Error(err)
	}
}

// 测试发送markdown
func TestSendMarkdown(t *testing.T) {
	endpoint := "http://127.0.0.1:8080"
	secret := "SEC277e2bf7f71ce33ed3ef83047ccf5e042f6a9ac4689367a4658734e99c22e385"
	accessToken := "7c58eafd7663d1b784062a58069538f8a9ed3daba75333cb558d5290fbf2c1a5"
	dpsClient, err := dps.NewDpsClient(endpoint, secret, accessToken)
	if err != nil {
		t.Error(err)
	}
	msg := &dps.MarkdownConfig{
		Markdown: dps.MarkdownFieldMarkdown{
			Title: "title",
			Text: `
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题

引用
> A man who stands for nothing will fall for anything.

文字加粗、斜体
**bold**
*italic*

链接
[this is a link](http://name.com)

图片
![](http://name.com/pic.jpg)

无序列表
- item1
- item2

有序列表
1. item1
2. item2
`,
		},
		At: dps.MarkdownFieldAt{},
	}

	err = dpsClient.Send(msg)
	if err != nil {
		t.Error(err)
	}
}

// 测试发送actionCard
func TestSendActionCard(t *testing.T) {
	endpoint := "http://127.0.0.1:8080"
	secret := "SEC277e2bf7f71ce33ed3ef83047ccf5e042f6a9ac4689367a4658734e99c22e385"
	accessToken := "7c58eafd7663d1b784062a58069538f8a9ed3daba75333cb558d5290fbf2c1a5"
	dpsClient, err := dps.NewDpsClient(endpoint, secret, accessToken)
	if err != nil {
		t.Error(err)
	}
	msg := &dps.ActionCardConfig{
		ActionCard: dps.ActionCardFieldActionCard{
			Title:          "title",
			Text:           "text",
			SingleTitle:    "",
			SingleUrl:      "",
			BtnOrientation: dps.BtnOrientationH,
			Btns: []dps.ActionCardFieldBtns{
				{
					Title:     "同意",
					ActionUrl: "#",
				},
				{
					Title:     "拒绝",
					ActionUrl: "#",
				},
			},
		},
	}

	err = dpsClient.Send(msg)
	if err != nil {
		t.Error(err)
	}
}

// 测试发送feedCard
func TestSendFeedCard(t *testing.T) {
	endpoint := "http://127.0.0.1:8080"
	secret := "SEC277e2bf7f71ce33ed3ef83047ccf5e042f6a9ac4689367a4658734e99c22e385"
	accessToken := "7c58eafd7663d1b784062a58069538f8a9ed3daba75333cb558d5290fbf2c1a5"
	dpsClient, err := dps.NewDpsClient(endpoint, secret, accessToken)
	if err != nil {
		t.Error(err)
	}
	msg := &dps.FeedCardConfig{
		FeedCard: dps.FeedCardFieldFeedCard{
			Links: []dps.FeedFieldLinks{
				{
					Title:      "title",
					MessageUrl: "#",
					PicUrl:     "#",
				},
				{
					Title:      "title",
					MessageUrl: "#",
					PicUrl:     "#",
				},
			},
		},
	}

	err = dpsClient.Send(msg)
	if err != nil {
		t.Error(err)
	}
}

// 测试两个机器人同时发text消息，且消息同时发出
func TestTwoBot(t *testing.T) {
	endpoint := "http://127.0.0.1:8080"
	secret := "SEC277e2bf7f71ce33ed3ef83047ccf5e042f6a9ac4689367a4658734e99c22e385"
	accessToken := "7c58eafd7663d1b784062a58069538f8a9ed3daba75333cb558d5290fbf2c1a5"

	secret2 := "SEC32e32b600250850196f67e49c3b7728df59b2fcef58c0159fe5154bd82a801c8"
	accessToken2 := "76d4769bf60a9e218221f08f9d2a50147cd03b09a9d8d949b1e8108eb04dbfea"
	dpsClient, err := dps.NewDpsClient(endpoint, secret, accessToken)
	dpsClient2, err := dps.NewDpsClient(endpoint, secret2, accessToken2)
	if err != nil {
		t.Error(err)
	}
	msg := &dps.TextConfig{
		Text: dps.TextFieldText{
			"sdk测试",
		},
		At: dps.TextFieldAt{},
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		err = dpsClient.Send(msg)
		if err != nil {
			t.Error(err)
		}
	}()
	go func() {
		defer wg.Done()
		err = dpsClient2.Send(msg)
		if err != nil {
			t.Error(err)
		}
	}()
	wg.Wait()
}

// 测试两个机器人同时发多条可合并的消息，且消息合并
func TestTwoBotMulti(t *testing.T) {
	endpoint := "http://127.0.0.1:8080"
	secret := "SEC277e2bf7f71ce33ed3ef83047ccf5e042f6a9ac4689367a4658734e99c22e385"
	accessToken := "7c58eafd7663d1b784062a58069538f8a9ed3daba75333cb558d5290fbf2c1a5"

	secret2 := "SEC32e32b600250850196f67e49c3b7728df59b2fcef58c0159fe5154bd82a801c8"
	accessToken2 := "76d4769bf60a9e218221f08f9d2a50147cd03b09a9d8d949b1e8108eb04dbfea"
	dpsClient, err := dps.NewDpsClient(endpoint, secret, accessToken)
	dpsClient2, err := dps.NewDpsClient(endpoint, secret2, accessToken2)
	if err != nil {
		t.Error(err)
	}
	msg := &dps.TextConfig{
		Text: dps.TextFieldText{
			"sdk测试",
		},
		At: dps.TextFieldAt{},
	}

	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		err = dpsClient.Send(msg)
		if err != nil {
			t.Error(err)
		}
	}()
	go func() {
		defer wg.Done()
		err = dpsClient.Send(msg)
		if err != nil {
			t.Error(err)
		}
	}()
	go func() {
		defer wg.Done()
		err = dpsClient2.Send(msg)
		if err != nil {
			t.Error(err)
		}
	}()
	go func() {
		defer wg.Done()
		err = dpsClient2.Send(msg)
		if err != nil {
			t.Error(err)
		}
	}()
	wg.Wait()
}

// 测试一个机器人同时发多条类型的消息，且各类型消息有间隔时间
func TestOneBotMulti(t *testing.T) {
	endpoint := "http://127.0.0.1:8080"
	secret := "SEC277e2bf7f71ce33ed3ef83047ccf5e042f6a9ac4689367a4658734e99c22e385"
	accessToken := "7c58eafd7663d1b784062a58069538f8a9ed3daba75333cb558d5290fbf2c1a5"

	dpsClient, err := dps.NewDpsClient(endpoint, secret, accessToken)
	if err != nil {
		t.Error(err)
	}
	msg := &dps.TextConfig{
		Text: dps.TextFieldText{
			"sdk测试",
		},
		At: dps.TextFieldAt{},
	}
	msg2 := &dps.MarkdownConfig{
		Markdown: dps.MarkdownFieldMarkdown{
			Title: "title",
			Text: `
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题

引用
> A man who stands for nothing will fall for anything.

文字加粗、斜体
**bold**
*italic*

链接
[this is a link](http://name.com)

图片
![](http://name.com/pic.jpg)

无序列表
- item1
- item2

有序列表
1. item1
2. item2
`,
		},
		At: dps.MarkdownFieldAt{},
	}

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		err = dpsClient.Send(msg)
		if err != nil {
			t.Error(err)
		}
	}()
	go func() {
		defer wg.Done()
		err = dpsClient.Send(msg)
		if err != nil {
			t.Error(err)
		}
	}()
	go func() {
		defer wg.Done()
		err = dpsClient.Send(msg2)
		if err != nil {
			t.Error(err)
		}
	}()
	wg.Wait()
}

// 测试循环发送markdown消息
func TestSendMarkdownLoop(t *testing.T) {
	endpoint := "http://127.0.0.1:8080"
	secret := "SEC277e2bf7f71ce33ed3ef83047ccf5e042f6a9ac4689367a4658734e99c22e385"
	accessToken := "7c58eafd7663d1b784062a58069538f8a9ed3daba75333cb558d5290fbf2c1a5"
	dpsClient, err := dps.NewDpsClient(endpoint, secret, accessToken)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 60; i++ {
		msg := &dps.MarkdownConfig{
			Markdown: dps.MarkdownFieldMarkdown{
				Title: "title",
				Text: `
# ` + strconv.Itoa(i) + `
`,
			},
			At: dps.MarkdownFieldAt{
				AtMobiles: []string{"130xxxxxxxx"},
			},
		}

		err = dpsClient.Send(msg)
		if err != nil {
			t.Error(err)
		}
		time.Sleep(time.Second)
	}

}
