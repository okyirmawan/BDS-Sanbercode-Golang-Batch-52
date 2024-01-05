package nilai

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
	table          = "nilai"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Nilai, error) {
	var ListNilai []models.Nilai
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
		var nilai models.Nilai
		if err = rowQuery.Scan(&nilai.ID,
			&nilai.Indeks,
			&nilai.Skor,
			&nilai.MahasiswaID,
			&nilai.MataKuliahID); err != nil {
			return nil, err
		}

		ListNilai = append(ListNilai, nilai)
	}
	return ListNilai, nil
}

// Insert Nilai Mahasiswa
func Insert(ctx context.Context, nilai models.Nilai) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (indeks, skor, mahasiswa_id, mata_kuliah_id) values('%v',%v,%v,%v)", table,
		nilai.Indeks,
		nilai.Skor,
		nilai.MahasiswaID,
		nilai.MataKuliahID)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Nilai
func Update(ctx context.Context, nilai models.Nilai, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set indeks ='%s', skor =%d, mahasiswa_id ='%d', mata_kuliah_id ='%d' where id = %s",
		table,
		nilai.Indeks,
		nilai.Skor,
		nilai.MahasiswaID,
		nilai.MataKuliahID,
		id,
	)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

// Delete Nilai
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
