<template>
  <div>
    <h1>Reporte de Rutas</h1>
    <div class="table-container">
      <table>
        <thead>
        <tr>
          <th>Ruta</th>
          <th>Paquetes en Ruta</th>
          <th>Paquetes Entregados</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(count, ruta) in paquetesEnRutaPorRuta" :key="ruta">
          <td>{{ ruta }}</td>
          <td>{{ count }}</td>
          <td>{{ paquetesEntregadosPorRuta[ruta] }}</td>
        </tr>
        </tbody>
      </table>
    </div>
    <button @click="obtenerReporte">Obtener Reporte</button>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      paquetesEnRutaPorRuta: {},
      paquetesEntregadosPorRuta: {}
    };
  },
  methods: {
    async obtenerReporte() {
      try {
        const response = await axios.get('http://localhost:8080/obtener_reporte_rutas');
        this.paquetesEnRutaPorRuta = response.data.paquetesEnRutaPorRuta;
        this.paquetesEntregadosPorRuta = response.data.paquetesEntregadosPorRuta;
      } catch (error) {
        console.error('Error al obtener el reporte de rutas:', error);
      }
    }
  }
}
</script>
<style scoped>
.table-container {
  border: 1px solid #000000;
  background-color: #f2f2f2;
  border-radius: 5px;
  padding: 10px;
  margin-top: 20px;
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
button {
  margin-top: 10px;
  padding: 10px 20px;
  background-color: #4CAF50;
  color: white;
  border: none;
  cursor: pointer;
}

button:hover {
  background-color: #45a049;
}
</style>