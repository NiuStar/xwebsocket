# xwebsocket

# 使用方法
    package main
    
    import (
        "nqc.cn/server"
        "golang.org/x/net/websocket"
        "nqc.cn/xwebsocket/MessageCenter"
        "nqc.cn/fmt"
    )
    
    func main() {
    
        ser := server.Default()
        center := MessageCenter.NewMsgCenter()
    
        imCenter := &IMCenter{}
        center.AddConnectRouter(imCenter.Connect)
        center.AddDisConnectRouter(imCenter.Disconncet)
        center.AddReadMsgRouter("IM", imCenter.ReadMsg)
    
        ser.HandfuncWebSocket("/wss",websocket.Handler(center.Center.EchoServer))
        ser.RunServer()
    }
    
    type IMCenter struct {
    
    }
    
    func (this *IMCenter)Connect(c *MessageCenter.Client,msg string) {
        fmt.Println("Connect")
    }
    
    func (this *IMCenter)Disconncet(c *MessageCenter.Client,msg string) {
        fmt.Println("Disconncet")
    }
    
    func (this *IMCenter)ReadMsg(c *MessageCenter.Client,msg string) {
    
        fmt.Println(msg)
    }
  
  连接：ws://127.0.0.1:3000/wss
  
  默认状态下：{"type":200}为心跳包
  
  {"order":"IM"}就可以进入ReadMsg方法中
  
  