package services

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// gorm.Model definition
type Carro struct {
	ID          uint   `gorm:"primarykey;unqiue;autoIncrement"`
	Placa       string `gorm:"unique"`
	Responsavel string `gorm:"not null"`
}

// gorm.Model definition
type Vaga struct {
	ID         uint   `gorm:"primarykey;autoIncrement"`
	Placa      string `gorm:"unique;foreignkey:placa;references:carro"`
	Disponivel bool
	Price      float64
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoCreateTime"`
}

// gorm.Model definition
type Admins struct {
	ID        int       `gorm:"primarykey;autoIncrement"`
	Nome      string    `gorm:"not null;size:40"`
	Token     string    `gorm:"not null;size:255;unique"`
	Valid     bool      `gorm:"not null;default:false"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt time.Time `gorm:"autoCreateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// gorm.Model definition
type Historys struct {
	ID        uint `gorm:"primarykey;unqiue;autoIncrement"`
	Placa     string
	Price     float64
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

var Db *gorm.DB
var err error

func DbStart() {
	dsn := "host=localhost user=postgres password=caua12 dbname=parking port=5432 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	Db.AutoMigrate(Vaga{}, Carro{}, Admins{}, Historys{})

	if err != nil {
		log.Println("Não foi possível se conectar ao DB")
	}
}
