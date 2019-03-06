package main

import (
	"database/sql"
	"log"
)

type Dao struct {
	DB *sql.DB
}

func (dao *Dao) connect(connectionString string) error {
	log.Printf("Connect to %v", connectionString)
	var err error
	dao.DB, err = a.sqlOpen("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Printf("DB Status %v", a.dao.DB.Stats())
	return nil
}
