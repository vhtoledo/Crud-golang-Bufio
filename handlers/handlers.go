package handlers

import (
	"bufio"
	"clase_3_mysql_driver/conectar"
	"clase_3_mysql_driver/modelos"
	"os"
	"strconv"

	//"text/scanner"

	//"database/sql"
	"fmt"
	"log"
)

func Listar(){
	conectar.Conectar()
	sql := "select id, nombre, correo, telefono from clientes order by id desc;"
	datos, err := conectar.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer conectar.CerrarConexion()
	clientes := modelos.Clientes{}
	for datos.Next(){
		dato := modelos.Cliente{}
		datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		clientes = append(clientes, dato)
	}
	fmt.Println(clientes)
}

func ListarPorId(id int) {
	conectar.Conectar()
	sql := "select id, nombre, correo, telefono from clientes where id =?;"
	datos, err := conectar.Db.Query(sql, id)
	if err != nil {
		fmt.Println(err)
	}
	for datos.Next(){
		var dato  modelos.Cliente
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dato)
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
	fmt.Println("Se creó el registro")
}

func Editar(cliente modelos.Cliente, id int) {
	conectar.Conectar()
	sql := "update clientes set nombre=?, correo, telefono=? where id =?;"
	result, err := conectar.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("Se actualizo el registro")
}

func Eliminar(id int) {
	conectar.Conectar()
	sql := "delete from clientes where id =?;"
	result, err := conectar.Db.Exec(sql, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("Se elimino el registro")
}

// ######################FUNCION DE TRABAJO
var ID int
var nombre, correo, telefono string

func Ejecutar() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Seleccione una opcion: \n\n")
	fmt.Println("1 - Listar clientes\n")
	fmt.Println("2 - Listar clientes por ID\n")
	fmt.Println("3 - Crear cliente\n")
	fmt.Println("4 - EDitar cliente\n")
	fmt.Println("5 - Eliminar clientes\n")
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
				fmt.Println("Ingrese email: \n")
				if scanner.Scan() {
					correo = scanner.Text()
				}
				fmt.Println("Ingrese telefono: \n")
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
				fmt.Println("Ingrese nombre: \n")
				if scanner.Scan() {
					nombre = scanner.Text()
				}
				fmt.Println("Ingrese email: \n")
				if scanner.Scan() {
					correo = scanner.Text()
				}
				fmt.Println("Ingrese telefono: \n")
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