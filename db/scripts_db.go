package db

import (
	"database/sql"
	"fmt"
)

func CreateTable(db *sql.DB, script string) error {

	_, err := db.Exec(script)
	if err != nil {

		return fmt.Errorf("não foi possivel criar a tabela: %w", err)

	}

	return nil

}

func ScriptsCreateTable() []string {

	scripts := []string{
		`CREATE TABLE IF NOT EXISTS type_pests(
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL
    )`,

		`CREATE TABLE IF NOT EXISTS pests(
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        type_of_pest_id INT REFERENCES type_pests(id) ON DELETE CASCADE
    )`,

		`CREATE TABLE IF NOT EXISTS soil_type(
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        description TEXT NOT NULL
    )`,

		`CREATE TABLE IF NOT EXISTS agriculture_culture(
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        cientific_name VARCHAR(255) NOT NULL,
        soil_type_id INT REFERENCES soil_type(id) ON DELETE CASCADE,
        ph_ideal_soil DECIMAL(2,2) NOT NULL,
        max_temperature DECIMAL(2,2) NOT NULL,
        min_temperature DECIMAL(2,2) NOT NULL,
        excellent_temperature DECIMAL(2,2) NOT NULL,
        weekly_water_requirement_max DECIMAL(2,2) NOT NULL,
        weeklly_water_requirement_min DECIMAL(2,2) NOT NULL,
        sunlight_requirement INT NOT NULL
        
    )`,

		`CREATE TABLE IF NOT EXISTS irrigation_recomended(
        id SERIAL PRIMARY KEY,
        phenological_phase VARCHAR(50) NOT NULL,
        phase_duration_days INT NOT NULL,
        irrigation_min DECIMAL(4,2) NOT NULL,
        irrigation_max DECIMAL(4,2) NOT NULL,
        unit VARCHAR(50) NOT NULL,
        description TEXT NOT NULL
    )`,

		`CREATE TABLE IF NOT EXISTS agriculture_culture_irrigation(
        id SERIAL PRIMARY KEY,
        id_agriculture_culture INT REFERENCES agriculture_culture(id) ON DELETE CASCADE,
        id_irrigation_recomended INT REFERENCES irrigation_recomended(id) ON DELETE CASCADE,
        UNIQUE (id_agriculture_culture, id_irrigation_recomended)
    )`,

		`CREATE TABLE IF NOT EXISTS sustainable_pest_control(
        id SERIAL PRIMARY KEY,
        description TEXT NOT NULL 
    )`,

		`CREATE TABLE IF NOT EXISTS pests_agriculture_culture(
        id SERIAL PRIMARY KEY,
        id_agriculture_culture INT REFERENCES agriculture_culture(id) ON DELETE CASCADE,
        id_pests INT REFERENCES pests(id) ON DELETE CASCADE,
        description TEXT NOT NULL
    )`,

		`CREATE TABLE IF NOT EXISTS agriculture_culture_pest_method(
        id SERIAL PRIMARY KEY,
        id_agriculture_culture INT REFERENCES agriculture_culture(id) ON DELETE CASCADE,
        id_pest  INT REFERENCES pests(id) ON DELETE CASCADE,
        id_sustainable_pest_control INT REFERENCE sustainable_pest_control(id) ON DELETE CASCADE
    )`,
	}

	return scripts
}

func VerifyIdExists(id uint, db *sql.DB, tableName string) (bool, error) {

	var exists *bool
	queryVerifyIdExists := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s.id = $1)", tableName, tableName)

	err := db.QueryRow(queryVerifyIdExists, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("erro ao verificar a existência do id")
	}

	return true, nil
}

func QueryVerifyIdExists(params ...string) string {

	var queryVerifyIdExists string = fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE %s.%s = ? AND %s.%s = ?)`, params[0], params[1],
		params[2], params[3], params[4])

	return queryVerifyIdExists
}
