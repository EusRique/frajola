<template>
  <div id="app">
    <nav>
      <router-link to="/">Lista de Documentos</router-link> |
      <router-link to="/about">Criar Documento</router-link>
    </nav>
    <router-view/>
  </div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #000000;
  text-decoration: none
}

nav a.router-link-exact-active {
  color: #f6a60c;
}
</style>
