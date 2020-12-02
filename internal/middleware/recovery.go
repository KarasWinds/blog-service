package middleware

import (
	"fmt"
	"time"

	"github.com/KarasWinds/blog-service/global"
	"github.com/KarasWinds/blog-service/pkg/app"
	"github.com/KarasWinds/blog-service/pkg/email"
	"github.com/KarasWinds/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recovery err: %v", err)
				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("例外拋出，發生時間: %d", time.Now().Unix()),
					fmt.Sprintf("錯誤訊息: %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err : %v", err)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
