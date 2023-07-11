package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type booking struct {
	ID      uint   `json:"id_booking"`
	Nama    string `json:"nama"`
	Tanggal string `json:"tanggal"`
	Email   string `json:"email"`
}

type pembayaran struct {
	ID      uint   `json:"id_pembayaran"`
	Nama    string `json:"nama"`
	Tanggal string `json:"tanggal"`
	Harga   string `json:"harga"`
}

type pendaftaran struct {
	ID            uint   `json:"id_pendaftaran"`
	Nama          string `json:"nama"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	Telepon       string `json:"telepon"`
	Alamat        string `json:"alamat"`
}

type user struct {
	ID       uint   `json:"id_user"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	// database connection
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dbs_vapi")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	// database connection

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Service API!")
	})

	// Mahasiswa
	e.GET("/tbl_booking", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM tbl_booking")

		defer res.Close()

		if err != nil {
			log.Fatal(err)
		}
		var mahasiswa []booking
		for res.Next() {
			var m booking
			_ = res.Scan(&m.ID, &m.Nama, &m.Tanggal, &m.Email)
			mahasiswa = append(mahasiswa, m)
		}

		return c.JSON(http.StatusOK, mahasiswa)
	})

	e.POST("/tbl_booking", func(c echo.Context) error {
		var mahasiswa booking
		c.Bind(&mahasiswa)

		sqlStatement := "INSERT INTO tbl_booking (id_booking,nama, tanggal,email)VALUES (?,?, ?, ?)"
		res, err := db.Query(sqlStatement, mahasiswa.ID, mahasiswa.Nama, mahasiswa.Tanggal, mahasiswa.Email)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.PUT("/tbl_booking/:id_booking", func(c echo.Context) error {
		var mahasiswa booking
		c.Bind(&mahasiswa)

		sqlStatement := "UPDATE tbl_booking SET nama=?,tanggal=?,email=? WHERE id_booking=?"
		res, err := db.Query(sqlStatement, mahasiswa.Nama, mahasiswa.Tanggal, mahasiswa.Email, c.Param("id"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.DELETE("/tbl_booking/:id_booking", func(c echo.Context) error {
		var mahasiswa pembayaran
		c.Bind(&mahasiswa)

		sqlStatement := "DELETE FROM tbl_booking WHERE id_booking=?"
		res, err := db.Query(sqlStatement, c.Param("id"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})
	// Mahasiswa

	// Mahasiswa
	e.GET("/tbl_pembayaran", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM tbl_pembayaran")

		defer res.Close()

		if err != nil {
			log.Fatal(err)
		}
		var mahasiswa []pembayaran
		for res.Next() {
			var m pembayaran
			_ = res.Scan(&m.ID, &m.Nama, &m.Tanggal, &m.Harga)
			mahasiswa = append(mahasiswa, m)
		}

		return c.JSON(http.StatusOK, mahasiswa)
	})

	e.POST("/tbl_pembayaran", func(c echo.Context) error {
		var mahasiswa pembayaran
		c.Bind(&mahasiswa)

		sqlStatement := "INSERT INTO tbl_pembayaran (id_pembayaran,nama, tanggal,alamat)VALUES (?,?, ?, ?)"
		res, err := db.Query(sqlStatement, mahasiswa.ID, mahasiswa.Nama, mahasiswa.Tanggal, mahasiswa.Harga)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.PUT("/tbl_pembayaran/:id_pembayaran", func(c echo.Context) error {
		var mahasiswa pembayaran
		c.Bind(&mahasiswa)

		sqlStatement := "UPDATE tbl_pembayaran SET nama=?,tanggal=?,harga=? WHERE id_pembayaran=?"
		res, err := db.Query(sqlStatement, mahasiswa.Nama, mahasiswa.Tanggal, mahasiswa.Harga, c.Param("id_pembayaran"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.DELETE("/tbl_pembayaran/:id_pembayaran", func(c echo.Context) error {
		var mahasiswa pembayaran
		c.Bind(&mahasiswa)

		sqlStatement := "DELETE FROM tbl_pembayaran WHERE id_pembayaran=?"
		res, err := db.Query(sqlStatement, c.Param("id_pembayaran"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})
	// pembayaran

	// pendaftaran
	e.GET("/tbl_pendaftaran", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM tbl_pendaftaran")

		defer res.Close()

		if err != nil {
			log.Fatal(err)
		}
		var mahasiswa []pendaftaran
		for res.Next() {
			var m pendaftaran
			_ = res.Scan(&m.ID, &m.Nama, &m.Tanggal_lahir, &m.Telepon, &m.Alamat)
			mahasiswa = append(mahasiswa, m)
		}

		return c.JSON(http.StatusOK, mahasiswa)
	})

	e.POST("/tbl_pendaftaran", func(c echo.Context) error {
		var mahasiswa pendaftaran
		c.Bind(&mahasiswa)

		sqlStatement := "INSERT INTO tbl_pendaftaran (id_pendaftaran,nama, tanggal_lahir, telepon,alamat)VALUES (?,?, ?, ?)"
		res, err := db.Query(sqlStatement, mahasiswa.ID, mahasiswa.Nama, mahasiswa.Tanggal_lahir, mahasiswa.Telepon, mahasiswa.Alamat)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.PUT("/tbl_pendaftaran/:id_pendaftaran", func(c echo.Context) error {
		var mahasiswa pendaftaran
		c.Bind(&mahasiswa)

		sqlStatement := "UPDATE tbl_pendaftaran SET nama=?,tanggal_lahir=?, telepon=?,alamat=? WHERE id_pendaftaran=?"
		res, err := db.Query(sqlStatement, mahasiswa.ID, mahasiswa.Nama, mahasiswa.Tanggal_lahir, mahasiswa.Telepon, mahasiswa.Alamat, c.Param("id"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.DELETE("/tbl_pendaftaran/:id_pendaftaran", func(c echo.Context) error {
		var mahasiswa pendaftaran
		c.Bind(&mahasiswa)

		sqlStatement := "DELETE FROM tbl_pendaftaran WHERE id_pendaftaran=?"
		res, err := db.Query(sqlStatement, c.Param("id_pendaftaran"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})
	// pendaftaran

	// user
	e.GET("/tbl_user", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM tbl_user")

		defer res.Close()

		if err != nil {
			log.Fatal(err)
		}
		var mahasiswa []user
		for res.Next() {
			var m user
			_ = res.Scan(&m.ID, &m.Username, &m.Password)
			mahasiswa = append(mahasiswa, m)
		}

		return c.JSON(http.StatusOK, mahasiswa)
	})

	e.POST("/tbl_user", func(c echo.Context) error {
		var mahasiswa user
		c.Bind(&mahasiswa)

		sqlStatement := "INSERT INTO tbl_user (id_user,username,password)VALUES (?,?, ?, ?)"
		res, err := db.Query(sqlStatement, mahasiswa.ID, mahasiswa.Username, mahasiswa.Password)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.PUT("/tbl_user/:id_user", func(c echo.Context) error {
		var mahasiswa user
		c.Bind(&mahasiswa)

		sqlStatement := "UPDATE tbl_user SET username=?,password=? WHERE id_user=?"
		res, err := db.Query(sqlStatement, mahasiswa.ID, mahasiswa.Username, mahasiswa.Password, c.Param("id_user"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.DELETE("/tbl_user/:id_user", func(c echo.Context) error {
		var mahasiswa user
		c.Bind(&mahasiswa)

		sqlStatement := "DELETE FROM tbl_user WHERE id_user=?"
		res, err := db.Query(sqlStatement, c.Param("id_user"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})
	// user

	e.Logger.Fatal(e.Start(":1323"))

}
