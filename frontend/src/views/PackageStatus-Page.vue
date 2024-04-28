<template>
  <div class="container">
    <h1 style="background-color: #212121; color: #fff; padding: 10px; border-radius: 5px;">Listar Paquetes por Estado</h1>
    <form @submit.prevent="listarPaquetes" class="package-list-form">
      <div class="form-group">
        <label for="estado">Estado del Paquete</label>
        <input type="text" id="estado" v-model="estado" required>
      </div>
      <button type="submit" class="btn-submit">Buscar Paquetes</button>
    </form>
    <div v-if="paquetes.length > 0">
      <h2>Paquetes encontrados</h2>
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
        <tr v-for="paquete in paquetes" :key="paquete.ID">
          <td>{{ paquete.ID }}</td>
          <td>{{ paquete.NombreCliente }}</td>
          <td>{{ paquete.PesoLibras }}</td>
          <td>{{ paquete.Destino }}</td>
          <td>{{ paquete.FechaIngreso }}</td>
          <td>{{ paquete.FechaSalida }}</td>
          <td>{{ paquete.Estado }}</td>
          <td>{{ paquete.PuntoControl }}</td>
          <td>{{ paquete.Ruta }}</td>
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
      estado: '',
      paquetes: []
    }
  },
  methods: {
    async listarPaquetes() {
      try {
        const response = await axios.get(`http://localhost:8080/paquetes/estado?estado=${this.estado}`);
        this.paquetes = response.data;
      } catch (error) {
        console.error(error);
        alert('Error al obtener la lista de paquetes');
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
  border: 1px solid #000000;
  border-radius: 5px;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  border: 1px solid #000000;
  padding: 8px;
}

th {
  background-color: #f2f2f2;
}

tr:nth-child(even) {
  background-color: #f2f2f2;
}
</style>
<script setup>
</script>