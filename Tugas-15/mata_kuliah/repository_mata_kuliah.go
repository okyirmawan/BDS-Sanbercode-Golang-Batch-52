package mata_kuliah

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-15/config"
	"github.com/okyirmawan/BDS-Sanbercode-Golang-Batch-52/Tugas-15/models"
	"log"
)

const (
	table          = "mata_kuliah"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.MataKuliah, error) {
	var ListNilai []models.MataKuliah
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var nilai models.MataKuliah
		if err = rowQuery.Scan(&nilai.ID,
			&nilai.Nama); err != nil {
			return nil, err
		}

		ListNilai = append(ListNilai, nilai)
	}
	return ListNilai, nil
}

// Insert Mata Kuliah
func Insert(ctx context.Context, nilai models.MataKuliah) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nama) values('%v')", table,
		nilai.Nama)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Mata Kuliah
func Update(ctx context.Context, nilai models.MataKuliah, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nama ='%s' where id = %s",
		table,
		nilai.Nama,
		id,
	)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

// Delete Mata Kuliah
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
