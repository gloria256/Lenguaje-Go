<template>
  <div id="app">
    <v-app id="inspire">
      <v-card
        class="mx-auto overflow-hidden"
        width="960px"
      >    
        <v-app-bar
          color="deep-purple accent-4"
          dark
          prominent
          height="70px"
        >
          <v-toolbar-title class="text-center">Gestion Compradores</v-toolbar-title>
        </v-app-bar>
        
        <v-card-title>
          <v-text-field
            v-model="search"
            append-icon="md-Search"
            label="Search"
            single-line
            hide-details
          ></v-text-field>
        </v-card-title>

        <v-data-table
          dense
          :headers="headers"
          :items="desserts"
          :search="search"
          item-key="Name"
          class="elevation-1"
        > 
          <v-card-actions>       
            <v-btn
              color="blue darken-1"
              text
            >
              Save
            </v-btn>
          </v-card-actions>
          
          <template v-slot:[`item.actions`]="{item}">
            <v-icon
              small
              class="mr-2"
              @click="getInfoCompradores(item.Id,item.Name)"
            >
              visibility
            </v-icon>
          </template>
        </v-data-table>

        <v-dialog v-model="dialog" max-width="700px">
          <v-card>
            <v-card-title>
              <span class="text-h5">{{ formTitle }}</span>
            </v-card-title>

            <v-container>
              <v-row>
                <v-col width="300px">
                        <template>
                          <v-simple-table fixed-header height="300px">
                            <template v-slot:default>
                              <thead>
                                <td colspan=3 class="text-center" style="background-color: lightgray;" height="40px">
                                  Comprador con la misma Ip
                                </td>
                                <tr>
                                  <th class="text-left">
                                    Ip
                                  </th>
                                  <th class="text-left">
                                    Dispositivo
                                  </th>
                                  <th class="text-left">
                                    Nombre
                                  </th>
                                </tr>
                              </thead>
                              <tbody>
                                <tr
                                  v-for="item in compradoresMismaIp"
                                  :key="item.Ip"
                                >
                                  <td>{{ item.Ip }}</td>
                                  <td>{{ item.Name }}</td>
                                  <td>{{ item.Dev }}</td>
                                </tr>
                              </tbody>
                            </template>
                          </v-simple-table>
                        </template>
                </v-col>

                <v-col width="300px">
                        <template>
                              <v-simple-table fixed-header height="300px">
                                <template v-slot:default>
                                  <thead>
                                    <td colspan=3 class="text-center" style="background-color: lightgray;" height="40px">
                                      Historial de Compras
                                    </td>
                                    <tr>
                                      <th class="text-left">
                                        Id producto
                                      </th>
                                      <th class="text-left">
                                        Nombre Producto
                                      </th>
                                      <th class="text-left">
                                        Precio
                                      </th>
                                    </tr>
                                  </thead>
                                  <tbody>
                                    <tr
                                      v-for="item in historialCompras"
                                      :key="item.Id_Producto"
                                    >
                                      <td>{{ item.Id_Producto }}</td>
                                      <td>{{ item.Name_Producto }}</td>
                                      <td>{{ item.Price_Producto }}</td>
                                    </tr>
                                  </tbody>
                                </template>
                              </v-simple-table>
                        </template>
                </v-col>
              </v-row>

              <v-row>
                <v-card class="mx-auto" width="600px" style="margin-top: 2px;">
                  <v-card-text>
                    <div style="padding: 12px;">
                      <h3>
                        Productos recomendados
                      </h3>
                      <v-space></v-space>
                      <div>
                      {{productosRecomendados}}
                      </div>
                    </div>
                  </v-card-text>
                </v-card>
              </v-row>
            </v-container>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="blue darken-1" text @click="close">
                  Cancel
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
        
      </v-card>
    </v-app>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data () {
    return {
      search: '',
      desserts: [],
      dialog: false,
      formTitle: '',
      sortable: false,
      headers: [
        {
          text: 'ID',
          align: 'start',
          sortable: false,
          value: 'Id',
        },
        { text: 'Comprador', value: 'Name' },
        { text: 'Información', value: 'actions',  sortable: false },
      ],
      historialCompras: null,
      compradoresMismaIp: null,
      productosRecomendados:"",
    }
  },
  created() {
    this.getListCompradores()
  },
  methods:{
    getListCompradores(){
      axios.get("http://localhost:3000/listaNombreCompradores").then((result) => {
      console.log(result.data);
      this.desserts = result.data
      })
    },

    getInfoCompradores(id,name){
      this.dialog = true
      this.formTitle = 'Informacion de: '+ name
      console.log('item: ',id,name)
      axios.get("http://localhost:1000/ConsultarCompradores").then((info) => {
        var datos
        var historialCompras
        var compradoresMismaIp
        var productosRecomendados
        datos = info.data[id]
        console.log("---",datos)
        if (datos[0] != null){
          historialCompras=datos[0]
          this.historialCompras=historialCompras
        }else{
          historialCompras=[{"Name": "",
                            "Id_Producto": "",
                            "Name_Producto": "",
                            "Price_Producto": 0}] 
          this.historialCompras=historialCompras
        }
        if (datos[1] != null){
          compradoresMismaIp=datos[1]
          this.compradoresMismaIp = compradoresMismaIp 
          console.log("---",compradoresMismaIp)  
        }else{
          compradoresMismaIp=[{"Ip": "",
                            "Name": "",
                            "Dev": ""}]
          this.compradoresMismaIp = compradoresMismaIp
        }
        if (datos[2] != null){
          productosRecomendados=datos[2] 
          this.productosRecomendados = ""
          for (let index = 0; index < productosRecomendados.length; index++) {
            const element = productosRecomendados[index].Name;
            this.productosRecomendados += ",  " + element
          }
          console.log("***",productosRecomendados) 
        }else{
          this.productosRecomendados = ""
          productosRecomendados="Ninguno" 
          this.productosRecomendados = productosRecomendados 
        }
      })
      
    },

    close () {
      this.dialog = false
    },
  }
};

</script>
