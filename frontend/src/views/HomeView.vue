<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/frajola.png">
    <Home msg="Lista de Documentos"/>
  </div>
</template>

<script>
// @ is an alias to /src
import Home from '@/components/Home.vue'

export default {
  name: 'HomeView',
  components: {
    Home
  }
}
</script>

<style>

.home img {
    max-width:200px;
    max-height:150px;
    width: auto;
    height: auto;
}

</style>
