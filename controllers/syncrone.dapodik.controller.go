package controllers

import (
	"backbone-dapodik/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListFile(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}
	files, errY := ioutil.ReadDir("./uploads/ptk")
	if errY != nil {
		fmt.Println("Error:", errY)
	}

	count := len(files)
	fmt.Printf("Jumlah Sekolah: %d\n", count)

	token, errT := models.GetTokenDapo()
	if errT != nil {
		log.Println("GET TOKEN")
		log.Println(errT.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: errT.Error()},
		)
	}

	root := "./uploads/ptk/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			file, errs := os.Open(path)
			if errs != nil {
				return c.JSON(http.StatusOK,
					models.Response{Status: 400, Message: fmt.Sprintf("Tidak ada file %s", path)},
				)
			}
			defer file.Close()
			var ptk []models.DapoPtk
			errs = json.NewDecoder(file).Decode(&ptk)

			// decoder := json.NewDecoder(file)
			// errD := decoder.Decode(&ptk)
			if errs != nil {
				// if path == "uploads/ptk/sekolah_10801502.json" {
				fmt.Println(errs.Error())
				fmt.Println(errs)
				// }
				fmt.Printf("Gagal Decode %s\n", path)
				parts := strings.Split(path, "_")
				numberString := parts[len(parts)-1]
				numberString = strings.TrimSuffix(numberString, ".json")
				// number, errT := strconv.Atoi(numberString)
				// if errT != nil {
				// 	panic(errT)
				// }
				// fmt.Println(number)

				reqs, err := http.NewRequest("GET", os.Getenv("URL_DAPO")+"/ptk?npsn="+numberString, nil)
				if err != nil {
					return c.JSON(http.StatusOK,
						models.Response{Status: 404, Message: err.Error()},
					)
				}

				// Menambahkan header Authorization dan X-API-Key
				reqs.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
				reqs.Header.Add("X-API-Key", os.Getenv("API_TOKEN_DAPO"))

				// Mengirim request
				clientt := &http.Client{}
				resps, err := clientt.Do(reqs)
				if err != nil {
					fmt.Println(err)
					return c.JSON(http.StatusOK,
						models.Response{Status: 404, Message: err.Error()},
					)
				}
				defer resps.Body.Close()

				if resps.StatusCode == 200 {

					// Menyimpan respon dari API
					bodys, err := ioutil.ReadAll(resps.Body)
					if err != nil {
						fmt.Println(err)
						return c.JSON(http.StatusOK,
							models.Response{Status: 404, Message: err.Error()},
						)
					}

					fmt.Println("Respon dari API untuk Sekolah", numberString, ":", "BERHASIL")

					err = ioutil.WriteFile(fmt.Sprintf("uploads/ptk/sekolah_%s.json", numberString), bodys, 0644)
					if err != nil {
						fmt.Println("Error:", err)
						return c.JSON(http.StatusOK,
							models.Response{Status: 404, Message: err.Error()},
						)
					}
				} else {
					fmt.Println("Respon dari API untuk Sekolah", numberString, ":", "GAGAL")
				}
			}

		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK,
		models.Response{Status: 200, Message: "successs"},
	)
}

func SyncWilayah(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	token, errT := models.GetTokenDapo()
	if errT != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: errT.Error()},
		)
	}

	result, err := models.GetWilayah(token.AccessToken)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: err.Error()},
		)
	}

	return c.JSON(http.StatusOK,
		models.Response{
			Status:  200,
			Message: "success",
			Data:    result,
		})
}

func SyncSekolahById(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	var kode string
	var sekId string

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil {
		fmt.Println("PARSING REQUEST FORM")
		kode = c.FormValue("kode_kecamatan")
		sekId = c.FormValue("id")
	} else {
		//json_map has the JSON Payload decoded into a map
		fmt.Println("PARSING REQUEST JSON")
		kode = fmt.Sprintf("%s", json_map["kode_kecamatan"])
		sekId = fmt.Sprintf("%s", json_map["id"])
	}

	if kode == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "Kode kecamatan tidak boleh kosong dan harus valid."},
		)
	}

	if sekId == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "Id sekolah tidak boleh kosong dan harus valid."},
		)
	}

	token, errT := models.GetTokenDapo()
	if errT != nil {
		log.Println("Error: ", errT.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: errT.Error()},
		)
	}

	result, err := models.GetSekolah(kode, token.AccessToken)
	if err != nil {
		log.Println("Error: ", err.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: err.Error()},
		)
	}

	return c.JSON(http.StatusOK,
		models.Response{
			Status:  200,
			Message: "success",
			Data:    result,
		})
}

func SyncSekolah(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	token, errT := models.GetTokenDapo()
	if errT != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: errT.Error()},
		)
	}

	file, err := os.Open("uploads/wilayah.json")
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "Gagal mengambil wilayah.json"},
		)
	}
	defer file.Close()

	var kecamatans []models.WilayahKecamatan
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&kecamatans)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "Gagal decode wilayah.json"},
		)
	}

	for _, k := range kecamatans {
		req, err := http.NewRequest("GET", os.Getenv("URL_DAPO")+"/sekolah?kode_kecamatan="+k.KodeKecamatan, nil)
		if err != nil {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: err.Error()},
			)
		}

		// Menambahkan header Authorization dan X-API-Key
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
		req.Header.Add("X-API-Key", os.Getenv("API_TOKEN_DAPO"))

		// Mengirim request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: err.Error()},
			)
		}
		defer resp.Body.Close()

		// Menyimpan respon dari API
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: err.Error()},
			)
		}

		fmt.Println("Respon dari API untuk Kecamatan", k.Kecamatan, ":", "BERHASIL")

		err = ioutil.WriteFile(fmt.Sprintf("uploads/sekolah/kecamatan_%s.json", k.KodeKecamatan), body, 0644)
		if err != nil {
			fmt.Println("Error:", err)
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: err.Error()},
			)
		}

		// result, err := models.GetSekolah(k.KodeKecamatan, token.AccessToken)
		// if err != nil {
		// 	return c.JSON(http.StatusOK,
		// 		models.Response{Status: 404, Message: err.Error()},
		// 	)
		// }
		// log.Println(result)
		// // Menampilkan hasil data dari API
		// fmt.Printf("Data Sekolah %s: %s\n", k.Kecamatan, "ada")
	}

	// result, err := models.GetSekolah(kode, token.AccessToken)
	// if err != nil {
	// 	return c.JSON(http.StatusOK,
	// 		models.Response{Status: 404, Message: err.Error()},
	// 	)
	// }

	return c.JSON(http.StatusOK,
		models.Response{
			Status:  200,
			Message: "success",
			Data:    nil,
		})
}

func SyncPtk(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	token, errT := models.GetTokenDapo()
	if errT != nil {
		log.Println("GET TOKEN")
		log.Println(errT.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: errT.Error()},
		)
	}

	file, err := os.Open("uploads/wilayah.json")
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "Gagal mengambil wilayah.json"},
		)
	}
	defer file.Close()

	var kecamatans []models.WilayahKecamatan
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&kecamatans)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "Gagal decode wilayah.json"},
		)
	}

	for i, k := range kecamatans {
		if i < 0 {
			fmt.Println("SKIP KECAMATAN ", i)
		} else {
			if i > 0 {
				token1, errT1 := models.GetTokenDapo()
				if errT1 != nil {
					log.Println("GET TOKEN")
					log.Println(errT.Error())
					return c.JSON(http.StatusOK,
						models.Response{Status: 404, Message: errT.Error()},
					)
				}
				token = token1
			}
			// if k.KodeKecamatan == "120225" || k.KodeKecamatan == "120226" || k.KodeKecamatan == "120227" || k.KodeKecamatan == "120228" {
			fileS, err := os.Open(fmt.Sprintf("uploads/sekolah/kecamatan_%s.json", k.KodeKecamatan))
			if err != nil {
				return c.JSON(http.StatusOK,
					models.Response{Status: 400, Message: fmt.Sprintf("Gagal mengambil sekolah kecamatan_%s.json", k.KodeKecamatan)},
				)
			}
			defer fileS.Close()

			var sekolahs []models.DapoSekolah
			decoders := json.NewDecoder(fileS)
			err = decoders.Decode(&sekolahs)
			if err != nil {
				fmt.Println(err)
				return c.JSON(http.StatusOK,
					models.Response{Status: 400, Message: fmt.Sprintf("Gagal decode sekolah kecamatan_%s.json", k.KodeKecamatan)},
				)
			}

			for _, sk := range sekolahs {
				reqs, errPtk := http.NewRequest("GET", os.Getenv("URL_DAPO")+"/ptk?npsn="+sk.Npsn, nil)
				if errPtk != nil {
					return c.JSON(http.StatusOK,
						models.Response{Status: 404, Message: errPtk.Error()},
					)
				}

				// Menambahkan header Authorization dan X-API-Key
				reqs.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
				reqs.Header.Add("X-API-Key", os.Getenv("API_TOKEN_DAPO"))

				// Mengirim request
				clientt := &http.Client{}
				resps, errPtk := clientt.Do(reqs)
				if errPtk != nil {
					fmt.Println(errPtk)
					return c.JSON(http.StatusOK,
						models.Response{Status: 404, Message: errPtk.Error()},
					)
				}
				defer resps.Body.Close()

				if resps.StatusCode == 200 {

					// Menyimpan respon dari API
					bodys, errPtk := ioutil.ReadAll(resps.Body)
					if errPtk != nil {
						fmt.Println(errPtk)
						return c.JSON(http.StatusOK,
							models.Response{Status: 404, Message: errPtk.Error()},
						)
					}
					var resPtk []models.DapoPtk

					err = json.Unmarshal(bodys, &resPtk)

					if err != nil {
						fmt.Println(err)
						return c.JSON(http.StatusOK,
							models.Response{Status: 404, Message: err.Error()},
						)
					}

					if len(resPtk) > 0 {

						for _, ss := range resPtk {
							models.InsertUpdatePtkDapodik(ss)
						}
					}

					fmt.Println("Respon dari API untuk Sekolah", sk.Npsn, ":", "BERHASIL")
				} else {
					fmt.Println("Respon dari API untuk Sekolah", sk.Npsn, ":", "GAGAL")
				}

				reqsPembelajaran, errPembelajaran := http.NewRequest("GET", os.Getenv("URL_DAPO")+"/pembelajaran?npsn="+sk.Npsn, nil)
				if errPembelajaran != nil {
					return c.JSON(http.StatusOK,
						models.Response{Status: 404, Message: err.Error()},
					)
				}

				// Menambahkan header Authorization dan X-API-Key
				reqsPembelajaran.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
				reqsPembelajaran.Header.Add("X-API-Key", os.Getenv("API_TOKEN_DAPO"))

				// Mengirim request
				clienttPembelaran := &http.Client{}
				respsPembelajaran, errPembelajaran := clienttPembelaran.Do(reqsPembelajaran)
				if errPembelajaran != nil {
					fmt.Println(errPembelajaran)
					return c.JSON(http.StatusOK,
						models.Response{Status: 404, Message: errPembelajaran.Error()},
					)
				}
				defer respsPembelajaran.Body.Close()

				if respsPembelajaran.StatusCode == 200 {

					// Menyimpan respon dari API
					bodysPembelajaran, errPembelajaran := ioutil.ReadAll(respsPembelajaran.Body)
					if errPembelajaran != nil {
						fmt.Println(errPembelajaran)
						return c.JSON(http.StatusOK,
							models.Response{Status: 404, Message: errPembelajaran.Error()},
						)
					}
					var resPembelajaran []models.PembelajaranDapodik

					errPembelajaran = json.Unmarshal(bodysPembelajaran, &resPembelajaran)
					if errPembelajaran != nil {
						fmt.Println(errPembelajaran)
						return c.JSON(http.StatusOK,
							models.Response{Status: 404, Message: errPembelajaran.Error()},
						)
					}
					if len(resPembelajaran) > 0 {
						for _, ssPembelajaran := range resPembelajaran {
							models.InsertUpdatePembelajaranDapodik(ssPembelajaran)
						}
					}

					fmt.Println("Respon dari API untuk Sekolah", sk.Npsn, ":", "BERHASIL")

				} else {
					fmt.Println("Respon dari API untuk Sekolah", sk.Npsn, ":", "GAGAL")
				}
				// }
			}
		}
		// } else {
		// 	fmt.Println("Respon dari API untuk WILAYAH: ", k.KodeKecamatan, ":", "SKIP")
		// }

		// result, err := models.GetSekolah(k.KodeKecamatan, token.AccessToken)
		// if err != nil {
		// 	return c.JSON(http.StatusOK,
		// 		models.Response{Status: 404, Message: err.Error()},
		// 	)
		// }
		// log.Println(result)
		// // Menampilkan hasil data dari API
		// fmt.Printf("Data Sekolah %s: %s\n", k.Kecamatan, "ada")
	}

	return c.JSON(http.StatusOK,
		models.Response{
			Status:  200,
			Message: "success",
			Data:    nil,
		})
}
