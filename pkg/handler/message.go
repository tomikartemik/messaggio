package handler

import (
	"github.com/gin-gonic/gin"
	"messaggio/pkg/model"
	"messaggio/pkg/utils"
	"net/http"
)

func (h *Handler) SaveNewMessage(c *gin.Context) {
	var newMessage model.Message
	if err := c.BindJSON(&newMessage); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.SaveNewMessage(newMessage)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, "Accepted!")
	}
}

func (h *Handler) ProcessedMessagesStats(c *gin.Context) {
	totalNumberOfProcessedMessages, err := h.services.ProcessedMessagesStats()
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"processed_messages": totalNumberOfProcessedMessages})
	}
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}
