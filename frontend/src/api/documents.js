import axios from '@/plugins/axios'

// eslint-disable-next-line quotes
const allDocuments = async payload => await axios.get(`/list-documents`)

export default {
  allDocuments
}
