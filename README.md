# API de Transações (PISMO)

API RESTful desenvolvida em Go para gerenciamento de contas e transações financeiras.

## Arquitetura e Boas Práticas

* **Clean Architecture**: Separação entre Domínio, Aplicação (UseCases), Apresentação (Handlers) e Infraestrutura (Repositórios).
* **Open/Closed Principle**: A lógica financeira é orientada por dados (`Multiplier` na entidade `OperationType`). Novos tipos de operações podem ser criados sem alterar o `TransactionUseCase`.
* **Validação usando Fail-Fast**: Utilização de *Struct Tags* via Gin/Validator na camada HTTP, blindando a aplicação contra requisições malformadas.
* **Padrão Repository & DI**: Implementações da Interface.

## Executar via Docker

```bash
git clone https://github.com/jonathantvrs/pchlng.git
cd pchlng
docker-compose up --build
```
