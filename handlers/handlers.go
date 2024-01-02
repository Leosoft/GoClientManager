package handlers

import (
	"bufio"
	"fmt"
	"go_mysql_driver/conectar"
	_ "go_mysql_driver/conectar"
	"go_mysql_driver/modelos"
	"log"
	"os"
	"strconv"
)

func Listar() {
	conectar.Conectar()

	sql := "select id,nombre,correo,telefono from clientes order by id desc;"
	datos, err := conectar.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer conectar.CerrarConexion()

	/*
		clientes := modelos.Clientes{}

		for datos.Next() {
			dato := modelos.Cliente{}
			datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
			clientes = append(clientes, dato)
		}
		fmt.Println(clientes)

	*/

	fmt.Println("\nListado de clientes")
	fmt.Println("--------------------------------------------------")
	for datos.Next() {

		var dato = modelos.Cliente{}
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Id: %v | Nombre: %v | E-mail: %s | Telefono %s \n", dato.Id, dato.Nombre, dato.Correo, dato.Telefono)
		fmt.Println("--------------------------------------------------")
	}
}

func ListarPorId(id int) {
	conectar.Conectar()

	sql := "select id,nombre,correo,telefono from clientes where id=?;"
	datos, err := conectar.Db.Query(sql, id)
	if err != nil {
		fmt.Println(err)
	}
	defer conectar.CerrarConexion()
	fmt.Println("\nListado de clientes por ID")
	fmt.Println("------------------------------------------------")
	for datos.Next() {
		var dato modelos.Cliente
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Id: %v | Nombre: %v | E-mail: %s | Telefono %s \n", dato.Id, dato.Nombre, dato.Correo, dato.Telefono)
		fmt.Println("--------------------------------------------------")

	}
}

func Insertar(cliente modelos.Cliente) {
	conectar.Conectar()
	sql := "insert into clientes values(null, ?,?,?);"
	result, err := conectar.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("Se creo el registro exitosamente")
}

func Editar(cliente modelos.Cliente, id int) {
	conectar.Conectar()
	sql := "update clientes set nombre=?, correo=?,telefono=? where id=?;"
	result, err := conectar.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("Se edito el registro exitosamente")
}

func Eliminar(id int) {
	conectar.Conectar()
	sql := "delete from clientes where id=?;"
	_, err := conectar.Db.Exec(sql, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Se elimino el registro exitosamente")

}

// ###############FUNCIONES DE TRABAJO
// bufio permite que el usuario vaya digitando la informacion

var ID int
var nombre, correo, telefono string

func Ejecutar() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Seleccione una opcion: \n\n")
	fmt.Println("1- Listar clientes\n")
	fmt.Println("2- Listar cliente por ID\n")
	fmt.Println("3- Crear cliente\n")
	fmt.Println("4- Editar cliente\n")
	fmt.Println("5- Eliminar cliente\n")
	if scanner.Scan() {
		for {
			if scanner.Text() == "1" {
				Listar()
				return
			}

			if scanner.Text() == "2" {
				fmt.Println("Ingrese el ID del cliente: \n")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
				}
				ListarPorId(ID)
				return
			}

			if scanner.Text() == "3" {

				fmt.Println("Ingrese nombre: \n")
				if scanner.Scan() {
					nombre = scanner.Text()
				}
				fmt.Println("Ingrese E-Mail: \n")
				if scanner.Scan() {
					correo = scanner.Text()
				}
				fmt.Println("Ingrese Telefono: \n")
				if scanner.Scan() {
					telefono = scanner.Text()
				}
				cliente := modelos.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
				Insertar(cliente)
				return
			}

			if scanner.Text() == "4" {

				fmt.Println("Ingrese ID cliente: \n")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
				}
				fmt.Println("Ingrese Nombre: \n")
				if scanner.Scan() {
					nombre = scanner.Text()
				}
				fmt.Println("Ingrese E-Mail: \n")
				if scanner.Scan() {
					correo = scanner.Text()
				}
				fmt.Println("Ingrese Telefono: \n")
				if scanner.Scan() {
					telefono = scanner.Text()
				}
				cliente := modelos.Cliente{Nombre: nombre, Correo: correo, Telefono: telefono}
				Editar(cliente, ID)
				return

			}

			if scanner.Text() == "5" {
				fmt.Println("Ingrese el ID del cliente: \n")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
				}
				Eliminar(ID)
				return
			}
		}
	}
}
