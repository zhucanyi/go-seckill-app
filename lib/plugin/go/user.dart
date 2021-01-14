import 'dart:async';

import 'package:flutter/services.dart';

class UserPlugin {
  // go-flutter插件中的包名，两者必须一致
  static const _channel = const MethodChannel("app.desktop.plugin.user");

  // go-flutter插件中的函数名，无参
  static Future<String> userinfo() async => _channel.invokeMethod("userinfo");

  // go-flutter插件中的函数名，有参
  static Future<String> message(String p) async => _channel.invokeMethod("message", p);

}