# Easy User GoLang
  
Essa aplicação se trata de um backend simples de gerenciamento de usuários em um banco de dados MySQL.
Este repositório contém uma aplicação Go que se conecta a um banco de dados MySQL, tudo orquestrado via Docker Compose.
Através de um Client como Insomnia, Postman ou outros, é possível acessar recursos como:

### POST: /v1/users

**Descrição:**  
Criar um novo usuário.

**Enviar:**  
```json
{
	"name": "Novo usuário", 
	"email": "usuario@email.com", 
	"age": 24
}
```

**Resposta:** 

```json
{
	"id": 1,
	"name": "Novo usuário",
	"email": "usuario@email.com",
	"age": 24
}
```

### GET:  /v1/users/{userId}

**Descrição:**  
Obter informações de um usuário específico.

**Enviar:**     
Um número identificador como {userId}, por exemplo: 1, 2, 3.

**Resposta:**

```json
{
    "id": 1,
    "name": "Novo usuário", 
    "email": "usuario@email.com", 
    "age": 24
}
```
### GET:  /v1/users?page={page}&limit={limit} (opcional)

**Descrição:**  
Listar usuários com paginação.

Paramêtros (Queries) **(opcional)**:

**page:** Qual página deve ser visualizada.     
**limit:** Quantidade de usuários por página.

**Resposta:**
```json
{
	"page": 1,
	"limit": 1,
	"totalPages": 1,
	"items": [
		{
			"id": 1,
			"name": "Novo usuário",
			"email": "usuario@email.com",
			"age": 24
		}
	]
}
```

### PUT:  /v1/users/

**Descrição:**      
Atualizar um usuário existente.

**Enviar:**

```json
{
    "id": 1,
    "name": "Novo usuário atualizado", 
    "email": "usuario2@email.com", 
    "age": 24
}
```

**Resposta:**
```json
{
    "id": 1,
	"name": "Novo usuário atualizado", 
	"email": "usuario2@email.com", 
	"age": 24
}
```

### DELETE: /v1/users/{userId}

**Descrição:**  
Remover um usuário específico.

**Enviar:**     
Um número identificador como {userId}, por exemplo: 1, 2, 3.

**Resposta:**   
```204 (NoContent) - Realizado com Sucesso```

## Regras e Validações
- Não é possível existir mais de um usuário com o mesmo e-mail.
- Os dados como idade, formato de e-mail precisam ser válidos.
- As operações CRUD estão sendo validadas.
- Os paramêtros nas requisições são validados.

## Pré-requisitos

Antes de começar, você precisa ter instalado em sua máquina:

- **Docker**: [Instruções de instalação do Docker](https://docs.docker.com/get-docker/)
- **Docker Compose**: [Instruções de instalação do Docker Compose](https://docs.docker.com/compose/install/)
- **Git** (opcional, para clonar o repositório): [Instalação do Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Passo a passo para rodar a aplicação

### 1. Clonar o repositório

Se você ainda não clonou o repositório, faça isso com o comando:

```bash
git clone https://github.com/rodrigosscode/easy-user.git
cd easy-user
```

### 2. Rodar o comando:
```bash
cd docker-compose
docker-compose up --build
```

Pronto! Se tudo ocorrer bem a sua aplicação estará de pé em http://localhost:8080

## Bônus

Na **V1** o projeto foi desenvolvido buscando os pontos especificos da **Arquitetura Limpa** sob a interação das camadas.    

Veja em:    
https://github.com/rodrigosscode/easy-user/tree/release/v1
