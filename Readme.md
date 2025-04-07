
# 🌱 AgroHero API

A **AgroPlusUltra API** é uma plataforma para o monitoramento inteligente de culturas agronômicas, construída em **Go** utilizando o framework **Gin**. A API oferece informações detalhadas sobre diversas culturas agrícolas e tipos de solo, além de integrar drones para análise e monitoramento das plantações.

## 🚀 Funcionalidades

### 🌾 Informações das Culturas
- **Nome científico e família botânica**
- **Tipo de solo recomendado**
- **Temperatura ideal para cultivo**
- **Época ideal de plantio**
- **Necessidade hídrica e exigência de luz solar**

### 🐛 Pragas e Doenças Comuns
- **Nome da praga**
- **Sintomas**
- **Métodos de controle** (químico, biológico, cultural)

### 🌿 Manejo e Nutrição
- **Necessidade de adubação** (NPK recomendado)
- **Rotação de culturas sugerida**
- **Densidade de plantio**

### 🌾 Produção e Colheita
- **Tempo médio para colheita**
- **Produção média esperada por hectare**
- **Métodos de colheita**

### 🚁 Monitoramento via Drones
- **Análise de saúde das plantas** (NDVI)
- **Detecção de pragas** via visão computacional
- **Mapeamento da fazenda** para otimizar irrigação e aplicação de fertilizantes
- **Reconhecimento de falhas no plantio** para replantio automático

## 🛠️ Tecnologias Utilizadas
- **Backend:** Go (Gin framework)
- **Banco de Dados:** PostgreSQL
- **IA:** Modelos de Hugging Face ou outro modelo de visão computacional
- **Drones:** Integração para análise de imagens e dados
- **Containerização:** Docker

## 📦 Instalação e Uso

### Requisitos:
- **Go 1.18+** instalado
- **Docker** instalado
- **PostgreSQL** configurado

### Passos:

1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/agrohero-api.git
   cd agrohero-api
   ```

2. Configure as variáveis de ambiente no arquivo `.env`.

3. Instale as dependências do Go:
   ```bash
   go mod tidy
   ```

4. Suba a aplicação com Docker:
   ```bash
   docker-compose up -d
   ```

5. Acesse a API em `http://localhost:8080`.

## 📡 Endpoints da API

## 🌱 Tipos de Solo

| Método | Rota                     | Descrição                          |
| ------ | ------------------------ | ---------------------------------- |
| GET    | `/v1/tipos-de-solo`      | Lista todos os tipos de solo       |
| GET    | `/v1/tipos-de-solo/{id}` | Detalha um tipo de solo específico |
| POST   | `/v1/tipos-de-solo`      | Cria um novo tipo de solo          |
| PUT    | `/v1/tipos-de-solo/{id}` | Atualiza um tipo de solo           |
| DELETE | `/v1/tipos-de-solo/{id}` | Deleta um tipo de solo             |

---

## 🌾 Culturas Agrícolas

| Método | Rota                          | Descrição                               |
| ------ | ----------------------------- | --------------------------------------- |
| GET    | `/v1/culturas-agricolas`      | Lista todas as culturas agrícolas       |
| GET    | `/v1/culturas-agricolas/{id}` | Detalha uma cultura agrícola específica |
| POST   | `/v1/culturas-agricolas`      | Cria uma nova cultura agrícola          |
| PUT    | `/v1/culturas-agricolas/{id}` | Atualiza uma cultura agrícola           |
| DELETE | `/v1/culturas-agricolas/{id}` | Deleta uma cultura agrícola             |

---

## 🐛 Tipos de Pragas

| Método | Rota                       | Descrição                           |
| ------ | -------------------------- | ----------------------------------- |
| GET    | `/v1/tipos-de-pragas`      | Lista todos os tipos de pragas      |
| GET    | `/v1/tipos-de-pragas/{id}` | Detalha um tipo de praga específico |
| POST   | `/v1/tipos-de-pragas`      | Cria um novo tipo de praga          |
| PUT    | `/v1/tipos-de-pragas/{id}` | Atualiza um tipo de praga           |
| DELETE | `/v1/tipos-de-pragas/{id}` | Deleta um tipo de praga             |

---

## 🐞 Pragas

| Método | Rota              | Descrição                    |
| ------ | ----------------- | ---------------------------- |
| GET    | `/v1/pragas`      | Lista todas as pragas        |
| GET    | `/v1/pragas/{id}` | Detalha uma praga específica |
| POST   | `/v1/pragas`      | Cria uma nova praga          |
| PUT    | `/v1/pragas/{id}` | Atualiza uma praga           |
| DELETE | `/v1/pragas/{id}` | Deleta uma praga             |

---

## 🌾🆚🐞 Relação Pragas x Culturas

| Método | Rota                                                             | Descrição                                       |
| ------ | ---------------------------------------------------------------- | ----------------------------------------------- |
| GET    | `/v1/pragas-das-culturas-agricolas`                              | Lista todas as relações entre pragas e culturas |
| GET    | `/v1/pragas-das-culturas-agricolas/relacao?pestId=?&cultureId=?` | Lista uma realação entre praga e cultura        |
| POST   | `/v1/pragas-das-culturas-agricolas`                              | Cria uma nova relação entre praga e cultura     |
| PUT    | `/v1/pragas-das-culturas-agricolas/relacao?pestId=?&cultureId=?` | Atualiza uma relação entre praga e cultura      |
| DELETE | `/v1/pragas-das-culturas-agricolas/relacao?pestId=?&cultureId=?` | Deleta uma relação entre praga e cultura        |

---

## 🌾🆚💧 Relação Irrigação x Culturas

| Método | Rota                                                       | Descrição                                                |
| ------ | ---------------------------------------------------------- | -------------------------------------------------------- |
| GET    | `/v1/irrigacao-cultura/?cultureId=?&irrigationId=?`         | Busca recomendação de irrigação associada à cultura       |
| POST   | `/v1/irrigacao-cultura/`                                    | Cria uma nova recomendação de irrigação para uma cultura  |
| PUT    | `/v1/irrigacao-cultura/?cultureId=?&irrigationId=?`         | Atualiza recomendação de irrigação para uma cultura       |
| DELETE | `/v1/irrigacao-cultura/?cultureId=?&irrigationId=?`         | Deleta recomendação de irrigação associada à cultura      |


### 🚁 Monitoramento via Drones
| Método | Rota                             | Descrição |
|--------|----------------------------------|-----------|
| `POST` | `/v1/drones/monitoramento/ndvi` | Envia imagem para análise de saúde das plantas (NDVI) |
| `POST` | `/v1/drones/monitoramento/pragas`| Envia imagem para detecção de pragas via visão computacional |
| `GET`  | `/v1/drones/monitoramento/irrigacao`| Obtém dados de monitoramento para otimização de irrigação |

