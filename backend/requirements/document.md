# Adicionar Documento

> ## Caso de sucesso

1. ✅ Recebe uma requisição do tipo **POST** na rota **/v1/api/create-document**
2. ✅ Valida dados obrigatórios
3. ✅ Cria uma compra
4. ✅ Retorna 201

</br>
</br>   

> ## Exceções
1.  ✅ Retorna erro **404** se a API não existir
2.  ✅ Retorna erro **400** se **os valores** não forem fornecidos pelo cliente
3.  ✅ Retorna erro **404** se der erro ao tentar criar a compra

</br>
</br>   

# Adicionar um documento
> ## APIs relacionadas a adicionar documento

POST **/v1/api/create-document** API para cadastrar uma compra

Exemplo CURL no Postman

```
curl --location --request POST 'localhost:8080/v1/api/create-document' \
--header 'Content-Type: application/json' \
--data-raw '{
    "document_number": "xxx.xxx.xxx-xx",
    "document_type": "CPF",
    "is_block_list": true
}'

```

O Retorno da rota

```
{
    "results": {
        "id": "04f2034b-df17-46cc-913e-9cdeea02ada3",
        "document": "xxxxxxxxxxx",
        "document_type": "CPF",
        "is_block_list": true,
        "created_at": "2022-07-21T17:45:41.949075526Z",
        "updated_at": "2022-07-21T17:45:41.949075526Z"
    }
}
```