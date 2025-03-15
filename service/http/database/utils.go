package database

import (
	"errors"
	"fmt"

	"github.com/Mrhb787/hospital-ward-manager/model"
)

func (s *service) GetUserById(userId uint32) (resp model.User, err error) {

	tx, txErr := s.client.DB.Begin()
	if txErr != nil {
		fmt.Println("Failed to begin transaction", txErr)
		return resp, txErr
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

	}()

	dbResp, dbErr := model.GetUserById(tx, int(userId))
	if dbErr != nil {
		return resp, dbErr
	}

	if dbResp == nil {
		return resp, errors.New("user not found")
	}

	resp = *dbResp
	return resp, nil
}
