package dto

type Bahan struct {
	Gula   float32 `json:"gula" validate:"required"`
	Tepung float32 `json:"tepung" validate:"required"`
	Coklat float32 `json:"coklat" validate:"required"`
}

type Snaki struct {
	Kebutuhan Bahan `json:"kebutuhan" validate:"required"`
	Tersedia  Bahan `json:"tersedia" validate:"required"`
}
