/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-11 21:27:25
 */
package v1

import (
	"mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {

	var listCategory service.CategoryService

	if err := c.ShouldBind(&listCategory); err == nil {

		res := listCategory.List(c.Request.Context())

		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
