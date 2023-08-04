package models

// https://www.calhoun.io/inserting-records-into-a-postgresql-database-with-gos-database-sql-package/
import (
	"fmt"
	"strings"

	"go.importar.dados.mssql/db"
)

type MunicipioIbge struct {
	MunCodigoIbge int    `json:codigo`
	MunNomeIbge   string `json:nome`
	MunUfIbge     string `json:uf`
}

func BuscaTodosOsMinicipios() ([]MunicipioIbge, error) {

	db := db.ConectaComSis()

	rows, err := db.Query("select mun_codigo_ibge, mun_nome_ibge, mun_uf_ibge from municipio_ibge order by mun_codigo_ibge asc")

	if err != nil {
		panic(err.Error())
		//return nil, err
	}

	defer rows.Close()

	municipios := []MunicipioIbge{}

	for rows.Next() {
		p := MunicipioIbge{}
		err := rows.Scan(&p.MunCodigoIbge, &p.MunNomeIbge, &p.MunUfIbge)
		municipios = append(municipios, p)

		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return nil, err
		}
	}

	//defer db.Close()

	return municipios, nil
}

func CreateMinicipiosVm(s *MunicipioIbge) (int64, error) {

	db := db.ConectaComVm()

	q := `
		SET QUOTED_IDENTIFIER OFF;
		INSERT INTO municipio_ibge (mun_codigo_ibge, mun_nome_ibge, mun_uf_ibge)
			VALUES ('%d','%s','%s');
		SET QUOTED_IDENTIFIER ON;
		select ID = convert(bigint, SCOPE_IDENTITY());`

	sql := fmt.Sprintf(q, s.MunCodigoIbge, strings.Replace(s.MunNomeIbge, "'", " ", -1), s.MunUfIbge)

	result, err := db.Query(sql)

	if err != nil {
		fmt.Println("Error inserting new row: " + err.Error())
		return -1, err
	}

	var lastInsertId1 int64
	for result.Next() {
		result.Scan(&lastInsertId1)
		//log.Printf("LastInsertId from SCOPE_IDENTITY(): %d\n", lastInsertId1)
	}
	return lastInsertId1, nil
}
