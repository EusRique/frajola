'use strict'

import axios from 'axios'
import store from '@/store/index'

const config = {
  baseURL: 'http://localhost:3000/api/v1',
  timeout: 60 * 1000
}

const _axios = axios.create(config)

_axios.interceptors.request.use(
  function (config) {
    // eslint-disable-next-line no-unneeded-ternary
    const isLoggedIn = store.getters['User/getUser'] ? true : false

    if (isLoggedIn) {
      config.headers.Authorization = `Bearer ${localStorage.getItem(
        'user_token'
      )}`
    }

    return config
  },

  function (error) {
    // Do something with request error
    return Promise.reject(error)
  }
)

// Add a response interceptor
_axios.interceptors.response.use(
  function (response) {
    // Do something with response data
    return response
  },
  function (error) {
    // Do something with request error
    if (error.response.data && error.response.data.errors) {
      store.dispatch('Validation/setErrors', error.response.data.error, {
        root: true
      })
    }
    // Do something with response error
    return Promise.reject(error)
  }
)

export default _axios
