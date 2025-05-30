# 📚 Sistema de Cadastro de Estudantes - School CLI

Este é um sistema de linha de comando (CLI) para cadastro, listagem e exclusão de estudantes. Ele foi desenvolvido com a linguagem **Go**, utilizando boas práticas como **DDD (Domain-Driven Design)** e **TDD (Test-Driven Development)**. O sistema possui suporte a dois tipos de repositórios: **em memória** e **SQLite**.

---

## ✅ O que o sistema faz?

O `School CLI` é uma aplicação de terminal que permite:

- ✅ Registrar estudantes com nome, email e idade;
- 📋 Listar todos os estudantes cadastrados;
- ❌ Remover estudantes pelo email (como identificador único).

---

## 🛠️ Tecnologias e Ferramentas Usadas

- **Go (Golang)** – Linguagem de programação principal;
- **Cobra CLI** – Biblioteca para criar comandos via terminal;
- **SQLite** – Banco de dados leve usado como persistência;
- **TDD (Test Driven Development)** – Aplicação testada desde o início;
- **DDD (Domain Driven Design)** – Arquitetura centrada na lógica de negócio.

---

## 🚀 Instalação

### 1. Instalar o Go

Baixe e instale o Go:  
➡️ https://go.dev/doc/install

Verifique se o Go foi instalado corretamente:

```bash
go version
```
### 2. Instalar Cobra
Use o seguinte comando para instalar a CLI Cobra globalmente (opcional):

```bash
go install github.com/spf13/cobra-cli@latest
```
Ou adicione no projecto
```bash
go get -u github.com/spf13/cobra
```

### 3. Instalar SQLite (no sistema)
**Linux (Debian/Ubuntu):**

```bash
sudo apt update
sudo apt install sqlite3 libsqlite3-dev
```

**macOS (via Homebrew):**

```bash
brew install sqlite3
```

**Windows:**
Baixe em: [https://www.sqlite.org/download.html](https://www.sqlite.org/download.html)

Caso queira instalar apenas no projecto
```bash
  go get github.com/mattn/go-sqlite3
```

## ⚙️ Compilando o projeto
Para rodar sem compilar:

```bash
go run .
```
Para compilar e gerar o binário:

```bash
go build -o School
```

## 💻 Como usar pelo terminal
**1. Listar todos os estudantes:**
```bash
./School student list
```

**2. Registrar um novo estudante:**
```bash
./School student register --nome "Kiala Emanuel" --email "kiala@gmail.com" --idade 22
```
**3. Deletar um estudante pelo email:**

```bash
./School student delete --email "kiala@gmail.com"
```
## 🧪 Metodologias Aplicadas
### 🧩 DDD (Domain-Driven Design)
O sistema é organizado por contexto, separando domínio, casos de uso, serviços e adaptadores (como CLI e persistência). Isso garante baixo acoplamento e alta coesão.

### ✅ TDD (Test-Driven Development)
O desenvolvimento começou pelos testes, garantindo que cada funcionalidade fosse validada desde o início. Exemplo: TestRegisterStudentInRepository.

## 🗃️ Repositórios
**🧠 Em Memória:** Útil para testes e simulações rápidas.

**💾 SQLite:** Persistência real dos dados entre execuções.

Você pode alternar entre eles modificando a criação do repositório no seu main.go ou adicionar outro a sua escolha.


© KIALA EMANUEL - 29 de Maio de 2025
📧 [kialamanuelkm@gmail.com](kialamanuelkm@gmail.com)
📱 WhatsApp:  [955649313](wa.me/244955649313?text=Olá%20Kiala%2C%20vi%20seu%20projeto%20School%20CLI%20e%20gostaria%20de%20saber%20mais.) / [972520308](wa.me/244972520308)
