package respository

import (
	"errors"
	"fmt"
	_ "github.com/lib/pq" // here
	"github.com/yagoazedias/dog-escaper/infraestruture"
	"github.com/yagoazedias/dog-escaper/model"
	"time"
)

type PortRepository struct {}

type PortRepositoryInterface interface {
	GetLastStatus() (*model.Port, error)
	UpdateLastStatus(status bool) (*model.Port, error)
}

func (r PortRepository) GetLastStatus() (*model.Port, error) {
	env := infraestruture.Postgres{}
	err := MustOpen(env.ConnectionString(), false)
	defer db.Close()

	if err != nil {
		return nil, err
	}

	port := model.Port{}

	result := db.Raw("SELECT * FROM port ORDER BY id DESC LIMIT 1").Scan(&port)

	if result.Error != nil {
		fmt.Println("Error connecting with database", result.Error.Error())
		return nil, errors.New("could not connect to database")
	}

	return &port, nil
}

func (r PortRepository) UpdateLastStatus(status bool) (*model.Port, error) {
	env := infraestruture.Postgres{}
	err := MustOpen(env.ConnectionString(), false)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	newPort := model.Port{
		IsOpen:    status,
		Timestamp: time.Now(),
	}
	if result := db.Create(newPort); result.Error != nil {
		return nil, errors.New(result.Error.Error())
	}
	return &newPort, nil
}