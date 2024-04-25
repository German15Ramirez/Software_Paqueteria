<template>
  <div class="container">
    <h1>Listar Puntos de Control</h1>
    <div class="table-container">
      <table>
        <thead>
        <tr>
          <th>ID</th>
          <th>Ruta ID</th>
          <th>Nombre</th>
          <th>Capacidad de Cola</th>
          <th>Tarifa Operacional</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="puntoDeControl in puntosDeControl" :key="puntoDeControl.ID">
          <td>{{ puntoDeControl.ID }}</td>
          <td>{{ puntoDeControl.RutaID }}</td>
          <td>{{ puntoDeControl.Nombre }}</td>
          <td>{{ puntoDeControl.CapacidadCola }}</td>
          <td>{{ puntoDeControl.TarifaOperacional }}</td>
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
      puntosDeControl: []
    };
  },
  mounted() {
    this.listarPuntosDeControl();
  },
  methods: {
    async listarPuntosDeControl() {
      try {
        const response = await axios.get('http://localhost:8080/listar_puntos_de_control');
        this.puntosDeControl = response.data;
      } catch (error) {
        console.error('Error al obtener puntos de control:', error);
      }
    }
  }
}
</script>

<style scoped>
.container {
  background-color: #f9f9f9;
  padding: 20px;
  border-radius: 5px;
}

.table-container {
  background-color: #fff;
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 20px;
  margin-top: 20px;
  box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
}

table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  border: 1px solid #ddd;
  padding: 12px;
}

th {
  background-color: #f2f2f2;
}

tr:nth-child(even) {
  background-color: #f2f2f2;
}
</style>