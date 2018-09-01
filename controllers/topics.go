package controllers

import (
	"github.com/zmisgod/blogApi/models"
)

type TopicsController struct {
	BaseController
}

//@router / [get]
func (h *TopicsController) Get() {
	var (
		err error
	)
	lists, err := models.GetTopicLists(h.page, h.pageSize)
	h.CheckError(err)
	h.SendData(lists, "successful")
}

//@router /:topicsId [get]
func (h *TopicsController) GetTopicsDetail() {
	var (
		err error
	)
	topicsID, err := h.GetInt(":topicsId")
	if err != nil {
		h.CheckError(err)
	}
	lists, err := models.GetTopicsArticleLists(topicsID, h.page, h.pageSize)
	h.CheckError(err)
	h.SendData(lists, "successful")
}
