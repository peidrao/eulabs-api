# Eulabs - API

## Pré-requisitos
Antes de começar, certifique-se de ter os seguintes requisitos instalados:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Considerações sobre o projeto

- Todos os requisitos foram atendidos
- O projeto foi dockerizado por completo.
- Foi adicionado um arquivo Makefile para termos uma maior agilidade em executarmos alguns comandos.

## Comandos pelo Makefile

| comando           | funcionalidade                                  |
| -------------     | ----------------------------------------------- |
| make docker-build | Cria um buid do projeto                         |
| make docker-run   | Realiza start                                   |
| make docker-stop  | Para containers e os remove                     |
| make docker-clean | Limpa os containers                             |
| make docker-logs  | Logs do projeto                                 |

## Como rodar o projeto

```bash
# Clone o repositorio
$ git clone https://github.com/peidrao/eulabs-api
# entre no diretorio
$ cd eulabs-api
# Star no projeto
$ make docker-run
# Verifizando logs
$ make logs
```

## Endpoints

> ⚠ **Atenção**
>
> A rota default para acessar endpoints **http://localhost:8080**.

- `POST /api/v1/product/` - Criar um novo produto

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "Produto A", "price": 29.99}' http://localhost:8080/api/v1/product/
```

- `GET /api/v1/product/{id_produto}` - Buscar um produto específico

```bash
curl http://localhost:8080/api/v1/product/1
```

- `PUT /api/v1/product/{id_produto}` - Atualizar informações de um produto

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name": "Produto B", "price": 39.99}' http://localhost:8080/api/v1/product/1
```

- `DELETE /api/v1/product/{id_produto}` - Remover um produto 

```bash
curl -X DELETE http://localhost:8080/api/v1/product/1
```
