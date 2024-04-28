<template>
  <div class="container">
    <h1>Listar Clientes</h1>
    <div class="table-container">
      <table>
        <thead>
        <tr>
          <th>ID</th>
          <th>Nombre</th>
          <th>NIT</th>
          <th>Dirección</th>
          <th>Teléfono</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="cliente in clientes" :key="cliente.ID">
          <td>{{ cliente.ID }}</td>
          <td>{{ cliente.Nombre }}</td>
          <td>{{ cliente.NIT }}</td>
          <td>{{ cliente.Direccion }}</td>
          <td>{{ cliente.Telefono }}</td>
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
      clientes: []
    };
  },
  mounted() {
    this.listarClientes();
  },
  methods: {
    async listarClientes() {
      try {
        const response = await axios.get('http://localhost:8080/listar_clientes');
        this.clientes = response.data;
      } catch (error) {
        console.error('Error al obtener clientes:', error);
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
  border: 1px solid #151515;
  border-radius: 5px;
  padding: 20px;
  margin-top: 20px;
}

table {
  width: 100%;
  border-collapse: collapse;
  font-size: 16px;
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