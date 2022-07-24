import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// Plugins
import vuetify from './plugins/vuetify'
import toasted from './plugins/toasted'
import mask from './plugins/mask'

Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,
  toasted,
  mask,
  render: h => h(App)
}).$mount('#app')
