<template>
  <div class="list-packages-container">
    <div class="header">
      <h1>Listar Paquetes en la Ruta</h1>
    </div>

    <div class="route-info">
      <p>Ruta: {{ routeName }}</p>
    </div>

    <div class="package-list">
      <h2>Lista de Paquetes</h2>
      <table>
        <thead>
        <tr>
          <th>ID</th>
          <th>Descripción</th>
          <th>Estado</th>
          <th>Acciones</th>
        </tr>
        </thead>
        <tbody>
        <!-- Aquí se listarán los paquetes -->
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      routeName: null
    };
  },
  mounted() {
    // Lógica para obtener el nombre de la ruta
    // Supongamos que obtienes el nombre de la ruta correctamente
    this.routeName = "Nombre de la Ruta obtenido desde el backend";
  }
}
</script>

<style scoped>
.list-packages-container {
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

.route-info {
  margin-top: 20px;
}

.route-info p {
  font-size: 1.2em; /* Texto más grande */
}

.package-list {
  margin-top: 20px;
}

.package-list h2 {
  font-size: 1.4em; /* Texto más grande */
  margin-bottom: 10px;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

th {
  background-color: #f2f2f2;
}
</style>