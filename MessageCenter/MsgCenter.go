package MessageCenter

import (
	"nqc.cn/xwebsocket"
	"encoding/json"
	"time"
	"fmt"
)

type MessageCenter struct {
	Clients map[*xwebsocket.WebSocketContext]*Client
	connectList []WebSocketEnvent
	readList map[string]WebSocketEnvent
	disconnectList []WebSocketEnvent
	Center *xwebsocket.XWebSocket
}

type Client struct {
	Conn *xwebsocket.WebSocketContext
	Code string
	StartTime int64
	Params map[string]interface{}
}

func NewMsgCenter() *MessageCenter {
	m := new (MessageCenter)
	m.Clients = make(map[*xwebsocket.WebSocketContext]*Client)
	m.Center = xwebsocket.NewXWebSocket(m.Connect,m.ReadMsg,m.DisConnect)
	return m
}

func (ms *MessageCenter)Connect(c *xwebsocket.WebSocketContext) {

	client := Client{Conn:c,StartTime:time.Now().Unix(),Params:make(map[string]interface{})}

	ms.Clients[c] = &client


	for _,value := range ms.connectList {
		value(&client,c.Message)
	}
	
}
func (ms *MessageCenter)ReadMsg(c *xwebsocket.WebSocketContext) {

	/*j1 := make(map[string]interface{})

	err := json.Unmarshal([]byte(c.Message),&j1)
	if err != nil || j1["order"] == nil{
		fmt.Println("error: ",err," ",j1["order"] == nil)
		return
	}
	ms.readList[j1["order"].(string)](ms.Clients[c],c.Message)*/
	j1 := make(map[string]interface{})

	err := json.Unmarshal([]byte(c.Message),&j1)
	if err != nil || j1["order"] == nil{
		if err == nil && j1["type"] != nil {
			if j1["type"].(float64) == 200 {
				message := `{"type":200,"state":true,"msg":"心跳包收到"}`
				c.WriteMsg(message)
			}
			return

		}
		fmt.Println("error: ",err," ",j1["order"] == nil)
		return
	}
	ms.readList[j1["order"].(string)](ms.Clients[c],c.Message)
}
func (ms *MessageCenter)DisConnect(c *xwebsocket.WebSocketContext) {
	fmt.Println("addr:",c.Conn.Request().RemoteAddr," connect from:",ms.Clients[c].StartTime," end Time:" ,time.Now().Unix())
	for _,value := range ms.disconnectList {
		value(ms.Clients[c],c.Message)
	}
}

type WebSocketEnvent func(c *Client,msg string)

func (ms *MessageCenter)AddConnectRouter(c WebSocketEnvent) {
	ms.connectList = append(ms.connectList,c)
}

func (ms *MessageCenter)AddReadMsgRouter(order string, c WebSocketEnvent) {
	if ms.readList == nil {
		ms.readList = make(map[string]WebSocketEnvent)
	}
	ms.readList[order] = c
	fmt.Println("ms:",ms.readList)
}

func (ms *MessageCenter)AddDisConnectRouter(c WebSocketEnvent) {
	ms.disconnectList = append(ms.disconnectList,c)
}


