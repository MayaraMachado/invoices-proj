# Invoices

Essa API permite a criação e listagem de todos os invoices criados na sessão. Como ainda não possui a integração com o banco de dados, os invoices são armazenados na sessão.

A API também possui validação JWT para poder realizar ações de GET e POST, por isso você deve realizar login antes de executar essas chamadas.

### Rotas disponíveis

- POST : https://mayara-invoices-api.herokuapp.com/login

O único usuário disponível é admin com a senha admin.

```
curl -d "username=admin&password=admin" -X POST https://mayara-invoices-api.herokuapp.com/login
```
- GET  : https://mayara-invoices-api.herokuapp.com/api/invoices
```sh
curl https://mayara-invoices-api.herokuapp.com/api/invoices
	-H "Accept: application/json"
    -H "Authorization: Bearer {token}"
```

- POST : https://mayara-invoices-api.herokuapp.com/api/invoices
```json
// Body da requisição
{
	"id" : 1,
	"createdAt" : "2020-05-02T15:00:00Z",
	"referenceMonth" : 4,
	"referenceYear" : 2020,
	"document" : "1234856",
	"description" : "short description",
	"amount" : 55
}
```

```sh
curl -X POST https://mayara-invoices-api.herokuapp.com/api/invoices
	 -H "Content-Type: application/json"
     -H "Authorization: Bearer {token}" 
     -H "Accept: application/json" 
     -d "{json_body}"
```

### Executar na máquina

Para executar na sua máquina local, clone esse repositório e execute:

``` sh
$ go run main.go 
```

*obs.: A aplicação utiliza a porta definida em env, mas você pode alterar em main.go para executar na porta desejada.*