package models

import (
	"backbone-dapodik/db"
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

type DapoPtk struct {
	PtkID                                         *string  `json:"ptk_id"`
	SemesterID                                    *string  `json:"semester_id"`
	SekolahID                                     *string  `json:"sekolah_id"`
	PtkTerdaftarID                                *string  `json:"ptk_terdaftar_id"`
	Nama                                          *string  `json:"nama"`
	Nip                                           *string  `json:"nip"`
	JenisKelamin                                  *string  `json:"jenis_kelamin"`
	TempatLahir                                   *string  `json:"tempat_lahir"`
	TanggalLahir                                  *string  `json:"tanggal_lahir"`
	Nik                                           *string  `json:"nik"`
	NoKk                                          *string  `json:"no_kk"`
	NiyNigk                                       *string  `json:"niy_nigk"`
	Nuptk                                         *string  `json:"nuptk"`
	Nuks                                          *string  `json:"nuks"`
	StatusKepegawaian                             *string  `json:"status_kepegawaian"`
	JenisPtkID                                    *string  `json:"jenis_ptk_id"`
	JenisPtk                                      *string  `json:"jenis_ptk"`
	PengawasBidangStudi                           *string  `json:"pengawas_bidang_studi"`
	Agama                                         *string  `json:"agama"`
	Kewarganegaraan                               *string  `json:"kewarganegaraan"`
	AlamatJalan                                   *string  `json:"alamat_jalan"`
	Rt                                            *int64   `json:"rt"`
	Rw                                            *int64   `json:"rw"`
	NamaDusun                                     *string  `json:"nama_dusun"`
	DesaKelurahan                                 *string  `json:"desa_kelurahan"`
	KodeKecamatan                                 *string  `json:"kode_kecamatan"`
	Kecamatan                                     *string  `json:"kecamatan"`
	KodeKabupaten                                 *string  `json:"kode_kabupaten"`
	Kabupaten                                     *string  `json:"kabupaten"`
	KodeProvinsi                                  *string  `json:"kode_provinsi"`
	Provinsi                                      *string  `json:"provinsi"`
	KodePos                                       *string  `json:"kode_pos"`
	Lintang                                       *float64 `json:"lintang"`
	Bujur                                         *float64 `json:"bujur"`
	NoTeleponRumah                                *string  `json:"no_telepon_rumah"`
	NoHp                                          *string  `json:"no_hp"`
	Email                                         *string  `json:"email"`
	StatusKeaktifan                               *string  `json:"status_keaktifan"`
	SkCpns                                        *string  `json:"sk_cpns"`
	TglCpns                                       *string  `json:"tgl_cpns"`
	SkPengangkatan                                *string  `json:"sk_pengangkatan"`
	TmtPengangkatan                               *string  `json:"tmt_pengangkatan"`
	LembagaPengangkat                             *string  `json:"lembaga_pengangkat"`
	PangkatGolongan                               *string  `json:"pangkat_golongan"`
	KeahlianLaboratorium                          *string  `json:"keahlian_laboratorium"`
	SumberGaji                                    *string  `json:"sumber_gaji"`
	NamaIbuKandung                                *string  `json:"nama_ibu_kandung"`
	StatusPerkawinan                              *string  `json:"status_perkawinan"`
	NamaSuamiIstri                                *string  `json:"nama_suami_istri"`
	NipSuamiIstri                                 *string  `json:"nip_suami_istri"`
	PekerjaanSuamiIstri                           *int64   `json:"pekerjaan_suami_istri"`
	TmtPns                                        *string  `json:"tmt_pns"`
	SudahLisensiKepalaSekolah                     *string  `json:"sudah_lisensi_kepala_sekolah"`
	JumlahSekolahBinaan                           *int64   `json:"jumlah_sekolah_binaan"`
	PernahDiklatKepengawasan                      *string  `json:"pernah_diklat_kepengawasan"`
	NmWp                                          *string  `json:"nm_wp"`
	StatusData                                    *int64   `json:"status_data"`
	Karpeg                                        *string  `json:"karpeg"`
	Karpas                                        *string  `json:"karpas"`
	MampuHandleKk                                 *string  `json:"mampu_handle_kk"`
	KeahlianBraille                               *string  `json:"keahlian_braille"`
	KeahlianBhsIsyarat                            *string  `json:"keahlian_bhs_isyarat"`
	Npwp                                          *string  `json:"npwp"`
	Bank                                          *string  `json:"bank"`
	RekeningBank                                  *string  `json:"rekening_bank"`
	RekeningAtasNama                              *string  `json:"rekening_atas_nama"`
	TahunAjaran                                   *string  `json:"tahun_ajaran"`
	NomorSuratTugas                               *string  `json:"nomor_surat_tugas"`
	TanggalSuratTugas                             *string  `json:"tanggal_surat_tugas"`
	TmtTugas                                      *string  `json:"tmt_tugas"`
	PtkInduk                                      *string  `json:"ptk_induk"`
	JenisKeluar                                   *string  `json:"jenis_keluar"`
	TglPtkKeluar                                  *string  `json:"tgl_ptk_keluar"`
	RiwayatKepangkatanPangkatGolongan             *string  `json:"riwayat_kepangkatan_pangkat_golongan"`
	RiwayatKepangkatanNomorSk                     *string  `json:"riwayat_kepangkatan_nomor_sk"`
	RiwayatKepangkatanTanggalSk                   *string  `json:"riwayat_kepangkatan_tanggal_sk"`
	RiwayatKepangkatanTmtPangkat                  *string  `json:"riwayat_kepangkatan_tmt_pangkat"`
	RiwayatKepangkatanMasaKerjaGolTahun           *int64   `json:"riwayat_kepangkatan_masa_kerja_gol_tahun"`
	RiwayatKepangkatanMasaKerjaGolBulan           *int64   `json:"riwayat_kepangkatan_masa_kerja_gol_bulan"`
	RiwayatGajiBerkalaPangkatGolongan             *string  `json:"riwayat_gaji_berkala_pangkat_golongan"`
	RiwayatGajiBerkalaNomorSk                     *string  `json:"riwayat_gaji_berkala_nomor_sk"`
	RiwayatGajiBerkalaTanggalSk                   *string  `json:"riwayat_gaji_berkala_tanggal_sk"`
	RiwayatGajiBerkalaTmtKgb                      *string  `json:"riwayat_gaji_berkala_tmt_kgb"`
	RiwayatGajiBerkalaMasaKerjaTahun              *int64   `json:"riwayat_gaji_berkala_masa_kerja_tahun"`
	RiwayatGajiBerkalaMasaKerjaBulan              *int64   `json:"riwayat_gaji_berkala_masa_kerja_bulan"`
	RiwayatGajiBerkalaGajiPokok                   *int64   `json:"riwayat_gaji_berkala_gaji_pokok"`
	InpassingPangkatGolongan                      *string  `json:"inpassing_pangkat_golongan"`
	InpassingNoSkInpassing                        *string  `json:"inpassing_no_sk_inpassing"`
	InpassingTglSkInpassing                       *string  `json:"inpassing_tgl_sk_inpassing"`
	InpassingTmtInpassing                         *string  `json:"inpassing_tmt_inpassing"`
	InpassingAngkaKredit                          *float64 `json:"inpassing_angka_kredit"`
	InpassingMasaKerjaTahun                       *int64   `json:"inpassing_masa_kerja_tahun"`
	InpassingMasaKerjaBulan                       *int64   `json:"inpassing_masa_kerja_bulan"`
	RiwayatSertifikasiBidangStudi                 *string  `json:"riwayat_sertifikasi_bidang_studi"`
	RiwayatSertifikasiJenisSertifikasi            *string  `json:"riwayat_sertifikasi_jenis_sertifikasi"`
	RiwayatSertifikasiTahunSertifikasi            *string  `json:"riwayat_sertifikasi_tahun_sertifikasi"`
	RiwayatSertifikasiNomorSertifikat             *string  `json:"riwayat_sertifikasi_nomor_sertifikat"`
	RiwayatSertifikasiNrg                         *string  `json:"riwayat_sertifikasi_nrg"`
	RiwayatSertifikasiNomorPeserta                *string  `json:"riwayat_sertifikasi_nomor_peserta"`
	RiwayatPendidikanFormalBidangStudi            *string  `json:"riwayat_pendidikan_formal_bidang_studi"`
	RiwayatPendidikanFormalJenjangPendidikan      *string  `json:"riwayat_pendidikan_formal_jenjang_pendidikan"`
	RiwayatPendidikanFormalGelarAkademik          *string  `json:"riwayat_pendidikan_formal_gelar_akademik"`
	RiwayatPendidikanFormalSatuanPendidikanFormal *string  `json:"riwayat_pendidikan_formal_satuan_pendidikan_formal"`
	RiwayatPendidikanFormalFakultas               *string  `json:"riwayat_pendidikan_formal_fakultas"`
	RiwayatPendidikanFormalKependidikan           *string  `json:"riwayat_pendidikan_formal_kependidikan"`
	RiwayatPendidikanFormalTahunMasuk             *int64   `json:"riwayat_pendidikan_formal_tahun_masuk"`
	RiwayatPendidikanFormalTahunLulus             *int64   `json:"riwayat_pendidikan_formal_tahun_lulus"`
	RiwayatPendidikanFormalNim                    *string  `json:"riwayat_pendidikan_formal_nim"`
	RiwayatPendidikanFormalStatusKuliah           *string  `json:"riwayat_pendidikan_formal_status_kuliah"`
	RiwayatPendidikanFormalSemester               *int64   `json:"riwayat_pendidikan_formal_semester"`
	RiwayatPendidikanFormalIpk                    *float64 `json:"riwayat_pendidikan_formal_ipk"`
	JumlahAnak                                    *int64   `json:"jumlah_anak"`
	TugasTambahanJabatanPtk                       *string  `json:"tugas_tambahan_jabatan_ptk"`
	TugasTambahanSekolah                          *string  `json:"tugas_tambahan_sekolah"`
	TugasTambahanJumlahJam                        *int64   `json:"tugas_tambahan_jumlah_jam"`
	TugasTambahanNomorSk                          *string  `json:"tugas_tambahan_nomor_sk"`
	TugasTambahanTmtTambahan                      *string  `json:"tugas_tambahan_tmt_tambahan"`
	TugasTambahanTstTambahan                      *string  `json:"tugas_tambahan_tst_tambahan"`
	RiwayatStrukturalJabatanPtk                   *string  `json:"riwayat_struktural_jabatan_ptk"`
	RiwayatStrukturalSkStruktural                 *string  `json:"riwayat_struktural_sk_struktural"`
	RiwayatStrukturalTmtJabatan                   *string  `json:"riwayat_struktural_tmt_jabatan"`
	RiwayatFungsionalJabatanFungsional            *string  `json:"riwayat_fungsional_jabatan_fungsional"`
	RiwayatFungsionalSkJabfung                    *string  `json:"riwayat_fungsional_sk_jabfung"`
	RiwayatFungsionalTmtJabatan                   *string  `json:"riwayat_fungsional_tmt_jabatan"`
	CreateDate                                    *string  `json:"create_date"`
	LastUpdate                                    *string  `json:"last_update"`
	SoftDeletePtkTerdaftar                        *int64   `json:"soft_delete_ptk_terdaftar"`
	SoftDeleteRwyKepangkatan                      *int64   `json:"soft_delete_rwy_kepangkatan"`
	SoftDeleteRiwayatGajiBerkala                  *int64   `json:"soft_delete_riwayat_gaji_berkala"`
	SoftDeleteInpassing                           *int64   `json:"soft_delete_inpassing"`
	SoftDeleteRwySertifikasi                      *int64   `json:"soft_delete_rwy_sertifikasi"`
	SoftDeleteRwyPendFormal                       *int64   `json:"soft_delete_rwy_pend_formal"`
	SoftDeleteTugasTambahan                       *int64   `json:"soft_delete_tugas_tambahan"`
	SoftDeleteRwyStruktural                       *int64   `json:"soft_delete_rwy_struktural"`
	SoftDeleteRwyFungsional                       *int64   `json:"soft_delete_rwy_fungsional"`
}

func GetPtk(npsn, token string) ([]DapoPtk, error) {
	var res []DapoPtk

	client := http.Client{Timeout: 30 * time.Second}
	// req, err := client.Get(os.Getenv("URL_DAPO") + "/sekolah?kode_wilayah=" + wilayah)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	return res, errors.New("Gagal memuat data.")
	// }

	req, err := http.NewRequest("GET", os.Getenv("URL_DAPO")+"/ptk?npsn="+npsn, nil)
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
	// file, err := os.Create("ptk_" + npsn + ".json")
	// if err != nil {
	// 	return res, errors.New("Gagal membuat file ptk_" + npsn + ".json")
	// }
	// defer file.Close()

	// encoder := json.NewEncoder(file)
	// err = encoder.Encode(dataByte)
	// if err != nil {
	// 	println("Gagal menulis data ke file ptk_" + npsn + ".json")
	// 	return res, errors.New("Gagal menulis data ke file ptk_" + npsn + ".json")
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
		return res, errors.New("agal memuat data")
	}

	// datas = append(datas, responseObject)
	// fmt.Print(string(bodyBytes))
	//  return json.Unmarshal(bodyBytes, &responseObject)

	if len(res) > 0 {
		output, _ := json.MarshalIndent(res, "", "  ")

		// Tulis data ke dalam file
		err := ioutil.WriteFile("ptk_"+npsn+".json", output, 0644)
		if err != nil {
			log.Println(err.Error())
			return res, errors.New("gagal menulis data ke file ptk_" + npsn + ".json")
		}
		return res, nil
	} else {
		return res, errors.New("data tidak ditemukan")
	}

}

func InsertUpdatePtkDapodik(ptk DapoPtk) error {

	con := db.CreateCon()

	tx, _ := con.Begin()

	// var uuid = uuid.New()
	// loca, _ := time.LoadLocation("Asia/Jakarta")

	// currentTime := time.Now().In(loca)

	// tglLahir, errTanggalL := time.Parse("2006-01-02", *ptk.TglLahir)
	// if errTanggalL != nil {
	// 	fmt.Println(errTanggalL)
	// }

	sqlStatement := "INSERT INTO _ptk_tb_dapodik (ptk_terdaftar_id, ptk_id, semester_id, sekolah_id, nama, nip, jenis_kelamin, tempat_lahir, tanggal_lahir, nik, no_kk, niy_nigk, nuptk, nuks, status_kepegawaian, jenis_ptk_id, jenis_ptk, pengawas_bidang_studi, agama, kewarganegaraan, alamat_jalan, rt, rw, nama_dusun, desa_kelurahan, kode_kecamatan, kecamatan, kode_kabupaten, kabupaten, kode_provinsi, provinsi, kode_pos, lintang, bujur, no_telepon_rumah, no_hp, email, status_keaktifan, sk_cpns, tgl_cpns, sk_pengangkatan, tmt_pengangkatan, lembaga_pengangkat, pangkat_golongan, keahlian_laboratorium, sumber_gaji, nama_ibu_kandung, status_perkawinan, nama_suami_istri, nip_suami_istri, pekerjaan_suami_istri, tmt_pns, sudah_lisensi_kepala_sekolah, jumlah_sekolah_binaan, pernah_diklat_kepengawasan, nm_wp, status_data, karpeg, karpas, mampu_handle_kk, keahlian_braille, keahlian_bhs_isyarat, npwp, bank, rekening_bank, rekening_atas_nama, tahun_ajaran, nomor_surat_tugas, tanggal_surat_tugas, tmt_tugas, ptk_induk, jenis_keluar, tgl_ptk_keluar, riwayat_kepangkatan_pangkat_golongan, riwayat_kepangkatan_nomor_sk, riwayat_kepangkatan_tanggal_sk, riwayat_kepangkatan_tmt_pangkat, riwayat_kepangkatan_masa_kerja_gol_tahun, riwayat_kepangkatan_masa_kerja_gol_bulan, riwayat_gaji_berkala_pangkat_golongan, riwayat_gaji_berkala_nomor_sk, riwayat_gaji_berkala_tanggal_sk, riwayat_gaji_berkala_tmt_kgb, riwayat_gaji_berkala_masa_kerja_tahun, riwayat_gaji_berkala_masa_kerja_bulan, riwayat_gaji_berkala_gaji_pokok, inpassing_pangkat_golongan, inpassing_no_sk_inpassing, inpassing_tgl_sk_inpassing, inpassing_tmt_inpassing, inpassing_angka_kredit, inpassing_masa_kerja_tahun, inpassing_masa_kerja_bulan, riwayat_sertifikasi_bidang_studi, riwayat_sertifikasi_jenis_sertifikasi, riwayat_sertifikasi_tahun_sertifikasi, riwayat_sertifikasi_nomor_sertifikat, riwayat_sertifikasi_nrg, riwayat_sertifikasi_nomor_peserta, riwayat_pendidikan_formal_bidang_studi, riwayat_pendidikan_formal_jenjang_pendidikan, riwayat_pendidikan_formal_gelar_akademik, riwayat_pendidikan_formal_satuan_pendidikan_formal, riwayat_pendidikan_formal_fakultas, riwayat_pendidikan_formal_kependidikan, riwayat_pendidikan_formal_tahun_masuk, riwayat_pendidikan_formal_tahun_lulus, riwayat_pendidikan_formal_nim, riwayat_pendidikan_formal_status_kuliah, riwayat_pendidikan_formal_semester, riwayat_pendidikan_formal_ipk, jumlah_anak, tugas_tambahan_jabatan_ptk, tugas_tambahan_sekolah, tugas_tambahan_jumlah_jam, tugas_tambahan_nomor_sk, tugas_tambahan_tmt_tambahan, tugas_tambahan_tst_tambahan, riwayat_struktural_jabatan_ptk, riwayat_struktural_sk_struktural, riwayat_struktural_tmt_jabatan, riwayat_fungsional_jabatan_fungsional, riwayat_fungsional_sk_jabfung, riwayat_fungsional_tmt_jabatan, create_date, last_update, soft_delete_ptk_terdaftar, soft_delete_rwy_kepangkatan, soft_delete_riwayat_gaji_berkala, soft_delete_inpassing, soft_delete_rwy_sertifikasi, soft_delete_rwy_pend_formal, soft_delete_tugas_tambahan, soft_delete_rwy_struktural, soft_delete_rwy_fungsional) VALUES (?, ?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE ptk_id=values (ptk_id), semester_id=values (semester_id), sekolah_id=values (sekolah_id), nama=values (nama), nip=values (nip), jenis_kelamin=values (jenis_kelamin), tempat_lahir=values (tempat_lahir), tanggal_lahir=values (tanggal_lahir), nik=values (nik), no_kk=values (no_kk), niy_nigk=values (niy_nigk), nuptk=values (nuptk), nuks=values (nuks), status_kepegawaian=values (status_kepegawaian), jenis_ptk_id=values (jenis_ptk_id), jenis_ptk=values (jenis_ptk), pengawas_bidang_studi=values (pengawas_bidang_studi), agama=values (agama), kewarganegaraan=values (kewarganegaraan), alamat_jalan=values (alamat_jalan), rt=values (rt), rw=values (rw), nama_dusun=values (nama_dusun), desa_kelurahan=values (desa_kelurahan), kode_kecamatan=values (kode_kecamatan), kecamatan=values (kecamatan), kode_kabupaten=values (kode_kabupaten), kabupaten=values (kabupaten), kode_provinsi=values (kode_provinsi), provinsi=values (provinsi), kode_pos=values (kode_pos), lintang=values (lintang), bujur=values (bujur), no_telepon_rumah=values (no_telepon_rumah), no_hp=values (no_hp), email=values (email), status_keaktifan=values (status_keaktifan), sk_cpns=values (sk_cpns), tgl_cpns=values (tgl_cpns), sk_pengangkatan=values (sk_pengangkatan), tmt_pengangkatan=values (tmt_pengangkatan), pangkat_golongan=values (pangkat_golongan), keahlian_laboratorium=values (keahlian_laboratorium), sumber_gaji=values (sumber_gaji), nama_ibu_kandung=values (nama_ibu_kandung), status_perkawinan=values (status_perkawinan), nama_suami_istri=values (nama_suami_istri), nip_suami_istri=values (nip_suami_istri), pekerjaan_suami_istri=values (pekerjaan_suami_istri), tmt_pns=values (tmt_pns), sudah_lisensi_kepala_sekolah=values (sudah_lisensi_kepala_sekolah), jumlah_sekolah_binaan=values (jumlah_sekolah_binaan), pernah_diklat_kepengawasan=values (pernah_diklat_kepengawasan), nm_wp=values (nm_wp), status_data=values (status_data), karpeg=values (karpeg), karpas=values (karpas), mampu_handle_kk=values (mampu_handle_kk), keahlian_braille=values (keahlian_braille), keahlian_bhs_isyarat=values (keahlian_bhs_isyarat), npwp=values (npwp), bank=values (bank), rekening_bank=values (rekening_bank), rekening_atas_nama=values (rekening_atas_nama), tahun_ajaran=values (tahun_ajaran), nomor_surat_tugas=values (nomor_surat_tugas), tanggal_surat_tugas=values (tanggal_surat_tugas), tmt_tugas=values (tmt_tugas), ptk_induk=values (ptk_induk), jenis_keluar=values (jenis_keluar), tgl_ptk_keluar=values (tgl_ptk_keluar), riwayat_kepangkatan_pangkat_golongan=values (riwayat_kepangkatan_pangkat_golongan), riwayat_kepangkatan_nomor_sk=values (riwayat_kepangkatan_nomor_sk), riwayat_kepangkatan_tanggal_sk=values (riwayat_kepangkatan_tanggal_sk), riwayat_kepangkatan_tmt_pangkat=values (riwayat_kepangkatan_tmt_pangkat), riwayat_kepangkatan_masa_kerja_gol_tahun=values (riwayat_kepangkatan_masa_kerja_gol_tahun), riwayat_kepangkatan_masa_kerja_gol_bulan=values (riwayat_kepangkatan_masa_kerja_gol_bulan), riwayat_gaji_berkala_pangkat_golongan=values (riwayat_gaji_berkala_pangkat_golongan), riwayat_gaji_berkala_nomor_sk=values (riwayat_gaji_berkala_nomor_sk), riwayat_gaji_berkala_tanggal_sk=values (riwayat_gaji_berkala_tanggal_sk), riwayat_gaji_berkala_tmt_kgb=values (riwayat_gaji_berkala_tmt_kgb), riwayat_gaji_berkala_masa_kerja_tahun=values (riwayat_gaji_berkala_masa_kerja_tahun), riwayat_gaji_berkala_masa_kerja_bulan=values (riwayat_gaji_berkala_masa_kerja_bulan), riwayat_gaji_berkala_gaji_pokok=values (riwayat_gaji_berkala_gaji_pokok), inpassing_pangkat_golongan=values (inpassing_pangkat_golongan), inpassing_no_sk_inpassing=values (inpassing_no_sk_inpassing), inpassing_tgl_sk_inpassing=values (inpassing_tgl_sk_inpassing), inpassing_tmt_inpassing=values (inpassing_tmt_inpassing), inpassing_angka_kredit=values (inpassing_angka_kredit), inpassing_masa_kerja_tahun=values (inpassing_masa_kerja_tahun), inpassing_masa_kerja_bulan=values (inpassing_masa_kerja_bulan), riwayat_sertifikasi_bidang_studi=values (riwayat_sertifikasi_bidang_studi), riwayat_sertifikasi_jenis_sertifikasi=values (riwayat_sertifikasi_jenis_sertifikasi), riwayat_sertifikasi_tahun_sertifikasi=values (riwayat_sertifikasi_tahun_sertifikasi), riwayat_sertifikasi_nomor_sertifikat=values (riwayat_sertifikasi_nomor_sertifikat), riwayat_sertifikasi_nrg=values (riwayat_sertifikasi_nrg), riwayat_sertifikasi_nomor_peserta=values (riwayat_sertifikasi_nomor_peserta), riwayat_pendidikan_formal_bidang_studi=values (riwayat_pendidikan_formal_bidang_studi), riwayat_pendidikan_formal_jenjang_pendidikan=values (riwayat_pendidikan_formal_jenjang_pendidikan), riwayat_pendidikan_formal_gelar_akademik=values (riwayat_pendidikan_formal_gelar_akademik), riwayat_pendidikan_formal_satuan_pendidikan_formal=values (riwayat_pendidikan_formal_satuan_pendidikan_formal), riwayat_pendidikan_formal_fakultas=values (riwayat_pendidikan_formal_fakultas), riwayat_pendidikan_formal_kependidikan=values (riwayat_pendidikan_formal_kependidikan), riwayat_pendidikan_formal_tahun_masuk=values (riwayat_pendidikan_formal_tahun_masuk), riwayat_pendidikan_formal_tahun_lulus=values (riwayat_pendidikan_formal_tahun_lulus), riwayat_pendidikan_formal_nim=values (riwayat_pendidikan_formal_nim), riwayat_pendidikan_formal_status_kuliah=values (riwayat_pendidikan_formal_status_kuliah), riwayat_pendidikan_formal_semester=values (riwayat_pendidikan_formal_semester), riwayat_pendidikan_formal_ipk=values (riwayat_pendidikan_formal_ipk), jumlah_anak=values (jumlah_anak), tugas_tambahan_jabatan_ptk=values (tugas_tambahan_jabatan_ptk), tugas_tambahan_sekolah=values (tugas_tambahan_sekolah), tugas_tambahan_jumlah_jam=values (tugas_tambahan_jumlah_jam), tugas_tambahan_nomor_sk=values (tugas_tambahan_nomor_sk), tugas_tambahan_tmt_tambahan=values (tugas_tambahan_tmt_tambahan), tugas_tambahan_tst_tambahan=values (tugas_tambahan_tst_tambahan), riwayat_struktural_jabatan_ptk=values (riwayat_struktural_jabatan_ptk), riwayat_struktural_sk_struktural=values (riwayat_struktural_sk_struktural), riwayat_struktural_tmt_jabatan=values (riwayat_struktural_tmt_jabatan), riwayat_fungsional_jabatan_fungsional=values (riwayat_fungsional_jabatan_fungsional), riwayat_fungsional_sk_jabfung=values (riwayat_fungsional_sk_jabfung), riwayat_fungsional_tmt_jabatan=values (riwayat_fungsional_tmt_jabatan), create_date=values (create_date), last_update=values (last_update), soft_delete_ptk_terdaftar=values (soft_delete_ptk_terdaftar), soft_delete_rwy_kepangkatan=values (soft_delete_rwy_kepangkatan), soft_delete_riwayat_gaji_berkala=values (soft_delete_riwayat_gaji_berkala), soft_delete_inpassing=values (soft_delete_inpassing), soft_delete_rwy_sertifikasi=values (soft_delete_rwy_sertifikasi), soft_delete_rwy_pend_formal=values (soft_delete_rwy_pend_formal), soft_delete_tugas_tambahan=values (soft_delete_tugas_tambahan), soft_delete_rwy_struktural=values (soft_delete_rwy_struktural), soft_delete_rwy_fungsional=values (soft_delete_rwy_fungsional)"
	var kode_kec string
	if ptk.KodeKecamatan != nil {
		kode_kec = strings.TrimSpace(*ptk.KodeKecamatan)
	}
	var kode_kab string
	if ptk.KodeKabupaten != nil {
		kode_kab = strings.TrimSpace(*ptk.KodeKabupaten)
	}
	var kode_prov string
	if ptk.KodeProvinsi != nil {
		kode_prov = strings.TrimSpace(*ptk.KodeProvinsi)
	}
	exe, err := tx.Exec(sqlStatement, ptk.PtkTerdaftarID, ptk.PtkID, ptk.SemesterID, ptk.SekolahID, ptk.Nama, ptk.Nip, ptk.JenisKelamin, ptk.TempatLahir, ptk.TanggalLahir, ptk.Nik, ptk.NoKk, ptk.NiyNigk, ptk.Nuptk, ptk.Nuks, ptk.StatusKepegawaian, ptk.JenisPtkID, ptk.JenisPtk, ptk.PengawasBidangStudi, ptk.Agama, ptk.Kewarganegaraan, ptk.AlamatJalan, ptk.Rt, ptk.Rw, ptk.NamaDusun, ptk.DesaKelurahan, kode_kec, ptk.Kecamatan, kode_kab, ptk.Kabupaten, kode_prov, ptk.Provinsi, ptk.KodePos, ptk.Lintang, ptk.Bujur, ptk.NoTeleponRumah, ptk.NoHp, ptk.Email, ptk.StatusKeaktifan, ptk.SkCpns, ptk.TglCpns, ptk.SkPengangkatan, ptk.TmtPengangkatan, ptk.LembagaPengangkat, ptk.PangkatGolongan, ptk.KeahlianLaboratorium, ptk.SumberGaji, ptk.NamaIbuKandung, ptk.StatusPerkawinan, ptk.NamaSuamiIstri, ptk.NipSuamiIstri, ptk.PekerjaanSuamiIstri, ptk.TmtPns, ptk.SudahLisensiKepalaSekolah, ptk.JumlahSekolahBinaan, ptk.PernahDiklatKepengawasan, ptk.NmWp, ptk.StatusData, ptk.Karpeg, ptk.Karpas, ptk.MampuHandleKk, ptk.KeahlianBraille, ptk.KeahlianBhsIsyarat, ptk.Npwp, ptk.Bank, ptk.RekeningBank, ptk.RekeningAtasNama, ptk.TahunAjaran, ptk.NomorSuratTugas, ptk.TanggalSuratTugas, ptk.TmtTugas, ptk.PtkInduk, ptk.JenisKeluar, ptk.TglPtkKeluar, ptk.RiwayatKepangkatanPangkatGolongan, ptk.RiwayatKepangkatanNomorSk, ptk.RiwayatKepangkatanTanggalSk, ptk.RiwayatKepangkatanTmtPangkat, ptk.RiwayatKepangkatanMasaKerjaGolTahun, ptk.RiwayatKepangkatanMasaKerjaGolBulan, ptk.RiwayatGajiBerkalaPangkatGolongan, ptk.RiwayatGajiBerkalaNomorSk, ptk.RiwayatGajiBerkalaTanggalSk, ptk.RiwayatGajiBerkalaTmtKgb, ptk.RiwayatGajiBerkalaMasaKerjaTahun, ptk.RiwayatGajiBerkalaMasaKerjaBulan, ptk.RiwayatGajiBerkalaGajiPokok, ptk.InpassingPangkatGolongan, ptk.InpassingNoSkInpassing, ptk.InpassingTglSkInpassing, ptk.InpassingTmtInpassing, ptk.InpassingAngkaKredit, ptk.InpassingMasaKerjaTahun, ptk.InpassingMasaKerjaBulan, ptk.RiwayatSertifikasiBidangStudi, ptk.RiwayatSertifikasiJenisSertifikasi, ptk.RiwayatSertifikasiTahunSertifikasi, ptk.RiwayatSertifikasiNomorSertifikat, ptk.RiwayatSertifikasiNrg, ptk.RiwayatSertifikasiNomorPeserta, ptk.RiwayatPendidikanFormalBidangStudi, ptk.RiwayatPendidikanFormalJenjangPendidikan, ptk.RiwayatPendidikanFormalGelarAkademik, ptk.RiwayatPendidikanFormalSatuanPendidikanFormal, ptk.RiwayatPendidikanFormalFakultas, ptk.RiwayatPendidikanFormalKependidikan, ptk.RiwayatPendidikanFormalTahunMasuk, ptk.RiwayatPendidikanFormalTahunLulus, ptk.RiwayatPendidikanFormalNim, ptk.RiwayatPendidikanFormalStatusKuliah, ptk.RiwayatPendidikanFormalSemester, ptk.RiwayatPendidikanFormalIpk, ptk.JumlahAnak, ptk.TugasTambahanJabatanPtk, ptk.TugasTambahanSekolah, ptk.TugasTambahanJumlahJam, ptk.TugasTambahanNomorSk, ptk.TugasTambahanTmtTambahan, ptk.TugasTambahanTstTambahan, ptk.RiwayatStrukturalJabatanPtk, ptk.RiwayatStrukturalSkStruktural, ptk.RiwayatStrukturalTmtJabatan, ptk.RiwayatFungsionalJabatanFungsional, ptk.RiwayatFungsionalSkJabfung, ptk.RiwayatFungsionalTmtJabatan, ptk.CreateDate, ptk.LastUpdate, ptk.SoftDeletePtkTerdaftar, ptk.SoftDeleteRwyKepangkatan, ptk.SoftDeleteRiwayatGajiBerkala, ptk.SoftDeleteInpassing, ptk.SoftDeleteRwySertifikasi, ptk.SoftDeleteRwyPendFormal, ptk.SoftDeleteTugasTambahan, ptk.SoftDeleteRwyStruktural, ptk.SoftDeleteRwyFungsional)
	// exe, err := tx.Exec(sqlStatement, strings.TrimSpace(*ptk.Id), strings.TrimSpace(*ptk.Id_ptk), strings.TrimSpace(*ptk.Email), *ptk.Nama, strings.TrimSpace(*ptk.Nik), strings.TrimSpace(*ptk.Nuptk), strings.TrimSpace(*ptk.Nip), strings.TrimSpace(*ptk.Nrg), strings.TrimSpace(*ptk.NoPeserta), *ptk.Npwp, *ptk.NoRekening, *ptk.CabangBank, *ptk.JenisKelamin, *ptk.TempatLahir, tglLahir, *ptk.StatusTugas, *ptk.TempatTugas, strings.TrimSpace(*ptk.Npsn), *ptk.Kecamatan, strings.TrimSpace(*ptk.IDKecamatan), strings.TrimSpace(*ptk.NoHp), *ptk.SkCpns, *ptk.TglCpns, *ptk.SkPengangkatan, *ptk.TmtPengangkatan, *ptk.JenisPtk, *ptk.Pendidikan, *ptk.BidangStudiPendidikan, *ptk.BidangStudiSertifikasi, *ptk.StatusKepegawaian, *ptk.MapelDiajarkan, *ptk.JamMengajarPerminggu, *ptk.JabatanKepsek, *ptk.PangkatGolongan, *ptk.NomorSkPangkat, *ptk.TglSkPangkat, *ptk.TmtPangkat, *ptk.MasaKerjaTahun, *ptk.MasaKerjaBulan, *ptk.GajiPokok, *ptk.SkKgb, *ptk.TglSkKgb, *ptk.TmtSkKgb, *ptk.MasaKerjaTahunKgb, *ptk.MasaKerjaBulanKgb, *ptk.GajiPokokKgb, *ptk.MengajarLainSatmikal, *ptk.NomorSkImpassing, *ptk.TglSkImpassing, *ptk.TmtSkImpassing, *ptk.JabatanAngkaKredit, *ptk.PangkatGolonganRuang, *ptk.MasaKerjaTahunImpassing, *ptk.MasaKerjaBulanImpassing, *ptk.JumlahTunjanganPokokImpassing, *ptk.LampiranImpassing, *ptk.LampiranFoto, *ptk.LampiranKarpeg, *ptk.LampiranKgb, *ptk.LampiranKtp, *ptk.LampiranNrg, *ptk.LampiranNuptk, *ptk.LampiranPangkat, *ptk.LampiranPernyataanIndividu, *ptk.IsCuti, *ptk.LampiranCuti, *ptk.IsPensiun, *ptk.LampiranPensiun, *ptk.LampiranSerdik, *ptk.LampiranKeteranganTambahanJam, *ptk.LampiranNpwp, *ptk.LampiranBukuRekening, *ptk.LampiranIjazah, *ptk.JenisTunjangan, *ptk.IsLocked, currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))

	if err != nil {
		// fmt.Println("Query error")
		fmt.Println(err.Error())
		tx.Rollback()
		return errors.New("gagal Insert Update PTK DAPO 1")
	}

	rows, err := exe.RowsAffected()
	if err != nil {
		// fmt.Println(err.Error())
		tx.Rollback()
		return errors.New("gagal insert update PTK DAPO 2")
	}

	if rows > 0 {
		tx.Commit()

		return nil
	}
	tx.Rollback()
	return errors.New("gagal insert update PTK DAPO 3")
}

func GetPtkBackbone(id_sekolah string) ([]DapoPtk, error) {
	var res []DapoPtk
	var obj DapoPtk

	con := db.CreateCon()

	sqlStatement := "SELECT * From _ptk_tb_dapodik WHERE sekolah_id = ?"

	respo, err := con.Query(sqlStatement, id_sekolah)
	if err != nil {
		return res, errors.New("gagal memuat data")
	}
	defer respo.Close()

	for respo.Next() {
		err = respo.Scan(
			&obj.PtkTerdaftarID,
			&obj.PtkID,
			&obj.SemesterID,
			&obj.SekolahID,
			&obj.Nama,
			&obj.Nip,
			&obj.JenisKelamin,
			&obj.TempatLahir,
			&obj.TanggalLahir,
			&obj.Nik,
			&obj.NoKk,
			&obj.NiyNigk,
			&obj.Nuptk,
			&obj.Nuks,
			&obj.StatusKepegawaian,
			&obj.JenisPtkID,
			&obj.JenisPtk,
			&obj.PengawasBidangStudi,
			&obj.Agama,
			&obj.Kewarganegaraan,
			&obj.AlamatJalan,
			&obj.Rt,
			&obj.Rw,
			&obj.NamaDusun,
			&obj.DesaKelurahan,
			&obj.KodeKecamatan,
			&obj.Kecamatan,
			&obj.KodeKabupaten,
			&obj.Kabupaten,
			&obj.KodeProvinsi,
			&obj.Provinsi,
			&obj.KodePos,
			&obj.Lintang,
			&obj.Bujur,
			&obj.NoTeleponRumah,
			&obj.NoHp,
			&obj.Email,
			&obj.StatusKeaktifan,
			&obj.SkCpns,
			&obj.TglCpns,
			&obj.SkPengangkatan,
			&obj.TmtPengangkatan,
			&obj.LembagaPengangkat,
			&obj.PangkatGolongan,
			&obj.KeahlianLaboratorium,
			&obj.SumberGaji,
			&obj.NamaIbuKandung,
			&obj.StatusPerkawinan,
			&obj.NamaSuamiIstri,
			&obj.NipSuamiIstri,
			&obj.PekerjaanSuamiIstri,
			&obj.TmtPns,
			&obj.SudahLisensiKepalaSekolah,
			&obj.JumlahSekolahBinaan,
			&obj.PernahDiklatKepengawasan,
			&obj.NmWp,
			&obj.StatusData,
			&obj.Karpeg,
			&obj.Karpas,
			&obj.MampuHandleKk,
			&obj.KeahlianBraille,
			&obj.KeahlianBhsIsyarat,
			&obj.Npwp,
			&obj.Bank,
			&obj.RekeningBank,
			&obj.RekeningAtasNama,
			&obj.TahunAjaran,
			&obj.NomorSuratTugas,
			&obj.TanggalSuratTugas,
			&obj.TmtTugas,
			&obj.PtkInduk,
			&obj.JenisKeluar,
			&obj.TglPtkKeluar,
			&obj.RiwayatKepangkatanPangkatGolongan,
			&obj.RiwayatKepangkatanNomorSk,
			&obj.RiwayatKepangkatanTanggalSk,
			&obj.RiwayatKepangkatanTmtPangkat,
			&obj.RiwayatKepangkatanMasaKerjaGolTahun,
			&obj.RiwayatKepangkatanMasaKerjaGolBulan,
			&obj.RiwayatGajiBerkalaPangkatGolongan,
			&obj.RiwayatGajiBerkalaNomorSk,
			&obj.RiwayatGajiBerkalaTanggalSk,
			&obj.RiwayatGajiBerkalaTmtKgb,
			&obj.RiwayatGajiBerkalaMasaKerjaTahun,
			&obj.RiwayatGajiBerkalaMasaKerjaBulan,
			&obj.RiwayatGajiBerkalaGajiPokok,
			&obj.InpassingPangkatGolongan,
			&obj.InpassingNoSkInpassing,
			&obj.InpassingTglSkInpassing,
			&obj.InpassingTmtInpassing,
			&obj.InpassingAngkaKredit,
			&obj.InpassingMasaKerjaTahun,
			&obj.InpassingMasaKerjaBulan,
			&obj.RiwayatSertifikasiBidangStudi,
			&obj.RiwayatSertifikasiJenisSertifikasi,
			&obj.RiwayatSertifikasiTahunSertifikasi,
			&obj.RiwayatSertifikasiNomorSertifikat,
			&obj.RiwayatSertifikasiNrg,
			&obj.RiwayatSertifikasiNomorPeserta,
			&obj.RiwayatPendidikanFormalBidangStudi,
			&obj.RiwayatPendidikanFormalJenjangPendidikan,
			&obj.RiwayatPendidikanFormalGelarAkademik,
			&obj.RiwayatPendidikanFormalSatuanPendidikanFormal,
			&obj.RiwayatPendidikanFormalFakultas,
			&obj.RiwayatPendidikanFormalKependidikan,
			&obj.RiwayatPendidikanFormalTahunMasuk,
			&obj.RiwayatPendidikanFormalTahunLulus,
			&obj.RiwayatPendidikanFormalNim,
			&obj.RiwayatPendidikanFormalStatusKuliah,
			&obj.RiwayatPendidikanFormalSemester,
			&obj.RiwayatPendidikanFormalIpk,
			&obj.JumlahAnak,
			&obj.TugasTambahanJabatanPtk,
			&obj.TugasTambahanSekolah,
			&obj.TugasTambahanJumlahJam,
			&obj.TugasTambahanNomorSk,
			&obj.TugasTambahanTmtTambahan,
			&obj.TugasTambahanTstTambahan,
			&obj.RiwayatStrukturalJabatanPtk,
			&obj.RiwayatStrukturalSkStruktural,
			&obj.RiwayatStrukturalTmtJabatan,
			&obj.RiwayatFungsionalJabatanFungsional,
			&obj.RiwayatFungsionalSkJabfung,
			&obj.RiwayatFungsionalTmtJabatan,
			&obj.CreateDate,
			&obj.LastUpdate,
			&obj.SoftDeletePtkTerdaftar,
			&obj.SoftDeleteRwyKepangkatan,
			&obj.SoftDeleteRiwayatGajiBerkala,
			&obj.SoftDeleteInpassing,
			&obj.SoftDeleteRwySertifikasi,
			&obj.SoftDeleteRwyPendFormal,
			&obj.SoftDeleteTugasTambahan,
			&obj.SoftDeleteRwyStruktural,
			&obj.SoftDeleteRwyFungsional,
		)

		if err != nil {
			return res, errors.New("gagal mengambil data")
		}

		res = append(res, obj)
	}

	return res, nil
}
