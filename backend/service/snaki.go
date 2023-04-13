package service

import (
	"github.com/MuShaf-NMS/zesthub_test/dto"
)

type Service interface {
	Snaki(snaki dto.Snaki) (int, dto.Bahan)
}

type service struct{}

func (s *service) Snaki(snaki dto.Snaki) (int, dto.Bahan) {
	total := 0
	bahan := []int{
		int(snaki.Tersedia.Gula / snaki.Kebutuhan.Gula),
		int(snaki.Tersedia.Tepung / snaki.Kebutuhan.Tepung),
		int(snaki.Tersedia.Coklat / snaki.Kebutuhan.Coklat),
	}
	for _, b := range bahan {
		if total == 0 {
			total = b
		} else if b < total {
			total = b
		}
	}
	sisa := dto.Bahan{
		Gula:   snaki.Tersedia.Gula - snaki.Kebutuhan.Gula*float32(total),
		Tepung: snaki.Tersedia.Tepung - snaki.Kebutuhan.Tepung*float32(total),
		Coklat: snaki.Tersedia.Coklat - snaki.Kebutuhan.Coklat*float32(total),
	}
	return total, sisa
}

func NewService() Service {
	return &service{}
}
