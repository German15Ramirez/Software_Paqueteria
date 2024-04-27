<template>
  <div class="login-container">
    <h1>Iniciar Sesión</h1>
    <form @submit.prevent="login" class="login-form">
      <div class="form-group">
        <label for="usuarioId">ID de Usuario:</label>
        <input type="text" id="usuarioId" v-model="userId" required>
      </div>
      <div class="form-group">
        <label for="password">Contraseña:</label>
        <input type="password" id="password" v-model="password" required>
      </div>
      <button type="submit" class="submit-button">Iniciar Sesión</button>
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
      error: ''
    };
  },
  methods: {
    async login() {
      try {
        if (!this.userId || !this.password) {
          this.error = 'Por favor, completa todos los campos.';
          return;
        }

        const response = await this.verificarCredenciales();

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
      } catch (error) {
        this.error = 'Error al iniciar sesión. Por favor, intenta de nuevo.';
        console.error('Error al iniciar sesión:', error);
      }
    },
    async verificarCredenciales() {
      try {
        const response = await axios.post('http://localhost:8080/verificar_credenciales', {
          usuario_id: parseInt(this.userId), // Convertir a entero
          contraseña: this.password // Mantener como cadena
        });

        return response.data;
      } catch (error) {
        console.error('Error al verificar las credenciales:', error);
        throw error;
      }
    }
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100vw;
  height: 100vh;
  background-color: #333;
}

.login-form {
  max-width: 400px;
  width: 80%;
  padding: 20px;
  border: 1px solid #555;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  background-color: #444;
}

h1 {
  font-size: 2em;
  margin-bottom: 20px;
  text-align: center;
  color: #fff;
}

.form-group {
  margin-bottom: 15px;
}

label {
  font-size: 1.2em;
  margin-bottom: 5px;
  color: #fff;
}

input {
  padding: 10px;
  font-size: 1em;
  width: 100%;
  border: 1px solid #555;
  border-radius: 3px;
  background-color: #666;
  color: #fff;
}

.submit-button {
  padding: 10px 20px;
  font-size: 1em;
  background-color: #42b983;
  color: #fff;
  border: none;
  border-radius: 3px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.submit-button:hover {
  background-color: #3aa872;
}

.error-message {
  color: #ff0000;
  margin-top: 10px;
}
</style>