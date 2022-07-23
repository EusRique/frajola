import Vue from 'vue'
import Vuex from 'vuex'

// Modules
import Documents from '@/store/Documents'
import Snackbar from '@/store/Snackbar'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    Documents,
    Snackbar
  }
})
