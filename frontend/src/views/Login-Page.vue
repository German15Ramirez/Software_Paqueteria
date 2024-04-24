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
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userId: '',
      password: ''
    };
  },
  methods: {
    login() {
      // Aquí puedes realizar la lógica de inicio de sesión
      console.log('Iniciando sesión con usuario:', this.userId);
      console.log('Contraseña:', this.password);
      // Reiniciar los campos después del inicio de sesión
      this.userId = '';
      this.password = '';
    }
  }
};
</script>

<style scoped>
/* Estilos específicos del componente Login */
.login-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100vw; /* Ocupa todo el ancho de la ventana */
  height: 100vh; /* Ocupa todo el alto de la ventana */
  background-color: #333;
}

.login-form {
  max-width: 400px;
  width: 80%; /* Utiliza el 80% del ancho del contenedor padre */
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
</style>