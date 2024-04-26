<template>
  <div class="container">
    <h1>Crear Punto de Control</h1>
    <form @submit.prevent="crearPuntoDeControl" class="checkpoint-form">
      <div class="form-group">
        <label for="id">ID del Punto de Control</label>
        <input type="number" id="id" v-model="puntoDeControl.ID" required>
      </div>
      <div class="form-group">
        <label for="rutaID">ID de la Ruta</label>
        <input type="number" id="rutaID" v-model="puntoDeControl.RutaID" required>
      </div>
      <div class="form-group">
        <label for="nombre">Nombre</label>
        <input type="text" id="nombre" v-model="puntoDeControl.Nombre" required>
      </div>
      <div class="form-group">
        <label for="capacidadCola">Capacidad de la Cola</label>
        <input type="number" id="capacidadCola" v-model="puntoDeControl.CapacidadCola" required>
      </div>
      <div class="form-group">
        <label for="tarifaOperacional">Tarifa Operacional</label>
        <input type="number" id="tarifaOperacional" v-model="puntoDeControl.TarifaOperacional" required>
      </div>
      <button type="submit" class="btn-submit">Crear Punto de Control</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      puntoDeControl: {
        ID: '',
        RutaID: '',
        Nombre: '',
        CapacidadCola: '',
        TarifaOperacional: ''
      }
    }
  },
  methods: {
    async crearPuntoDeControl() {
      try {
        const response = await axios.post('http://localhost:8080/crear_punto_de_control', this.puntoDeControl);
        console.log(response.data);
        if (response.data.message) {
          alert(response.data.message);
          this.limpiarCampos();
        }
      } catch (error) {
        console.error(error);
        alert('Error al crear el punto de control');
      }
    },
    limpiarCampos() {
      this.puntoDeControl = {
        ID: '',
        RutaID: '',
        Nombre: '',
        CapacidadCola: '',
        TarifaOperacional: ''
      };
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

input {
  width: 100%;
  padding: 0.7em;
  border: 1px solid #ccc;
  border-radius: 5px;
}

button.btn-submit {
  padding: 0.7em 1em;
  background-color: #007BFF;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
</style>