package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	// 导入我们刚才生成的go代码所在的包，注意你们自己的项目路径，可能跟本例子不一样

	"github.com/gin-gonic/examples/protobuf/result"
	"github.com/gin-gonic/examples/protobuf/score_server"
)

type User struct {
	Id    int
	Name  string
	Age   int
	score []int
}

func main() {
	// 初始化消息
	score_info := &score_server.BaseScoreInfoT{}
	score_info.WinCount = 10
	score_info.LoseCount = 1
	score_info.ExceptionCount = 2
	score_info.KillCount = 2
	score_info.DeathCount = 1
	score_info.AssistCount = 3
	score_info.Rating = 120
	score_info.AssistPlayerId = []int32{4, 6, 8}
	resA := &result.SearchRequest{
		Query:         "test",
		PageNumber:    1,
		ResultPerPage: 10,
	}
	anyA, _ := proto.Marshal(resA)

	score_info.Details = &anypb.Any{
		TypeUrl: "type.googleapis.com/result.SearchRequest",
		// TypeUrl: "SearchRequest A",
		Value: anyA,
	}
	fmt.Printf("--socresize--%+v\n", proto.Size(score_info))
	// 以字符串的形式打印消息
	// fmt.Println("score_info--", score_info.String())
	// fmt.Printf("--score_info--%v\n", score_info)

	// encode, 转换成二进制数据
	data, err := proto.MarshalOptions{Deterministic: true}.Marshal(score_info)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--data--%+v\n", string(data))

	// dj, _ := protojson.Marshal(score_info)
	// dj, _ := protojson.MarshalOptions{
	// 	Multiline:       false,
	// 	Indent:          "  ",
	// 	UseProtoNames:   true,
	// 	EmitUnpopulated: true,
	// }.Marshal(score_info)

	// fmt.Printf("--data json--%+v\n", string(dj))

	// decode, 将二进制数据转换成struct对象
	new_score_info := score_server.BaseScoreInfoT{}

	err = proto.Unmarshal(data, &new_score_info)
	if err != nil {
		panic(err)
	}

	// 以字符串的形式打印消息
	details := new_score_info.GetDetails()
	anytype := details.GetTypeUrl()
	anystr := details.String()
	anyval := details.GetValue()
	anymess, _ := details.UnmarshalNew()

	fmt.Printf("--anytype--%+v\n", anytype)
	fmt.Println("--anystr-", anystr)
	fmt.Printf("--anyval--%+v\n", anyval)
	fmt.Printf("--anymess--%+v-- type:%T\n", anymess.(*result.SearchRequest).PageNumber, anymess)

	// new_score_json := score_server.BaseScoreInfoT{}
	// _ = protojson.Unmarshal(dj, &new_score_json)
	// fmt.Printf("--new_score_json--%+v\n", new_score_json.String())

	noticeReq := &score_server.NoticeReaderRequest{
		Msg: "hello notice",
		NoticeWay: &score_server.NoticeReaderRequest_Email{
			Email: "xiaozhaofu@test.com",
		},
	}

	email := noticeReq.GetEmail()
	msg := noticeReq.GetMsg()
	way := noticeReq.GetNoticeWay()
	switch v := noticeReq.NoticeWay.(type) {
	case *score_server.NoticeReaderRequest_Email:
		fmt.Println("----email:", v.Email)
	case *score_server.NoticeReaderRequest_Phone:
		fmt.Println("----phone:", v.Phone)
	default:
		fmt.Println("unknown way")
	}
	fmt.Println("--email--", email)
	fmt.Println("--msg--", msg)
	fmt.Println("--way--", way)

	name := proto.MessageName(noticeReq)
	fmt.Println("--name--", name)

}

type osError string

func (err osError) Error() string {
	return string(err)
}

const (
	ErrInvalid    = osError("invalid argument")
	ErrPermission = osError("permission denied")
	ErrExist      = osError("file already exists")
	ErrNotExist   = osError("file does not exist")
	ErrClosed     = osError("file already closed")
)

func doerr() {
	fmt.Println(ErrInvalid)
}
