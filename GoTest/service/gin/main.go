package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//指定默认值
		//http://localhost:8080/user 才会打印出来默认的值
		name := c.DefaultQuery("name", "枯藤")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		if file.Size > 1024*1024*2 {
			fmt.Println("文件太大了")
			c.String(http.StatusBadRequest, "文件太大了")
			return
		}
		if file.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传图片")
			c.String(http.StatusBadRequest, "只允许上传图片")

			return
		}
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})

	// 创建路由，默认使用2个中间件Logger(), Recovery()
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload/batch", upload_m)
	routes_group(r)
	r.Run(":8080")
}

func upload_m(c *gin.Context) {
	fmt.Println("进入批量上传")
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		return
	}
	files := form.File["files"]
	for _, file := range files {
		if err := c.SaveUploadedFile(file, "./uploadFiles/"+file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("upload ok %d files", len(files)))
}

// 路由组
func routes_group(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
