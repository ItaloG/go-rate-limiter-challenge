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
- Para realizar teste execute
```sh
    curl http://localhost:8080
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

- Para realizar teste execute
```sh
    curl http://localhost:8080 -H "API_KEY: abc123"
```

``Rate limit por TOKEN e IP com docker``

É possível passar tanto um token quando um ip para realizar o bloqueao de rate limit, mas caso a a request possua o token para ser bloquado as configuração de token serão sobrepostas a de ip. Por exemplo caso o token seja **abc123** e o ip da request **192.168.0.1** e nas configurações esteja para 6 tentativas em um segundo para o ip **192.168.0.1** e configurado 2 tentativas em tres segundo para o token **abc123**, a requisição sera bloqueada ao atingir o limite de token.

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
  **BLOCK_IP_TIME** \
  **BLOCK_TOKEN** \
  **BLOCK_TOKEN_LIMIT** \
  **BLOCK_TOKEN_TIME**

- O token é lido através do headers API_KEY

```sh
  docker run -p "8080:8080" --network redis-network \
    -e REDIS_HOST=my_redis:6379 \
    -e BLOCK_IP=127.0.0.1 \
    -e BLOCK_IP_LIMIT=2 \
    -e BLOCK_IP_TIME=3 \
    -e BLOCK_TOKEN=abc123 \
    -e BLOCK_TOKEN_LIMIT=2 \
    -e BLOCK_TOKEN_TIME=3 \
    --rm italog09/go-ratelimit
```

- Para realizar teste execute
```sh
    curl http://localhost:8080 -H "API_KEY: abc123"
```

``Rate limit por IP rodando local``

Para realizar o rate limit por ip com o docker você deve possuir o [docker](https://www.docker.com/) na sua máquina, você pode utilizar o seu servidor redis ou executar os seguintes comandos para criar um servidor redis localmente.

- Primeiro iremos iniciar um container redis.
```sh
  docker run -d -p "6379:6379" --name my_redis redis:latest
```

- Para executar a aplicação você deve criar um arquivo .env passar as seguintes variáveis (você pode seguir o exemplo do arquivo .env.example) \
  **REDIS_HOST** \
  **BLOCK_IP** \
  **BLOCK_IP_LIMIT** \
  **BLOCK_IP_TIME** 

```sh
  go run cmd/main.go
```

- Para realizar teste execute
```sh
    curl http://localhost:8080
```

🚧Atenção o projeto carece de algumas validações, por exemplo não existe validação para caso você passe um IP em formato inválido.🚧

``Rate limit por TOKEN rodando local``

Para realizar o rate limit por token com o docker você deve possuir o [docker](https://www.docker.com/) na sua máquina, você pode utilizar o seu servidor redis ou executar os seguintes comandos para criar um servidor redis localmente.

- Primeiro iremos iniciar um container redis.
```sh
  docker run -d -p "6379:6379" --name my_redis redis:latest
```

- Para executar a aplicação você deve criar um arquivo .env passar as seguintes variáveis (você pode seguir o exemplo do arquivo .env.example) \
  **REDIS_HOST** \
  **BLOCK_TOKEN** \
  **BLOCK_TOKEN_LIMIT** \
  **BLOCK_TOKEN_TIME**

- O token é lido através do headers API_KEY

```sh
  go run cmd/main.go
```

- Para realizar teste execute
```sh
    curl http://localhost:8080 -H "API_KEY: abc123"
```

``Rate limit por TOKEN e IP rodando local``

É possível passar tanto um token quando um ip para realizar o bloqueao de rate limit, mas caso a a request possua o token para ser bloquado as configuração de token serão sobrepostas a de ip. Por exemplo caso o token seja **abc123** e o ip da request **192.168.0.1** e nas configurações esteja para 6 tentativas em um segundo para o ip **192.168.0.1** e configurado 2 tentativas em tres segundo para o token **abc123**, a requisição sera bloqueada ao atingir o limite de token.

- Primeiro iremos iniciar um container redis.
```sh
  docker run -d -p "6379:6379" --name my_redis redis:latest
```

- Para executar a aplicação você deve criar um arquivo .env passar as seguintes variáveis (você pode seguir o exemplo do arquivo .env.example)s \
  **REDIS_HOST** \
  **BLOCK_IP** \
  **BLOCK_IP_LIMIT** \
  **BLOCK_IP_TIME** \
  **BLOCK_TOKEN** \
  **BLOCK_TOKEN_LIMIT** \
  **BLOCK_TOKEN_TIME**

- O token é lido através do headers API_KEY

```sh
  go run cmd/main.go
```

- Para realizar teste execute
```sh
    curl http://localhost:8080 -H "API_KEY: abc123"
```
