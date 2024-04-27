<template>
  <div class="recep-container">
    <div class="header" style="background-color: #212121; padding: 20px; border-radius: 5px; text-align: center;">
      <h1 style="font-size: 1.8em; color: #fff; margin: 0;">Bienvenido de nuevo Operador</h1>
    </div>
    <div class="options" style="background-color: #f2f2f2; padding: 10px; border-radius: 5px;">
      <h2 class="option-title">Opciones de operador:</h2>
      <ul>
        <li><router-link to="/oper/process">Procesar paquetes</router-link></li>
      </ul>
    </div>
  </div>
</template>

<script>

</script>

<style scoped>
.oper-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.header {
  background-color: #212121; /* Color de fondo del encabezado */
  padding: 20px;
  text-align: center;
}

.header h1 {
  font-size: 1.8em; /* Texto más grande */
  color: #fff;
  margin: 0;
}

.options {
  margin-top: 20px;
}

.option-title {
  font-size: 1.4em; /* Texto más grande */
  margin-bottom: 10px;
}

ul {
  list-style: none;
  padding: 0;
}

li {
  font-size: 1.2em; /* Texto más grande */
  margin-bottom: 10px;
}

a {
  text-decoration: none;
  color: #007bff; /* Color azul de enlaces */
  font-size: 1.2em; /* Texto más grande */
}

a:hover {
  text-decoration: underline;
}

</style>