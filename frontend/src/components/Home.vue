<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <div id="app">
      <v-app id="inspire">
        <v-data-table
          :headers="headers"
          :items="desserts"
          sort-by="calories"
          class="elevation-1"
        >
          <template v-slot:top>
            <v-toolbar flat>
              <v-spacer></v-spacer>
              <v-dialog v-model="dialog" max-width="500px">
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                    color="orange"
                    dark
                    class="mb-2"
                    v-bind="attrs"
                    v-on="on"
                  >
                    Novo Documento
                  </v-btn>
                </template>
                <v-card>
                  <v-card-title>
                    <span class="text-h5">{{ formTitle }}</span>
                  </v-card-title>

                  <v-card-text >
                    <v-form ref="form" class="mx-2" lazy-validation>
                      <v-container>
                        <v-row>
                          <v-col cols="12" sm="8" md="8">
                            <v-text-field
                              ref="document"
                              v-mask="['###.###.###-##', '##.###.###/####-##']"
                              v-model="editedItem.document"
                              :rules="[() => !!editedItem.document || 'Campo obrigatório']"
                              :error-messages="errorMessages"
                              required
                              label="Número do Documento"
                            ></v-text-field>
                          </v-col>
                        </v-row>
                        <v-row>
                          <v-col cols="12" sm="8" md="6">
                            <v-select
                              ref="document_type"
                              :items="typeDocument"
                              v-model="editedItem.document_type"
                              :rules="[() => !!editedItem.document_type || 'Campo obrigatório']"
                              :error-messages="errorMessages"
                              required
                              label="Tipo de Documento"
                            ></v-select>
                          </v-col>
                          <v-col cols="12" sm="6" md="4">
                            <v-select
                              ref="is_block_list"
                              :items="blocked"
                              v-model="editedItem.is_block_list"
                              :rules="[() => !!editedItem.is_block_list || 'Campo obrigatório']"
                              :error-messages="errorMessages"
                              required
                              label="Bloqueado"
                            ></v-select>
                          </v-col>
                        </v-row>
                      </v-container>
                    </v-form>
                  </v-card-text>

                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="orange" text @click="close">
                      Cancelar
                    </v-btn>
                    <v-btn color="orange" text @click="save">
                      Cadastrar
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
              <v-dialog v-model="dialogDelete" max-width="500px">
                <v-card>
                  <v-card-title class="text-h6"
                    >Tem certeza de que deseja excluir este item?</v-card-title
                  >
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="orange" text @click="closeDelete"
                      >Cancelar</v-btn
                    >
                    <v-btn color="orange" text @click="deleteItemConfirm"
                      >Excluir</v-btn
                    >
                    <v-spacer></v-spacer>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </v-toolbar>
          </template>
          <template v-slot:[`item.is_block_list`]="{ item }">
            <v-chip
              :color="getColor(item.is_block_list)"
              dark
            >
              {{ item.is_block_list }}
            </v-chip>
          </template>
          <template v-slot:[`item.actions`]="{ item }">
            <v-icon small class="mr-2" @click="editItem(item)">
              mdi-pencil
            </v-icon>
            <v-icon small @click="deleteItem(item)"> mdi-delete </v-icon>
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
      dialog: false,
      dialogDelete: false,
      formHasErrors: false,
      errorMessages: '',
      headers: [
        {
          text: 'id',
          align: 'start',
          sortable: false,
          value: 'id'
        },
        { text: 'Número Documento', value: 'document' },
        { text: 'Tipo de Documento', value: 'document_type' },
        { text: 'Bloqueado', value: 'is_block_list' },
        { text: 'Actions', value: 'actions', sortable: false }
      ],
      desserts: [],
      blocked: ['Não', 'Sim'],
      typeDocument: ['CPF', 'CNPJ'],
      editedIndex: -1,
      editedItem: {
        document: '',
        document_type: '',
        is_block_list: ''
      },
      defaultItem: {
        document: '',
        document_type: '',
        is_block_list: ''
      },
      nameRules: [
        v => !!v || 'Campos obrigatórios'
      ]
    }
  },

  computed: {
    formTitle () {
      return this.editedIndex === -1 ? 'Novo Documento' : 'Alterar Documento'
    },

    form () {
      return {
        document: this.editedItem.document,
        document_type: this.editedItem.document_type,
        is_block_list: this.editedItem.is_block_list
      }
    }
  },

  watch: {
    dialog (val) {
      val || this.close()
    },

    dialogDelete (val) {
      val || this.closeDelete()
    },

    document () {
      this.errorMessages = ''
    }
  },

  created () {
    this.listAllDocuments()
  },

  methods: {
    ...mapActions('Documents', {
      allDocuments: 'allDocuments',
      createDocument: 'createDocument',
      updateDocument: 'updateDocument',
      deleteDocument: 'deleteDocument'
    }),

    async listAllDocuments () {
      const data = await this.allDocuments()
      const newListAll = data.data.results.map(document => {
        let isBlocked
        if (document.is_block_list === false) isBlocked = 'Não'
        else isBlocked = 'Sim'

        const listDocuments = {
          id: document.id,
          document: this.formatCpfOrCnpj(document),
          document_type: document.document_type,
          is_block_list: isBlocked
        }

        return listDocuments
      })

      this.desserts = newListAll
    },

    formatCpfOrCnpj (document) {
      let documentFormated
      switch (document.document_type.toUpperCase()) {
        case 'CPF':
          documentFormated = document.document.replace(/^(\d{3})(\d{3})(\d{3})(\d{2})/, '$1.$2.$3-$4')
          break
        case 'CNPJ':
          documentFormated = document.document.replace(/^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})/, '$1.$2.$3/$4-$5')
          break
        default:
          documentFormated = document.document
      }

      return documentFormated
    },

    getColor (block) {
      if (block === 'Não') return 'green'
      else return 'red'
    },

    editItem (item) {
      this.editedIndex = this.desserts.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialog = true
    },

    deleteItem (item) {
      this.editedIndex = this.desserts.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialogDelete = true
    },

    async deleteItemConfirm () {
      const data = await this.deleteDocument(this.editedItem)
      console.log(data)
      if (data.error === false) {
        this.desserts.splice(this.editedIndex, 1)
      }
      this.closeDelete()
    },

    close () {
      this.dialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    closeDelete () {
      this.dialogDelete = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    resetForm () {
      this.errorMessages = []
      this.formHasErrors = false

      Object.keys(this.form).forEach(f => {
        this.$refs[f].reset()
      })
    },

    async save () {
      this.formHasErrors = false

      Object.keys(this.form).forEach(f => {
        if (!this.form[f]) this.formHasErrors = true

        this.$refs[f].validate(true)
      })

      const newDocument = {
        id: this.editedItem.id,
        document_number: this.editedItem.document,
        document_type: this.editedItem.document_type,
        is_block_list: this.editedItem.is_block_list === 'Sim'
      }

      if (this.editedIndex > -1) {
        const data = await this.updateDocument(newDocument)
        if (data.error === false) Object.assign(this.desserts[this.editedIndex], this.editedItem)
        this.close()
        this.resetForm()
      } else {
        let data
        if (this.formHasErrors === false) {
          data = await this.createDocument(newDocument)
          if (data.error === false) this.listAllDocuments()
          this.close()
          this.resetForm()
        }
      }
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
