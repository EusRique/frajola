## Aplicação de validação de CPF/CNPJ

Desenvolver uma aplicação de validação de CPF/CNPJ que deve conter uma interface (UI) para gerenciamento de CPF/CNPJ (CRUD) com a possibilidade filtros, ordenação e marcação de alguns em uma blocklist.


## Como executar
É utilizado Docker para que todos os serviços utilizados fiquem disponíveis.

- Faça o clone do projeto
- Tendo o docker instalado em sua máquina apenas execute:
`docker-compose up -d`


### Como executar a aplicação
- Acesse o container da aplicação executando: `docker exec -it backend_app_1 bash`
- Para rodar a aplicação rode `go run main.go rest`
- Execute a URL de teste abaixo para rodar as migrations e verificar se a aplicação está rodando
    
    - `curl --location --request GET 'localhost:8080/v1/api/alive'`
  
### Serviços utilizados ao executar o docker-compose

- Aplicação principal
- Postgres
- PgAdmin


## **Abaixo os links para a documentação da API**

1. [Adicionar um documento](./requirements/document.md)