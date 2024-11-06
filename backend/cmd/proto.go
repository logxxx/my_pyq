package main

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

type ApplyTmplReq struct {
	Elems []interface{} `json:"data"`
}

type ApplyTmplElem struct {
}
