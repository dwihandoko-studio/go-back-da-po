package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type WilayahKecamatan struct {
	KodeKecamatan string `json:"kode_kecamatan"`
	Kecamatan     string `json:"kecamatan"`
}

// var client *http.Client

func GetWilayah(token string) ([]WilayahKecamatan, error) {
	var res []WilayahKecamatan

	client := http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("GET", os.Getenv("URL_DAPO")+"/wilayah_akses_kecamatan", nil)
	if err != nil {
		fmt.Print(err.Error())
		return nil, errors.New("gagal memuat data")
	}

	req.Header = http.Header{
		"Authorization": {"Bearer " + token},
		"X-API-Key":     {os.Getenv("API_TOKEN_DAPO")},
	}

	ress, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		return res, err
	}
	defer ress.Body.Close()

	// var dataByte []byte
	// _, err = ress.Body.Read(dataByte)
	// if err != nil {
	// 	return res, errors.New("Gagal membaca data dari API")
	// }

	// // Simpan data ke file JSON
	// file, err := os.Create("wilayah.json")
	// if err != nil {
	// 	return res, errors.New("Gagal membuat file data.json")
	// }
	// defer file.Close()

	// encoder := json.NewEncoder(file)
	// err = encoder.Encode(dataByte)
	// if err != nil {
	// 	println("Gagal menulis data ke file data.json")
	// 	return res, errors.New("Gagal menulis data ke file wilayah.json")
	// }

	bodyBytes, err := ioutil.ReadAll(ress.Body)
	if err != nil {
		fmt.Print(err.Error())
		return res, errors.New("gagal memuat data")
	}

	err = json.Unmarshal(bodyBytes, &res)

	if err != nil {
		fmt.Print(err.Error())
		return res, errors.New("gagal memuat data")
	}

	// datas = append(datas, responseObject)
	// fmt.Print(string(bodyBytes))
	//  return json.Unmarshal(bodyBytes, &responseObject)

	if len(res) > 0 {

		for i, kecamatan := range res {
			res[i].KodeKecamatan = strings.TrimSpace(kecamatan.KodeKecamatan)
		}

		responseJSON, _ := json.MarshalIndent(res, "", "  ")

		// output, _ := json.MarshalIndent(responseJSON, "", "  ")

		// Tulis data ke dalam file
		err := ioutil.WriteFile("uploads/wilayah.json", responseJSON, 0644)
		if err != nil {
			log.Println(err.Error())
			return res, errors.New("gagal menulis data ke file wilayah.json")
		}

		return res, nil
	} else {
		return res, errors.New("data tidak ditemukan")
	}

}
