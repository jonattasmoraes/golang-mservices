**[Click here for the English version](#english_version)**

# 📦 StellarCommerce: The Sales Microservice

![License](https://img.shields.io/badge/license-MIT-blue.svg) ![Go](https://img.shields.io/badge/Go-1.22-blue.svg) ![gRPC](https://img.shields.io/badge/gRPC-Enabled-brightgreen) ![Kafka](https://img.shields.io/badge/Kafka-Enabled-brightgreen) ![Docker](https://img.shields.io/badge/Docker-Enabled-blue) ![Build](https://img.shields.io/badge/Build-Passing-brightgreen) ![Tests](https://img.shields.io/badge/Tests-Passing-brightgreen) ![Coverage](https://img.shields.io/badge/Coverage-56%25-orange)

## 📝 Descrição

**StellarCommerce** é uma aplicação de vendas robusta e escalável, composta por quatro microserviços, cada um nomeado como um satélite de Saturno. Este projeto utiliza gRPC para comunicação entre os microserviços de usuários, produtos e vendas, e Kafka para transmitir informações de vendas para o serviço de relatórios. A API possui um gateway simples e está documentada com Swagger.

## 🚀 Microserviços

- **Titan:** Microserviço de usuários
  - [Repositório no GitHub](https://github.com/jonattasmoraes/titan)
- **Mimas:** Microserviço de produtos
  - [Repositório no GitHub](https://github.com/jonattasmoraes/mimas)
- **Telesto:** Microserviço de vendas
  - [Repositório no GitHub](https://github.com/jonattasmoraes/telesto)
- **Dione:** Microserviço de relatórios
  - [Repositório no GitHub](https://github.com/jonattasmoraes/dione)

## 🌐 Monorepo

Todos os microserviços estão organizados em um monorepo:
- [Monorepo no GitHub](https://github.com/jonattasmoraes/golang-mservices)

## 🛠️ Tecnologias Utilizadas

- **Golang 1.22**
- **gRPC**
- **Kafka**
- **Docker**
- **Docker Compose**

## 📂 Estrutura do Projeto

```plaintext
StellarCommerce/
├── titan/
├── mimas/
├── telesto/
├── dione/
├── kafka/
├── gateway/
├── makefile
└── README.md
```


## ⚙️ Configuração e Execução

### Pré-requisitos

- **Docker** e **Docker Compose** instalados.
- **Golang 1.22** instalado.
- **Make** instalado.

### Passo a Passo

1. **Clone o monorepo:**
git clone https://github.com/jonattasmoraes/stellar_commerce.git
cd stellar_commerce

2. **Suba os containers com Docker Compose:**

```markdown
make run
```

3. **Acesse os microserviços:**
- **Titan (Usuários):** `http://localhost:8081`
- **Mimas (Produtos):** `http://localhost:8082`
- **Telesto (Vendas):** `http://localhost:8083`
- **Dione (Relatórios):** `http://localhost:8084`

ou acesse a API Gateway:
- **GateWay:** `http://localhost:3000`

## 📖 Documentação da API

Cada microserviço possui sua própria documentação da API, disponível nos respectivos repositórios:

- [Titan](https://github.com/jonattasmoraes/titan)
- [Mimas](https://github.com/jonattasmoraes/mimas)
- [Telesto](https://github.com/jonattasmoraes/telesto)
- [Dione](https://github.com/jonattasmoraes/dione)

## 🧪 Testes

Para executar os testes, navegue até o diretório do microserviço desejado e utilize o comando:

```bash
go test ./...
```

ou execute o makefile na raiz do monorepo:

```markdown
make test
```

## 📄 Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 🤝 Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

## 📧 Contato

- **Email:** jonattasmoraes@hotmail.com
- **LinkedIn:** [Jonattas-Moraes](https://www.linkedin.com/in/jonattas-moraes/)

<div id='english_version'/>
---

# 📦 StellarCommerce: The Sales Microservice (English Version)

![License](https://img.shields.io/badge/license-MIT-blue.svg) ![Go](https://img.shields.io/badge/Go-1.22-blue.svg) ![gRPC](https://img.shields.io/badge/gRPC-Enabled-brightgreen) ![Kafka](https://img.shields.io/badge/Kafka-Enabled-brightgreen) ![Docker](https://img.shields.io/badge/Docker-Enabled-blue) ![Build](https://img.shields.io/badge/Build-Passing-brightgreen) ![Tests](https://img.shields.io/badge/Tests-Passing-brightgreen) ![Coverage](https://img.shields.io/badge/Coverage-56%25-orange)

## 📝 Description

**StellarCommerce** is a robust and scalable sales application composed of four microservices, each named after a satellite of Saturn. This project uses gRPC for communication between user, product, and sales microservices, and Kafka to transmit sales information to the reporting service. The API has a simple gateway and is documented with Swagger.

## 🚀 Microservices

- **Titan:** User microservice
  - [GitHub Repository](https://github.com/jonattasmoraes/titan)
- **Mimas:** Product microservice
  - [GitHub Repository](https://github.com/jonattasmoraes/mimas)
- **Telesto:** Sales microservice
  - [GitHub Repository](https://github.com/jonattasmoraes/telesto)
- **Dione:** Reports microservice
  - [GitHub Repository](https://github.com/jonattasmoraes/dione)

## 🌐 Monorepo

All microservices are organized in a monorepo:
- [Monorepo on GitHub](https://github.com/jonattasmoraes/golang-mservices)

## 🛠️ Technologies Used

- **Golang 1.22**
- **gRPC**
- **Kafka**
- **Docker**
- **Docker Compose**

## 📂 Project Structure

```plaintext
StellarCommerce/
├── titan/
├── mimas/
├── telesto/
├── dione/
├── kafka/
├── gateway/
├── makefile
└── README.md
```

## ⚙️ Setup and Execution

### Prerequisites

- **Docker** e **Docker Compose** installed.
- **Golang 1.22** installed.
- **Make** installed.

### Steps

1. **Clone the monorepo:**
git clone https://github.com/jonattasmoraes/golang-mservices.git
cd stellar_commerce

2. **Bring up the containers with Docker Compose:**

```bash
make run
```

3. **Access the microservices:**
- **Titan (Users):** `http://localhost:8081`
- **Mimas (Products):** `http://localhost:8082`
- **Telesto (Sales):** `http://localhost:8083`
- **Dione (Reports):** `http://localhost:8084`

or access the API Gateway:
- **GateWay:** `http://localhost:3000`

## 📖 API Documentation

Each microservice has its own API documentation available in their respective repositories:

- [Titan](https://github.com/jonattasmoraes/titan)
- [Mimas](https://github.com/jonattasmoraes/mimas)
- [Telesto](https://github.com/jonattasmoraes/telesto)
- [Dione](https://github.com/jonattasmoraes/dione)

## 🧪 Tests

To run tests, navigate to the desired microservice directory and use the command:

```bash
go test ./...
```

or run the makefile at the root of the monorepo:

```markdown
make test
```

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) for details.

## 🤝 Contribution

Contributions are welcome! Feel free to open issues and pull requests.

## 📧 Contact

- **Email:** jonattasmoraes@hotmail.com
- **LinkedIn:** [Jonattas-Moraes](https://www.linkedin.com/in/jonattas-moraes/)
