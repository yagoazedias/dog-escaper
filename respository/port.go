package respository

import (
	"errors"
	"fmt"
	"github.com/yagoazedias/dog-escaper/infraestruture"
	"github.com/yagoazedias/dog-escaper/model"
	"time"
)

type PortRepository struct {}

type PortRepositoryInterface interface {
	GetLastStatus() (bool, error)
	UpdateLastStatus(status bool) (*model.Port, error)
}

func (r PortRepository) GetLastStatus() (bool, error) {
	return false, nil
}

func (r PortRepository) UpdateLastStatus(status bool) (*model.Port, error) {
	env := infraestruture.Postgres{}
	fmt.Printf(env.ConnectionString())
	err := MustOpen(env.ConnectionString(), false)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	newPort := model.Port{
		IsOpen:    status,
		Timestamp: time.Now().Format(time.RFC850),
	}
	if result := db.Create(newPort); result.Error != nil {
		return nil, errors.New(result.Error.Error())
	}
	return &newPort, nil
}