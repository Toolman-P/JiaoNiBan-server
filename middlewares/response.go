package middlewares

import (
	"JiaoNiBan-server/databases"

	"github.com/gin-gonic/gin"
)

func WebsiteDesc(c *gin.Context) {
	var q DescQuery
	c.Bind(&q)
	desc := databases.GetDesc(q.Author, q.Page)
	data := gin.H{"sum": len(desc), "data": desc}
	c.JSON(200, data)
}

func WebsiteContent(c *gin.Context) {
	var q ContentQuery
	c.Bind(&q)
	data := databases.GetContent(q.Author, q.Hash)
	c.JSON(200, data)
}

func WebsiteIndex(c *gin.Context) {
	var q IndexQuery
	c.Bind(&q)
	l, s := databases.GetIndex(q.Author)
	c.JSON(200, gin.H{"latest": l, "sum": s})
}

func Latest(c *gin.Context) {
	dat := databases.GetLatest()
	data := gin.H{"data": dat, "sum": len(dat)}
	c.JSON(200, data)
}
