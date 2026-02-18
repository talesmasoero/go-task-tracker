# Task Tracker

### Funcionalidades iniciais do programa (o que o usuário deve conseguir fazer)

- Inicialmente: O usuário deve conseguir criar e visualizar tasks
- Depois: Finalizar com edição e deleção

### Domínio/Entidades do programa

Task

- Inicialmente: Apenas ID e Description
- Depois: Adicionar os campos de Status ("enum"), createdAt e updatedAt

### Rotas

- `POST /tasks`: criar/adicionar task
- `GET /tasks`: listar todas as tasks

### Como o programa deve funcionar/rodar (onde o usuário irá usá-lo)

- Inicialmente: Salvando dados apenas em memória (enquanto o programa roda), funcionando como uma API (localhost)
- Depois (opcional): Refatorar o código (organizar pacotes, funções e métodos), aplicar DI e DIP, fazer modelo CLI, salvar dados em JSON, etc
