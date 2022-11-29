# dingtalk-push-golang-sdk

[dingtalk-push](https://github.com/xiiiew/dingtalk-push) golang SDK, 封装调用[dingtalk-push](https://github.com/xiiiew/dingtalk-push)服务的方法。

支持推送钉钉text、link、markdown、actionCard及feedCard消息类型，并且能根据每个机器人限制推送频率和合并推送消息。

#### Usage

```shell
go get github.com/xiiiew/dingtalk-push-golang-sdk
```

```go
func TestSendText(t *testing.T) {
    endpoint := "http://127.0.0.1:28080"  // dingtalk-push服务地址
    secret := "SEC××××"     // 钉钉机器人secret
    accessToken := "****"   // 钉钉机器人access_token
    // dpsClient, err := dps.NewDpsClient(endpoint, secret, accessToken)   // 使用默认http client创建dpsClient
    dpsClient, err := dps.NewDpsClientWithHTTPClient(endpoint, secret, accessToken, yourClient) // 使用自定义http client创建dpsClient
    if err != nil {
        t.Error(err)
    }
    // 定义消息类型
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
    // 发送消息
    err = dpsClient.Send(msg)
    if err != nil {
        t.Error(err)
    }
}
```

#### 支持钉钉消息类型

* text类型
```go
dps.TextConfig{
    Text: dps.TextFieldText{
        "sdk测试",    // 消息内容
    },
    At: dps.TextFieldAt{
        AtMobiles: nil,     // 需要at的用户的手机号列表 
        AtUserIds: nil,     // 需要at的用户的id列表
        IsAtAll:   false,   // 是否需要at所有人
    },
}
```

* link类型
```go
dps.LinkConfig{
    Link: dps.LinkFieldLink{
        Text:       "text",             // 链接文本
        Title:      "title",            // 链接标题
        PicUrl:     "www.baidu.com",    // 图片地址
        MessageUrl: "www.baidu.com",    // 链接地址
    },
}
```

* markdown类型
```go
dps.MarkdownConfig{
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
    At: dps.TextFieldAt{
        AtMobiles: nil,     // 需要at的用户的手机号列表 
        AtUserIds: nil,     // 需要at的用户的id列表
        IsAtAll:   false,   // 是否需要at所有人
    },
}
```

* actionCard类型
```go
dps.ActionCardConfig{
    ActionCard: dps.ActionCardFieldActionCard{
        Title:          "title",    // 标题
        Text:           "text",     // 内容
        SingleTitle:    "",         // 单个按钮的标题。(设置此项和singleURL后，btns无效。)
        SingleUrl:      "",         // 单个按钮链接
        BtnOrientation: dps.BtnOrientationH,    // dps.BtnOrientationV：按钮竖直排列 dps.BtnOrientationH：按钮横向排列
        Btns: []dps.ActionCardFieldBtns{        // 多按钮列表
            {
                Title:     "同意",       // 按钮标题
                ActionUrl: "#",         // 跳转链接
            },
            {
                Title:     "拒绝",
                ActionUrl: "#",
            },
        },
    },
}
```

* feedCard类型
```go
dps.FeedCardConfig{
    FeedCard: dps.FeedCardFieldFeedCard{
        Links: []dps.FeedFieldLinks{    // 链接列表
            {
                Title:      "title",    // 链接标题
                MessageUrl: "#",        // 跳转链接
                PicUrl:     "#",        // 图片地址
            },
            {
                Title:      "title",
                MessageUrl: "#",
                PicUrl:     "#",
            },
        },
    },
}
```

#### 可合并消息类型

仅支持`text`,`markdown`,`feedCard`类型的消息合并, `link`,`actionCard`类型的消息不支持合并。

#### 消息限频

所有类型的消息都将执行限频策略。每个secret对应的机器人在单位频率区间内只会发送一条消息。单位频率区间收到的同一可合并类型的消息将会合并成一条消息，并会在某个频率区间统一发送到钉钉。
