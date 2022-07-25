/* eslint-disable quotes */
import axios from '@/plugins/axios'

// eslint-disable-next-line quotes
const allDocuments = async payload => await axios.get(`/list-documents?document_type=${payload.document_type}&document_number=${payload.document_number}&is_block_list=${payload.is_block_list}`)

const createDocument = async payload => await axios.post(`/create-document`, { ...payload })

const updateDocument = async payload => await axios.put(`/update-document/${payload.id}`, { ...payload })

const deleteDocument = async payload => await axios.delete(`/delete-document/${payload.id}`)

export default {
  allDocuments,
  createDocument,
  updateDocument,
  deleteDocument
}
