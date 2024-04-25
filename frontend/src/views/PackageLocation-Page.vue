<template>
  <div class="package-location-container">
    <div class="header">
      <h1>Localización del Paquete</h1>
    </div>

    <div class="package-info">
      <p>Paquete ID: {{ packageId }}</p>
      <p>Ubicación: {{ location }}</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      packageId: null,
      location: null
    };
  },
  mounted() {
    // Lógica para obtener la localización del paquete utilizando la ruta /paquetes/:id/localizacion
    const packageId = this.$route.params.id;

    // Aquí deberías hacer una petición al backend para obtener la localización del paquete
    // Supongamos que obtienes la localización correctamente
    this.packageId = packageId;
    this.location = "Ubicación del paquete obtenida desde el backend";
  }
}
</script>

<style scoped>
.package-location-container {
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

.package-info {
  margin-top: 20px;
}

.package-info p {
  font-size: 1.2em; /* Texto más grande */
}
</style>