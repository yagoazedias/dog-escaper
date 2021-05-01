package respository

import (
	"github.com/jinzhu/gorm"
	"log"
)

var (
	db *gorm.DB
)

func MustOpen(connectionURL string, logMode bool) error {
	var err error
	db, err = gorm.Open("postgres", connectionURL)
	if err != nil {
		log.Print("could not open database connection", "err", err)
		return err
	}
	db.LogMode(logMode)
	return nil
}

func Close() {
	_ = db.Close()
}
