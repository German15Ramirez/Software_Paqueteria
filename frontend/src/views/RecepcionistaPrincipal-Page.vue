<template>
  <div class="recep-container">
    <div class="header">
      <h1>Bienvenido de nuevo Recepcionista</h1>
    </div>

    <div class="options">
      <h2 class="option-title">Opciones de recepcionista:</h2>
      <ul>
        <li><router-link to="/recep/packages">Listar paquetes</router-link></li>
        <li><router-link to="/recep/locations">Localización de paquetes</router-link></li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      recepName: '' // Aquí se almacenará el nombre del recepcionista obtenido de la conexión SQL
    };
  },
  mounted() {
    // Aquí puedes realizar una consulta SQL para obtener el nombre del recepcionista
    // Ejemplo: this.recepName = consultarNombreRecepcionista();
    this.recepName = "Recepcionista"; // Simulando la obtención del nombre del recepcionista
  }
}
</script>

<style scoped>
.recep-container {
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