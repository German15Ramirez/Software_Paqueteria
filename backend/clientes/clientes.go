package clientes
import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Cliente struct {
	ID        int
	Nombre    string
	NIT       int
	Direccion string
	Telefono  int
}

// ClienteHandler maneja las operaciones relacionadas con los clientes
type ClienteHandler struct {
	db *sql.DB
}

// NewClienteHandler crea un nuevo manejador de clientes
func NewClienteHandler(db *sql.DB) *ClienteHandler {
	return &ClienteHandler{db: db}
}

// CrearCliente crea un nuevo cliente
func (h *ClienteHandler) CrearCliente(c *gin.Context) {
	var cliente Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de cliente inválidos"})
		return
	}

	// Verificar si se proporcionó un valor para cliente_id
	if cliente.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere un valor para cliente_id"})
		return
	}

	_, err := h.db.Exec("INSERT INTO clientes (cliente_id, nombre, nit, direccion, telefono) VALUES (?, ?, ?, ?, ?)",
		cliente.ID, cliente.Nombre, cliente.NIT, cliente.Direccion, cliente.Telefono)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el cliente"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente creado exitosamente"})
}

// ListarClientes lista todos los clientes
func (h *ClienteHandler) ListarClientes(c *gin.Context) {
	rows, err := h.db.Query("SELECT cliente_id, nombre, nit, direccion, telefono FROM clientes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	clientes := make([]Cliente, 0)
	for rows.Next() {
		var cliente Cliente
		if err := rows.Scan(&cliente.ID, &cliente.Nombre, &cliente.NIT, &cliente.Direccion, &cliente.Telefono); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		clientes = append(clientes, cliente)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clientes)
}

// ActualizarCliente actualiza los datos de un cliente existente
func (h *ClienteHandler) ActualizarCliente(c *gin.Context) {
	var cliente Cliente

	// Decodificar el cuerpo JSON de la solicitud en la estructura del cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de cliente inválidos"})
		return
	}

	// Verificar si el cliente que se intenta actualizar existe
	if !h.existeClienteID(cliente.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El cliente no existe"})
		return
	}

	// Actualizar los datos del cliente en la base de datos
	_, err := h.db.Exec("UPDATE clientes SET nombre=?, nit=?, direccion=?, telefono=? WHERE cliente_id=?",
		cliente.Nombre, cliente.NIT, cliente.Direccion, cliente.Telefono, cliente.ID)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL de actualización: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el cliente"})
		return
	}

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Cliente actualizado exitosamente"})
}

// EliminarCliente elimina un cliente existente
func (h *ClienteHandler) EliminarCliente(c *gin.Context) {
	// Obtener el ID del cliente de los parámetros de la ruta
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cliente inválido"})
		return
	}

	// Verificar si el cliente que se intenta eliminar existe
	if !h.existeClienteID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El cliente no existe"})
		return
	}

	// Eliminar el cliente de la base de datos
	_, err = h.db.Exec("DELETE FROM clientes WHERE cliente_id = ?", id)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL de eliminación: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el cliente"})
		return
	}

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Cliente eliminado exitosamente"})
}

// existeClienteID verifica si ya existe un cliente con el mismo ID en la base de datos
func (h *ClienteHandler) existeClienteID(id int) bool {
	var count int
	err := h.db.QueryRow("SELECT COUNT(*) FROM clientes WHERE cliente_id = ?", id).Scan(&count)
	if err != nil {
		log.Printf("Error al verificar la existencia del cliente:  %v", err)
		return true
	}
	return count > 0
}
