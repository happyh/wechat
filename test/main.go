package main

import (
	"fmt"
	"github.com/happyh/wechat"
	"github.com/happyh/wechat/cache"
	"github.com/happyh/wechat/menu"
)

func main() {
	//配置微信参数
	memcache := cache.NewMemoryCache("127.0.0.1:11211")
	config := &wechat.Config{
		AppID:          "wx5ea56cd113c8a5b4",
		AppSecret:      "3c17e5664dd3f9f05f960c49b0aca5f0",
		Token:          "xxxx",
		EncodingAESKey: "xxxx",
		Cache:          memcache,
	}
	wc := wechat.NewWechat(config)
	mu := wc.GetMenu()

	buttons := make([]*menu.Button, 1)
	btn := new(menu.Button)

	//创建click类型菜单
	btn.SetClickButton("name", "key123")
	buttons[0] = btn

	//设置btn为二级菜单
	btn2 := new(menu.Button)
	btn2.SetSubButton("subButton", buttons)

	buttons2 := make([]*menu.Button, 1)
	buttons2[0] = btn2

	//发送请求
	err := mu.SetMenu(buttons2)
	if err != nil {
		fmt.Printf("err= %v", err)
		return
	}
}
