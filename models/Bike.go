package models

type Bike struct {
	ID                     uint    `gorm:"primaryKey;autoIncrement"`
	Marca                  string  `gorm:"type:text;not null"`
	Modelo                 string  `gorm:"type:text;not null"`
	Año                    int     `gorm:"not null"`
	Tipo                   string  `gorm:"type:text;not null"`
	CilindrajeCC           int     `gorm:"not null"`
	PotenciaHP             int     `gorm:"not null"`
	TorqueNm               float64 `gorm:"not null"`
	Motor                  string  `gorm:"type:text;not null"`
	Refrigeracion          string  `gorm:"type:text;not null"`
	MedidaNeumaticoDelantero int  `gorm:"not null"`
	MedidaNeumaticoTrasero int    `gorm:"not null"`
	Transmision            string  `gorm:"type:text;not null"`
	PesoKg                 int     `gorm:"not null"`
	CapacidadCombustibleL  float64 `gorm:"not null"`
	ImagenUrl              string  `gorm:"type:text;not null"`
	Descripcion            string  `gorm:"type:text;not null"`
}
func (Bike) TableName() string {
	return "Bikes"
}
