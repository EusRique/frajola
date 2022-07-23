<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <div id="app">
      <v-app id="inspire">
        <v-data-table
          :headers="headers"
          :items="desserts"
          class="elevation-1"
        >
          <template v-slot:[`item.block`]="{ item }">
            <v-chip
              :color="getColor(item.block)"
              dark
            >
              {{ item.block }}
            </v-chip>
          </template>
        </v-data-table>
      </v-app>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },

  data () {
    return {
      documents: [],
      headers: [
        {
          text: 'id',
          align: 'start',
          sortable: false,
          value: 'id'
        },
        { text: 'Número Documento', value: 'documentNumber' },
        { text: 'Tipo de Documento', value: 'documentType' },
        { text: 'Bloqueado (g)', value: 'block' }
      ],
      desserts: [
        {
          id: '2ec22f69-1f4a-46d4-acc8-87d8875ef11a',
          documentNumber: '159-000-000-11',
          documentType: 'CPF',
          block: 'Sim'
        },
        {
          id: '2ec22f69-1f4a-46d4-acc8-87d8875ef11b',
          documentNumber: '237-000-000-11',
          documentType: 'CPF',
          block: 'Sim'
        },
        {
          id: '2ec22f69-1f4a-46d4-acc8-87d8875ef11c',
          documentNumber: '262-000-000-11',
          documentType: 'CPF',
          block: 'Sim'
        },
        {
          id: '2ec22f69-1f4a-46d4-acc8-87d8875ef11d',
          documentNumber: '305-000-000-11',
          documentType: 'CPF',
          block: 'Sim'
        },
        {
          id: '2ec22f69-1f4a-46d4-acc8-87d8875ef11e',
          documentNumber: '356-000-000-11',
          documentType: 'CPF',
          block: 'Sim'
        },
        {
          id: '2ec22f69-1f4a-46d4-acc8-87d8875ef11f',
          documentNumber: '375-000-000-11',
          documentType: 'CPF',
          block: 'Não'
        },
        {
          id: '2ec22f69-1f4a-46d4-acc8-87d8875ef11g',
          documentNumber: '392-000-000-11',
          documentType: 'CPF',
          block: 'Sim'
        },
        {
          id: '2ec22f69-1f4a-46d4-acc8-87d8875ef11h',
          documentNumber: '408-000-000-11',
          documentType: 'CPF',
          block: 'Sim'
        },
        {
          id: '2ec22f69-1f4a-46d4-acc8-87d8875ef11i',
          documentNumber: '452-000-000-11',
          documentType: 'CPF',
          block: 'Não'
        },
        {
          id: '2ec22f69-1f4a-46d4-acc8-87d8875ef11j',
          documentNumber: '518-000-000-11',
          documentType: 'CPF',
          block: 'Sim'
        }
      ]
    }
  },

  mounted () {
    this.listAllDocuments()
  },

  methods: {
    ...mapActions('Documents', {
      allDocuments: 'allDocuments'
    }),

    async listAllDocuments () {
      const data = await this.allDocuments()
      console.log('Lista carregada!!!')
      console.log(data)
    },

    getColor (block) {
      if (block === 'Não') return 'green'
      else return 'red'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
