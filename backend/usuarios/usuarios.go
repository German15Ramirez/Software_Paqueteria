package usuarios

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

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
	rows, err := h.db.Query("SELECT usuario_id, nombre, correo, telefono, contraseña, rol FROM usuarios ORDER BY usuario_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	usuarios := make([]Usuario, 0)
	for rows.Next() {
		var usuario Usuario
		if err := rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Correo, &usuario.Telefono, &usuario.Contraseña, &usuario.Rol); err != nil {
			log.Printf("Error al escanear fila: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar sobre filas: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Usuarios obtenidos de la base de datos: %v", usuarios)

	c.JSON(http.StatusOK, usuarios)
}
// ActualizarUsuario actualiza los datos de un usuario existente
func (h *UsuarioHandler) ActualizarUsuario(c *gin.Context) {
	var usuario Usuario

	// Decodificar el cuerpo JSON de la solicitud en la estructura del usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario inválidos"})
		return
	}

	// Verificar si el usuario que se intenta actualizar existe
	if !h.existeUsuarioID(usuario.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El usuario no existe"})
		return
	}

	// Actualizar los datos del usuario en la base de datos
	_, err := h.db.Exec("UPDATE usuarios SET nombre=?, correo=?, telefono=?, contraseña=?, rol=? WHERE usuario_id=?",
		usuario.Nombre, usuario.Correo, usuario.Telefono, usuario.Contraseña, usuario.Rol, usuario.ID)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL de actualización: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario"})
		return
	}

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente"})
}

// EliminarUsuario elimina un usuario existente
func (h *UsuarioHandler) EliminarUsuario(c *gin.Context) {
	// Obtener el ID del usuario de los parámetros de la ruta
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	// Verificar si el usuario que se intenta eliminar existe
	if !h.existeUsuarioID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El usuario no existe"})
		return
	}

	// Eliminar el usuario de la base de datos
	_, err = h.db.Exec("DELETE FROM usuarios WHERE usuario_id = ?", id)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL de eliminación: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
		return
	}

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
}

// existeUsuarioID verifica si ya existe un usuario con el mismo ID en la base de datos
func (h *UsuarioHandler) existeUsuarioID(id int) bool {
	var count int
	err := h.db.QueryRow("SELECT COUNT(*) FROM usuarios WHERE usuario_id = ?", id).Scan(&count)
	if err != nil {
		log.Printf("Error al verificar la existencia del usuario:  %v", err)
		return true
	}
	return count > 0
}

func (h *UsuarioHandler) VerificarCredenciales(c *gin.Context) {
	var credenciales struct {
		UsuarioID   int    `json:"usuario_id"`
		Contraseña  string `json:"contraseña"` // Cambiar a "contraseña" en lugar de "contraseña" para que coincida con el JSON del frontend
	}

	// Decodificar los datos JSON del cuerpo de la solicitud
	if err := c.ShouldBindJSON(&credenciales); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de credenciales inválidos"})
		return
	}

	// Verificar las credenciales del usuario
	rol, err := h.VerificarCredencialesBD(c, credenciales.UsuarioID, credenciales.Contraseña)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	// Devolver el rol del usuario
	c.JSON(http.StatusOK, gin.H{"rol": rol})
}

// VerificarCredencialesBD verifica las credenciales del usuario en la base de datos
func (h *UsuarioHandler) VerificarCredencialesBD(c *gin.Context, usuarioID int, contraseña string) (string, error) {
	var rol string
	err := h.db.QueryRow("SELECT rol FROM usuarios WHERE usuario_id = ? AND contraseña = ?", usuarioID, contraseña).Scan(&rol)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("Usuario o contraseña incorrectos")
		}
		return "", err // Otro error
	}
	return rol, nil
}