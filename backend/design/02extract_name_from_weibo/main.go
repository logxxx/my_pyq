package main

import (
	"fmt"
	"github.com/logxxx/utils/fileutil"
	"github.com/logxxx/utils/netutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	os.Chdir("design/02extract_name_from_weibo")

	apiResp := &Resp{}
	err := fileutil.ReadJsonFile("source.json", apiResp)
	if err != nil {
		panic(err)
	}

	users := []string{}
	oldUsersMap := map[string]string{}
	fileutil.ReadByLine("users.txt", func(s string) error {
		if s != "" {
			users = append(users, s)
			oldUsersMap[s] = "old"
		}
		return nil
	})

	newUsersMap := map[string]string{}
	for _, e := range apiResp.Status {
		if oldUsersMap[fmt.Sprintf("%v_%v", e.User.IDStr, e.User.ScreenName)] != "" {
			continue
		}
		newUsersMap[fmt.Sprintf("%v_%v", e.User.IDStr, e.User.ScreenName)] = e.User.Avatar
	}

	log.Printf("get %v new users", len(newUsersMap))

	if len(newUsersMap) <= 0 {
		return
	}

	for k, v := range newUsersMap {
		users = append(users, k)
		_, err := url.Parse(v)
		if err == nil {
			_, avatarData, err := netutil.HttpGetRaw(v)
			if err == nil && len(avatarData) > 10*1024 {
				fileutil.WriteToFile(avatarData, filepath.Join("avatars", fmt.Sprintf("%v.jpg", k)))
			}
		}

	}

	log.Printf("get %v apiResp", users)

	fileutil.WriteToFile([]byte(strings.Join(users, "\n")), "users.txt")
}

type Resp struct {
	Status []struct {
		User struct {
			IDStr      string `json:"idstr"`
			ScreenName string `json:"screen_name"`
			Avatar     string `json:"avatar_hd"`
		} `json:"user"`
	} `json:"statuses"`
}
