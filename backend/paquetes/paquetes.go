package paquetes
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
// obtenerListaDePaquetes obtiene la lista de paquetes desde la base de datos.
func (h *PaqueteHandler) obtenerListaDePaquetes() ([]Paquete, error) {
	fmt.Println("Obteniendo la lista de paquetes desde la base de datos...")
	rows, err := h.db.Query("SELECT paquete_id, cliente_id, peso_libras, destino, fecha_ingreso, fecha_salida, estado FROM paquetes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paquetes []Paquete
	for rows.Next() {
		var paquete Paquete
		if err := rows.Scan(&paquete.ID, &paquete.ClienteID, &paquete.PesoLibras, &paquete.Destino, &paquete.FechaIngreso, &paquete.FechaSalida, &paquete.Estado); err != nil {
			return nil, err
		}
		paquetes = append(paquetes, paquete)
	}

	// Agregar mensaje para mostrar la cantidad de paquetes obtenidos
	fmt.Println("Se han obtenido", len(paquetes), "paquetes desde la base de datos")

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return paquetes, nil
}
// ReporteRutas genera un reporte de las rutas y la cantidad de paquetes en cada estado por ruta.
func (h *PaqueteHandler) ReporteRutas(c *gin.Context) {
	// Obtener la lista de paquetes
	paquetes, err := h.obtenerListaDePaquetes()
	if err != nil {
		log.Printf("Error al obtener la lista de paquetes: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la lista de paquetes"})
		return
	}

	// Mapas para almacenar la cantidad de paquetes por estado para cada ruta
	paquetesEnRutaPorRuta := make(map[string]int)
	paquetesEntregadosPorRuta := make(map[string]int)

	// Iterar sobre los paquetes y contar la cantidad por estado para cada ruta
	for _, paquete := range paquetes {
		if paquete.Estado == "En Ruta" {
			paquetesEnRutaPorRuta[paquete.Ruta]++
		} else if paquete.Estado == "Entregado" {
			paquetesEntregadosPorRuta[paquete.Ruta]++
		}
	}
	responseData := map[string]interface{}{
		"paquetesEnRutaPorRuta":    paquetesEnRutaPorRuta,
		"paquetesEntregadosPorRuta": paquetesEntregadosPorRuta,
	}

	c.JSON(http.StatusOK, responseData)
}

// ProcesarPaquete maneja el procesamiento de un paquete
func (h *PaqueteHandler) ProcesarPaquete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de paquete inválido"})
		return
	}

	// Estructura para almacenar el cuerpo de la solicitud
	var data struct {
		TiempoPasado float64 `json:"tiempo_pasado"`
	}

	// Decodificar el cuerpo JSON de la solicitud en la estructura
	if err := c.BindJSON(&data); err != nil {
		log.Printf("Error al decodificar el cuerpo JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar si se proporcionó el tiempo pasado
	if data.TiempoPasado == 0 {
		log.Println("No se proporcionó el tiempo pasado")
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se proporcionó el tiempo pasado"})
		return
	}
	log.Printf("ID del paquete: %d", id)
	log.Printf("Tiempo pasado recibido: %f", data)
	// Obtener el punto de control actual del paquete
	var puntoControlActual int
	err = h.db.QueryRow("SELECT punto_de_control_id FROM paquetes WHERE paquete_id = ?", id).Scan(&puntoControlActual)
	if err != nil {
		log.Printf("Error al obtener el punto de control actual del paquete: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el punto de control actual del paquete"})
		return
	}

	// Incrementar el punto de control actual en 1
	puntoControlActual++

	fmt.Printf("Nuevo punto de control: %d\n", puntoControlActual)

	// Actualizar el punto de control actual del paquete en la base de datos
	_, err = h.db.Exec("UPDATE paquetes SET punto_de_control_id = ? WHERE paquete_id = ?", puntoControlActual, id)
	if err != nil {
		log.Printf("Error al mover el paquete al siguiente punto de control: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al mover el paquete al siguiente punto de control"})
		return
	}

	// Registrar el tiempo que el paquete pasó en el punto de control actual en la base de datos
	_, err = h.db.Exec("INSERT INTO tiempos_paquete (paquete_id, tiempo_pasado) VALUES (?, ?)", id, data.TiempoPasado)
	if err != nil {
		log.Printf("Error al registrar el tiempo que el paquete pasó en el punto de control: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el tiempo que el paquete pasó en el punto de control"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Procesamiento del paquete completado exitosamente"})
}
// ListarPaquetesListosEnDestino lista los paquetes listos en su destino
func (h *PaqueteHandler) ListarPaquetesListosEnDestino(c *gin.Context) {
	// Obtener el último punto de control de cada ruta
	rows, err := h.db.Query("SELECT ruta_id, MAX(puntos_de_control_id) FROM puntos_de_control GROUP BY ruta_id")
	if err != nil {
		log.Printf("Error al obtener los últimos puntos de control de cada ruta: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los puntos de control"})
		return
	}
	defer rows.Close()

	// Mapa para almacenar el último punto de control de cada ruta
	ultimosPuntosControl := make(map[int]int)

	for rows.Next() {
		var rutaID, ultimoPuntoControl int
		if err := rows.Scan(&rutaID, &ultimoPuntoControl); err != nil {
			log.Printf("Error al escanear filas: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los puntos de control"})
			return
		}
		ultimosPuntosControl[rutaID] = ultimoPuntoControl
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al escanear filas: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los puntos de control"})
		return
	}

	// Filtrar los paquetes que están listos en su destino y tienen el último punto de control
	rows, err = h.db.Query("SELECT paquete_id, cliente_id, destino, punto_de_control_id, ruta_id FROM paquetes WHERE estado = 'En Ruta'")
	if err != nil {
		log.Printf("Error al obtener los paquetes en ruta: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los paquetes"})
		return
	}
	defer rows.Close()

	var paquetesListos []Paquete
	for rows.Next() {
		var paquete Paquete
		if err := rows.Scan(&paquete.ID, &paquete.ClienteID, &paquete.Destino, &paquete.PuntoControl, &paquete.ID); err != nil {
			log.Printf("Error al escanear filas: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escanear filas"})
			return
		}

		// Convertir PuntoControl a int
		puntoControlInt, err := strconv.Atoi(paquete.PuntoControl)
		if err != nil {
			log.Printf("Error al convertir PuntoControl a int: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al convertir PuntoControl a int"})
			return
		}

		// Verificar si el punto de control del paquete es el último de su ruta
		if puntoControlInt == ultimosPuntosControl[paquete.ID] {
			paquetesListos = append(paquetesListos, paquete)
		}
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al escanear filas: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los paquetes"})
		return
	}

	c.JSON(http.StatusOK, paquetesListos)
}

// CrearPaquete crea un nuevo paquete para un cliente
func (h *PaqueteHandler) CrearPaquete(c *gin.Context) {
	// Estructura para almacenar el cuerpo de la solicitud
	var paquete Paquete
	if err := c.BindJSON(&paquete); err != nil {
		log.Printf("Error al decodificar el cuerpo JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Insertar el nuevo paquete en la base de datos
	result, err := h.db.Exec("INSERT INTO paquetes (cliente_id, nombrecliente, peso_libras, destino, fecha_ingreso, fecha_salida, estado, punto_de_control_id, ruta) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", paquete.ClienteID, paquete.NombreCliente, paquete.PesoLibras, paquete.Destino, paquete.FechaIngreso, paquete.FechaSalida, paquete.Estado, paquete.PuntoControl, paquete.Ruta)
	if err != nil {
		log.Printf("Error al insertar el nuevo paquete en la base de datos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar el nuevo paquete en la base de datos"})
		return
	}

	// Obtener el ID del paquete insertado
	paqueteID, _ := result.LastInsertId()

	// Asignar el ID al paquete
	paquete.ID = int(paqueteID)

	c.JSON(http.StatusOK, gin.H{"message": "Paquete creado exitosamente", "paquete": paquete})
}

// ListarPaquetesPorEstado lista los paquetes filtrados por estado.
func (h *PaqueteHandler) ListarPaquetesPorEstado(c *gin.Context) {
	estado := c.Query("estado") // Obtener el parámetro de consulta "estado"

	var rows *sql.Rows
	var err error

	if estado != "" {
		rows, err = h.db.Query("SELECT pa.paquete_id, pa.cliente_id, cl.nombre AS nombre_cliente, pa.peso_libras, pa.destino, pa.fecha_ingreso, pa.fecha_salida, pa.estado, pc.nombre AS punto_control, r.destino AS ruta FROM paquetes pa JOIN puntos_de_control pc ON pa.punto_de_control_id = pc.puntos_de_control_id JOIN rutas r ON pc.ruta_id = r.ruta_id JOIN clientes cl ON pa.cliente_id = cl.cliente_id WHERE pa.estado = ?", estado)
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