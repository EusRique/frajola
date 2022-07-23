import Vue from 'vue'

export function showMessageError (e) {
  Vue.toasted.global.defaultError({ msg: e })
}

export function showMessageSuccess (e) {
  Vue.toasted.global.defaultSuccess({ msg: e })
}

export default { showMessageError, showMessageSuccess }
