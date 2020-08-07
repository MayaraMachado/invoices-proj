# Título do Projeto

A API Invoices é uma API desenvolvida utilizando o framework Gin Gonic, utilizando a linguagem Golang. Ela implementa um CRUD básico onde é possível realizar a gestão de um objeto Invoice. Um exemplo do objeto invoice pode ser visto abaixo:

```
{
      "id": 10,
      "created_at": "2020-08-01T14:21:44.225Z",
      "reference_month": 3,
      "reference_year": 2015,
      "document": "EWS140",
      "description": "a short description",
      "amount": 15.21,
      "is_active": 1
}
```


## Guia

Através dessas instruções você irá obter uma cópia do projeto e poderá executá-lo em sua máquina, para propósitos de desenvolvimento e teste. Leia as notas de deploy para descobrir como fazer deploy em produção.

### Pré-requisitos

É necessário possuir a linguagem Go instalado em sua máquina, e também ter o banco de dados postgres instalado na sua máquina.

### Instalação

#### Configurar Base de Dados

Primeiro configure o banco de dados em sua máquina, crie uma base de dados no postgre e referencie a URL de acesso na variável de ambiente da sua máquina com o nome *DATABASE_URL*. Caso a aplicação não encontre esta variável, ela irá tentar se conectar com uma url default: *postgres://postgres:123@localhost/postgres?sslmode=disable*.

Após a criação da conexão PostgreSQL, execute os scripts especificado no tópico [Configurando a Base de Dados Postgres](https://github.com/MayaraMachado/invoices-proj/wiki/Configurando-a-Base-de-Dados-Postgres) da documentação.


#### Executando a aplicação

Clone este repositório na sua máquina, e dentro do diretório onde se encontra o arquivo *main.go* execute o comando.

```
go run main.go
```


### Rotas da aplicação

A aplicação conta com as seguintes rotas:

- GET: Ping, verifica disponibilidade da aplicação
- POST: Resgistar usuário
- POST: Login de usuário
- POST: Criar invoice
- GET: Listar invoices
- PUT: Alterar Invoice
- DEL: Remover logicamente um invoice

Acesse a página [Requisições](https://github.com/MayaraMachado/invoices-proj/wiki/Requisi%C3%A7%C3%B5es) da wiki do projeto para obter a documentação mais específica de como utilizar essas rotas.

A documentação completa desta API se encontra no Wiki do repositório.

## Deploy

Esta API se encontra disponível no Heroku na url: https://mayara-invoices-api.herokuapp.com.

## Construído com

* [Golang](https://golang.org/) - Linguagem utilizada.
* [Gin Gonic](https://github.com/gin-gonic/gin) - Framework Rest.
* [PostgreSQL](https://www.postgresql.org/) - Uasdo como base de dados.
* [JWT-Go](github.com/dgrijalva/jwt-go) - Usado para implementar a autenticação JWT.

