# Exemplo de implementação de rate limit com Redis em Golang
### O objetivo do projeto é realizar o rate limit baseado em um IP ou TOKEN enviado pelo headers 



``Rate limit por IP com docker``

Para realizar o rate limit por ip com o docker você deve possuir o [docker](https://www.docker.com/) na sua máquina, você pode utilizar o seu servidor redis ou executar os seguintes comandos para criar um servidor redis localmente.

- Primeiro precisamos criar uma network para realizar a comunicação entre o redis e a aplicação
```sh
  docker network create redis-network
```
- Agora iremos iniciar um container redis.
```sh
  docker run -d --name my_redis --network redis-network redis:latest
```

- Para executar a aplicação você deve passar as seguintes variáveis \
  **REDIS_HOST** \
  **BLOCK_IP** \
  **BLOCK_IP_LIMIT** \
  **BLOCK_IP_TIME**

```sh
  docker run -p "8080:8080" --network redis-network \
    -e REDIS_HOST=my_redis:6379 \
    -e BLOCK_IP=127.0.0.1 \
    -e BLOCK_IP_LIMIT=2 \
    -e BLOCK_IP_TIME=3 \
    --rm italog09/go-ratelimit
```

🚧Atenção o projeto carece de algumas validações, por exemplo não existe validação para caso você passe um IP em formato inválido.🚧

``Rate limit por TOKEN com docker``

Para realizar o rate limit por token com o docker você deve possuir o [docker](https://www.docker.com/) na sua máquina, você pode utilizar o seu servidor redis ou executar os seguintes comandos para criar um servidor redis localmente.

- Primeiro precisamos criar uma network para realizar a comunicação entre o redis e a aplicação
```sh
  docker network create redis-network
```
- Agora iremos iniciar um container redis.
```sh
  docker run -d --name my_redis --network redis-network redis:latest
```

- Para executar a aplicação você deve passar as seguintes variáveis \
  **REDIS_HOST** \
  **BLOCK_TOKEN** \
  **BLOCK_TOKEN_LIMIT** \
  **BLOCK_TOKEN_TIME**

- O token é lido através do headers API_KEY

```sh
  docker run -p "8080:8080" --network redis-network \
    -e REDIS_HOST=my_redis:6379 \
    -e BLOCK_TOKEN=abc123 \
    -e BLOCK_TOKEN_LIMIT=2 \
    -e BLOCK_TOKEN_TIME=3 \
    --rm italog09/go-ratelimit
```

docker run -p "8080:8080" --network redis-network \
  -e REDIS_HOST=my_redis:6379 \
  -e BLOCK_IP=172.30.0.1 \
  -e BLOCK_IP_LIMIT=2 \
  -e BLOCK_IP_TIME=3 \
  --rm italog09/go-ratelimit

O projeto utiliza de migrations criadas usando [golang-migrate](https://github.com/golang-migrate/migrate) e [docker](https://www.docker.com/) para criar containers do mysql e rabbitmq

antes de iniciar o projeto execute os seguintes comandos
- passo 1
```sh
  docker compose up -d
```
- passo 2
```sh
  make migrate ou migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
```

Para rodar o projeto execute o comando a baixo dentro do diretório cmd/ordersystem

```sh
  go run main.go wire_gen.go
```

Para testar via HTTP use 

``Criar pedido``
```sh
    curl -X POST http://localhost:8000/order -H "Content-Type: application/json" -d "{"id":"a","price": 100.5,"tax": 0.5}"
```

``Listar pedidos``
```sh
    curl http://localhost:8000/order
```

Para testar via GraphQl playground

``Criar pedido``
>mutation createOrder {
  createOrder (input: {id: "graph_id", Price: 100.0, Tax: 0.2}) {
    id
    Price
    Tax
    FinalPrice
  }
>}

``Listar pedidos``
>query ListOrders {
  orders {
  id
  Price
  Tax
  FinalPrice
 }
>}

Para testar via gRpc utilizando o **evans** selecione o package *pb*, service *OrderService* e chame o *CreateOrder* para criar um pedido, ou *ListOrders* para listar pedidos