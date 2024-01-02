README (Español)

Sistema de Gestión de Clientes por Consola

Este programa proporciona una interfaz de consola para gestionar datos de clientes. Incluye opciones para listar todos los clientes, listar clientes por ID, crear nuevos clientes, modificar registros de clientes existentes y eliminar registros de clientes.

Paquete: handlers
El paquete handlers contiene funciones para manejar diversas operaciones en los datos de clientes.

Función: Listar()

go
Copy code
func Listar()
Lista todos los clientes en orden descendente según sus IDs.
Muestra detalles del cliente, incluidos ID, nombre, correo electrónico y número de teléfono.
Función: ListarPorId(id int)

go
Copy code
func ListarPorId(id int)
Lista un cliente según el ID proporcionado.
Muestra detalles del cliente, incluidos ID, nombre, correo electrónico y número de teléfono.
Función: Insertar(cliente modelos.Cliente)

go
Copy code
func Insertar(cliente modelos.Cliente)
Inserta un nuevo registro de cliente en la base de datos.
Requiere un objeto modelos.Cliente como entrada, que contiene detalles del cliente (nombre, correo electrónico, teléfono).
Función: Editar(cliente modelos.Cliente, id int)

go
Copy code
func Editar(cliente modelos.Cliente, id int)
Edita un registro de cliente existente según el ID proporcionado.
Requiere un objeto modelos.Cliente como entrada, que contiene detalles actualizados del cliente (nombre, correo electrónico, teléfono).
Función: Eliminar(id int)

go
Copy code
func Eliminar(id int)
Elimina un registro de cliente según el ID proporcionado.
Uso: Ejecutar()
La función Ejecutar() permite a los usuarios interactuar con el programa a través de la consola. Presenta un menú con opciones para realizar diversas acciones:

Listar todos los clientes
Listar un cliente por ID
Crear un nuevo cliente
Editar un cliente existente
Eliminar un cliente
Los usuarios pueden ingresar el número de opción correspondiente y seguir las indicaciones para ejecutar la operación deseada.