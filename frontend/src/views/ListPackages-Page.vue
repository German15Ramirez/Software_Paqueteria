<template>
  <div style="background-color: #f2f2f2; padding: 20px; border-radius: 5px;">
    <h1 style="background-color: #212121; color: #fff; padding: 10px; border-radius: 5px;">Listar Paquetes</h1>
    <div class="search-container">
      <input type="text" v-model="destino" placeholder="Ingrese el destino">
      <button @click="buscarPaquetes">Buscar</button>
    </div>
    <div class="table-container">
      <table>
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
      destino: '',
      paquetes: []
    };
  },
  methods: {
    async buscarPaquetes() {
      try {
        const response = await axios.get('http://localhost:8080/listar_paquetes', {
          params: { destino: this.destino }
        });
        this.paquetes = response.data;
      } catch (error) {
        console.error('Error al obtener paquetes:', error);
      }
    }
  }
}
</script>

<style scoped>
.search-container {
  margin-bottom: 20px;
}

.search-container input {
  width: 200px;
  padding: 5px;
  margin-right: 10px;
}

.search-container button {
  padding: 5px 10px;
  background-color: #4CAF50;
  color: white;
  border: none;
  cursor: pointer;
}

.search-container button:hover {
  background-color: #45a049;
}

.table-container {
  border: 1px solid #000000;
  background-color: #ccc;
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
</style>