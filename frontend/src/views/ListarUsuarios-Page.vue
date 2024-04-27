<template>
  <div class="container">
    <h1>Listar Usuarios</h1>
    <div class="table-container">
      <table>
        <thead>
        <tr>
          <th>ID</th>
          <th>Nombre</th>
          <th>Correo</th>
          <th>Teléfono</th>
          <th>Rol</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="usuario in usuarios" :key="usuario.ID">
          <td>{{ usuario.ID }}</td>
          <td>{{ usuario.Nombre }}</td>
          <td>{{ usuario.Correo }}</td>
          <td>{{ usuario.Telefono }}</td>
          <td>{{ usuario.Rol }}</td>
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
      usuarios: []
    };
  },
  mounted() {
    this.listarUsuarios();
  },
  methods: {
    async listarUsuarios() {
      try {
        const response = await axios.get('http://localhost:8080/listar_usuarios');
        this.usuarios = response.data;
      } catch (error) {
        console.error('Error al obtener usuarios:', error);
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