# Remover um documento Adicionado


> ## Caso de sucesso

1. ✅ Recebe uma requisição do tipo **DELETE** na rota **/v1/api/delete-document/:document_id**
2. ✅ Valida o parâmetro **document_id**
3. ✅ Valida se o registro a ser excluído existe
4. ✅ Retorna **200** com a remoção do registro


</br>
</br>   

> ## Exceções

1. ✅ Retorna erro **404** se a API não existir
2. ✅ Retorna erro **400** se o document_id passado na URL for inválido
3. ✅ Retorna erro **404** se o payload enviado pelo for inválido
4. ✅ Retorna erro **404** se der erro ao tentar atualizar o documento
5. ✅ Retorna erro **400** se der erro ao tentar carregar o registro

</br>
</br>   

# Remover um documento

> ## APIs relacionadas a remoção de um documento cadastrado

GET **/v1/api/delete-document/:document_id** API para remoção de um documento cadastrado

Exemplo CURL no Postman

```
curl --location --request DELETE 'localhost:8080/v1/api/delete-document/04f2034b-df17-46cc-913e-9cdeea02ada3'
```

O Retorno é o documento removido

```
{
    "results": "Document deleted"
}

```