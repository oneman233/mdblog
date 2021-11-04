package utils

import (
	"errors"
	"io/ioutil"
	"mdblog/models"
	"os"
	"strings"
)

var articles []models.Article
var havArticles = false

func GetMarkdowns() ([]models.Article, error) {
	if havArticles {
		return articles, nil
	}
	files, err := ioutil.ReadDir(MarkdownPath)
	if err != nil {
		return nil, err
	}

	counter := 1

	for _, file := range files {
		f, err := os.Open(MarkdownPath + "/" + file.Name())
		if err != nil {
			f.Close()
			return nil, err
		}

		stat, err := f.Stat()
		if err != nil {
			return nil, err
		}

		data := make([]byte, stat.Size())
		_, err = f.Read(data)
		if err != nil {
			return nil, err
		}

		fileName := strings.TrimSuffix(file.Name(), ".md") // 去除 .md 后缀
		str := strings.Split(fileName, "-")                // 分割字符串

		dataString := string(data)
		L := strings.Index(dataString, "<!--") + len("<!--")
		R := strings.LastIndex(dataString, "-->")
		abstract := dataString[L:R]

		article := models.Article{
			ID:       counter,
			Title:    str[0],
			Time:     str[1],
			Writer:   str[2],
			Abstract: abstract,
			Content:  dataString,
		}

		articles = append(articles, article)
		counter++
	}

	havArticles = true

	return articles, nil
}

func GetMarkdownByID(ID int) (models.Article, error) {
	if !havArticles {
		_, err := GetMarkdowns()
		if err != nil {
			return models.Article{}, err
		}
	}

	ret := models.Article{}
	hav := false

	for _, article := range articles {
		if article.ID == ID {
			ret = article
			hav = true
			break
		}
	}

	if !hav {
		return models.Article{}, errors.New("can't find the article")
	}

	return ret, nil
}
