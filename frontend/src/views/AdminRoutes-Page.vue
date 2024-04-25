<template>
  <div class="admin-container">
    <div class="header">
      <h1>Administrar Rutas</h1>
    </div>

    <div class="options">
      <h2 class="option-title">Opciones:</h2>
      <ul>
        <li><router-link to="/admin/routes/crear">Crear Ruta</router-link></li>
        <li><router-link to="/admin/routes/listar">Listar Rutas</router-link></li>
        <li><router-link to="/admin/routes/actualizar">Actualizar Ruta</router-link></li>
        <li><router-link to="/admin/routes/eliminar">Eliminar Ruta</router-link></li>
        <li><router-link to="/admin/routes/tarifa_operacional">Tarifa Operacional</router-link></li>
        <li><router-link to="/admin/routes/actualizar_tarifa_global">Actualizar Tarifa Global</router-link></li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {

    };
  }
}
</script>

<style scoped>
.admin-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.header {
  background-color: #212121; /* Color de fondo del encabezado */
  padding: 20px;
  text-align: center;
}

.header h1 {
  font-size: 1.8em; /* Texto más grande */
  color: #fff;
  margin: 0;
}

.options {
  margin-top: 20px;
}

.option-title {
  font-size: 1.4em; /* Texto más grande */
  margin-bottom: 10px;
}

ul {
  list-style: none;
  padding: 0;
}

li {
  font-size: 1.2em; /* Texto más grande */
  margin-bottom: 10px;
}

a {
  text-decoration: none;
  color: #007bff; /* Color azul de enlaces */
  font-size: 1.2em; /* Texto más grande */
}

a:hover {
  text-decoration: underline;
}
</style>