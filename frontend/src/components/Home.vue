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

                  <v-card-text>
                    <v-container>
                      <v-row>
                        <v-col cols="12" sm="8" md="8">
                          <v-text-field
                            v-model="editedItem.document_number"
                            label="Número do Documento"
                          ></v-text-field>
                        </v-col>
                      </v-row>
                      <v-row>
                        <v-col cols="12" sm="8" md="6">
                          <v-select
                            item-text="block"
                            :items="typeDocument"
                            v-model="editedItem.document_type"
                            label="Tipo de Documento"
                          ></v-select>
                        </v-col>
                        <v-col cols="12" sm="6" md="4">
                          <v-select
                            item-text="block"
                            :items="blocked"
                            v-model="editedItem.is_block_list"
                            label="Bloqueado"
                          ></v-select>
                        </v-col>
                      </v-row>
                    </v-container>
                  </v-card-text>

                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="close">
                      Cancelar
                    </v-btn>
                    <v-btn color="blue darken-1" text @click="save">
                      Cadastrar
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
              <v-dialog v-model="dialogDelete" max-width="500px">
                <v-card>
                  <v-card-title class="text-h5"
                    >Are you sure you want to delete this item?</v-card-title
                  >
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="closeDelete"
                      >Cancel</v-btn
                    >
                    <v-btn color="blue darken-1" text @click="deleteItemConfirm"
                      >OK</v-btn
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
              {{ item.is_block_list ? 'Não' : 'Sim' }}
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
        document_number: '',
        document_type: '',
        is_block_list: ''
      },
      defaultItem: {
        document_number: '',
        document_type: '',
        is_block_list: ''
      }
    }
  },

  computed: {
    formTitle () {
      return this.editedIndex === -1 ? 'Novo Documento' : 'Alterar Documento'
    }
  },

  watch: {
    dialog (val) {
      val || this.close()
    },
    dialogDelete (val) {
      val || this.closeDelete()
    }
  },

  created () {
    this.listAllDocuments()
  },

  methods: {
    ...mapActions('Documents', {
      allDocuments: 'allDocuments',
      createDocument: 'createDocument'
    }),

    async listAllDocuments () {
      const data = await this.allDocuments()
      this.desserts = data.data.results
    },

    getColor (block) {
      if (block) return 'green'
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

    deleteItemConfirm () {
      this.desserts.splice(this.editedIndex, 1)
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

    async save () {
      if (this.editedIndex > -1) {
        console.log('SAVE IF')
        Object.assign(this.desserts[this.editedIndex], this.editedItem)
      } else {
        const newDocument = {
          document_number: this.editedItem.document_number,
          document_type: this.editedItem.document_type,
          is_block_list: this.editedItem.is_block_list === 'Sim'
        }
        const data = await this.createDocument(newDocument)

        if (data.error === false) {
          this.desserts.push(data.data.data.results)
        }
      }
      this.close()
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
