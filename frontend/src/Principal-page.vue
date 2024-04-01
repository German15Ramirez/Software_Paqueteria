<template>
  <div class="app-container">
    <div class="header">
      <h1 class="title">Gestión de Paqueteria IPC2</h1>
    </div>

    <!-- Botón de acceso al sistema -->
    <button @click="abrirFormularioSesion" class="btn">Acceder</button>
  </div>
</template>

<script>
export default {
  name: 'PrincipalPage',
  methods: {
    abrirFormularioSesion() {
      // Abrir una nueva pestaña o ventana con el formulario de inicio de sesión
      const nuevaPestana = window.open('', '_blank');
      // Verificar si la nueva pestaña se abrió correctamente
      if (nuevaPestana) {
        // Contenido HTML del formulario de inicio de sesión
        const contenidoFormulario = "<html>" +
            "<head>" +
            "<title>Inicio de Sesión</title>" +
            "<style>" +
            "body {" +
            "font-family: Arial, sans-serif;" +
            "padding: 20px;" +
            "}" +
            "h1 {" +
            "font-size: 2em;" +
            "margin-bottom: 20px;" +
            "}" +
            "form {" +
            "display: flex;" +
            "flex-direction: column;" +
            "gap: 10px;" +
            "max-width: 400px;" +
            "margin: 0 auto;" +
            "}" +
            "label {" +
            "font-size: 1.2em;" +
            "}" +
            "input, button {" +
            "padding: 10px;" +
            "font-size: 1em;" +
            "}" +
            "</style>" +
            "</head>" +
            "<body>" +
            "<h1>Inicio de Sesión</h1>" +
            "<form>" +
            "<label for='usuarioId'>ID de usuario o cliente:</label>" +
            "<input type='text' id='usuarioId' required>" +
            "<label for='password'>Contraseña:</label>" +
            "<input type='password' id='password' required>" +
            "<button type='submit'>Iniciar sesión</button>" +
            "</form>" +
            "</body>" +
            "</html>";

        // Escribir el contenido en la nueva pestaña
        nuevaPestana.document.write(contenidoFormulario);
        nuevaPestana.document.close(); // Cerrar la escritura del documento
      } else {
        alert('Error al abrir la nueva pestaña.');
      }
    }
  }
}
</script>

<style>
html, body {
  margin: 0;
  padding: 0;
  height: 100%;
}

.app-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh; /* Ocupa toda la altura de la ventana */
}

.header {
  background-image: url('../public/img/fondo_principal.jpg');
  background-size: cover;
  background-position: center;
  text-align: center;
  padding: 100px 0;
  flex-grow: 1;
  width: 100%;
  box-sizing: border-box;
}

.title {
  font-size: 4.5em;
  color: black;
}

.btn {
  margin-top: 20px;
  padding: 20px 30px;
  font-size: 2.2em;
  background-color: #42b983;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
</style>