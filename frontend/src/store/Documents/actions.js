import documents from '@/api/documents.js'
import messages from '@/plugins/messages.js'

export default {
  async allDocuments ({ dispatch }, payload) {
    try {
      const data = await documents.allDocuments()
      dispatch(
        'Snackbar/setSnackbar',
        {
          status: true,
          text: messages.showMessageSuccess('Lista de documentos carregada com sucesso!!!'),
          type: 'success'
        },
        {
          root: true
        }
      )

      return data
    } catch (error) {
      dispatch(
        'Snackbar/setSnackbar',
        {
          status: true,
          text: messages.showMessageError(`Erro ao carregar lista - ${error.response.data.error}`),
          type: 'error'
        },
        {
          root: true
        }
      )

      return error
    }
  }
}
