package main

import (
	"github.com/logxxx/utils"
	"github.com/logxxx/utils/fileutil"
	"path/filepath"
	"strings"
	"time"
)

var (
	memosTmpl = &Tmpl{}
	profile   = &TmplProfile{}

	choreDir = "C:\\Users\\Administrator\\Desktop\\wx_source\\chore"
)

func init() {
	err := fileutil.ReadYamlFile(filepath.Join(choreDir, "memos.yml"), memosTmpl)
	if err != nil {
		panic(err)
	}

	if len(memosTmpl.Data) <= 0 {
		panic("no memosTmpl")
	}

	go func() {
		for {
			time.Sleep(3 * time.Second)
			fileutil.ReadYamlFile(filepath.Join(choreDir, "memos.yml"), memosTmpl)
			if len(memosTmpl.Data) <= 0 {
				panic("no memosTmpl")
			}
		}
	}()

	fileutil.ReadYamlFile(filepath.Join(choreDir, "profile.yml"), profile)

	if profile.Name == "" {
		panic("no profile Tmpl")
	}
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
