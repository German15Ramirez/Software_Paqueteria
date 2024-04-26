package paquetes
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Paquete representa la estructura de datos de un paquete.
type Paquete struct {
	ID            int
	ClienteID     int
	NombreCliente string
	PesoLibras    float64
	Destino       string
	FechaIngreso  string
	FechaSalida   string
	Estado        string
	PuntoControl  string
	Ruta          string
}

// PaqueteHandler maneja las operaciones relacionadas con los paquetes.
type PaqueteHandler struct {
	db *sql.DB
}

// NewPaqueteHandler crea un nuevo manejador de paquetes.
func NewPaqueteHandler(db *sql.DB) *PaqueteHandler {
	return &PaqueteHandler{db: db}
}

// ObtenerLocalizacionPaquete obtiene la localización de un paquete por su ID.
func (h *PaqueteHandler) ObtenerLocalizacionPaquete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		fmt.Println("Error: Se requiere el parámetro 'id'")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere el parámetro 'id'"})
		return
	}

	var paquete Paquete
	row := h.db.QueryRow("SELECT pa.paquete_id, pa.cliente_id, cl.nombre AS nombre_cliente, pa.peso_libras, pa.destino, pa.fecha_ingreso, pa.fecha_salida, pa.estado, pc.nombre AS punto_control, r.destino AS ruta FROM paquetes pa JOIN puntos_de_control pc ON pa.punto_de_control_id = pc.puntos_de_control_id JOIN rutas r ON pc.ruta_id = r.ruta_id JOIN clientes cl ON pa.cliente_id = cl.cliente_id WHERE pa.paquete_id = ?", idStr)
	err := row.Scan(&paquete.ID, &paquete.ClienteID, &paquete.NombreCliente, &paquete.PesoLibras, &paquete.Destino, &paquete.FechaIngreso, &paquete.FechaSalida, &paquete.Estado, &paquete.PuntoControl, &paquete.Ruta)
	if err == sql.ErrNoRows {
		fmt.Println("Error: Paquete no encontrado")
		c.JSON(http.StatusNotFound, gin.H{"error": "Paquete no encontrado"})
		return
	} else if err != nil {
		fmt.Println("Error al consultar la base de datos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar la base de datos"})
		return
	}

	c.JSON(http.StatusOK, paquete)
}
// ListarPaquetesPorDestino lista los paquetes filtrados por destino.
func (h *PaqueteHandler) ListarPaquetesPorDestino(c *gin.Context) {
	destino := c.Query("destino") // Obtener el parámetro de consulta "destino"

	var rows *sql.Rows
	var err error

	if destino != "" {
		rows, err = h.db.Query("SELECT pa.paquete_id, pa.cliente_id, cl.nombre AS nombre_cliente, pa.peso_libras, pa.destino, pa.fecha_ingreso, pa.fecha_salida, pa.estado, pc.nombre AS punto_control, r.destino AS ruta FROM paquetes pa JOIN puntos_de_control pc ON pa.punto_de_control_id = pc.puntos_de_control_id JOIN rutas r ON pc.ruta_id = r.ruta_id JOIN clientes cl ON pa.cliente_id = cl.cliente_id WHERE pa.destino = ?", destino)
	} else {
		rows, err = h.db.Query("SELECT pa.paquete_id, pa.cliente_id, cl.nombre AS nombre_cliente, pa.peso_libras, pa.destino, pa.fecha_ingreso, pa.fecha_salida, pa.estado, pc.nombre AS punto_control, r.destino AS ruta FROM paquetes pa JOIN puntos_de_control pc ON pa.punto_de_control_id = pc.puntos_de_control_id JOIN rutas r ON pc.ruta_id = r.ruta_id JOIN clientes cl ON pa.cliente_id = cl.cliente_id")
	}

	if err != nil {
		log.Printf("Error al consultar la base de datos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar la base de datos"})
		return
	}
	defer rows.Close()

	paquetes := make([]Paquete, 0)
	for rows.Next() {
		var paquete Paquete
		if err := rows.Scan(&paquete.ID, &paquete.ClienteID, &paquete.NombreCliente, &paquete.PesoLibras, &paquete.Destino, &paquete.FechaIngreso, &paquete.FechaSalida, &paquete.Estado, &paquete.PuntoControl, &paquete.Ruta); err != nil {
			log.Printf("Error al escanear filas: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escanear filas"})
			return
		}
		paquetes = append(paquetes, paquete)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al escanear filas: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escanear filas"})
		return
	}

	c.JSON(http.StatusOK, paquetes)
}