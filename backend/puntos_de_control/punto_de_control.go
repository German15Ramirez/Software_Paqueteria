package puntos_de_control

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
type PuntoDeControl struct {
	ID                int
	RutaID            int
	Nombre            string
	CapacidadCola     int
	TarifaOperacional float64
}

// PuntoDeControlHandler maneja las operaciones relacionadas con los puntos de control
type PuntoDeControlHandler struct {
	db *sql.DB
}

// NewPuntoDeControlHandler crea un nuevo manejador de puntos de control
func NewPuntoDeControlHandler(db *sql.DB) *PuntoDeControlHandler {
	return &PuntoDeControlHandler{db: db}
}

// CrearPuntoDeControl crea un nuevo punto de control
func (h *PuntoDeControlHandler) CrearPuntoDeControl(c *gin.Context) {
	// Estructura para almacenar los datos del punto de control
	var puntoDeControl PuntoDeControl

	// Decodificar el cuerpo JSON de la solicitud en la estructura del punto de control
	if err := c.ShouldBindJSON(&puntoDeControl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de punto de control inválidos"})
		return
	}

	// Insertar los datos del punto de control en la base de datos
	_, err := h.db.Exec("INSERT INTO puntos_de_control (puntos_de_control_id, ruta_id, nombre, capacidad_cola, tarifa_operacional) VALUES (?, ?, ?, ?, ?)",
		puntoDeControl.ID, puntoDeControl.RutaID, puntoDeControl.Nombre, puntoDeControl.CapacidadCola, puntoDeControl.TarifaOperacional)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el punto de control"})
		return
	}

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Punto de control creado exitosamente"})
}

// ListarPuntosDeControl lista todos los puntos de control
func (h *PuntoDeControlHandler) ListarPuntosDeControl(c *gin.Context) {
	// Consulta SQL para seleccionar los puntos de control
	rows, err := h.db.Query("SELECT puntos_de_control_id, ruta_id, nombre, capacidad_cola, tarifa_operacional FROM puntos_de_control")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Slice para almacenar los puntos de control
	puntosDeControl := make([]PuntoDeControl, 0)

	// Iterar sobre las filas y escanear los datos en la estructura PuntoDeControl
	for rows.Next() {
		var pc PuntoDeControl
		if err := rows.Scan(&pc.ID, &pc.RutaID, &pc.Nombre, &pc.CapacidadCola, &pc.TarifaOperacional); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		puntosDeControl = append(puntosDeControl, pc)
	}

	// Verificar si hubo un error durante el escaneo de las filas
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver la lista de puntos de control como respuesta JSON
	c.JSON(http.StatusOK, puntosDeControl)
}