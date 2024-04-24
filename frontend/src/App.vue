<template>
  <div class="app-container">
    <!-- Contenido de la página principal -->
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'App',
  created() {
    this.$router.push({ name: 'Home' }); // Redirigir automáticamente a la página de inicio
  }
}
</script>

<style>
/* Estilos para el cuerpo de la página */
html, body {
  margin: 0;
  padding: 0;
  height: 100%;
}

/* Estilos para el contenedor principal */
.app-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh; /* Ocupa toda la altura de la ventana */
  background-color: #121212; /* Color de fondo oscuro */
  background-image: url('../public/img/principal_image.jpg'); /* Ruta de la imagen de fondo */
  background-size: cover; /* Ajustar la imagen al tamaño del contenedor */
  background-position: center; /* Centrar la imagen */
}
</style>