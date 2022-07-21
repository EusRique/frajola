# Atualizar Documento Adicionado


> ## Caso de sucesso

1. ✅ Recebe uma requisição do tipo **PUT** na rota **/v1/api/update-document/:document_id**
2. ✅ Valida o parâmetro **document_id**
3. ✅ Valida se os **campos** tem uma resposta válida
4. ✅ Valida se o registro a ser atualizado existe
5. ✅ **Atualiza** um registro com os dados fornecidos caso já tenha um registro
6. ✅ Retorna **200** com os dados do resultado do documento


</br>
</br>   

> ## Exceções

1. ✅ Retorna erro **404** se a API não existir
2. ✅ Retorna erro **400** se o document_id passado na URL for inválido
3. ✅ Retorna erro **404** se o payload enviado pelo for inválido
4. ✅ Retorna erro **404** se der erro ao tentar atualizar o documento

</br>
</br>   

# Atualizar um documento

> ## APIs relacionadas a atualização de um documento cadastrado

GET **/v1/api/update-document/:document_id** API para atualização de um documento cadastrado

Exemplo CURL no Postman

```
curl --location --request PUT 'localhost:8080/v1/api/update-document/dcf590d8-be26-4ffb-88e0-562b2aa7ea5b' \
--header 'Content-Type: application/json' \
--data-raw '{
    "document_number": "xxxxxxxxxxx",
    "document_type": "CPF",
    "is_block_list": false
}'
```

O Retorno é o documento atualizado

```
{
    "id": "04f2034b-df17-46cc-913e-9cdeea02ada3",
    "document": "xxxxxxxxxxx",
    "document_type": "CPF",
    "is_block_list": true,
    "created_at": "2022-07-21T17:45:41.949076Z",
    "updated_at": "2022-07-21T17:45:41.949076Z"
}

```