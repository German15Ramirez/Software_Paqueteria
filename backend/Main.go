package main
import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

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

	// Rutas de Gin para rutas
	rutaHandler := rutas.NewRutaHandler(db)
	router.POST("/crear_ruta", rutaHandler.CrearRuta)
	router.GET("/listar_rutas", rutaHandler.ListarRutas)

	// Rutas de Gin para usuarios
	usuarioHandler := usuarios.NewUsuarioHandler(db)
	router.POST("/crear_usuario", usuarioHandler.CrearUsuario)
	router.GET("/listar_usuarios", usuarioHandler.ListarUsuarios)

	// Iniciar el servidor Gin en el puerto 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor Gin:", err)
	}
}
