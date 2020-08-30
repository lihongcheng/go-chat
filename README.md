# go-chat
golang实现的简易版聊天室

使用方法
1、分别将src/client/和src/msgserver下的文件进行go build分别生成client和msgserver的可执行文件
2、打开终端执行msgserver
3、再打开两个终端分别执行client，用来测试聊天效果。

聊天室效果如下：

服务端：
lhcwifi:msgserver lihongcheng$ ./msgserver

收到消息，游客92f5说：我是游客小明
收到消息，游客5c3b说：我是游客小红

客户端1：
lhcwifi:client lihongcheng$ ./client

欢迎来到聊天室！下面可以畅所欲言了！
我是游客小明
游客5c3b说:我是游客小红

客户端2：
lhcwifi:client lihongcheng$ ./client

欢迎来到聊天室！下面可以畅所欲言了！
游客92f5说:我是游客小明
我是游客小红
