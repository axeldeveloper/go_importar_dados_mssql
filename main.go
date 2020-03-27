package main
 
import (
	"fmt"
	mssql "github.com/axeldeveloper/go.importar.dados.mssql/mssql"

)

func main() {
	fmt.Printf("iniciando \n")
	mssql.Run();
	//todosOsProdutos := models.BuscaTodosOsMinicipios()
}