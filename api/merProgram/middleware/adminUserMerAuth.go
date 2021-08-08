package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/api/merProgram/cache"
	"github.com/ququgou-shop/api/merProgram/db"
	merchantModel "github.com/ququgou-shop/modules/merchant/model"
	userModel "github.com/ququgou-shop/modules/user/model"
	"net/http"
)

//超级管理员认证
func AdminUserMerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err error
		)

		v, _ := c.Get("claims")

		if v == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":   401,
				"status": -1,
				"msg":    "授权已过期",
			})
			c.Abort()
			return
		}
		va := v.(*CustomClaims)

		//通过缓存获取

		//查询用户是否有商户权限，没有跳转到申请页面
		var userMer merchantModel.MerchantUser
		var modelUser userModel.User
		//通过缓存获取
		if bol := cache.WebCache.GetData(cache.User+va.Guid, &modelUser); !bol {
			err = db.MysqlConn().Where("guid =? ", va.Guid).First(&modelUser).Error
			if err != nil {
				//没有查询到 重新登录
				c.JSON(http.StatusOK, gin.H{
					"code":   401,
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}

			//存入缓存
			go cache.WebCache.SetData(cache.User+va.Guid, modelUser, 0)
		}

		err = db.MysqlConn().Where("user_id =? ", modelUser.ID).First(&userMer).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				//没有查询到 返回无权限
				c.JSON(http.StatusOK, gin.H{
					"code":   403,
					"status": -1,
					"msg":    "无权限",
				})
				c.Abort()
				return
			}
			//错误返回
			c.JSON(http.StatusOK, gin.H{
				"code":   401,
				"status": -1,
				"msg":    err.Error(),
			})
		}
		//继续交由下一个路由处理
	}
}
