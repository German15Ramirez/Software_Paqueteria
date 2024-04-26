<template>
  <div class="container">
    <h1>Localizar Paquete</h1>
    <form @submit.prevent="obtenerLocalizacion" class="package-location-form">
      <div class="form-group">
        <label for="id">ID del Paquete</label>
        <input type="text" id="id" v-model="paqueteID" required>
      </div>
      <button type="submit" class="btn-submit">Localizar Paquete</button>
    </form>
    <div v-if="localizacionPaquete">
      <h2>Localización del Paquete</h2>
      <table class="table-container">
        <thead>
        <tr>
          <th>ID</th>
          <th>Cliente</th>
          <th>Peso (libras)</th>
          <th>Destino</th>
          <th>Fecha de Ingreso</th>
          <th>Fecha de Salida</th>
          <th>Estado</th>
          <th>Punto de Control</th>
          <th>Ruta</th>
        </tr>
        </thead>
        <tbody>
        <tr>
          <td>{{ localizacionPaquete.ID }}</td>
          <td>{{ localizacionPaquete.ClienteID }}</td>
          <td>{{ localizacionPaquete.PesoLibras }}</td>
          <td>{{ localizacionPaquete.Destino }}</td>
          <td>{{ localizacionPaquete.FechaIngreso }}</td>
          <td>{{ localizacionPaquete.FechaSalida }}</td>
          <td>{{ localizacionPaquete.Estado }}</td>
          <td>{{ localizacionPaquete.PuntoControl }}</td>
          <td>{{ localizacionPaquete.Ruta }}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      paqueteID: '',
      localizacionPaquete: null
    }
  },
  methods: {
    async obtenerLocalizacion() {
      try {
        const response = await axios.get(`http://localhost:8080/paquetes/${this.paqueteID}/localizacion`);
        this.localizacionPaquete = response.data;
      } catch (error) {
        console.error(error);
        alert('Error al obtener la localización del paquete');
      }
    }
  }
}
</script>

<style scoped>
.container {
  max-width: 800px;
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

.table-container {
  margin-top: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  border: 1px solid #ddd;
  padding: 8px;
}

th {
  background-color: #f2f2f2;
}

tr:nth-child(even) {
  background-color: #f2f2f2;
}
</style>