package utils

// import (
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// 	"strconv"
// )

// func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		page, _ := strconv.Atoi(c.Query("page"))
// 		if page == 0 {
// 			page = 1
// 		}

// 		pageSize, _ := strconv.Atoi(c.Query("page_size"))
// 		switch {
// 		case pageSize > 100:
// 			pageSize = 100
// 		case pageSize <= 0:
// 			pageSize = 10
// 		}

// 		offset := (page - 1) * pageSize
// 		return db.Offset(offset).Limit(pageSize)
// 	}
// }

// db.Scopes(Paginate(c)).Find(&users)
// db.Scopes(Paginate(c)).Find(&articles)