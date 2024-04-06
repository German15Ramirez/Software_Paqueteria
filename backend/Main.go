package main
import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"gestionpaquetes/clientes"
	"gestionpaquetes/paquetes"
	"gestionpaquetes/puntos_de_control"
	"gestionpaquetes/rutas"
	"gestionpaquetes/usuarios"
)

func main() {
	// Configurar Gin
	router := gin.Default()

	// Configurar conexión a la base de datos
	db, err := sql.Open("mysql", "root:23al15@tcp(localhost:3306)/gestion_paqueteria")
	if err != nil {
		log.Fatal("Error al abrir la conexión a la base de datos:", err)
	}
	defer db.Close() // Cierra la conexión al finalizar la función

	// Rutas de Gin para puntos de control
	puntoControlHandler := puntos_de_control.NewPuntoDeControlHandler(db)
	router.POST("/crear_punto_de_control", puntoControlHandler.CrearPuntoDeControl)
	router.GET("/listar_puntos_de_control", puntoControlHandler.ListarPuntosDeControl)
	router.PUT("/actualizar_punto_de_control", puntoControlHandler.ActualizarPuntoDeControl)
	router.DELETE("/eliminar_punto_de_control/:id", puntoControlHandler.EliminarPuntoDeControl)
	// Rutas de Gin para rutas
	rutaHandler := rutas.NewRutaHandler(db)
	router.POST("/crear_ruta", rutaHandler.CrearRuta)
	router.GET("/listar_rutas", rutaHandler.ListarRutas)
	router.PUT("/actualizar_ruta", rutaHandler.ActualizarRuta)
	router.DELETE("/eliminar_ruta/:id", rutaHandler.EliminarRuta)

	//especie de modificacion de CRUD de tarifas
	router.GET("/rutas/:id/tarifa_operacional", rutaHandler.ObtenerTarifaOperacion)
	router.PUT("/rutas/:id/actualizar_tarifa_global", rutaHandler.ActualizarTarifaGlobal)

	// Rutas de Gin para usuarios
	usuarioHandler := usuarios.NewUsuarioHandler(db)
	router.POST("/crear_usuario", usuarioHandler.CrearUsuario)
	router.GET("/listar_usuarios", usuarioHandler.ListarUsuarios)
	router.PUT("/actualizar_usuario", usuarioHandler.ActualizarUsuario)
	router.DELETE("/eliminar_usuario", usuarioHandler.EliminarUsuario)

	// Rutas de Gin para clientes
	clienteHandler := clientes.NewClienteHandler(db)
	router.POST("/crear_cliente", clienteHandler.CrearCliente)
	router.GET("/listar_clientes", clienteHandler.ListarClientes)
	router.PUT("/actualizar_cliente", clienteHandler.ActualizarCliente)
	router.DELETE("/eliminar_cliente/:id", clienteHandler.EliminarCliente)


	// Rutas de Gin para la localización de paquetes
	paqueteHandler := paquetes.NewPaqueteHandler(db)
	router.GET("/paquetes/:id/localizacion", paqueteHandler.ObtenerLocalizacionPaquete)
	// Iniciar el servidor Gin en el puerto 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor Gin:", err)
	}
}
