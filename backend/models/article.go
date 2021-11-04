package models

import "strconv"

type Article struct {
	ID       int    `json:"id"`       // 文章编号
	Title    string `json:"title"`    // 文章标题
	Time     string `json:"time"`     // 发布时间
	Writer   string `json:"writer"`   // 文章作者
	Abstract string `json:"abstract"` // 文章摘要
	Content  string `json:"content"`  // 文章内容
}

func (article Article) String() string {
	str := ""
	str += "ID: " + strconv.Itoa(article.ID) + "\n"
	str += "Title: " + article.Title + "\n"
	str += "Time: " + article.Time + "\n"
	str += "Writer: " + article.Writer + "\n"
	str += "Abstract: " + article.Abstract + "\n"
	return str
}
