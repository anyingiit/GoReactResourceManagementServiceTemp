package globalVars

import (
	"errors"

	"gorm.io/gorm"
)

type db struct {
	val   *gorm.DB
	isSet bool
}

var Db = &db{
	val:   nil,
	isSet: false,
}

func (d *db) SetDb(dbObj *gorm.DB) error {
	if d.isSet {
		return errors.New("db is already set")
	}
	d.val = dbObj
	d.isSet = true
	return nil
}

func (d *db) GetDb() (*gorm.DB, error) {
	if !d.isSet {
		return nil, errors.New("db is not set")
	}
	return d.val, nil
}
