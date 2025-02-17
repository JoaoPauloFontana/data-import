# Como Rodar o Projeto

Este guia fornece instruções para executar o projeto usando Docker e Docker Compose.

## Pré-requisitos

- **Docker**: Certifique-se de que o Docker está instalado em sua máquina. [Instruções de instalação](https://docs.docker.com/get-docker/)
- **Docker Compose**: Certifique-se de que o Docker Compose está instalado. [Instruções de instalação](https://docs.docker.com/compose/install/)

## Configuração do Projeto

1. **Clone o Repositório**

   Primeiro, clone o repositório para sua máquina local:

   ```bash
   git clone https://github.com/JoaoPauloFontana/data-import.git
   cd data-import

2. **Construa e Inicie os Containers**

   Navegue até o diretório raiz do projeto e execute o seguinte comando para construir e iniciar os containers:

   ```bash
   docker-compose up -d --build
   ```

   ## OBS: Espere o import finalizar o processo, pois será preciso inserir os dados no banco para continuar

3. **Acesse o Frontend**

   Abra seu navegador e vá para:

   ```bash
   http://localhost:3000
   ```

   ## Se necessário acessar a api, a url é:

   ```bash
   http://localhost:8200
   ```

   ## Endpoints:

   - Login:
   ```bash
   curl --location 'http://localhost:8200/login' --header 'Content-Type: application/json' --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE5MzYzMzMsInVzZXJuYW1lIjoidGVzdCJ9.6CBKTy8ePMlCzVsLNJNRRRk0bFDbKzawwlqbZ9oV8Ss' --data '{"username":"test","password":"password"}'
   ```

   - Records:
   ```bash
   curl --location 'http://localhost:8200/api/records' --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE5MzYzMzMsInVzZXJuYW1lIjoidGVzdCJ9.6CBKTy8ePMlCzVsLNJNRRRk0bFDbKzawwlqbZ9oV8Ss'
   ```

5. **Credenciais de Login**

   Use as seguintes credenciais para fazer login:

   - Login: 'test'
   - Password: 'password'
