<template>
  <div class="container">
    <h1>Iniciar Sesión</h1>
    <form @submit.prevent="login" class="login-form">
      <div class="form-group">
        <label for="userId">Usuario:</label>
        <input type="text" id="userId" v-model="userId" required>
      </div>
      <div class="form-group">
        <label for="password">Contraseña:</label>
        <div class="password-input">
          <input :type="showPassword ? 'text' : 'password'" id="password" v-model="password" required>
          <button @click.prevent="togglePasswordVisibility">{{ showPassword ? 'Ocultar' : 'Mostrar' }}</button>
        </div>
      </div>
      <button type="submit" class="btn-submit">Iniciar Sesión</button>
      <p v-if="error" class="error-message">{{ error }}</p>
    </form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      userId: '',
      password: '',
      error: '',
      showPassword: false
    };
  },
  methods: {
    togglePasswordVisibility() {
      this.showPassword = !this.showPassword;
    },
    async login() {
      if (!this.userId || !this.password) {
        this.error = 'Por favor, completa todos los campos.';
        return;
      }
      try {
        const response = await this.verificarCredenciales();
        this.handleLoginResponse(response);
      } catch (error) {
        this.error = 'Error al iniciar sesión. Por favor, intenta de nuevo.';
        console.error('Error al iniciar sesión:', error);
      }
    },
    async verificarCredenciales() {
      try {
        const response = await axios.post('http://localhost:8080/verificar_credenciales', {
          usuario_id: parseInt(this.userId),
          contraseña: this.password
        });
        return response.data;
      } catch (error) {
        console.error('Error al verificar las credenciales:', error);
        throw error;
      }
    },
    handleLoginResponse(response) {
      switch (response.rol) {
        case 'Administrador':
          this.$router.push('/AdministradorPrincipal-Page');
          break;
        case 'Recepcionista':
          this.$router.push('/RecepcionistaPrincipal-Page');
          break;
        case 'Operador':
          this.$router.push('/OperadorPrincipal-Page');
          break;
        default:
          this.error = 'Rol desconocido';
          console.error('Rol desconocido:', response.rol);
      }
    }
  }
};
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
  margin-top: 1em;
  padding: 0.7em 1em;
  background-color: #14b642;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
</style>