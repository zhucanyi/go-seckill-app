package user

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// go-flutter插件需要声明包名和函数名
// dart代码中调用时需要指定相应的包名和函数名
const (
	channelName = "app.desktop.plugin.user"
	userinfo    = "userinfo"
	message     = "message"
)

// 声明插件结构体
type UserPlugin struct{}

// 指定为go-flutter插件
var _ flutter.Plugin = &UserPlugin{}

// 初始化插件
func (UserPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc(userinfo, userInfoFunc)
	return nil
}

// 具体的逻辑处理函数，无参数传递
func userInfoFunc(arguments interface{}) (reply interface{}, err error) {

	return "Hello, zcy2", nil
}
