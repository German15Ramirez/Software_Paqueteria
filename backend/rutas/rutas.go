package rutas
import (
	"database/sql"
	"log"
	"net/http"

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
	result, err := h.db.Exec("INSERT INTO rutas (ruta_id, destino, tarifa_operacion, capacidad_maxima) VALUES (?, ?, ?, ?)",
		ruta.ID, ruta.Destino, ruta.TarifaOperacion, ruta.CapacidadMaxima)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la ruta"})
		return
	}

	// Verificar el número de filas afectadas por la inserción
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("No se insertaron filas")
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