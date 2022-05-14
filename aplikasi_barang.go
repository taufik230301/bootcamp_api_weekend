package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Iipkoko@34"
	dbname   = "aplikasi_barang"
)

type M map[string]interface{}

type barang struct {
	Id_Barang     string `json:"p_id_barang"`
	Nama_Barang   string `json:"p_nama_barang"`
	Stok          int    `json:"p_stok"`
	Harga_Barang  int    `json:"p_harga_barang"`
	Tanggal_Masuk string `json:"p_tanggal_masuk"`
	Status_Diskon string `json:"p_status_diskon"`
}

type barang_array struct {
	barangs []barang `json:"barang_data"`
}

func main() {

	r := echo.New()

	r.GET("/test", func(ctx echo.Context) error {
		data := "Hello World"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/barang/getData", func(ctx echo.Context) error {
		psqlcom := fmt.Sprintf("host=%s port=%d user =%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlcom)

		if err != nil {
			data := M{"pesan": err.Error(), "Status": 400}
			return ctx.JSON(http.StatusOK, data)
		}
		defer db.Close()
		readStmt := "select * from read_all_barang()"
		rows, error := db.Query(readStmt)

		if error != nil {
			panic(error)
		}

		defer rows.Close()

		barang_result := barang_array{}

		for rows.Next() {
			barang_data := barang{}

			error2 := rows.Scan(&barang_data.Id_Barang, &barang_data.Nama_Barang, &barang_data.Stok, &barang_data.Harga_Barang, &barang_data.Tanggal_Masuk, &barang_data.Status_Diskon)

			if error2 != nil {
				panic(error2)
			}

			barang_result.barangs = append(barang_result.barangs, barang_data)

		}
		data := M{"Data": barang_result.barangs, "pesan": "Berhasil", "Status": 200}
		return ctx.JSON(http.StatusOK, data)
	})

	r.Start(":8888")
}
