import documents from '@/api/documents.js'
import messages from '@/plugins/messages.js'

export default {
  async allDocuments ({ dispatch }, payload) {
    try {
      const data = await documents.allDocuments(payload)

      return data
    } catch (error) {
      dispatch(
        'Snackbar/setSnackbar',
        {
          status: true,
          text: messages.showMessageError(`Erro ao carregar lista - ${error.message}`),
          type: 'error'
        },
        {
          root: true
        }
      )

      return error
    }
  },

  async createDocument ({ dispatch }, payload) {
    try {
      const data = await documents.createDocument(payload)
      dispatch(
        'Snackbar/setSnackbar',
        {
          status: true,
          text: messages.showMessageSuccess('Documento criado com sucesso'),
          type: 'success'
        },
        {
          root: true
        }
      )

      const response = {
        data: data,
        error: false
      }

      return response
    } catch (error) {
      dispatch(
        'Snackbar/setSnackbar',
        {
          status: true,
          text: messages.showMessageError(`Erro ao criar documento - ${error.response.data.message}`),
          type: 'error'
        },
        {
          root: true
        }
      )

      const response = {
        data: error,
        error: true
      }

      return response
    }
  },

  async updateDocument ({ dispatch }, payload) {
    try {
      const data = await documents.updateDocument(payload)
      dispatch(
        'Snackbar/setSnackbar',
        {
          status: true,
          text: messages.showMessageSuccess('Documento atualizado com sucesso'),
          type: 'success'
        },
        {
          root: true
        }
      )

      const response = {
        data: data,
        error: false
      }

      return response
    } catch (error) {
      dispatch(
        'Snackbar/setSnackbar',
        {
          status: true,
          text: messages.showMessageError(`Erro ao atualizar documento - ${error.response.data.message}`),
          type: 'error'
        },
        {
          root: true
        }
      )

      const response = {
        data: error,
        error: true
      }

      return response
    }
  },

  async deleteDocument ({ dispatch }, payload) {
    try {
      const data = await documents.deleteDocument(payload)
      dispatch(
        'Snackbar/setSnackbar',
        {
          status: true,
          text: messages.showMessageSuccess('Documento removido com sucesso'),
          type: 'success'
        },
        {
          root: true
        }
      )

      const response = {
        data: data,
        error: false
      }

      return response
    } catch (error) {
      dispatch(
        'Snackbar/setSnackbar',
        {
          status: true,
          text: messages.showMessageError(`Erro ao remover documento - ${error.response.data.message}`),
          type: 'error'
        },
        {
          root: true
        }
      )

      const response = {
        data: error,
        error: true
      }

      return response
    }
  }
}
