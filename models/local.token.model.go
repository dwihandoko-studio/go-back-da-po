package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func GetTokenLocalDapo() (TokenDapo, error) {
	var err error
	var data TokenDapo

	file, err := ioutil.ReadFile("token.json")
	if err != nil {
		return data, errors.New("gagal membaca data dari TOKEN LOCAL")
	}

	// Melakukan decoding data JSON ke dalam interface
	err = json.Unmarshal(file, &data)
	if err != nil {
		return data, errors.New("gagal membaca data dari TOKEN LOCAL")
	}

	return data, nil
}
