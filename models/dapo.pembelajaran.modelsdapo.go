package models

import (
	"backbone-dapodik/db"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type PembelajaranDapodik struct {
	PembelajaranID         *string `json:"pembelajaran_id"`
	SemesterID             *string `json:"semester_id"`
	SekolahID              *string `json:"sekolah_id"`
	RombonganBelajarID     *string `json:"rombongan_belajar_id"`
	MataPelajaranID        *int64  `json:"mata_pelajaran_id"`
	NamaMataPelajaran      *string `json:"nama_mata_pelajaran"`
	PtkTerdaftarID         *string `json:"ptk_terdaftar_id"`
	PtkID                  *string `json:"ptk_id"`
	NamaPTK                *string `json:"nama_ptk"`
	SKMengajar             *string `json:"sk_mengajar"`
	TanggalSKMengajar      *string `json:"tanggal_sk_mengajar"`
	JamMengajarPerMinggu   *int64  `json:"jam_mengajar_per_minggu"`
	StatusDiKurikulum      *int64  `json:"status_di_kurikulum"`
	KetStatDiKurikulum     *string `json:"ket_stat_di_kurikulum"`
	NamaMataPelajaranLokal *string `json:"nama_mata_pelajaran_lokal"`
	IndukPembelajaranID    *string `json:"induk_pembelajaran_id"`
	CreateDate             *string `json:"create_date"`
	LastUpdate             *string `json:"last_update"`
	SoftDelete             *int64  `json:"soft_delete"`
}

func GetPembelajaran(npsn, token string) ([]PembelajaranDapodik, error) {
	var res []PembelajaranDapodik

	client := http.Client{Timeout: 30 * time.Second}
	// req, err := client.Get(os.Getenv("URL_DAPO") + "/sekolah?kode_wilayah=" + wilayah)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	return res, errors.New("Gagal memuat data.")
	// }

	req, err := http.NewRequest("GET", os.Getenv("URL_DAPO")+"/pembelajaran?npsn="+npsn, nil)
	if err != nil {
		// fmt.Print(err.Error())
		return res, errors.New("gagal memuat data. HIT")
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
		// fmt.Print(err.Error())
		return res, err
	}
	defer ress.Body.Close()

	if ress.StatusCode == 200 {

		bodyBytes, err := ioutil.ReadAll(ress.Body)
		if err != nil {
			// fmt.Print(err.Error())
			return res, errors.New("gagal memuat data. HIT 1")
		}
		// data, err := json.Unmarshal(bodyBytes, &responseObject)
		// var responseObject []DapoSekolah

		err = json.Unmarshal(bodyBytes, &res)

		if err != nil {
			// fmt.Print(err.Error())
			return res, errors.New("gagal memuat data. HIT 2")
		}

		// datas = append(datas, responseObject)
		// fmt.Print(string(bodyBytes))
		//  return json.Unmarshal(bodyBytes, &responseObject)

		if len(res) > 0 {
			return res, nil
		} else {
			return res, errors.New("data tidak ditemukan")
		}
	} else {
		return res, errors.New("gagal Terkoneksi dengan Server Dapodik")
	}

}

func InsertUpdatePembelajaranDapodik(pembelajaran PembelajaranDapodik) error {

	con := db.CreateCon()

	tx, _ := con.Begin()

	sqlStatement := "INSERT INTO _pembelajaran_dapodik (pembelajaran_id, semester_id, sekolah_id, rombongan_belajar_id, mata_pelajaran_id, nama_mata_pelajaran, ptk_terdaftar_id, ptk_id, nama_ptk, sk_mengajar, tanggal_sk_mengajar, jam_mengajar_per_minggu, status_di_kurikulum, ket_stat_di_kurikulum, nama_mata_pelajaran_lokal, induk_pembelajaran_id, create_date, last_update, soft_delete) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE semester_id=values (semester_id), sekolah_id=values (sekolah_id), rombongan_belajar_id=values (rombongan_belajar_id), mata_pelajaran_id=values (mata_pelajaran_id), nama_mata_pelajaran=values (nama_mata_pelajaran), ptk_terdaftar_id=values (ptk_terdaftar_id), ptk_id=values (ptk_id), nama_ptk=values (nama_ptk), sk_mengajar=values (sk_mengajar), tanggal_sk_mengajar=values (tanggal_sk_mengajar), jam_mengajar_per_minggu=values (jam_mengajar_per_minggu), status_di_kurikulum=values (status_di_kurikulum), ket_stat_di_kurikulum=values (ket_stat_di_kurikulum), nama_mata_pelajaran_lokal=values (nama_mata_pelajaran_lokal), induk_pembelajaran_id=values (induk_pembelajaran_id), create_date=values (create_date), last_update=values (last_update), soft_delete=values (soft_delete)"

	exe, err := tx.Exec(sqlStatement, pembelajaran.PembelajaranID, pembelajaran.SemesterID, pembelajaran.SekolahID, pembelajaran.RombonganBelajarID, pembelajaran.MataPelajaranID, pembelajaran.NamaMataPelajaran, pembelajaran.PtkTerdaftarID, pembelajaran.PtkID, pembelajaran.NamaPTK, pembelajaran.SKMengajar, pembelajaran.TanggalSKMengajar, pembelajaran.JamMengajarPerMinggu, pembelajaran.StatusDiKurikulum, pembelajaran.KetStatDiKurikulum, pembelajaran.NamaMataPelajaranLokal, pembelajaran.IndukPembelajaranID, pembelajaran.CreateDate, pembelajaran.LastUpdate, pembelajaran.SoftDelete)

	if err != nil {
		// fmt.Println("Query error")
		// fmt.Println(err.Error())
		tx.Rollback()
		return errors.New("gagal Insert Update PEMBELAJARAN DAPO 1")
	}

	rows, err := exe.RowsAffected()
	if err != nil {
		// fmt.Println(err.Error())
		tx.Rollback()
		return errors.New("gagal insert update PEMBELAJARAN DAPO 2")
	}

	if rows > 0 {
		tx.Commit()

		return nil
	}
	tx.Rollback()
	return errors.New("gagal insert update PEMBELAJARAN DAPO 3")
}

func GetPembelajaranBackbone(id_sekolah string) ([]PembelajaranDapodik, error) {
	var res []PembelajaranDapodik
	var obj PembelajaranDapodik

	con := db.CreateCon()

	sqlStatement := "SELECT * From _pembelajaran_dapodik WHERE sekolah_id = ?"

	respo, err := con.Query(sqlStatement, id_sekolah)
	if err != nil {
		return res, errors.New("gagal memuat data")
	}
	defer respo.Close()

	for respo.Next() {
		err = respo.Scan(
			&obj.PembelajaranID,
			&obj.SemesterID,
			&obj.SekolahID,
			&obj.RombonganBelajarID,
			&obj.MataPelajaranID,
			&obj.NamaMataPelajaran,
			&obj.PtkTerdaftarID,
			&obj.PtkID,
			&obj.NamaPTK,
			&obj.SKMengajar,
			&obj.TanggalSKMengajar,
			&obj.JamMengajarPerMinggu,
			&obj.StatusDiKurikulum,
			&obj.KetStatDiKurikulum,
			&obj.NamaMataPelajaranLokal,
			&obj.IndukPembelajaranID,
			&obj.CreateDate,
			&obj.LastUpdate,
			&obj.SoftDelete,
		)

		if err != nil {
			return res, errors.New("gagal mengambil data")
		}

		res = append(res, obj)
	}

	return res, nil
}
