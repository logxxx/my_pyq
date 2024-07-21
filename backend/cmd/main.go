package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/logxxx/utils"
	"github.com/logxxx/utils/reqresp"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var (
	avatar = "avatar.jpg"
	cover  = "cover2.jpg"
)

type GetMemosResp struct {
	Memos []Memo `json:"memos"`
}

type Profile struct {
	AvatarUrl string `json:"avatarUrl,omitempty"`
	NickName  string `json:"nickname,omitempty"`
	Cover     string `json:"cover,omitempty"`
}

type Memo struct {
	ID       int       `json:"id"`
	User     User      `json:"user"`
	Content  string    `json:"content"`
	Place    string    `json:"place"`
	Images   []string  `json:"images,omitempty"`
	Video    string    `json:"video,omitempty"`
	ShowTime string    `json:"show_time,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
	Likes    []string  `json:"likes"`
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

	//os.Chdir("chores")

	wd, _ := os.Getwd()
	log.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>> wd:%v", wd)

	g := gin.Default()

	g.Use(reqresp.Cors())

	g.GET("/backend/v1/profile", func(c *gin.Context) {
		reqresp.MakeResp(c, getProfile())
	})

	g.GET("/backend/v1/file", func(c *gin.Context) {
		id := c.Query("id")
		//log.Infof("get file:%v", id)

		if id == "" {
			reqresp.MakeErrMsg(c, errors.New("empty id"))
			return
		}

		filePath := getFilePathByID(id)

		filePath = filepath.Join("/home/xunlei/sync/chore", filePath)

		log.Infof("get file:%v", filePath)

		//if !utils.HasFile(filePath) {
		//	reqresp.MakeErrMsg(c, fmt.Errorf("file not found:%v", filePath))
		//	return
		//}

		c.File(filePath)
	})

	g.GET("/backend/v1/memos", func(c *gin.Context) {
		resp := GetMemosResp{}

		/*
			m1 := &Memo{
				User: User{
					NickName:  "英孚在线教育",
					AvatarUrl: "r1/avatar.jpg",
				},
				Content:  "为何会有如此说法呢？首先不得不承认，数学对于思维能力的确有着颇高的要求，正因如此，很多家长都觉得学习数学很大程度地依赖于天赋和智商。",
				Images:   []string{"r1/1.jpg", "r1/2.jpg", "r1/3.jpg", "r1/4.jpg"},
				Video:    "",
				ShowTime: "昨天",
				Place:    "长沙市·芒果TV(总部)",
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
				Likes: []string{"张三", "里斯", "王二", "码字"},
			}
			resp.Memos = append(resp.Memos, m1)

			m2 := &Memo{
				User: User{
					NickName:  "阿布003",
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
				Likes: []string{"张三", "里斯", "王二", "码字"},
			}
			resp.Memos = append(resp.Memos, m2)

			m3 := &Memo{
				User: User{
					NickName:  "木木木木啊",
					AvatarUrl: "r3/avatar.jpg",
				},
				Content: "#原相机 #纯欲 #御姐",
				//Images:   []string{"r2/1.jpg", "r2/2.jpg", "r2/3.jpg", "r2/4.jpg"},
				Video:    "r3/1.mp4",
				ShowTime: "3小时前",
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
				Likes: []string{"张三", "里斯", "王二", "码字"},
			}
			resp.Memos = append(resp.Memos, m3)

			m4 := &Memo{
				User: User{
					NickName:  "梳子在发光",
					AvatarUrl: "r4/avatar.jpg",
				},
				Content: "据说摄影师这组视频点赞过万了",
				//Images:   []string{"r2/1.jpg", "r2/2.jpg", "r2/3.jpg", "r2/4.jpg"},
				Video:    "r4/1.mp4",
				ShowTime: "一个月前",
				Comments: []Comment{
					{
						UserName: "kenijiba",
						Content:  "我的白月光",
					},
					{
						UserName: "大苹果",
						Content:  "我在突厥斯坦",
					},
					{
						UserName: "Danny",
						Content:  "拍出了艺术气息~",
					},
				},
				Likes: nil,
			}
			resp.Memos = append(resp.Memos, m4)

			m5 := &Memo{
				User: User{
					NickName:  "种草安利鹅",
					AvatarUrl: "r5/avatar.jpg",
				},
				Content:  "雨后初晴，深圳双彩虹！收录了几张朋友们发到群里的照片，真的太美啦！",
				Images:   []string{"r5/1.jpg"},
				Video:    "",
				ShowTime: "5天前",
				Comments: []Comment{
					{
						UserName: "卷只小贝奶酪",
						Content:  "这小众的赛道居然还是被人发现并且进去了",
					},
					{
						UserName: "就养一只洋",
						Content:  "越到后面直接就更加疯批了",
					},
					{
						UserName: "抱着胖杉碎大觉",
						Content:  "没见过这么惨的，苏容华只是个纯爱战神罢了~",
					},
				},
				//Likes: []string{"h宋书颻_yx", "柳柳shiny", "意钟人VIN", "要做快乐的草莓熊", "德云社王筱阁", "拾拾拾拾拾年", "吃了这块皂", "再熬夜我就明天早点睡", "蝶熙YoYo", "强牌慢打故作姿态", "甜杏丷"},
				Likes: []string{"h宋书颻_yx"},
			}

			resp.Memos = append(resp.Memos, m5)

			for i, m := range resp.Memos {
				m.ID = i + 1
				m.User.AvatarUrl = utils.B64(m.User.AvatarUrl)
				for i := range m.Images {
					m.Images[i] = utils.B64(m.Images[i])
				}
				if m.Video != "" {
					m.Video = utils.B64(m.Video)
				}
			}
		*/

		resp.Memos = getMemos()

		reqresp.MakeResp(c, resp)
	})

	g.Run("0.0.0.0:8090")

}

func getFilePathByID(id string) string {
	return utils.B64To(id)
}
