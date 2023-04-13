package handler

import (
	"github.com/MuShaf-NMS/zesthub_test/dto"
	"github.com/MuShaf-NMS/zesthub_test/helper"
	"github.com/MuShaf-NMS/zesthub_test/service"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Index(ctx *gin.Context)
	Snaki(ctx *gin.Context)
}

func (h *handler) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "Wellcome"})
}

func (h *handler) Snaki(ctx *gin.Context) {
	var snaki dto.Snaki
	ctx.BindJSON(&snaki)
	if err := helper.Validate(snaki); err != nil {
		ctx.JSON(400, gin.H{"msg": "value harus lebih besar dari 0"})
		return
	}
	snaki.Tersedia = dto.Bahan{
		Gula:   snaki.Tersedia.Gula * 1000,
		Tepung: snaki.Tersedia.Tepung * 1000,
		Coklat: snaki.Tersedia.Coklat * 1000,
	}
	total, sisa := h.service.Snaki(snaki)
	ctx.JSON(200, gin.H{"total": total, "sisa": sisa})
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &handler{
		service: service,
	}
}
