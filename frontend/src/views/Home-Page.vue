<template>
  <div class="home-container">
    <div class="header">
      <div class="title-container">
        <h1 class="title">Gestión de Paquetería IPC2</h1>
        <p class="description">Bienvenido a la plataforma de gestión de paquetería IPC2. Inicia sesión para comenzar.</p>
      </div>
    </div>

    <div class="content">
      <router-link to="/login" class="btn">Iniciar Sesión</router-link>
      <router-link to="/about" class="btn about-btn">Acerca de</router-link>
    </div>
  </div>
</template>

<script>

</script>

<style>
.home-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.header {
  background-color: #212121;
  padding: 20px;
  width: 100%;
  text-align: center;
}

.title-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.title {
  font-size: 2.5em;
  color: #fff;
  margin: 0;
}

.description {
  font-size: 1.2em;
  color: #fff;
  margin-top: 10px;
}

.content {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.btn {
  padding: 15px 30px;
  font-size: 1.2em;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  text-decoration: none;
  margin-right: 10px;
  transition: background-color 0.3s ease;
}

.btn:nth-child(odd) {
  background-color: #7e57c2; /* Lila */
}

.btn:nth-child(even) {
  background-color: #42b983; /* Verde de Iniciar Sesión */
}

.btn:hover {
  background-color: #5c437f; /* Lila Oscuro */
}

.about-btn {
  background-color: #66bb6a; /* Verde */
}

.about-btn:hover {
  background-color: #539156; /* Verde Oscuro */
}
</style>