README (English)

Console-based Customer Management System

This program provides a console-based interface to manage customer data. It includes options to list all customers, list customers by ID, create new customers, modify existing customer records, and delete customer records.

Package: handlers
The handlers package contains functions for handling various operations on customer data.

Function: Listar()

go
Copy code
func Listar()
Lists all customers in descending order of their IDs.
Displays customer details including ID, name, email, and phone number.
Function: ListarPorId(id int)

go
Copy code
func ListarPorId(id int)
Lists a customer based on the provided ID.
Displays customer details including ID, name, email, and phone number.
Function: Insertar(cliente modelos.Cliente)

go
Copy code
func Insertar(cliente modelos.Cliente)
Inserts a new customer record into the database.
Requires a modelos.Cliente object as input, containing customer details (name, email, phone).
Function: Editar(cliente modelos.Cliente, id int)

go
Copy code
func Editar(cliente modelos.Cliente, id int)
Edits an existing customer record based on the provided ID.
Requires a modelos.Cliente object as input, containing updated customer details (name, email, phone).
Function: Eliminar(id int)

go
Copy code
func Eliminar(id int)
Deletes a customer record based on the provided ID.
Usage: Ejecutar()
The Ejecutar() function allows users to interact with the program through the console. It presents a menu with options to perform various actions:

List all customers
List a customer by ID
Create a new customer
Edit an existing customer
Delete a customer
Users can input the corresponding option number and follow the prompts to execute the desired operation.

