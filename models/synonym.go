package models

import "github.com/oezg/keymatch-api/db"

type Synonym struct {
	Id      int64
	Domain  string `binding:"required"`
	Url     string `binding:"required"`
	Caption string `binding:"required"`
}

func (synonym *Synonym) Save() error {
	query := `
	INSERT INTO synonym(domain, url, caption)
	VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(synonym.Domain, synonym.Url, synonym.Caption)
	if err != nil {
		return err
	}

	synonym.Id, err = result.LastInsertId()
	return err
}

func GetSynonyms() ([]Synonym, error) {
	query := "SELECT * FROM synonym"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var synonyms []Synonym
	for rows.Next() {
		var synonym Synonym
		err = rows.Scan(&synonym.Id, &synonym.Domain, &synonym.Url, &synonym.Caption)
		if err != nil {
			return nil, err
		}
		synonyms = append(synonyms, synonym)
	}
	return synonyms, nil
}

func GetSynonym(synonym_id int64) (Synonym, error) {
	query := "SELECT domain, url, caption FROM synonym WHERE id = ?"
	row := db.DB.QueryRow(query, synonym_id)

	synonym := Synonym{Id: synonym_id}
	err := row.Scan(&synonym.Domain, &synonym.Url, &synonym.Caption)
	return synonym, err
}
