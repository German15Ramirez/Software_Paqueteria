<template>
  <div class="process-package-container">
    <div class="header">
      <h1>Procesar Paquete</h1>
    </div>

    <div class="package-info">
      <p>Paquete ID: {{ packageId }}</p>
      <p>Tiempo en el punto de control actual: {{ currentTime }}</p>
    </div>

    <div class="form">
      <label for="newTime">Nuevo tiempo en el punto de control:</label>
      <input type="text" id="newTime" v-model="newTime">

      <button @click="updateTime">Actualizar Tiempo</button>
    </div>

    <div class="next-checkpoint">
      <button @click="moveToNextCheckpoint">Pasar al Siguiente Punto de Control</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      packageId: null,
      currentTime: null,
      newTime: null
    };
  },
  mounted() {
    // Lógica para obtener la información del paquete y su tiempo actual en el punto de control
    const packageId = this.$route.params.id;

    // Supongamos que obtienes la información del paquete y su tiempo actual correctamente
    this.packageId = packageId;
    this.currentTime = "Tiempo actual obtenido desde el backend";
  },
  methods: {
    updateTime() {
      // Lógica para actualizar el tiempo del paquete en el punto de control actual
      console.log("Nuevo tiempo:", this.newTime);
      // Aquí deberías hacer una petición al backend para actualizar el tiempo del paquete
    },
    moveToNextCheckpoint() {
      // Lógica para pasar al siguiente punto de control
      console.log("Pasar al siguiente punto de control");
      // Aquí deberías hacer una petición al backend para mover el paquete al siguiente punto de control
    }
  }
}
</script>

<style scoped>
.process-package-container {
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

.form {
  margin-top: 20px;
}

.form label {
  font-size: 1.2em; /* Texto más grande */
  margin-bottom: 10px;
}

.form input {
  width: 100%;
  padding: 8px;
  font-size: 1.2em; /* Texto más grande */
  margin-bottom: 10px;
}

.next-checkpoint {
  margin-top: 20px;
}

.next-checkpoint button {
  font-size: 1.2em; /* Texto más grande */
  padding: 8px 16px;
  cursor: pointer;
}
</style>