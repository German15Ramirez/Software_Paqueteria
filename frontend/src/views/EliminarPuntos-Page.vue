<template>
  <div class="container">
    <h1>Eliminar Punto de Control</h1>
    <form @submit.prevent="eliminarPuntoDeControl" class="punto-de-control-delete-form">
      <div class="form-group">
        <label for="id">ID del Punto de Control a Eliminar</label>
        <input type="text" id="id" v-model="puntoDeControlID" required>
      </div>
      <button type="submit" class="btn-submit">Eliminar Punto de Control</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      puntoDeControlID: ''
    }
  },
  methods: {
    async eliminarPuntoDeControl() {
      try {
        const response = await axios.delete(`http://localhost:8080/eliminar_punto_de_control/${this.puntoDeControlID}`);
        alert(response.data.message); // Mostrar alerta con el mensaje de éxito o error
        this.puntoDeControlID = ''; // Limpiar el campo de ID del punto de control
      } catch (error) {
        alert('No se pudo eliminar el punto de control, motivo [desconocido]'); // Mostrar alerta en caso de error
        console.error(error);
        // Manejar el error y mostrar un mensaje al usuario si la eliminación falla
      }
    }
  }
}
</script>

<style scoped>
.container {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #e0e0e0;
  border-radius: 5px;
  background-color: #f9f9f9;
}

.form-group {
  margin-bottom: 1.5em;
}

label {
  display: block;
  margin-bottom: 0.5em;
  font-weight: bold;
}

input, select {
  width: 100%;
  padding: 0.7em;
  border: 1px solid #ccc;
  border-radius: 5px;
}

button.btn-submit {
  margin-top: 1em;
  padding: 0.7em 1em;
  background-color: #DC3545;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
</style>