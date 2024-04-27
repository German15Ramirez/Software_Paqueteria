package rutas
import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Ruta struct {
	ID              int
	Destino         string
	TarifaOperacion float64
	CapacidadMaxima int
}

// RutaHandler maneja las operaciones relacionadas con las rutas
type RutaHandler struct {
	db *sql.DB
}

// NewRutaHandler crea un nuevo manejador de rutas
func NewRutaHandler(db *sql.DB) *RutaHandler {
	return &RutaHandler{db: db}
}

// CrearRuta crea una nueva ruta
func (h *RutaHandler) CrearRuta(c *gin.Context) {
	var ruta Ruta
	if err := c.ShouldBindJSON(&ruta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de ruta inválidos"})
		return
	}

	// Verificar si se proporcionó un valor para ruta_id
	if ruta.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere un valor para ruta_id"})
		return
	}

	// Insertar los datos de la ruta en la base de datos
	_, err := h.db.Exec("INSERT INTO rutas (ruta_id, destino, tarifa_operacion, capacidad_maxima) VALUES (?, ?, ?, ?)",
		ruta.ID, ruta.Destino, ruta.TarifaOperacion, ruta.CapacidadMaxima)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la ruta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ruta creada exitosamente"})
}

// ListarRutas lista todas las rutas
func (h *RutaHandler) ListarRutas(c *gin.Context) {
	rows, err := h.db.Query("SELECT ruta_id, destino, tarifa_operacion, capacidad_maxima FROM rutas")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	rutas := make([]Ruta, 0)
	for rows.Next() {
		var ruta Ruta
		if err := rows.Scan(&ruta.ID, &ruta.Destino, &ruta.TarifaOperacion, &ruta.CapacidadMaxima); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		rutas = append(rutas, ruta)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rutas)
}

// existeRutaID verifica si ya existe una ruta con el mismo ID en la base de datos
func (h *RutaHandler) existeRutaID(id int) bool {
	var count int
	err := h.db.QueryRow("SELECT COUNT(*) FROM rutas WHERE ruta_id = ?", id).Scan(&count)
	if err != nil {
		log.Printf("Error al verificar la existencia de la ruta:  %v", err)
		return true
	}
	return count > 0
}

// ActualizarRuta actualiza los datos de una ruta existente
func (h *RutaHandler) ActualizarRuta(c *gin.Context) {
	// Estructura para almacenar los datos actualizados de la ruta
	var ruta Ruta

	// Decodificar el cuerpo JSON de la solicitud en la estructura de la ruta
	if err := c.ShouldBindJSON(&ruta); err != nil {
		log.Printf("Error al decodificar el cuerpo JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de ruta inválidos"})
		return
	}

	// Agregar mensajes de depuración para conocer los datos que llegan
	log.Printf("Datos recibidos para actualización - ID: %v, Destino: %v, TarifaOperacion: %v, CapacidadMaxima: %v",
		ruta.ID, ruta.Destino, ruta.TarifaOperacion, ruta.CapacidadMaxima)

	// Verificar si la ruta que se intenta actualizar existe
	if !h.existeRutaID(ruta.ID) {
		log.Printf("La ruta con ID %v no existe", ruta.ID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "La ruta no existe"})
		return
	}

	// Actualizar los datos de la ruta en la base de datos
	_, err := h.db.Exec("UPDATE rutas SET destino=?, tarifa_operacion=?, capacidad_maxima=? WHERE ruta_id=?",
		ruta.Destino, ruta.TarifaOperacion, ruta.CapacidadMaxima, ruta.ID)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL de actualización: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la ruta"})
		return
	}

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Ruta actualizada exitosamente"})
}

// EliminarRuta elimina una ruta existente
func (h *RutaHandler) EliminarRuta(c *gin.Context) {
	// Obtener el ID de la ruta de los parámetros de la ruta
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de ruta inválido"})
		return
	}

	// Verificar si la ruta que se intenta eliminar existe
	if !h.existeRutaID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La ruta no existe"})
		return
	}

	// Eliminar la ruta de la base de datos
	_, err = h.db.Exec("DELETE FROM rutas WHERE ruta_id = ?", id)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL de eliminación: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la ruta"})
		return
	}

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Ruta eliminada exitosamente"})
}

func (h *RutaHandler) ObtenerTarifaOperacion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de ruta inválido"})
		return
	}

	// Consultar la base de datos para obtener la tarifa de operación de la ruta con el ID dado
	var tarifaOperacion float64
	err = h.db.QueryRow("SELECT tarifa_operacion FROM rutas WHERE ruta_id = ?", id).Scan(&tarifaOperacion)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ruta no encontrada"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la tarifa de operación"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tarifa_operacion": tarifaOperacion})
}

// ActualizarTarifaGlobal actualiza la tarifa global de una ruta.
func (h *RutaHandler) ActualizarTarifaGlobal(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de ruta inválido"})
		return
	}

	var nuevaTarifa float64
	if err := c.ShouldBindJSON(&nuevaTarifa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de tarifa inválidos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tarifa global actualizada exitosamente"})
}
