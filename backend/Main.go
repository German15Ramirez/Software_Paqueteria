package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	//herramientas necesarias para la conexion con mysql
)

func main() {
	// Configuración de las credenciales de conexión a la base de datos de MySQL
	dbUser := "root"
	dbPassword := "23al15"
	dbName := "gestion_paqueteria"
	dbHost := "localhost"
	dbPort := "3306" // Puerto por defecto de MySQL

	// Formar la cadena de conexión
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Abrir una conexión a la base de datos
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Error al abrir la conexión a la base de datos:", err)
	}
	defer db.Close() // Cierra la conexión al finalizar la función

	// Intentar realizar una conexión ping para verificar que la conexión sea exitosa
	if err := db.Ping(); err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	// Si la conexión fue exitosa se muestra este mensaje
	log.Println("Conexión exitosa a la base de datos MySQL")

	// Solicitar ID y contraseña
	var usuarioID int
	var contraseña string
	fmt.Print("Ingrese su ID de usuario: ")
	fmt.Scanln(&usuarioID)
	fmt.Print("Ingrese su contraseña: ")
	fmt.Scanln(&contraseña)

	// Consultar si el usuario y contraseña son válidos
	var rol string
	err = db.QueryRow("SELECT rol FROM usuarios WHERE usuario_id = ? AND contraseña = ?", usuarioID, contraseña).Scan(&rol)
	if err != nil {
		log.Fatal("Error al verificar el usuario:", err)
	}

	if rol == "Administrador" {
		// Menú para el usuario administrador
		fmt.Println("\nMenú Administrador:")
		fmt.Println("1. Fijar tarifa de operación")
		fmt.Println("2. Fijar cuota de destino")
		fmt.Println("3. Desactivar rutas")
		fmt.Println("4. Agregar/quitar puntos de control")
		fmt.Println("5. Modificar datos de punto de control")
		fmt.Println("6. Desactivar usuarios")
		fmt.Println("7. Ver reportes")
		fmt.Println("0. Salir")

		var opcion int
		fmt.Print("\nSeleccione una opción (0-7): ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Opción 1 seleccionada: Fijar tarifa de operación")
			// Implementar lógica para esta opción
		case 2:
			fmt.Println("Opción 2 seleccionada: Fijar cuota de destino")
			// Implementar lógica para esta opción
		case 3:
			fmt.Println("Opción 3 seleccionada: Desactivar rutas")
			// Implementar lógica para esta opción
		case 4:
			fmt.Println("Opción 4 seleccionada: Agregar/quitar puntos de control")
			// Implementar lógica para esta opción
		case 5:
			fmt.Println("Opción 5 seleccionada: Modificar datos de punto de control")
			// Implementar lógica para esta opción
		case 6:
			fmt.Println("Opción 6 seleccionada: Desactivar usuarios")
			// Implementar lógica para esta opción
		case 7:
			fmt.Println("Opción 7 seleccionada: Ver reportes")
			// Implementar lógica para esta opción
		case 0:
			fmt.Println("Saliendo del programa...")
		default:
			fmt.Println("Opción no válida.")
		}
	} else {
		fmt.Println("Usuario no autorizado.")
	}
}