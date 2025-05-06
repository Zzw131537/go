/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-09 16:41:48
 */
package v1

import (
	"mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCarousel(c *gin.Context) {

	var listCarousel service.CarouselService

	if err := c.ShouldBind(&listCarousel); err == nil {

		res := listCarousel.List(c.Request.Context())

		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
