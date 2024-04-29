package rutas
import (
	"database/sql"
	"fmt"
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

type RutaHandler struct {
	db           *sql.DB
	tarifaGlobal float64 // Tarifa global
}

func NewRutaHandler(db *sql.DB) *RutaHandler {
	return &RutaHandler{db: db}
}

func (h *RutaHandler) CrearRuta(c *gin.Context) {
	var ruta Ruta
	if err := c.ShouldBindJSON(&ruta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de ruta inválidos"})
		return
	}

	if ruta.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere un valor para ruta_id"})
		return
	}

	_, err := h.db.Exec("INSERT INTO rutas (ruta_id, destino, tarifa_operacion, capacidad_maxima) VALUES (?, ?, ?, ?)",
		ruta.ID, ruta.Destino, ruta.TarifaOperacion, ruta.CapacidadMaxima)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la ruta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ruta creada exitosamente"})
}

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

func (h *RutaHandler) existeRutaID(id int) bool {
	var count int
	err := h.db.QueryRow("SELECT COUNT(*) FROM rutas WHERE ruta_id = ?", id).Scan(&count)
	if err != nil {
		log.Printf("Error al verificar la existencia de la ruta:  %v", err)
		return true
	}
	return count > 0
}

func (h *RutaHandler) ActualizarRuta(c *gin.Context) {
	var ruta Ruta
	if err := c.ShouldBindJSON(&ruta); err != nil {
		log.Printf("Error al decodificar el cuerpo JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de ruta inválidos"})
		return
	}

	log.Printf("Datos recibidos para actualización - ID: %v, Destino: %v, TarifaOperacion: %v, CapacidadMaxima: %v",
		ruta.ID, ruta.Destino, ruta.TarifaOperacion, ruta.CapacidadMaxima)

	if !h.existeRutaID(ruta.ID) {
		log.Printf("La ruta con ID %v no existe", ruta.ID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "La ruta no existe"})
		return
	}

	_, err := h.db.Exec("UPDATE rutas SET destino=?, tarifa_operacion=?, capacidad_maxima=? WHERE ruta_id=?",
		ruta.Destino, ruta.TarifaOperacion, ruta.CapacidadMaxima, ruta.ID)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL de actualización: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la ruta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ruta actualizada exitosamente"})
}

func (h *RutaHandler) EliminarRuta(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de ruta inválido"})
		return
	}

	if !h.existeRutaID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La ruta no existe"})
		return
	}

	_, err = h.db.Exec("DELETE FROM rutas WHERE ruta_id = ?", id)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL de eliminación: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la ruta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ruta eliminada exitosamente"})
}

func (h *RutaHandler) ObtenerTarifaOperacion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de ruta inválido"})
		return
	}

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

// SetTarifaGlobalHandler es un controlador para establecer la tarifa global
// SetTarifaGlobalHandler es un controlador para establecer la tarifa global
func (h *RutaHandler) SetTarifaGlobalHandler(c *gin.Context) {
	var body struct {
		Tarifa float64 `json:"tarifa"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de tarifa inválidos"})
		return
	}

	h.SetTarifaGlobal(body.Tarifa)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Tarifa global establecida: %.2f", body.Tarifa)})
}

// GetTarifaGlobalHandler es un controlador para obtener la tarifa global
func (h *RutaHandler) GetTarifaGlobalHandler(c *gin.Context) {
	tarifa := h.ObtenerTarifaGlobal()
	log.Printf("Tarifa global obtenida: %.2f", tarifa)
	c.JSON(http.StatusOK, gin.H{"tarifa_global": tarifa})
}

// SetTarifaGlobal establece la tarifa global y la guarda en la variable tarifaGlobal
func (h *RutaHandler) SetTarifaGlobal(tarifa float64) {
	h.tarifaGlobal = tarifa
	log.Printf("Tarifa global establecida: %.2f", tarifa)
}

// ObtenerTarifaGlobal devuelve la tarifa global almacenada en la variable tarifaGlobal
func (h *RutaHandler) ObtenerTarifaGlobal() float64 {
	tarifa := h.tarifaGlobal
	log.Printf("Tarifa global obtenida: %.2f", tarifa)
	return tarifa
}