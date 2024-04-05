package usuarios
import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	ID        int
	Nombre    string
	Correo    string
	Telefono  int
	Contraseña string
	Rol       string
}

// UsuarioHandler maneja las operaciones relacionadas con los usuarios
type UsuarioHandler struct {
	db *sql.DB
}

// NewUsuarioHandler crea un nuevo manejador de usuarios
func NewUsuarioHandler(db *sql.DB) *UsuarioHandler {
	return &UsuarioHandler{db: db}
}

// CrearUsuario crea un nuevo usuario
func (h *UsuarioHandler) CrearUsuario(c *gin.Context) {
	var usuario Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario inválidos"})
		return
	}

	// Verificar si se proporcionó un valor para usuario_id
	if usuario.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere un valor para usuario_id"})
		return
	}

	_, err := h.db.Exec("INSERT INTO usuarios (usuario_id, nombre, correo, telefono, contraseña, rol) VALUES (?, ?, ?, ?, ?, ?)",
		usuario.ID, usuario.Nombre, usuario.Correo, usuario.Telefono, usuario.Contraseña, usuario.Rol)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario creado exitosamente"})
}

// ListarUsuarios lista todos los usuarios
func (h *UsuarioHandler) ListarUsuarios(c *gin.Context) {
	rows, err := h.db.Query("SELECT usuario_id, nombre, correo, telefono, contraseña, rol FROM usuarios")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	usuarios := make([]Usuario, 0)
	for rows.Next() {
		var usuario Usuario
		if err := rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Correo, &usuario.Telefono, &usuario.Contraseña, &usuario.Rol); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}