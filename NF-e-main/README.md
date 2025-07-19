# API de NFe - Sistema de Nota Fiscal Eletrônica

Esta é uma API REST desenvolvida em Go para gerenciar destinatários e criar Notas Fiscais Eletrônicas (NFe).

## 🚀 Funcionalidades

- **Gestão de Remetentes**: CRUD completo para remetentes/emissores de NFes
- **Gestão de Destinatários**: CRUD completo para destinatários de produtos/serviços
- **Gestão de NFes**: Criação, consulta, atualização e remoção de NFes
- **Controle de Status**: Autorização e cancelamento de NFes
- **Validações**: Validação de CPF/CNPJ, campos obrigatórios e regras de negócio
- **Cálculos Automáticos**: Cálculo automático de impostos e valores totais
- **Integração SEFAZ**: Geração de XML e envio para SEFAZ em homologação

## 📋 Pré-requisitos

- Go 1.24.5 ou superior
- Git

## 🛠️ Instalação

1. Clone o repositório:
```bash
git clone <url-do-repositorio>
cd go
```

2. Instale as dependências:
```bash
go mod tidy
```

3. Execute a aplicação:
```bash
go run cmd/main.go
```

A API estará disponível em `http://localhost:8000`

## 📚 Endpoints da API

### Health Check
```
GET /ping
```

### Remetentes

#### Criar Remetente
```
POST /api/remetentes
```

**Body:**
```json
{
  "nome": "Empresa XYZ Ltda",
  "cpf_cnpj": "12345678000199",
  "tipo_pessoa": "J",
  "email": "contato@xyz.com",
  "telefone": "(11) 3333-4444",
  "endereco": {
    "logradouro": "Av. Paulista",
    "numero": "1000",
    "complemento": "Sala 1001",
    "bairro": "Bela Vista",
    "cidade": "São Paulo",
    "estado": "SP",
    "cep": "01310-100",
    "pais": "Brasil"
  },
  "inscricao_estadual": "123456789012",
  "inscricao_municipal": "987654321098",
  "regime_tributario": "3",
  "razao_social": "Empresa XYZ Ltda",
  "nome_fantasia": "XYZ",
  "website": "www.xyz.com",
  "banco": "Banco do Brasil",
  "agencia": "1234",
  "conta": "12345-6"
}
```

#### Listar Remetentes
```
GET /api/remetentes
```

#### Buscar Remetente
```
GET /api/remetentes/:id
```

#### Atualizar Remetente
```
PUT /api/remetentes/:id
```

#### Remover Remetente
```
DELETE /api/remetentes/:id
```

### Destinatários

#### Criar Destinatário
```
POST /api/destinatarios
```

**Body:**
```json
{
  "nome": "João Silva",
  "cpf_cnpj": "12345678901",
  "tipo_pessoa": "F",
  "email": "joao@email.com",
  "telefone": "(11) 99999-9999",
  "endereco": {
    "logradouro": "Rua das Flores",
    "numero": "123",
    "complemento": "Apto 45",
    "bairro": "Centro",
    "cidade": "São Paulo",
    "estado": "SP",
    "cep": "01234-567",
    "pais": "Brasil"
  },
  "inscricao_estadual": "123456789",
  "inscricao_municipal": "987654321"
}
```

#### Listar Destinatários
```
GET /api/destinatarios
```

#### Buscar Destinatário
```
GET /api/destinatarios/:id
```

#### Atualizar Destinatário
```
PUT /api/destinatarios/:id
```

#### Remover Destinatário
```
DELETE /api/destinatarios/:id
```

### NFes

#### Criar NFe
```
POST /api/nfes
```

**Body:**
```json
{
  "numero": "001",
  "serie": "1",
  "data_emissao": "2024-01-15T10:00:00Z",
  "data_saida": "2024-01-15T10:00:00Z",
  "remetente": {
    "nome": "Empresa XYZ Ltda",
    "cpf_cnpj": "12345678000199",
    "tipo_pessoa": "J",
    "email": "contato@xyz.com",
    "telefone": "(11) 3333-4444",
    "endereco": {
      "logradouro": "Av. Paulista",
      "numero": "1000",
      "complemento": "Sala 1001",
      "bairro": "Bela Vista",
      "cidade": "São Paulo",
      "estado": "SP",
      "cep": "01310-100",
      "pais": "Brasil"
    },
    "inscricao_estadual": "123456789012",
    "inscricao_municipal": "987654321098",
    "regime_tributario": "3",
    "razao_social": "Empresa XYZ Ltda",
    "nome_fantasia": "XYZ",
    "website": "www.xyz.com",
    "banco": "Banco do Brasil",
    "agencia": "1234",
    "conta": "12345-6"
  },
  "destinatario": {
    "nome": "João Silva",
    "cpf_cnpj": "12345678901",
    "tipo_pessoa": "F",
    "email": "joao@email.com",
    "telefone": "(11) 99999-9999",
    "endereco": {
      "logradouro": "Rua das Flores",
      "numero": "123",
      "complemento": "Apto 45",
      "bairro": "Centro",
      "cidade": "São Paulo",
      "estado": "SP",
      "cep": "01234-567",
      "pais": "Brasil"
    }
  },
  "itens": [
    {
      "codigo": "PROD001",
      "descricao": "Produto Teste",
      "quantidade": 2,
      "unidade_medida": "UN",
      "valor_unitario": 100.00,
      "aliquota_icms": 18.00,
      "aliquota_ipi": 0.00,
      "aliquota_pis": 1.65,
      "aliquota_cofins": 7.60
    }
  ]
}
```

#### Listar NFes
```
GET /api/nfes
```

#### Buscar NFe
```
GET /api/nfes/:id
```

#### Atualizar NFe
```
PUT /api/nfes/:id
```

#### Remover NFe
```
DELETE /api/nfes/:id
```

#### Autorizar NFe
```
POST /api/nfes/:id/autorizar
```

#### Cancelar NFe
```
POST /api/nfes/:id/cancelar
```

### SEFAZ

#### Enviar NFe para SEFAZ
```
POST /api/sefaz/enviar
```

**Body:**
```json
{
  "nfe_id": "20241201120001"
}
```

#### Consultar Status SEFAZ
```
GET /api/sefaz/status
```

#### Gerar XML da NFe
```
GET /api/sefaz/xml/:id
```

## 📊 Estrutura do Projeto

```
go/
├── cmd/
│   └── main.go                 # Ponto de entrada da aplicação
├── internal/
│   ├── models/
│   │   ├── destinatario.go     # Modelos de dados do destinatário
│   │   └── nfe.go             # Modelos de dados da NFe
│   ├── services/
│   │   ├── destinatario_service.go  # Lógica de negócio dos destinatários
│   │   └── nfe_service.go           # Lógica de negócio das NFes
│   ├── handlers/
│   │   ├── destinatario_handler.go  # Handlers HTTP dos destinatários
│   │   └── nfe_handler.go           # Handlers HTTP das NFes
│   └── routes/
│       └── routes.go               # Configuração das rotas
├── go.mod
├── go.sum
└── README.md
```

## 🔧 Validações

### Remetente
- Nome obrigatório
- CPF/CNPJ obrigatório e válido (11 dígitos para CPF, 14 para CNPJ)
- Tipo de pessoa obrigatório ("F" para Física, "J" para Jurídica)
- Email obrigatório e válido
- Endereço completo obrigatório
- Inscrição estadual obrigatória
- Regime tributário obrigatório ("1" - Simples Nacional, "2" - Simples Nacional - excesso, "3" - Regime Normal)

### Destinatário
- Nome obrigatório
- CPF/CNPJ obrigatório e válido (11 dígitos para CPF, 14 para CNPJ)
- Tipo de pessoa obrigatório ("F" para Física, "J" para Jurídica)
- Email obrigatório e válido
- Endereço completo obrigatório

### NFe
- Número e série obrigatórios
- Datas de emissão e saída obrigatórias
- Remetente obrigatório
- Destinatário obrigatório
- Pelo menos um item obrigatório
- Quantidade e valor unitário dos itens devem ser maiores que zero

## 📈 Status das NFes

- **rascunho**: NFe criada mas não autorizada (pode ser editada/excluída)
- **autorizada**: NFe autorizada (não pode ser editada)
- **cancelada**: NFe cancelada (não pode ser editada)

## 🧪 Exemplos de Uso

### Criar um Destinatário
```bash
curl -X POST http://localhost:8000/api/destinatarios \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "Maria Santos",
    "cpf_cnpj": "98765432100",
    "tipo_pessoa": "F",
    "email": "maria@email.com",
    "telefone": "(11) 88888-8888",
    "endereco": {
      "logradouro": "Av. Paulista",
      "numero": "1000",
      "bairro": "Bela Vista",
      "cidade": "São Paulo",
      "estado": "SP",
      "cep": "01310-100",
      "pais": "Brasil"
    }
  }'
```

### Criar uma NFe
```bash
curl -X POST http://localhost:8000/api/nfes \
  -H "Content-Type: application/json" \
  -d '{
    "numero": "001",
    "serie": "1",
    "data_emissao": "2024-01-15T10:00:00Z",
    "data_saida": "2024-01-15T10:00:00Z",
    "destinatario": {
      "nome": "Maria Santos",
      "cpf_cnpj": "98765432100",
      "tipo_pessoa": "F",
      "email": "maria@email.com",
      "endereco": {
        "logradouro": "Av. Paulista",
        "numero": "1000",
        "bairro": "Bela Vista",
        "cidade": "São Paulo",
        "estado": "SP",
        "cep": "01310-100",
        "pais": "Brasil"
      }
    },
    "itens": [
      {
        "codigo": "PROD001",
        "descricao": "Notebook Dell",
        "quantidade": 1,
        "unidade_medida": "UN",
        "valor_unitario": 3500.00,
        "aliquota_icms": 18.00,
        "aliquota_pis": 1.65,
        "aliquota_cofins": 7.60
      }
    ]
  }'
```

## 🚨 Observações

- Esta é uma implementação de demonstração que usa armazenamento em memória
- Para produção, recomenda-se implementar persistência em banco de dados
- As validações de CPF/CNPJ são básicas (apenas quantidade de dígitos)
- Os cálculos de impostos são simplificados para demonstração

## 📝 Licença

Este projeto está sob a licença MIT. 