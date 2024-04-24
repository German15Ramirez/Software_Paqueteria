<template>
  <div class="admin-container">
    <div class="header">
      <h1>Bienvenido de nuevo administrador: {{ adminName }}</h1>
    </div>

    <div class="options">
      <h2>Opciones de administrador:</h2>
      <ul>
        <li><router-link to="/admin/users">Administrar usuarios</router-link></li>
        <li><router-link to="/admin/packages">Administrar paquetes</router-link></li>
        <li><router-link to="/admin/reports">Ver reportes</router-link></li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      adminName: '' // Aquí se almacenará el nombre del administrador obtenido de la conexión SQL
    };
  },
  mounted() {
    // Aquí puedes realizar una consulta SQL para obtener el nombre del administrador
    // Ejemplo: this.adminName = consultarNombreAdmin();
    this.adminName = "Administrador"; // Simulando la obtención del nombre del administrador
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
  font-size: 1.5em;
  color: #fff;
  margin: 0;
}

.options {
  margin-top: 20px;
}

.options h2 {
  font-size: 1.2em;
  margin-bottom: 10px;
}

ul {
  list-style: none;
  padding: 0;
}

li {
  margin-bottom: 10px;
}

a {
  text-decoration: none;
  color: #007bff; /* Color azul de enlaces */
}

a:hover {
  text-decoration: underline;
}
</style>