package puntos_de_control
import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

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
	// Verificar si ya existe un punto de control con el mismo ID
	if h.existePuntoDeControlID(puntoDeControl.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID de punto de control ya está en uso"})
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
// existePuntoDeControlID verifica si ya existe un punto de control con el mismo ID en la base de datos
func (h *PuntoDeControlHandler) existePuntoDeControlID(id int) bool {
	var count int
	err := h.db.QueryRow("SELECT COUNT(*) FROM puntos_de_control WHERE puntos_de_control_id = ?", id).Scan(&count)
	if err != nil {
		log.Printf("Error al verificar la existencia del punto de control:  %v", err)
		return true // Por precaución, asumimos que el punto de control ya existe
	}
	return count > 0
}
// ActualizarPuntoDeControl actualiza los datos de un punto de control existente
func (h *PuntoDeControlHandler) ActualizarPuntoDeControl(c *gin.Context) {
	// Estructura para almacenar los datos actualizados del punto de control
	var puntoDeControl PuntoDeControl

	// Decodificar el cuerpo JSON de la solicitud en la estructura del punto de control
	if err := c.ShouldBindJSON(&puntoDeControl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de punto de control inválidos"})
		return
	}

	// Verificar si el punto de control que se intenta actualizar existe
	if !h.existePuntoDeControlID(puntoDeControl.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El punto de control no existe"})
		return
	}

	// Actualizar los datos del punto de control en la base de datos
	_, err := h.db.Exec("UPDATE puntos_de_control SET ruta_id=?, nombre=?, capacidad_cola=?, tarifa_operacional=? WHERE puntos_de_control_id=?",
		puntoDeControl.RutaID, puntoDeControl.Nombre, puntoDeControl.CapacidadCola, puntoDeControl.TarifaOperacional, puntoDeControl.ID)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL de actualización: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el punto de control"})
		return
	}

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Punto de control actualizado exitosamente"})
}
// EliminarPuntoDeControl elimina un punto de control existente
func (h *PuntoDeControlHandler) EliminarPuntoDeControl(c *gin.Context) {
	// Obtener el ID del punto de control de los parámetros de ruta
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de punto de control inválido"})
		return
	}

	// Verificar si el punto de control que se intenta eliminar existe
	if !h.existePuntoDeControlID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El punto de control no existe"})
		return
	}

	// Verificar si el punto de control tiene paquetes asociados
	if h.tienePaquetesAsociados(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se puede eliminar el punto de control porque tiene paquetes asociados"})
		return
	}

	// Eliminar el punto de control de la base de datos
	_, err = h.db.Exec("DELETE FROM puntos_de_control WHERE puntos_de_control_id = ?", id)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL de eliminación: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el punto de control"})
		return
	}

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Punto de control eliminado exitosamente"})
}
// tienePaquetesAsociados verifica si un punto de control tiene paquetes asociados
func (h *PuntoDeControlHandler) tienePaquetesAsociados(id int) bool {
	var count int
	err := h.db.QueryRow("SELECT COUNT(*) FROM paquetes WHERE punto_de_control_id = ?", id).Scan(&count)
	if err != nil {
		log.Printf("Error al verificar la existencia de paquetes asociados: %v", err)
		return true
	}
	return count > 0
}