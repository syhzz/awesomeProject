package models

import (
	"BeegoDemo/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"strconv"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
	//Status int //Status=0为正常，1为删除，2为冻结
}

// ---------数据处理-----------
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	return i, err
}

// -----------数据库操作---------------
// 插入一篇文章
func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}

// 根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	num, _ := beego.AppConfig.Int("articlePageNum")
	page--
	fmt.Println("-------------->page", page)
	return QueryArticleWithPage(page, num)
}

func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticleWithCon(sql)
}

func QueryArticleWithCon(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	fmt.Println(sql)
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList, nil
}

// 存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var artcileRowsNum = 0

// 只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum() int {
	if artcileRowsNum == 0 {
		artcileRowsNum = QueryArticleRowNum()
	}
	return artcileRowsNum
}

// 查询文章的总条数
func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

// 设置页数
func SetArticleRowsNum() {
	artcileRowsNum = QueryArticleRowNum()
}

func QueryArticleWithId(id int) Article {
	sql := "select id,title,tags,short,content,author,createtime from article where id=" + strconv.Itoa(id)
	row := utils.QueryRowDB(sql)
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}

func UpdateArticle(article Article) (int64, error) {
	sql := "update article set title=?,tags=?,short=?,content=? where id=?"
	return utils.ModifyDB(sql, article.Title, article.Tags, article.Short, article.Content, article.Id)
}

func DeleteArticle(artId int64) (int64, error) {
	i, err := deleteArticleWithId(artId)
	SetArticleRowsNum()
	return i, err
}
func deleteArticleWithId(artId int64) (int64, error) {
	return utils.ModifyDB("delete from article where id = ?", artId)
}

// 查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	rows, err := utils.QueryDB(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}

// --------------按照标签查询--------------
func QueryArticlesWithTag(tag string, page int) ([]Article, error) {
	num, _ := beego.AppConfig.Int("articlePageNum")
	page--
	sql := " where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	sql += " limit " + strconv.Itoa(page*num) + "," + strconv.Itoa(num)
	return QueryArticleWithCon(sql)
}
