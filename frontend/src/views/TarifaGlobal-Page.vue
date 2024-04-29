<template>
  <div class="container">
    <h1>Configuración de Tarifa Global</h1>
    <form @submit.prevent="setTarifaGlobal" class="tarifa-form">
      <div class="form-group">
        <label for="tarifa">Ingrese la tarifa global:</label>
        <input type="number" id="tarifa" v-model="tarifa" required>
      </div>
      <button type="submit" class="btn-submit">Establecer Tarifa Global</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      tarifa: null
    };
  },
  methods: {
    setTarifaGlobal() {
      axios.post('http://localhost:8080/tarifa-global', { tarifa: this.tarifa })
          .then(response => {
            alert(response.data.message); // Mostrar alerta con el mensaje de éxito
            this.tarifa = null; // Limpiar el campo de tarifa después de establecerla
          })
          .catch(error => {
            console.error('Error al establecer la tarifa global:', error);
            // Manejar el error y mostrar un mensaje al usuario si la configuración falla
          });
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

input, select {
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