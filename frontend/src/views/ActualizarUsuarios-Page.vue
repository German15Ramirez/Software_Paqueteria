<template>
  <div class="container">
    <h1>Actualizar Usuario</h1>
    <form @submit.prevent="actualizarUsuario" class="user-update-form">
      <div class="form-group">
        <label for="id">ID del Usuario</label>
        <input type="number" id="id" v-model="usuarioID" required>
      </div>
      <div class="form-group">
        <label for="nombre">Nombre</label>
        <input type="text" id="nombre" v-model="nombre" required>
      </div>
      <div class="form-group">
        <label for="correo">Correo Electrónico</label>
        <input type="email" id="correo" v-model="correo" required>
      </div>
      <div class="form-group">
        <label for="telefono">Teléfono</label>
        <input type="number" id="telefono" v-model="telefono" required>
      </div>
      <div class="form-group">
        <label for="contraseña">Contraseña</label>
        <input type="password" id="contraseña" v-model="contraseña" required>
      </div>
      <div class="form-group">
        <label for="rol">Rol</label>
        <select id="rol" v-model="rol" required>
          <option value="admin">Administrador</option>
          <option value="user">Recepcionista</option>
          <option value="admin">Operador</option>
        </select>
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
      usuarioID: 0,
      nombre: '',
      correo: '',
      telefono: 0, // Cambiado a tipo entero (int)
      contraseña: '',
      rol: ''
    }
  },
  methods: {
    async actualizarUsuario() {
      try {
        const response = await axios.put(`http://localhost:8080/actualizar_usuario`, {
          usuarioID: this.usuarioID,
          nombre: this.nombre,
          correo: this.correo,
          telefono: parseInt(this.telefono), // Convertir a entero antes de enviar al servidor
          contraseña: this.contraseña,
          rol: this.rol
        });
        console.log(response.data.message); // Mostrar mensaje de éxito en la consola
        // También podrías mostrar un mensaje de éxito en la interfaz
      } catch (error) {
        console.error(error);
        // Manejar el error y mostrar un mensaje al usuario si la actualización falla
      }
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

input, select {
  width: 100%;
  padding: 0.7em;
  border: 1px solid #ccc;
  border-radius: 5px;
}

button.btn-submit {
  margin-top: 1em;
  padding: 0.7em 1em;
  background-color: #007BFF;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
</style>