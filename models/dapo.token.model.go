package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
)

type TokenDapo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiredDate string `json:"expired_date"`
}

func GetTokenDapo() (TokenDapo, error) {
	var err error
	var client = &http.Client{}
	var data TokenDapo

	var param = url.Values{}
	param.Set("username", "backbone_lampungtengah")
	param.Set("password", "^2022!lampungtengah#")
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", os.Getenv("URL_DAPO")+`/token`, payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("X-API-Key", os.Getenv("API_TOKEN_DAPO"))

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	var dataByte []byte
	_, err = response.Body.Read(dataByte)
	if err != nil {
		return data, errors.New("gagal membaca data dari api")
	}

	// Simpan data ke file JSON
	file, err := os.Create("uploads/token.json")
	if err != nil {
		return data, errors.New("token.json")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		println("Gagal menulis data ke file token.json")
		return data, errors.New("gagal menulis data ke file token.json")
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}
