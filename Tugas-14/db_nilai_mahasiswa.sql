CREATE TABLE `db_mahasiswa`.`nilai_mahasiswa`
(
    `id`           INT NOT NULL AUTO_INCREMENT,
    `nama`         VARCHAR(50) NULL,
    `mata_kuliah`  VARCHAR(50) NULL,
    `indeks_nilai` VARCHAR(1) NULL,
    `nilai`        INT(3) NULL,
    `created_at`   DATETIME NULL,
    `updated_at`   DATETIME NULL,
    PRIMARY KEY (`id`)
);
