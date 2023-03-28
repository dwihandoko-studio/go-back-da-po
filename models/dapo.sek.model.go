package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type DapoSekolah struct {
	SekolahID                string  `json:"sekolah_id"`
	SemesterID               string  `json:"semester_id"`
	Nama                     string  `json:"nama"`
	NamaNomenklatur          string  `json:"nama_nomenklatur"`
	Nss                      string  `json:"nss"`
	Npsn                     string  `json:"npsn"`
	BentukPendidikanID       int32   `json:"bentuk_pendidikan_id"`
	BentukPendidikan         string  `json:"bentuk_pendidikan"`
	AlamatJalan              string  `json:"alamat_jalan"`
	Rt                       int32   `json:"rt"`
	Rw                       int32   `json:"rw"`
	NamaDusun                string  `json:"nama_dusun"`
	KodeWilayah              string  `json:"kode_wilayah"`
	KodeDesaKelurahan        string  `json:"kode_desa_kelurahan"`
	DesaKelurahan            string  `json:"desa_kelurahan"`
	KodeKecamatan            string  `json:"kode_kecamatan"`
	Kecamatan                string  `json:"kecamatan"`
	KodeKabupaten            string  `json:"kode_kabupaten"`
	Kabupaten                string  `json:"kabupaten"`
	KodeProvinsi             string  `json:"kode_provinsi"`
	Provinsi                 string  `json:"provinsi"`
	KodePos                  string  `json:"kode_pos"`
	Lintang                  float64 `json:"lintang"`
	Bujur                    float64 `json:"bujur"`
	NomorTelepon             string  `json:"nomor_telepon"`
	NomorFax                 string  `json:"nomor_fax"`
	Email                    string  `json:"email"`
	Website                  string  `json:"website"`
	KebutuhanKhususID        int     `json:"kebutuhan_khusus_id"`
	KebutuhanKhusus          string  `json:"kebutuhan_khusus"`
	StatusSekolahID          string  `json:"status_sekolah_id"`
	StatusSekolah            string  `json:"status_sekolah"`
	SkPendirianSekolah       string  `json:"sk_pendirian_sekolah"`
	TanggalSkPendirian       string  `json:"tanggal_sk_pendirian"`
	StatusKepemilikanID      int     `json:"status_kepemilikan_id"`
	StatusKepemilikan        string  `json:"status_kepemilikan"`
	YayasanID                string  `json:"yayasan_id"`
	Yayasan                  string  `json:"yayasan"`
	SkIzinOperasional        string  `json:"sk_izin_operasional"`
	TanggalSkIzinOperasional string  `json:"tanggal_sk_izin_operasional"`
	NoRekening               string  `json:"no_rekening"`
	NamaBank                 string  `json:"nama_bank"`
	CabangKcpUnit            string  `json:"cabang_kcp_unit"`
	RekeningAtasNama         string  `json:"rekening_atas_nama"`
	Mbs                      int     `json:"mbs"`
	LuasTanahMilik           int     `json:"luas_tanah_milik"`
	LuasTanahBukanMilik      int     `json:"luas_tanah_bukan_milik"`
	KodeRegistrasi           string  `json:"kode_registrasi"`
	Npwp                     string  `json:"npwp"`
	NmWp                     string  `json:"nm_wp"`
	Keaktifan                string  `json:"keaktifan"`
	Flag                     string  `json:"flag"`
	DayaListrik              int     `json:"daya_listrik"`
	KontinuitasListrik       string  `json:"kontinuitas_listrik"`
	JarakListrik             int     `json:"jarak_listrik"`
	WilayahTerpencil         string  `json:"wilayah_terpencil"`
	WilayahPerbatasan        string  `json:"wilayah_perbatasan"`
	WilayahTransmigrasi      string  `json:"wilayah_transmigrasi"`
	WilayahAdatTerpencil     string  `json:"wilayah_adat_terpencil"`
	WilayahBencanaAlam       string  `json:"wilayah_bencana_alam"`
	WilayahBencanaSosial     string  `json:"wilayah_bencana_sosial"`
	PartisipasiBos           string  `json:"partisipasi_bos"`
	WaktuPenyelenggaraanID   int     `json:"waktu_penyelenggaraan_id"`
	WaktuPenyelenggaraan     string  `json:"waktu_penyelenggaraan"`
	SumberListrikID          int     `json:"sumber_listrik_id"`
	SumberListrik            string  `json:"sumber_listrik"`
	SertifikasiIsoID         int     `json:"sertifikasi_iso_id"`
	SertifikasiIso           string  `json:"sertifikasi_iso"`
	AksesInternetID          int     `json:"akses_internet_id"`
	AksesInternet            string  `json:"akses_internet"`
	AksesInternet2ID         int     `json:"akses_internet_2_id"`
	AksesInternet2           string  `json:"akses_internet_2"`
	Akreditasi               string  `json:"akreditasi"`
	CreateDate               string  `json:"create_date"`
	LastUpdate               string  `json:"last_update"`
}

// var client *http.Client

func GetSekolah(wilayah, token string) ([]DapoSekolah, error) {
	var res []DapoSekolah

	client := http.Client{Timeout: 30 * time.Second}
	// req, err := client.Get(os.Getenv("URL_DAPO") + "/sekolah?kode_wilayah=" + wilayah)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	return res, errors.New("Gagal memuat data.")
	// }

	req, err := http.NewRequest("GET", os.Getenv("URL_DAPO")+"/sekolah?kode_kecamatan="+wilayah, nil)
	if err != nil {
		fmt.Print(err.Error())
		return res, errors.New("gagal memuat data")
	}

	req.Header = http.Header{
		"Authorization": {"Bearer " + token},
		"X-API-Key":     {os.Getenv("API_TOKEN_DAPO")},
	}

	ress, err := client.Do(req)

	// req.Header.Add("Accept", "application/json")
	// req.Header.Add("Content-Type", "application/json")
	// resp, err := client.Do(req)
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
	// file, err := os.Create("sekolah_" + wilayah + ".json")
	// if err != nil {
	// 	return res, errors.New("Gagal membuat file sekolah_" + wilayah + ".json")
	// }
	// defer file.Close()

	// encoder := json.NewEncoder(file)
	// err = encoder.Encode(dataByte)
	// if err != nil {
	// 	println("Gagal menulis data ke file sekolah_" + wilayah + ".json")
	// 	return res, errors.New("Gagal menulis data ke file sekolah_" + wilayah + ".json")
	// }

	bodyBytes, err := ioutil.ReadAll(ress.Body)
	if err != nil {
		fmt.Print(err.Error())
		return res, errors.New("gagal memuat data")
	}
	// data, err := json.Unmarshal(bodyBytes, &responseObject)
	// var responseObject []DapoSekolah

	err = json.Unmarshal(bodyBytes, &res)

	if err != nil {
		fmt.Print(err.Error())
		return res, errors.New("gagal memuat data")
	}

	// datas = append(datas, responseObject)
	// fmt.Print(string(bodyBytes))
	//  return json.Unmarshal(bodyBytes, &responseObject)

	if len(res) > 0 {
		output, _ := json.MarshalIndent(res, "", "  ")

		// Tulis data ke dalam file
		err := ioutil.WriteFile("sekolah_"+wilayah+".json", output, 0644)
		if err != nil {
			log.Println(err.Error())
			return res, errors.New("Gagal menulis data ke file sekolah_" + wilayah + ".json")
		}
		return res, nil
	} else {
		return res, errors.New("data tidak ditemukan")
	}

}
