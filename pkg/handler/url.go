package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/shortener-url-backend/pkg/urlCheck"
)

// @Summary		AddUrl
// @Tags			url
// @Description	add url
// @ID				add-url
// @Accept			json
// @Produce		json
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/a [post]
func (h *Handler) addUrl(c *gin.Context) { // добавление url
	var url string = c.Query("url")
	resIsUrl, err := urlCheck.CheckIsUrl(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"result":  "",
		})
		return
	}
	if !resIsUrl {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в query",
			"result":  "",
		})
		return
	}
	if resIsUrl {
		resAddUrl, statusCode, err := h.services.AddUrl(url)
		if err != nil {
			c.JSON(statusCode, map[string]interface{}{
				"status":  statusCode,
				"message": err.Error(),
				"result":  "",
			})
			return
		}
		if len(resAddUrl) > 0 {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "успешное добавление ссылки",
				"result":  resAddUrl,
			})
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  http.StatusOK,
				"message": "ошибка добавления ссылки",
				"result":  resAddUrl,
			})
		}
	}
}
