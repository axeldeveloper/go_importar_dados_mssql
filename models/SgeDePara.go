package models

import (
	"database/sql"
	"fmt"
	db "github.com/axeldeveloper/go.importar.dados.mssql/db"
)

type SgeDePara struct {
	
}

func BuscaTodosOsSgeDePara() ([]SgeDePara, error) {

	db := db.ConectaComSisged();

	q := `SELECT EmpCod,
			Cod,
			NomeRh,
			EstruturaIDSge,
			ColecaoIDSge,
			RegistroEstruturaID,
			NomeSge,
			CodHierarquia,
			CodConsist,
			EstruturaUems
		FROM SgeDePara`
	
	
	sql := fmt.Sprintf(q)
	
	rows, err := db.Query(sql)
	
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return nil, err
	}
	
	defer rows.Close()

	count := 0
	
	list := []SgeDePara{}

	for rows.Next() {
		m := SgeDePara{}
		err := rows.Scan(&m.EmpCod, 
			&m.Cod, 
			&m.NomeRh, 
			&m.EstruturaIDSge,
			&m.ColecaoIDSge,
			&m.RegistroEstruturaID,
			&m.NomeSge,
			&m.CodHierarquia,
			&m.CodConsist, 
			&m.EstruturaUems)

			list = append(list, m);

			if err != nil {
				fmt.Println("Error reading rows: " + err.Error())
				return nil, err
			}	
		count++
	}

	//defer rows.Close()
	return list, nil
}

// ExecuteAggregateStatement output summary of prices
func ExecuteAggregateStatement() {
	
	db := db.ConectaComSisged();
	
	result, err := db.Prepare("SELECT SUM(Price) as sum FROM Table_with_5M_rows")
	if err != nil {
		fmt.Println("Error preparing query: " + err.Error())
	}

	row := result.QueryRow()
	var sum string
	err = row.Scan(&sum)
	fmt.Printf("Sum: %s\n", sum)
}

// Create an SgeDeParaVm
func CreateSgeDeParaVm() (int64, error) {
	
	db := db.ConectaComSis();

	s := SgeDePara{}
	q := `INSERT INTO SgeDePara
					(EmpCod
					,Cod
					,NomeRh
					,EstruturaIDSge
					,ColecaoIDSge
					,RegistroEstruturaID
					,NomeSge
					,CodHierarquia
					,CodConsist
					,EstruturaUems)
					VALUES ('%d','%d','%s','%d', '%d','%d','%s','%s','%s','%d') 
					select ID = convert(bigint, SCOPE_IDENTITY());`
					
	sql := fmt.Sprintf(q,
					s.EmpCod.Int32, 
					s.Cod.Int32,
					s.NomeRh.String, 
					s.EstruturaIDSge.Int32,
					s.ColecaoIDSge.Int32,
					s.RegistroEstruturaID.Int32,
					s.NomeSge.String, 
					s.CodHierarquia.String, 
					s.CodConsist.String, 
					s.EstruturaUems.Int32)

					
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

// Read all ReadSgeDePara
func ReadSgeDePara() (int, error) {
	db := db.ConectaComSis();
	
	q := `SELECT EmpCod,
			Cod,
			NomeRh,
			EstruturaIDSge,
			ColecaoIDSge,
			RegistroEstruturaID,
			NomeSge,
			CodHierarquia,
			CodConsist,
			EstruturaUems
		FROM SgeDePara`
	
	sql := fmt.Sprintf(q)
	
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return -1, err
	}
	defer rows.Close()
	count := 0
	
	
	for rows.Next() {
		m := SgeDePara{}
		err := rows.Scan(&m.EmpCod, 
			&m.Cod, 
			&m.NomeRh, 
			&m.EstruturaIDSge,
			&m.ColecaoIDSge,
			&m.RegistroEstruturaID,
			&m.NomeSge,
			&m.CodHierarquia,
			&m.CodConsist, 
			&m.EstruturaUems)

		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return -1, err
		}
		
		//CreateDev(&m);
		
		//fmt.Println( m.EmpCod.Int32 );
		//fmt.Println( m.NomeRh.String );
		//fmt.Printf("%T", &m.NomeRh.String)
		//fmt.Printf("EmpCod: %d, Cod: %d, NomeRh: %s  \n", &m.EmpCod, &m.Cod, &m.NomeRh)
		
		count++
	}
	return count, nil
}
