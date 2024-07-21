package main

import (
	"github.com/logxxx/utils"
	"github.com/logxxx/utils/fileutil"
	"strings"
)

var (
	memosTmpl = &Tmpl{}
	profile   = &TmplProfile{}
)

func init() {
	fileutil.ReadYamlFile("/home/xunlei/sync/chore/memos.yml", memosTmpl)

	if len(memosTmpl.Data) <= 0 {
		panic("no memosTmpl")
	}

	fileutil.ReadYamlFile("/home/xunlei/sync/chore/profile.yml", profile)

	if profile.Name == "" {
		panic("no profile Tmpl")
	}
}

type Tmpl struct {
	Data []TmplData `yaml:"data"`
}
type TmplComment struct {
	Name    string `yaml:"name"`
	Content string `yaml:"content"`
	ReplyTo string `yaml:"reply_to"`
}
type TmplData struct {
	Name     string        `yaml:"name"`
	Avatar   string        `yaml:"avatar"`
	Pics     []string      `yaml:"pics"`
	Video    string        `yaml:"video"`
	Content  string        `yaml:"content"`
	Place    string        `yaml:"place"`
	ShowTime string        `yaml:"show_time"`
	Likes    []string      `yaml:"likes"`
	Comment  []TmplComment `yaml:"comments"`
}

type TmplProfile struct {
	Name   string `yaml:"name"`
	Avatar string `yaml:"avatar"`
	Cover  string `yaml:"cover"`
}

func getProfile() (resp Profile) {
	resp = Profile{
		AvatarUrl: utils.B64(profile.Avatar),
		NickName:  profile.Name,
		Cover:     utils.B64(profile.Cover),
	}
	return
}

func getMemos() (resp []Memo) {

	for i, elem := range memosTmpl.Data {

		if strings.ReplaceAll(elem.Name, " ", "") == "" {
			continue
		}

		respElem := Memo{
			ID: i + 1,
			User: User{
				AvatarUrl: utils.B64(elem.Avatar),
				NickName:  elem.Name,
			},
			Content:  elem.Content,
			Place:    elem.Place,
			Video:    utils.B64(strings.ReplaceAll(elem.Video, " ", "")),
			ShowTime: elem.ShowTime,
			Likes:    utils.RemoveEmpty(elem.Likes),
		}

		for _, pic := range elem.Pics {
			if strings.ReplaceAll(pic, " ", "") == "" {
				continue
			}
			respElem.Images = append(respElem.Images, utils.B64(pic))
		}

		for _, c := range elem.Comment {
			if strings.ReplaceAll(c.Name, " ", "") == "" {
				continue
			}
			respElem.Comments = append(respElem.Comments, Comment{
				UserName: c.Name,
				ReplyTo:  c.ReplyTo,
				Content:  c.Content,
			})
		}

		resp = append(resp, respElem)

	}

	return
}
