<template>
  <div class="container">
    <h1>Actualizar Ruta</h1>
    <form @submit.prevent="actualizarRuta" class="ruta-update-form">
      <div class="form-group">
        <label for="id">ID de la Ruta a Actualizar</label>
        <input type="text" id="id" v-model="ruta.ID" required>
      </div>
      <div class="form-group">
        <label for="destino">Destino</label>
        <input type="text" id="destino" v-model="ruta.Destino" required>
      </div>
      <div class="form-group">
        <label for="tarifa">Tarifa de Operación</label>
        <input type="text" id="tarifa" v-model="ruta.TarifaOperacion" required>
      </div>
      <div class="form-group">
        <label for="capacidad">Capacidad Máxima</label>
        <input type="text" id="capacidad" v-model="ruta.CapacidadMaxima" required>
      </div>
      <button type="submit" class="btn-submit">Actualizar Ruta</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      ruta: {
        ID: 0,
        Destino: '',
        TarifaOperacion: 0,
        CapacidadMaxima: 0
      }
    }
  },
  methods: {
    async actualizarRuta() {
      try {
        // Convertir el valor de ID a un número entero
        this.ruta.ID = parseInt(this.ruta.ID);
        this.ruta.TarifaOperacion = parseFloat(this.ruta.TarifaOperacion);
        this.ruta.CapacidadMaxima = parseInt(this.ruta.CapacidadMaxima);
        const response = await axios.put(`http://localhost:8080/actualizar_ruta`, {
          ID: this.ruta.ID,
          Destino: this.ruta.Destino,
          TarifaOperacion: this.ruta.TarifaOperacion,
          CapacidadMaxima: this.ruta.CapacidadMaxima
        });
        alert(response.data.message); // Mostrar alerta con el mensaje de éxito
      } catch (error) {
        alert('No se pudo actualizar la ruta'); // Mostrar alerta en caso de error
        console.error(error);
        // Manejar el error y mostrar un mensaje al usuario si la actualización falla
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
  background-color: #007BFF;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
</style>