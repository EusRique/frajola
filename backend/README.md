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
    
    - `curl --location --request GET 'localhost:3000/v1/api/alive'`
  
### Como rodar os teste unitários
- Dentro da pasta backend: `ginkgo cover cmd domain/model/`
- Dentro da pasta backend: `ginkgo --cover cmd domain/model/`
    
    - `curl --location --request GET 'localhost:3000/v1/api/alive'`
  
### Como executar a interface gráfica do PgAdmin
- Serviço disponível em `localhost:9000'`
- Login e senha: `admin@user.com | 123456`

No dashbord botão direito em **Serves** > **Register** > **Server**

Dados de Conexão aba connection:
- **Host:** `172.17.0.1` 
- **username:** `postgres` 
- **senha:** `root` 
- **database:** frajola
  
### Serviços utilizados ao executar o docker-compose

- Aplicação principal
- Postgres
- PgAdmin


## **Abaixo os links para a documentação da API**

1. [Adicionar um registro](./requirements/document.md)
2. [Listar todos os registros](./requirements/load-document.md)
2. [Atualizar um registros](./requirements/update-document.md)
2. [Deletar um registros](./requirements/delete-document.md)


Checklist dos Requisitos Backend:
---

- [x] Validação do número do CPF/CNPJ (dígito verificador) na consulta e inclusão
- [x] Filtros, ordenação e marcação em blocklist
- [x] Interface REST para integração com o front-end
- [ ] Rota de suporte que retorne as informações de up-time do servidor e quantidade de consultas realizadas desde o start
- [x] PostgreSQL
- [x] Docker
- [x] Testes unitários;
- [x] GO;