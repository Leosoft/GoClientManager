package conectar

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

// Funcion para conectarme a la BD
var Db *sql.DB

func Conectar() {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}
	connection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))

	if err != nil {
		panic(err)
	}
	Db = connection

}

// Cerrar la conexion
func CerrarConexion() {
	Db.Close()
}

//go get github.com/go-sql/driver/mysql
//go get github.com/joho/godotenv

/*
CREATE TABLE `clientes`(
`id` int NOT NULL AUTO_INCREMENT,
`nombre` varchar(100) NOT NULL,
`correo` varchar(100) NOT NULL,
`telefono` varchar(20) NOT NULL,
`fecha` datetime NOT NULL,
PRIMARY_KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COLLATE= utf8mb4_0900_ai_ci
*/
