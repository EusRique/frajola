# Listar Documentos Adicionados


> ## Caso de sucesso

1. ✅ Recebe uma requisição do tipo **GET** na rota **/v1/api/list-documents**
2. ✅ Retorna **201** se não tiver nenhum registro
3. ✅ Retorna **201** com os dados dos registros

</br>
</br>   

> ## Exceções

1. ✅ Retorna erro **404** se a API não existir
2. ✅ Retorna erro **404** se der erro ao tentar listar os documentos

</br>
</br>   

# Adicionar um documento

> ## APIs relacionadas a listar todos os documentos cadastradas

GET **/v1/api/list-documents** API para listar todos os documentos cadastrados

Parâmetros utilizados para ordenação e filtros

    - page_size = Quantidade de registros a serem mostrados
    - page = Página a ser mostrada
    - sort = Ordenação ASC ou DESC
    - document_number = Filtra por número de documentação (xxxxxxxxxxx)
    - document_type = Filtra por tipo de documento (CPF ou CNPJ)
    - is_block_list = Filtra por Verdadeiro ou Falso (true ou false)

Exemplo CURL no Postman lista padrão

```
curl --location --request GET 'localhost:8080/v1/api/list-documents'
```

Exemplo CURL no Postman lista com filtros

```
curl --location --request GET 'localhost:8080/v1/api/list-documents?param1&param2'
```

O Retorno é todos os documentos cadastrados

```
{
    "pagination": {
        "limit": 50,
        "page": 1,
        "sort": "id ",
        "total_rows": 2,
        "total_pages": 1
    },
    "results": [
        {
            "id": "04f2034b-df17-46cc-913e-9cdeea02ada3",
            "document": "xxxxxxxxxxx",
            "document_type": "CPF",
            "is_block_list": true,
            "created_at": "2022-07-21T17:45:41.949076Z",
            "updated_at": "2022-07-21T17:45:41.949076Z"
        },
        {
            "id": "6a581dff-4f69-47d3-98a4-36c71de8b9b0",
            "document": "xxxxxxxxxxx",
            "document_type": "CPF",
            "is_block_list": true,
            "created_at": "2022-07-21T17:31:48.347996Z",
            "updated_at": "2022-07-21T17:31:48.347996Z"
        }
    ]
}

```