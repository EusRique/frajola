<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/frajola.png">
    <AllDocuments msg="Lista de Documentos"/>
  </div>
</template>

<script>
// @ is an alias to /src
import AllDocuments from '@/components/AllDocuments.vue'

export default {
  name: 'HomeView',
  components: {
    AllDocuments
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
