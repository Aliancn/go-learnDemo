package utils

import (
	"gin-jwt-learn/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetPage 传入页数，返回过滤的结果数
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
