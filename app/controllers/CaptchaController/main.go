package captchacontroller

// https://developer.aliyun.com/article/719480
// https://github.com/dchest/captcha
// https://book.aikaiyuan.com/golang/www.topgoer.com/gin%E6%A1%86%E6%9E%B6/%E5%85%B6%E4%BB%96/gin%E9%AA%8C%E8%AF%81%E7%A0%81.html
// https://www.codeleading.com/article/6387979921/

import (
	"bytes"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"`
	ImageUrl  string `json:"imageUrl"`
}

type CaptchaRequest struct {
	CaptchaId string `json:"captchaId" form:"captchaId"  binding:"required,min=4"`
	PngCode   string `json:"pngCode" form:"pngCode" binding:"required,min=4"`
}

func NewCaptchaString(c *gin.Context) {
	captcha := CaptchaString(captcha.DefaultLen)
	c.JSON(http.StatusOK, captcha)
}

// reload: 有直的話, 可以重新產生圖片(captchaId復用)
func GetCaptcha(c *gin.Context) {
	captchaId := c.Param("captchaId")
	fmt.Println("GetCaptchaPng : " + captchaId)
	ServeHTTP(c.Writer, c.Request)
}

// 產生 lenth長度驗證碼
func CaptchaString(length int) CaptchaResponse {
	if length <= 0 {
		length = captcha.DefaultLen
	}
	captchaId := captcha.NewLen(length)
	var captcha CaptchaResponse
	captcha.CaptchaId = captchaId
	captcha.ImageUrl = "/captcha/" + captchaId + ".png"
	return captcha
}

// 驗證captcha, id 跟 輸入
func VerifyString(cap CaptchaRequest) bool {
	// The function deletes the captcha with the given id from the internal storage, so that the same captcha can't be verified anymore.
	return captcha.VerifyString(cap.CaptchaId, cap.PngCode)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	fmt.Println("dir : " + dir)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	fmt.Println("file : " + file)
	fmt.Println("ext : " + ext)
	fmt.Println("id : " + id)
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("reload : " + r.FormValue("reload"))
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if Serve(w, r, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

// 測試middleware用
func VerifyCaptcha(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "okoko",
	})
}
