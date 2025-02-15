# Lab Deploy com Cloud Run - GoExpert Weather 

 <br>

## Sobre o projeto
Este é o repositório destinado ao laboratório de Deploy com Cloud Run do curso Pós Goexpert da faculdade FullCycle. O projeto permite ao usuário visualizar a temperatura atual da cidade indicada, em diversas unidades de medida de temperatura.

   <br>

## Funcionalidades
- Receber um CEP pela url;
- Consultar a API ViaCEP para identificação do cep indicado;
- Consultar a temperatura na localização indicada, pela API WeatherAPI;
- Visualizar a temperatura recebida da API WeatherAPI, em outras unidades de medida.
  
   <br>

## Como executar o projeto

### Pré-requisitos
Antes de começar, você vai precisar ter instalado em sua máquina as seguintes ferramentas:
- [Git](https://git-scm.com)

- [VSCode](https://code.visualstudio.com/)

- [Rest Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)

 - [Docker](https://www.docker.com/)

-  [WEATHER_API_KEY](https://www.weatherapi.com/) (Necessário cadastro para gerar key)

 <br>

#### Acessando o repositório
```bash

# Clone este repositório

$ git clone https://github.com/pedrogutierresbr/lab-weather-api-pos-goexpert.git

```

 <br>
  

#### Executando a aplicação localmente
```bash

# Crie um arquivo .env no diretório ./cmd

# Adicione ao arquivo .env sua api key. WEATHER_API_KEY={valor api key}

# Abrir um terminal

# 1 - Importar os pacotes

$ go mod tidy

# 2 - Acessar diretório cmd

$ cd cmd/

# 3 - Executar a aplicação

$ go run main.go

```

 <br>

#### Executando a aplicação com docker-compose
```bash

# No arquivo docker-compose.yml, inserir sua api key no campo WEATHER_API_KEY: "valor api key"

# Abrir um terminal

# 1 - Importar os pacotes

$ go mod tidy

# 2 - Executar docker compose

$ docker-compose up -d

# 3 - Para pausar a aplicação

$ docker-compose down

```

 <br>

#### Realizando requisição localmente
Serviços estarão disponíveis na seguinte porta:

- Web Server : 8000
 
##### Web Server
```bash

# Abra a pasta API

# O arquivo get_weather.http possui diversos exemplos de como fazer uma consulta via url do navegador

# Alternativa via terminal 

$ curl http://localhost:8080/weather?cep=38050600

```

 <br>

#### Realizando requisição via endereço ativo
Aplicação está disponível no Google Cloud Run no seguinte endereço:

https://lab-weather-api-510231468538.us-central1.run.app

```bash
Para fazer um requisição, basta adicionar /weather?cep={cep} na url.

Exemplo:

https://lab-weather-api-510231468538.us-central1.run.app/weather?cep=12220810
```

<br>

## Testes automatizados
```bash

# Abra um terminal

# Execute o seguinte comando

$ go test ./...

```

 <br>

## Licença

Este projeto esta sobe a licença [MIT](./LICENSE).

Feito por Pedro Gutierres [Entre em contato!](https://www.linkedin.com/in/pedrogabrielgutierres/)