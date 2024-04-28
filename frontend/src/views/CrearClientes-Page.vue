<template>
  <div class="container">
    <h1>Crear Cliente</h1>
    <form @submit.prevent="crearCliente" class="client-form">
      <div class="form-group">
        <label for="id">ID</label>
        <input type="number" id="id" v-model="cliente.ID" required>
      </div>
      <div class="form-group">
        <label for="nombre">Nombre</label>
        <input type="text" id="nombre" v-model="cliente.Nombre" required>
      </div>
      <div class="form-group">
        <label for="nit">NIT</label>
        <input type="number" id="nit" v-model="cliente.NIT" required>
      </div>
      <div class="form-group">
        <label for="direccion">Dirección</label>
        <input type="text" id="direccion" v-model="cliente.Direccion" required>
      </div>
      <div class="form-group">
        <label for="telefono">Teléfono</label>
        <input type="number" id="telefono" v-model="cliente.Telefono" required>
      </div>
      <button type="submit" class="btn-submit">Crear Cliente</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      cliente: {
        ID: '',
        Nombre: '',
        NIT: '',
        Direccion: '',
        Telefono: ''
      }
    }
  },
  methods: {
    async crearCliente() {
      try {
        const response = await axios.post('http://localhost:8080/crear_cliente', this.cliente);
        console.log(response.data);
        if (response.data.message) {
          alert(response.data.message);
          this.limpiarCampos();
        }
      } catch (error) {
        console.error(error);
        alert('Error al crear el cliente');
      }
    },
    limpiarCampos() {
      this.cliente = {
        ID: '',
        Nombre: '',
        NIT: '',
        Direccion: '',
        Telefono: ''
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