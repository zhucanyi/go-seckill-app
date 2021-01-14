package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	"go-plugins/user"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(400, 600),
	// 添加插件
	flutter.AddPlugin(&user.UserPlugin{}),
}
