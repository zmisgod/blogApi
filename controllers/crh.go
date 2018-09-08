package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/zmisgod/blogApi/models"
)

//CrhController CHSR项目专用
type CrhController struct {
	BaseController
}

//@router /line.generate [post]
func (h *CrhController) LineGenerate() {
	res := h.Ctx.Request.Body
	requestJSON, err := ioutil.ReadAll(res)
	if err != nil {
		h.SendError(err.Error())
	} else {
		var chsr models.CHSR
		err = json.Unmarshal(requestJSON, &chsr)
		if err != nil {
			h.SendError(err.Error())
		} else {
			content := models.CRHGenerate(chsr)
			h.SendData(content, "ok")
		}
	}
}
