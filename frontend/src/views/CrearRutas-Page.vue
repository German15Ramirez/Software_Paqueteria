<template>
  <div class="container">
    <h1>Crear Ruta</h1>
    <form @submit.prevent="crearRuta" class="route-form">
      <div class="form-group">
        <label for="id">ID de la Ruta</label>
        <input type="number" id="id" v-model="ruta.ID" required>
      </div>
      <div class="form-group">
        <label for="destino">Destino</label>
        <input type="text" id="destino" v-model="ruta.Destino" required>
      </div>
      <div class="form-group">
        <label for="tarifaOperacion">Tarifa de Operación</label>
        <input type="number" id="tarifaOperacion" v-model="ruta.TarifaOperacion" required>
      </div>
      <div class="form-group">
        <label for="capacidadMaxima">Capacidad Máxima</label>
        <input type="number" id="capacidadMaxima" v-model="ruta.CapacidadMaxima" required>
      </div>
      <button type="submit" class="btn-submit">Crear Ruta</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      ruta: {
        ID: '',
        Destino: '',
        TarifaOperacion: '',
        CapacidadMaxima: ''
      }
    }
  },
  methods: {
    async crearRuta() {
      try {
        const response = await axios.post('http://localhost:8080/crear_ruta', this.ruta);
        console.log(response.data);
        if (response.data.message) {
          alert(response.data.message);
          this.limpiarCampos();
        }
      } catch (error) {
        console.error(error);
        alert('Error al crear la ruta');
      }
    },
    limpiarCampos() {
      this.ruta = {
        ID: '',
        Destino: '',
        TarifaOperacion: '',
        CapacidadMaxima: ''
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