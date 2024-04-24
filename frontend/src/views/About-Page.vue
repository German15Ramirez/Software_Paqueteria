<template>
  <div class="about-container">
    <h2>Acerca de</h2>
    <p>Esta aplicación de Gestión de Paquetería IPC2 te permite realizar seguimiento de envíos, gestionar clientes y mucho más.</p>
    <p>¡Esperamos que disfrutes usando nuestra aplicación!</p>
    <img src="../assets/Tux.png" alt="Logo de la aplicación">
  </div>
</template>

<script>
export default {
  name: 'AboutPage'
}
</script>

<style scoped>
.about-container {
  width: 100%;
  padding: 20px 10%;
  box-sizing: border-box;
  background-color: #333; /* Fondo oscuro */
  color: #fff; /* Texto blanco */
  min-height: 100vh; /* Cubre toda la altura de la ventana */
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
h2 {
  font-size: 2em;
  margin-bottom: 10px;
}
p {
  font-size: 1.2em;
  margin-bottom: 15px;
}
img {
  max-width: 100%;
  margin-top: 20px;
}
</style>