<template>
  <div class="container">
    <h1>Listar Rutas</h1>
    <div class="table-container">
      <table>
        <thead>
        <tr>
          <th>ID</th>
          <th>Destino</th>
          <th>Tarifa de Operación</th>
          <th>Capacidad Máxima</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="ruta in rutas" :key="ruta.ID">
          <td>{{ ruta.ID }}</td>
          <td>{{ ruta.Destino }}</td>
          <td>{{ ruta.TarifaOperacion }}</td>
          <td>{{ ruta.CapacidadMaxima }}</td>
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
      rutas: []
    };
  },
  mounted() {
    this.listarRutas();
  },
  methods: {
    async listarRutas() {
      try {
        const response = await axios.get('http://localhost:8080/listar_rutas');
        this.rutas = response.data;
      } catch (error) {
        console.error('Error al obtener rutas:', error);
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
  border: 1px solid #000000;
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
  border: 1px solid #000000;
  padding: 12px;
}

th {
  background-color: #f2f2f2;
}

tr:nth-child(even) {
  background-color: #f2f2f2;
}
</style>