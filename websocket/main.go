// @Author xiaozhaofu 2022/10/12 11:48:00
package main

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 8192

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Time to wait before force close on connection.
	closeGracePeriod = 10 * time.Second
)

type Connection struct {
	conn            *websocket.Conn
	mutex           sync.Mutex
	inChan, outChan chan []byte        // 收到的消息, 发出的消息
	ctx             context.Context    // Connection的上下文
	cancelFunc      context.CancelFunc // 取消Connection的方法
	sync.RWMutex
}

func main() {
	news()
}

func news() {
	conn := newWsConn()
	defer func() {
		_ = conn.conn.Close()
		FsWarn("--------new websocket 退出-------")
	}()

	// 申明定时器2s，设置心跳时间为2
	// ticker := time.NewTicker(pingPeriod)
	ticker := time.NewTicker(3 * time.Second)
	// 保持心跳
	go conn.heartbeat(ticker)

	conn.Listen()

	time.Sleep(closeGracePeriod)
}

func (c *Connection) Listen() {
	var wg sync.WaitGroup
	wg.Add(3)

	go c.readLoop(&wg)
	go c.processLoop(&wg)
	go c.writeLoop(&wg)

	wg.Wait()
}

// @Title readLoop
// @Description 读取信息
// @Author xiaozhaofu 2022-10-13 11:29:51
// @Param ctx
// @Param wg
func (c *Connection) readLoop(wg *sync.WaitGroup) {
	fmt.Println("----start readLoop--------")
	defer func() {
		wg.Done()
	}()
	go readMsg(c)

	select {
	case <-c.ctx.Done():
		fmt.Println("-------exit readLoop------")
		return
	}
}

func readMsg(c *Connection) {
	defer func() {
		fmt.Println("----readLoop cancel-----")
		c.cancelFunc()
	}()
	// c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Time{})
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Time{})
		return nil
	})

	for {
		messageType, message, err := c.conn.ReadMessage()

		if err != nil {
			if e, ok := err.(*websocket.CloseError); ok {
				log.Println("-----CloseError-----", e.Text, e.Code)
			}
			log.Printf("2-----%+v\n", err)
			break
		}
		if messageType == websocket.CloseMessage {
			log.Println("----CloseMessage Type---")
			return
		}
		// 目前忽略非文本的数据
		if messageType != websocket.TextMessage {
			continue
		}
		select {
		case <-c.ctx.Done():
			return
		case c.inChan <- message:

		}

	}
}

func (c *Connection) processLoop(wg *sync.WaitGroup) {
	fmt.Println("----start processLoop--------")
	defer func() {
		wg.Done()
	}()
	for {
		select {
		case <-c.ctx.Done():
			fmt.Println("-------exit processLoop------")
			return
		case msg := <-c.inChan:
			go processMsg(c, msg)
		}
	}
}

func processMsg(c *Connection, msg []byte) {
	var (
		wsmsg WsMsg
		wscon WsContent
	)

	if err := json.Unmarshal(msg, &wsmsg); err != nil {
		c.cancelFunc()
	}
	fmt.Printf("-----ws msg struct: %+v\n", wsmsg)

	if err := json.Unmarshal([]byte(wsmsg.Message.Content), &wscon); err != nil {
		c.cancelFunc()
		fmt.Println(err)
	}

	fee := strconv.Itoa(wscon.RefundFee / 100)
	fsinfo := "订单号: " + wscon.Tid + "; 退款金额: " + fee
	// 发送飞书消息
	go FsWarn(fsinfo)

	sendmsg := SendWsNews{
		ID:          wsmsg.ID,
		CommandType: "Ack",
		Time:        wsmsg.Time,
		SendTime:    wsmsg.SendTime,
		Type:        "trace_isv_notify",
	}
	sms := StructToJsonByte(sendmsg)
	if sms == nil {
		return
	}
	fmt.Println("--------ack sms----", string(sms))
	select {
	case c.outChan <- sms:
	case <-c.ctx.Done():
		fmt.Println("-------exit processMsg------")
		return

	}

}

// @Title writeLoop
// @Description 处理并返回信息
// @Author xiaozhaofu 2022-10-13 11:30:10
// @Param ctx
// @Param wg
func (c *Connection) writeLoop(wg *sync.WaitGroup) {
	fmt.Println("----start writeLoop--------")
	defer func() {
		wg.Done()
	}()
	for {
		select {
		case <-c.ctx.Done():
			fmt.Println("-------exit processMsg------")
			return
		case sms := <-c.outChan:

			if err := c.SendMessage(websocket.TextMessage, sms); err != nil {
				FsWarn("pdd ws WriteMessage failed:" + err.Error())
				c.cancelFunc()
				log.Println("3-----", err)
				return
			}

			log.Println("pdd ws WriteMessage success")

		}
	}
}

func newWsConn() *Connection {
	// 跳过证书验证
	tlsConfig := tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            nil,
	}
	dial := websocket.Dialer{TLSClientConfig: &tlsConfig}
	mill := getmill()
	dig := getdigest(mill)

	u := url.URL{
		Scheme: "wss",
		Host:   "message-api-test.pinduoduo.com",
		Path:   "/message/1034463058b34ee6b3e61b794308af44/" + mill + "/" + dig,
	}
	log.Printf("connecting to %s", u.String())

	c, _, err := dial.Dial(u.String(), nil)
	if err != nil {
		log.Fatalln("1-----", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	return &Connection{
		conn:       c,
		inChan:     make(chan []byte, 1000),
		outChan:    make(chan []byte, 1000),
		ctx:        ctx,
		cancelFunc: cancel,
	}
}

func (c *Connection) heartbeat(ticker *time.Ticker) {
	var hb HeartBeat

	i := 0
	defer ticker.Stop()
	for {
		i++
		hb = HeartBeat{
			CommandType: "HeartBeat",
			Time:        int(time.Now().UnixMilli()),
			Id:          i,
		}
		hbbyte := StructToJsonByte(hb)
		fmt.Println(string(hbbyte))
		select {
		case <-ticker.C:
			c.mutex.Lock()
			if err := c.SendMessage(websocket.PingMessage, hbbyte); err != nil {
				FsWarn("pdd news 心跳发送失败" + err.Error())
				log.Printf("ping error: %s\n", err.Error())
			}
			log.Println("---ping msg:", websocket.PingMessage)
			c.mutex.Unlock()

		case <-c.ctx.Done():
			fmt.Println("-------end heart----")
			return
		}

	}

}

// 发送消息，请统一调用本函数进行发送
// 消息发送时增加互斥锁，加强并发情况下程序稳定性
// 提醒：开发者发送消息时，不要调用 c.Conn.WriteMessage(messageType, []byte(message)) 直接发送消息
func (c *Connection) SendMessage(messageType int, data []byte) error {
	c.Lock()
	defer func() {
		c.Unlock()
	}()
	// 发送消息时，必须设置本次消息的最大允许时长(秒)
	if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		// 日志
		return err
	}
	if err := c.conn.WriteMessage(messageType, data); err != nil {
		return err
	} else {
		fmt.Println("发送消息成功")
		return nil
	}
}

func getdigest(mill string) string {
	cid := "1034463058b34ee6b3e61b794308af44"
	csec := "4d596f655ac072564b237929863a386b2e6f0786"
	md5str := Md5(cid + mill + csec)
	str := base64.StdEncoding.EncodeToString([]byte(md5str))
	return str
}

func getmill() string {
	now := time.Now()
	nano := now.UnixMilli()
	return strconv.Itoa(int(nano))
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func FsWarn(msg string) {
	fsurl := "https://open.feishu.cn/open-apis/bot/v2/hook/6d04948c-0dfd-4ef0-a2b2-bc27ebcb107a"
	contentType := "application/json"

	info := &fsInfo{
		MsgType: "text",
		Content: map[string]string{
			"text": msg,
		},
	}

	data, _ := json.Marshal(info)

	resp, err := http.Post(fsurl, contentType, strings.NewReader(string(data)))
	if err != nil {
		log.Printf("post failed, err:%v\n", err)
		return
	}

	resp.Body.Close()

}

type fsInfo struct {
	MsgType string            `json:"msg_type"`
	Content map[string]string `json:"content"`
}

type WsMsg struct {
	ID          int64   `json:"id"`
	CommandType string  `json:"commandType"`
	Time        int64   `json:"time"`
	SendTime    int64   `json:"sendTime"`
	Message     Message `json:"message"`
}
type Message struct {
	Type    string `json:"type"`
	MallID  int64  `json:"mallID"`
	Content string `json:"content"`
}
type WsContent struct {
	MallID    int64  `json:"mall_id"`
	RefundFee int    `json:"refund_fee"`
	BillType  int    `json:"bill_type"`
	Modified  int64  `json:"modified"`
	RefundID  int    `json:"refund_id"`
	Operation int    `json:"operation"`
	Tid       string `json:"tid"`
}

type SendWsNews struct {
	ID          int64  `json:"id"`
	CommandType string `json:"commandType"`
	Time        int64  `json:"time"`
	SendTime    int64  `json:"sendTime"`
	Type        string `json:"type"`
}

type HeartBeat struct {
	CommandType string `json:"commandType"`
	Time        int    `json:"time"`
	Id          int    `json:"id"`
}

func StructToJsonByte(e interface{}) []byte {
	if b, err := json.Marshal(e); err == nil {
		return b
	} else {
		return nil
	}
}
