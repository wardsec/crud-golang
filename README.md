
# Loja Golang - CRUD

Este projeto foi feito para testar conceitos de crud na linaguem Go, utlizando o padrao MVC.


## Conectando ao banco:
Para conectar em nosso banco de dados utilizamos o arquivo db.go localizado na pasta db.
``````
conexao := "user=postgres dbname=loja_golang password= host=localhost sslmode=disable"

``````
## Como utilizar
Primeiro de tudo para utilizar devemos criar nosso modelo de banco, neste caso utilizei o postgresql:
``````
CREATE TABLE produtos(
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
  quantidade integer)
``````
Para rodar o programa utilizamos o seguinte comando: go run main.go
ou se preferir um live server utilize o comando fresh


