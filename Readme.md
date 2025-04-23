
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

### ✅ Exemplo de Request: `POST /v1/tipos-de-solo`

```json
{
  "name": "Argiloso",
  "description": "Solo com alta capacidade de retenção de água e nutrientes."
}
```

---

## 🌾 Culturas Agrícolas

| Método | Rota                          | Descrição                               |
| ------ | ----------------------------- | --------------------------------------- |
| GET    | `/v1/culturas-agricolas`      | Lista todas as culturas agrícolas       |
| GET    | `/v1/culturas-agricolas/{id}` | Detalha uma cultura agrícola específica |
| POST   | `/v1/culturas-agricolas`      | Cria uma nova cultura agrícola          |
| PUT    | `/v1/culturas-agricolas/{id}` | Atualiza uma cultura agrícola           |
| DELETE | `/v1/culturas-agricolas/{id}` | Deleta uma cultura agrícola             |

### ✅ Exemplo de Request: `POST /v1/culturas-agricolas`

```json
{
  "name": "Milho",
  "variety": "Milho Doce",
  "soil_type_id": 1,
  "region": "SOUTHEAST",
  "use_type": "ALIMENTACAO_HUMANA",
  "ph_ideal_soil": 6.5,
  "max_temperature": 35.0,
  "min_temperature": 10.0,
  "excellent_temperature": 25.0,
  "weekly_water_requirement_max": 50.0,
  "weekly_water_requirement_min": 30.0,
  "sunlight_requirement": 8
}
```

---

## 🐛 Tipos de Pragas

| Método | Rota                       | Descrição                           |
| ------ | -------------------------- | ----------------------------------- |
| GET    | `/v1/tipos-de-pragas`      | Lista todos os tipos de pragas      |
| GET    | `/v1/tipos-de-pragas/{id}` | Detalha um tipo de praga específico |
| POST   | `/v1/tipos-de-pragas`      | Cria um novo tipo de praga          |
| PUT    | `/v1/tipos-de-pragas/{id}` | Atualiza um tipo de praga           |
| DELETE | `/v1/tipos-de-pragas/{id}` | Deleta um tipo de praga             |

### ✅ Exemplo de Request: `POST /v1/tipos-de-pragas`

```json
{
  "name": "Inseto"
}
```
---

## 🐞 Pragas

| Método | Rota              | Descrição                    |
| ------ | ----------------- | ---------------------------- |
| GET    | `/v1/pragas`      | Lista todas as pragas        |
| GET    | `/v1/pragas/{id}` | Detalha uma praga específica |
| POST   | `/v1/pragas`      | Cria uma nova praga          |
| PUT    | `/v1/pragas/{id}` | Atualiza uma praga           |
| DELETE | `/v1/pragas/{id}` | Deleta uma praga             |

### ✅ Exemplo de Request: `POST /v1/pragas` 

```json
{
  "name": "Lagarta do cartucho",
  "type_pest_id": 1
}

```

---

## 🌾🆚🐞 Relação Pragas x Culturas

| Método | Rota                                                             | Descrição                                       |
| ------ | ---------------------------------------------------------------- | ----------------------------------------------- |
| GET    | `/v1/pragas-das-culturas-agricolas`                              | Lista todas as relações entre pragas e culturas |
| GET    | `/v1/pragas-das-culturas-agricolas/relacao?pestId=?&cultureId=?` | Lista uma realação entre praga e cultura        |
| POST   | `/v1/pragas-das-culturas-agricolas`                              | Cria uma nova relação entre praga e cultura     |
| PUT    | `/v1/pragas-das-culturas-agricolas/relacao?pestId=?&cultureId=?` | Atualiza uma relação entre praga e cultura      |
| DELETE | `/v1/pragas-das-culturas-agricolas/relacao?pestId=?&cultureId=?` | Deleta uma relação entre praga e cultura        |

### ✅ Exemplo de Request: `POST /v1/pragas-das-culturas-agricolas`

```json
{
  "agriculture_culture_id": 1,
  "pest_id": 2,
  "description": "Causa danos nas folhas e reduz o rendimento da cultura.",
  "image": "https://exemplo.com/imagem-praga.jpg"
}
```

### ✅ Exemplo de Reponse: `GET /v1/pragas-das-culturas-agricolas/relacao?pestId=2&cultureId=lagartadocartucho?`
```json
{
  "agriculture_culture_name": "Milho",
  "pest_name": "Lagarta do cartucho",
  "description": "Causa danos severos nas folhas e espigas do milho, reduzindo a produtividade.",
  "image_url": "https://exemplo.com/imagens/lagarta-do-cartucho.jpg"
}
``` 
---
## 🌾🆚💧 Irrigação

| Método | Rota                                                       | Descrição                                                 |
| ------ | ---------------------------------------------------------- | --------------------------------------------------------  |
| GET    | `/v1/irrigação`                                            | Lista todas as irrigações                                 |
| GET    | `/v1/irrigação/id`                                         | Detalha uma irrigação específica                          |
| POST   | `/v1/irrigacao`                                            | Cria uma nova irrigação                                   |
| PUT    | `/v1/irrigacao/id`                                         | Atualiza irrigação                                        |
| DELETE | `/v1/irrigacao/id`                                         | Deleta irrigação                                          |

### ✅ Exemplo de Request: `POST /v1/irrigacao` 

```json
{
  "phenological_phase": "Floração",
  "phase_duration_days": 20,
  "irrigation_max": 60.0,
  "irrigation_min": 40.0,
  "description": "Durante a floração, recomenda-se irrigação moderada.",
  "unit": "mm"
}
```
---

## 🌾🆚💧 Relação Irrigação x Culturas

| Método | Rota                                                       | Descrição                                                 |
| ------ | ---------------------------------------------------------- | --------------------------------------------------------  |
| GET    | `/v1/irrigacao-cultura/?cultureId=?`                       | Busca recomendação de irrigação associada à cultura       |
| POST   | `/v1/irrigacao-cultura/`                                   | Cria uma nova recomendação de irrigação para uma cultura  |
| PUT    | `/v1/irrigacao-cultura/?cultureId=?&irrigationId=?`        | Atualiza recomendação de irrigação para uma cultura       |
| DELETE | `/v1/irrigacao-cultura/?cultureId=?&irrigationId=?`        | Deleta recomendação de irrigação associada à cultura      |

### ✅ Exemplo de Request: `POST irrigacao-cultura`

```json
{
  "agriculture_culture_id": 1,
  "irrigation_recomended_id": 1
}
```

### ✅ Exemplo de Response: `GET /v1/irrigacao-cultura/?cultureId=milho`

```json
[
  {
    "name": "Milho",
    "pheneological_phase": "Fase de floração",
    "phase_duration_days": 20,
    "irrigation_max": 30.0,
    "irrigation_min": 18.0,
    "unit": "mm/dia"
  },

  {
  "name": "Milho",
  "pheneological_phase": "Fase vegetativa",
  "phase_duration_days": 25,
  "irrigation_max": 25.0,
  "irrigation_min": 15.0,
  "unit": "mm/dia"
  }
]
```

---

## 🌿🦟🛡️ Relação Métodos Sustentáveis x Pragas x Culturas

| Método | Rota                                                                                                       | Descrição                                                                                  |
|--------|------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------|
| GET    | `/v1/controle-de-praga-agricultura`                                                                        | Lista todas as relações entre cultura, praga e método sustentável de controle              |
| GET    | `/v1/controle-de-praga-agricultura?agricultureCultureName=?&pestName=?&sustainablePestControlMethod=?`     | Retorna uma relação específica filtrada por cultura, praga e método sustentável            |
| POST   | `/v1/controle-de-praga-agricultura`                                                                        | Cria uma nova relação entre cultura, praga e método sustentável de controle                |
| PUT    | `/v1/controle-de-praga-agricultura?agricultureCultureName=?&pestName=?&sustainablePestControlMethod=?`     | Atualiza uma relação entre cultura, praga e método sustentável de controle                 |
| DELETE | `/v1/controle-de-praga-agricultura?agricultureCultureName=?&pestName=?&sustainablePestControlMethod=?`     | Deleta uma relação entre cultura, praga e método sustentável de controle                   |

### ✅ Exemplo de Request: ` POST /v1/controle-de-praga-agricultura`

```json
{
  "agriculture_culture_id": 1,
  "pest_id": 2,
  "sustainable_pest_control_id": 3,
  "description": "Uso de inimigos naturais para controle da praga."
}
```
### ✅ Exemplo de Response: `GET /v1/controle-de-praga-agricultura?agricultureCultureName=soja&pestName=percevejomarrom&sustainablePestControlMethod=biologico` 

```json
{
  "agriculture_culture_name": "Soja",
  "pest_name": "Percevejo-marrom",
  "sustainable_pest_control_method": "Controle biológico com parasitoides",
  "description": "Aplicação de vespas parasitoides para controle natural da população de percevejos."
}
```
---

## 📦🌱 Batchs (Lotes Agrícolas)

| Método | Rota                | Descrição                                           | Status esperado |
|--------|---------------------|-----------------------------------------------------|-----------------|
| POST   | `/v1/batchs/`       | Cria um novo lote agrícola                          | `201 Created`   |
| GET    | `/v1/batchs/`       | Lista todos os lotes agrícolas                      | `200 OK`        |
| GET    | `/v1/batchs/:id`    | Busca um lote agrícola pelo ID                      | `200 OK`        |
| PUT    | `/v1/batchs/:id`    | Atualiza os dados de um lote agrícola pelo ID       | `200 OK`        |
| DELETE | `/v1/batchs/:id`    | Deleta um lote agrícola pelo ID                     | `204 No Content`|

---

### 📤 Exemplo de Request (POST / PUT)

```json
{
  "name": "Lote Norte",
  "area": 12.5,
  "unit": "hectare"
}
```
---

### 📥 Exemplo de Response (GET /v1/batchs/:id)

```json
{
  "id": 1,
  "name": "Lote Norte",
  "area": 12.5,
  "unit": "hectare"
}
```
---

### 📥 Exemplo de Response (GET /v1/batchs)

```json
[
  {
    "id": 1,
    "name": "Lote Norte",
    "area": 12.5,
    "unit": "hectare"
  },
  {
    "id": 2,
    "name": "Lote Sul",
    "area": 8.3,
    "unit": "hectare"
  }
]
```
---


## 💰🌱 Custos de Produção

| Método | Rota                                     | Descrição                                 |
|--------|------------------------------------------|-------------------------------------------|
| GET    | `/v1/custos-plantacoes`                  | Lista todos os custos de produção         |
| GET    | `/v1/custos-plantacoes/:id`              | Detalha um custo de produção específico   |
| POST   | `/v1/custos-plantacoes`                  | Cria um novo custo de produção            |
| PUT    | `/v1/custos-plantacoes/:id`              | Atualiza um custo de produção             |
| DELETE | `/v1/custos-plantacoes/:id`              | Deleta um custo de produção               |

---

### ✅ Exemplo de Request: `POST /v1/custos-plantacoes`

```json
{
  "planting_id": 1,
  "item_name": "Adubo NPK",
  "unit": "kg",
  "quantity": 50.0,
  "cost_per_unit": 2.5,
  "cost_date": "2024-07-10T00:00:00Z"
}
```
---


## 💰🌱 Venda de Plantação

| Método | Rota                                     | Descrição                                 |
|--------|------------------------------------------|-------------------------------------------|
| GET    | `/v1/vendas-plantacoes`                  | Lista todos as vendas de plantações       |
| GET    | `/v1/vendas-plantacoes/id`               | Detalha uma venda de uma plantação        |
| POST   | `/v1/vendas-plantacoes`                  | Cria uma nova venda de plantação          |
| PUT    | `/v1/vendas-plantacoes/id`               | Atualiza uma venda de plantação           |
| DELETE | `/v1/vendas-plantacoes/id`               | Deleta uma venda de uma plantação         |

---

### ✅ Exemplo de Request: `POST /v1/vendas-plantacoes`

```json
{
  "planting_id": 1,
  "value_sale": 150.75
}

```

---

## 💰🌱 Lucro

| Método | Rota                                     | Descrição                                     |
|--------|------------------------------------------|-----------------------------------------------|
| GET    | `/v1/lucro/:id`                          | Calcula o lucro de uma plantação em específico|

### ✅ Exemplo de Response: `GET /v1/lucro/1`

```json
{
    "value_sale_plantiation": 301.5,
    "total_cost": 1530,
    "profit": -1228.5,
    "profit_margen": -12.285
}
```

---

# 📊🌾 Performance de Plantação


| Método | Rota                                         | Descrição                                          |
|--------|----------------------------------------------|----------------------------------------------------|
| GET    | `/v1/performances-das-plantacoes`             | Lista todas as performances de plantações          |
| GET    | `/v1/performances-das-plantacoes/:id`         | Detalha a performance de uma plantação específica  |
| POST   | `/v1/performances-das-plantacoes`             | Cria um registro de performance de plantação       |
| PUT    | `/v1/performances-das-plantacoes/:id`         | Atualiza uma performance de plantação              |
| DELETE | `/v1/performances-das-plantacoes/:id`         | Remove uma performance de plantação                |

---

### ✅ Exemplo de Response: `GET /v1/performances-das-plantacoes/:id`

```json
{
  "planting": {
    "batch_name": "Lote 04",
    "is_planting": false,
    "agriculture_culture_name": "Manga",
    "start_date_planting": "2025-04-23T12:21:53.399681Z"
  },
  "id": 1,
  "production_obtained": 1500.5,
  "production_obtained_formated": "1500.5kg",
  "harvested_area": 2.5,
  "harvested_area_formated": "2.5ha",
  "harvested_date": "2025-04-23T12:21:53.399681Z"
}
```

### ✅ Exemplo de Request: `POST /v1/performances-das-plantacoes/`

```json
{
  "planting_id": 1,
  "production_obtained": 1500.5,
  "unit_production_obtained": "kg",
  "harvested_area": 2.5,
  "unit_harvested_area": "ha",
  "harvested_date": 2.5
}

```

---

### 🚁 Monitoramento via Drones
| Método | Rota                             | Descrição |
|--------|----------------------------------|-----------|
| `POST` | `/v1/drones/monitoramento/ndvi` | Envia imagem para análise de saúde das plantas (NDVI) |
| `POST` | `/v1/drones/monitoramento/pragas`| Envia imagem para detecção de pragas via visão computacional |
| `GET`  | `/v1/drones/monitoramento/irrigacao`| Obtém dados de monitoramento para otimização de irrigação |

