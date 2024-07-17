package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/logxxx/utils"
	"github.com/logxxx/utils/reqresp"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	avatar = "/data/hehanyang/mytest/my_pyq/backend/chores/avatar.jpg"
	cover  = "/data/hehanyang/mytest/my_pyq/backend/chores/cover2.jpg"
)

type GetMemosResp struct {
	Memos []*Memo `json:"memos"`
}

type Profile struct {
	AvatarUrl string `json:"avatarUrl,omitempty"`
	NickName  string `json:"nickname,omitempty"`
	Cover     string `json:"cover,omitempty"`
}

type Memo struct {
	ID       int64     `json:"id"`
	User     User      `json:"user"`
	Content  string    `json:"content"`
	Images   []string  `json:"images,omitempty"`
	Video    string    `json:"video,omitempty"`
	ShowTime string    `json:"show_time,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
	Likes    []string  `json:"likes,omitempty"`
}

type User struct {
	ID        int64  `json:"id"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	NickName  string `json:"nickname,omitempty"`
}

type Comment struct {
	ID       int64  `json:"id"`
	MemoID   int64  `json:"memoId"`
	UserName string `json:"username,omitempty"`
	ReplyTo  string `json:"reply_to,omitempty"`
	Content  string `json:"content,omitempty"`
}

func main() {

	log.Infof("hehe")

	os.Chdir("/data/hehanyang/mytest/my_pyq/backend/chores/")

	wd, _ := os.Getwd()
	_ = wd

	g := gin.Default()

	g.Use(reqresp.Cors())

	g.GET("/backend/v1/profile", func(c *gin.Context) {
		resp := &Profile{
			AvatarUrl: utils.B64(avatar),
			NickName:  "杨蕊",
			Cover:     utils.B64(cover),
		}
		reqresp.MakeResp(c, resp)
	})

	g.GET("/backend/v1/file", func(c *gin.Context) {
		id := c.Query("id")
		//log.Infof("get file:%v", id)

		if id == "" {
			reqresp.MakeErrMsg(c, errors.New("empty id"))
			return
		}

		filePath := getFilePathByID(id)

		log.Infof("get file:%v", filePath)

		c.File(filePath)

		return

	})

	g.GET("/backend/v1/memos", func(c *gin.Context) {
		resp := GetMemosResp{}

		m1 := &Memo{
			User: User{
				AvatarUrl: "r1/avatar.jpg",
			},
			Content:  "为何会有如此说法呢？首先不得不承认，数学对于思维能力的确有着颇高的要求，正因如此，很多家长都觉得学习数学很大程度地依赖于天赋和智商。",
			Images:   []string{"r1/1.jpg", "r1/2.jpg", "r1/3.jpg", "r1/4.jpg"},
			Video:    "",
			ShowTime: "昨天",
			Comments: []Comment{
				{
					UserName: "桐宝and胖桓",
					Content:  "谢谢，拿走咯~",
				},
				{
					UserName: "阳光蓝色20170",
					Content:  "好好好",
				},
				{
					UserName: "用户9720105860",
					Content:  "转发微博",
				},
			},
			Likes: nil,
		}
		resp.Memos = append(resp.Memos, m1)

		m2 := &Memo{
			User: User{
				AvatarUrl: "r2/avatar.jpg",
			},
			Content:  "深圳能见度太高了，天儿特别好看",
			Images:   []string{"r2/1.jpg", "r2/2.jpg", "r2/3.jpg", "r2/4.jpg"},
			Video:    "",
			ShowTime: "1小时前",
			Comments: []Comment{
				{
					UserName: "一起散步吗Q",
					Content:  "快回重庆玩！",
				},
				{
					UserName: "万万又漫漫",
					Content:  "真心换真心[泪]真的好幸福",
				},
				{
					UserName: "_钙铁欣",
					Content:  "快来快来",
				},
			},
			Likes: nil,
		}
		resp.Memos = append(resp.Memos, m2)

		for _, m := range resp.Memos {
			m.User.AvatarUrl = utils.B64(m.User.AvatarUrl)
			for i := range m.Images {
				m.Images[i] = utils.B64(m.Images[i])
			}
		}

		/*
			resp.Memos = append(resp.Memos, Memo{
				ID: 1,
				User: User{
					ID:        11,
					AvatarUrl: utils.B64(avatar),
					NickName:  "杨蕊",
				},
				Content:  "洗了还是人了",
				Images:   []string{utils.B64("/data/hehanyang/mytest/moments/backend/chores/5.jpg"), utils.B64("/data/hehanyang/mytest/moments/backend/chores/6.jpg")},
				Video:    "",
				ShowTime: "昨天",
				Comments: []Comment{
					{MemoID: 222, ID: 233, UserName: "D.", ReplyTo: "", Content: "好多水啊"},
					{MemoID: 222, ID: 234, UserName: "惘", ReplyTo: "", Content: "给粉丝抽了"},
					{MemoID: 222, ID: 235, UserName: "珍惜缘分", ReplyTo: "", Content: "塞我嘴里 别洗也别扔"},
					{MemoID: 222, ID: 236, UserName: ".7", ReplyTo: "", Content: "我一个朋友说要当球衣收藏"},
				},
				Likes: []string{"惘", "珍惜缘分", "光", "四百", "Xu01", "呆子", "昵称个行将", "A 赦", "Z."},
			})

			resp.Memos = append(resp.Memos, Memo{
				ID: 2,
				User: User{
					ID:        11,
					AvatarUrl: utils.B64("/data/hehanyang/mytest/moments/backend/chores/avatar.jpg"),
					NickName:  "杨蕊",
				},
				Content:  "我真的好孕",
				Images:   []string{utils.B64("/data/hehanyang/mytest/moments/backend/chores/7.jpg"), utils.B64("/data/hehanyang/mytest/moments/backend/chores/8.jpg")},
				Video:    "",
				ShowTime: "昨天",
				Comments: []Comment{
					{MemoID: 222, ID: 233, UserName: "张三", ReplyTo: "", Content: "笑死"},
					{MemoID: 222, ID: 234, UserName: "李四", ReplyTo: "", Content: "抽象"},
					{MemoID: 222, ID: 235, UserName: "王二", ReplyTo: "", Content: "你是真的6"},
					{MemoID: 222, ID: 236, UserName: "麻子", ReplyTo: "", Content: "你好无聊啊"},
				},
				Likes: []string{"小猫", "大熊", "还带", "还珠格格", "皇阿玛"},
			})

		*/
		reqresp.MakeResp(c, resp)
	})

	g.Run(":8090")

}

func getFilePathByID(id string) string {
	return utils.B64To(id)
}
