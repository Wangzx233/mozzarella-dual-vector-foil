package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"time"
)

const (
	accessKey = "GwdirJ9DSwB5axgMcRL6N1O4TGv4dV7VBg6FxUe9"
	secretKey = "f8exhrTKYC-LEABlTKHSeclhScm2tSSxIavw6l_k"
	bucket    = "cqupt-post"
)

type res struct {
	Token string    `json:"token"`
	Exp   time.Time `json:"exp"`
}

func GetUpToken(c *gin.Context) {

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	r := res{
		Token: upToken,
		Exp:   time.Now().Add(time.Hour * 2),
	}
	c.JSON(200, gin.H{
		"status": 10000,
		"info":   "success",
		"data":   r,
	})
}
