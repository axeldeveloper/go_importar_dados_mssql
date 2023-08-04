package mssql

import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	models "go.importar.dados.mssql/models"
)

func ReadSgeDePara() {
	members, err := models.BuscaTodosOsSgeDePara()
	if err != nil {
		log.Fatal("ReadSgeDePara failed:", err.Error())
	}
	fmt.Printf("Read %d rows successfully.\n", len(members))
	for _, value := range members {
		//fmt.Println(value.EmpCod.Int32)
		fmt.Println(value)
	}
}

func ReadMinicipios() {
	minicipios, err := models.BuscaTodosOsMinicipios()
	if err != nil {
		log.Fatal("Minicipios failed:", err.Error())
	}
	fmt.Printf("Read %d rows successfully.\n", len(minicipios))
	for _, value := range minicipios {
		///fmt.Println(value.MunCodigoIbge)

		minicipioID, err := models.CreateMinicipiosVm(&value)
		if err != nil {
			log.Fatal("Erro ao inseir os Minicipios:", err.Error())
		}

		fmt.Printf("insert %d vm rows successfully.\n", minicipioID)
	}

}

func Run() {

	//ReadSgeDePara()

	// ReadMinicipios();
}
