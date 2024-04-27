<template>
  <div class="container">
    <h1>Editar Usuario</h1>
    <form @submit.prevent="editarUsuario" class="user-form">
      <div class="form-group">
        <label for="id">ID</label>
        <input type="number" id="id" v-model="usuario.ID" required>
      </div>
      <div class="form-group">
        <label for="nombre">Nombre</label>
        <input type="text" id="nombre" v-model="usuario.Nombre" required>
      </div>
      <div class="form-group">
        <label for="correo">Correo</label>
        <input type="email" id="correo" v-model="usuario.Correo" required>
      </div>
      <div class="form-group">
        <label for="telefono">Teléfono</label>
        <input type="number" id="telefono" v-model="usuario.Telefono" required>
      </div>
      <div class="form-group">
        <label for="contraseña">Contraseña</label>
        <input type="password" id="contraseña" v-model="usuario.Contraseña" required>
      </div>
      <div class="form-group">
        <label for="rol">Rol</label>
        <input type="text" id="rol" v-model="usuario.Rol" required>
      </div>
      <button type="submit" class="btn-submit">Actualizar Usuario</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      usuario: {
        ID: '',
        Nombre: '',
        Correo: '',
        Telefono: '',
        Contraseña: '',
        Rol: ''
      }
    }
  },
  methods: {
    async editarUsuario() {
      try {
        const response = await axios.put('http://localhost:8080/actualizar_usuario', this.usuario);
        console.log(response.data);
        if (response.data.message) {
          alert(response.data.message);
          this.limpiarCampos();
        }
      } catch (error) {
        console.error(error);
        alert('Error al actualizar el usuario');
      }
    },
    limpiarCampos() {
      this.usuario = {
        ID: '',
        Nombre: '',
        Correo: '',
        Telefono: '',
        Contraseña: '',
        Rol: ''
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