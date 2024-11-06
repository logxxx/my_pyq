package main

func main() {

}

type Resp struct {
	Status []struct {
		User struct {
			IDStr      string `json:"idstr"`
			ScreenName string `json:"screen_name"`
			Avatar     string `json:"avatar_hd"`
		} `json:"user"`
		TextRaw string `json:"text_raw"`
	} `json:"statuses"`
	PageInfo PageInfo   `json:"page_info,omitempty"` //视频
	PicInfos PicInfoMap `json:"pic_infos,omitempty"` //图片

}

type PicInfoMap map[string]struct { //key: id
	Large struct {
		Url string `json:"url,omitempty"`
	} `json:"large,omitempty"`
	Largest struct {
		Url string `json:"url,omitempty"`
	} `json:"largest,omitempty"`
}

type PageInfo struct {
	PageID    string    `json:"page_id,omitempty"`
	PageTitle string    `json:"page_title,omitempty"`
	MediaInfo MediaInfo `json:"media_info,omitempty"`
}

type MediaInfo struct {
	Url         string `json:"mp4_hd_url,omitempty"`
	StreamUrlHD string `json:"stream_url_hd,omitempty"`
}
