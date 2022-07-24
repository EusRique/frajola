/* eslint-disable quotes */
import axios from '@/plugins/axios'

// eslint-disable-next-line quotes
const allDocuments = async payload => await axios.get(`/list-documents`)

const createDocument = async payload => await axios.post(`/create-document`, { ...payload })

const updateDocument = async payload => await axios.put(`/update-document/${payload.id}`, { ...payload })

export default {
  allDocuments,
  createDocument,
  updateDocument
}
