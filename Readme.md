
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

- **Docker** instalado
- **MQTT Explorer** instalado

### Passos:

Claro! Aqui está o README completo, sem a parte sobre as imagens:

---

# Projeto Agrohero - Ambiente Completo

Este documento explica como configurar e rodar todo o ambiente local do projeto Agrohero, que inclui:

* API principal Agrohero
* API de autenticação separada
* Banco de dados PostgreSQL + PGAdmin
* MinIO (armazenamento)
* Kafka + Zookeeper + Kafdrop (UI Kafka)
* Mosquitto MQTT Broker
* Como instalar o MQTT Explorer (AppImage) para testes

---

## Pré-requisitos

* Docker & Docker Compose
* Git
* Linux com suporte a FUSE (para rodar AppImage)

---

## Passos para rodar o projeto completo

### 1. Criar a pasta do projeto e clonar os repositórios

No terminal, escolha a pasta onde quer trabalhar e faça:

```bash
mkdir agrohero-full
cd agrohero-full

git clone https://github.com/ericsanto/apiAgroPlusUltraV1.git
git clone https://github.com/ericsanto/api_authentication.git
git clone https://github.com/ericsanto/kafka_minio_python.git
```

---

### 2. Criar os arquivos `.env` em cada pasta clonada

Configure as variáveis de ambiente para cada API:

* `apiAgroPlusUltraV1/.env`
* `api_authentication/.env`

Exemplo apiAgroPlusUltraV1:

```
JWT_SECRET_KEY= 
ACCESS_KEY_ID_MINIO= 
SECRET_KEY_MINIO=
BUCKET_NAME= 
ENDPOINT=minio:9000 
OPEN_WEATHER_API_KEY=
DEEPSEEK_API_KEY=
URL_BROKER_MOSQUITTO=ws://mosquitto:9005
PASSWORD_BROKER_MOSQUITTO=
USERNAME_BROKER_MOSQUITTO=
```

Exemplo api_authentication:

```
JWT_SECRET_KEY=
```

Exemplo python_minio_kafka:

```
MINIO_ACCESS_KEY=
MINIO_SECRET_KEY=
```

---

### 3. Salvar o arquivo `docker-compose.yml` no diretório raiz `agrohero-full`

Crie um arquivo `docker-compose.yml` com o conteúdo:

```yaml
services:
  app1:
    build:
      context: ./apiAgroPlusUltraV1
      dockerfile: Dockerfile
    ports: 
      - "8080:8080"
    depends_on:
      - db
    networks:
      - app-network
    command: air
    env_file:
      - ./apiAgroPlusUltraV1/.env

  app2:
    build:
      context: ./api_authentication
    ports:
      - "8081:8080"
    depends_on:
      - db
    networks:
      - app-network
    env_file:
      - ./api_authentication/.env

  db:
    image: postgres:latest
    container_name: postgres-db
    environment:
      POSTGRES_USER: go 
      POSTGRES_PASSWORD: go
      POSTGRES_DB: go
    volumes:
      - pgdata:/var/lib/postgresql/data
    ports:
      - "5432:5432"
    networks:
      - app-network

  pgadmin:
    image: dpage/pgadmin4
    container_name: pgadmin
    environment:
      PGADMIN_DEFAULT_EMAIL: admin@admin.com
      PGADMIN_DEFAULT_PASSWORD: admin123
    ports:
      - "5050:80"
    depends_on:
      - db
    networks:
      - app-network

  minio:
    image: minio/minio
    container_name: minio
    ports:
      - "9000:9000"  # Porta da API
      - "9001:9001"  # Porta do painel administrativo
    environment:
      MINIO_ROOT_USER: minioadmin
      MINIO_ROOT_PASSWORD: minioadmin123
    command: server /data --console-address ":9001"
    volumes:
      - minio_data:/data
    networks:
      - app-network

  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    networks:
      - app-network
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000

  kafka:
    image: confluentinc/cp-kafka:7.2.1
    networks:
      - app-network
    depends_on:
      - zookeeper
    ports:
      - 9092:9092
      - 29092:29092
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka:9092,PLAINTEXT_HOST://localhost:29092
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1

  kafdrop:
    image: obsidiandynamics/kafdrop:latest
    networks:
      - app-network
    depends_on:
      - kafka
    ports:
      - 19000:9000
    environment:
      KAFKA_BROKERCONNECT: kafka:9092

  mosquitto:
    image: eclipse-mosquitto:latest
    ports:
      - "8883:8883"   # TLS MQTT
      - "9005:9005"   # TLS WebSocket
    volumes:
      - ./mosquitto/config/mosquitto.conf:/mosquitto/config/mosquitto.conf
      - ./mosquitto/data:/mosquitto/data
      - ./mosquitto/log:/mosquitto/log
      - ./mosquitto/config/pwfile:/mosquitto/config/pwfile
      - ./mosquitto/certs:/mosquitto/certs
    networks:
      - app-network


volumes:
  pgdata:
  minio_data:

networks:
  app-network:
    driver: bridge
```

---

### 4. Subir todo o ambiente

No terminal:

```bash
docker-compose up -d --build
```

---

### 5. Acessar serviços importantes

* API Agrohero: [http://localhost:8080](http://localhost:8080)
* API Authentication: [http://localhost:8081](http://localhost:8081)
* PGAdmin: [http://localhost:5050](http://localhost:5050)
* MinIO Console: [http://localhost:9001](http://localhost:9001)
* Kafka UI (Kafdrop): [http://localhost:19000](http://localhost:19000)
* MQTT Broker (Mosquitto): Porta TLS MQTT `8883` e TLS WebSocket `9005`

---

### 6. Executar servico Python

```bash

cd kafka_minio_python

python -m venv venv

source venv/bin/activate

pip install -r requirements.txt

python3 main.py
```

## Como usar o MQTT Explorer no Linux

### 1. Baixar o AppImage

Acesse a página oficial ou baixe direto pelo terminal:

```bash
wget https://github.com/thomasnordquist/MQTT-Explorer/releases/latest/download/MQTT-Explorer.AppImage
```

---

### 2. Tornar o arquivo executável

```bash
chmod +x MQTT-Explorer.AppImage
```

---

### 3. Instalar dependências necessárias (FUSE)

O AppImage precisa do FUSE para rodar. Instale conforme sua distro:

* **Debian/Ubuntu**

  ```bash
  sudo apt update
  sudo apt install fuse libfuse2
  ```

* **Arch Linux/Manjaro**

  ```bash
  sudo pacman -S fuse2
  ```

* **Fedora**

  ```bash
  sudo dnf install fuse
  ```

> ⚠️ Se ainda não executar, reinicie o sistema.

---

### 4. Executar o MQTT Explorer

```bash
./MQTT-Explorer.AppImage
```

> Se der erro, tente:

```bash
./MQTT-Explorer.AppImage --no-sandbox
```

---

## Dicas para rodar e debugar

* Para ver logs do Agrohero API:

  ```bash
  docker-compose logs -f app1
  ```

* Parar o ambiente:

  ```bash
  docker-compose down
  ```

---


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

## 🚜 Fazenda

| Método | Rota                     | Descrição                             |
| ------ | ------------------------ | --------------------------------------|
| GET    | `/v1/fazenda/`           | Lista todas as fazendas de um usuario |
| GET    | `/v1/fazenda/{id}`       | Detalha uma fazenda de do usuario     |
| POST   | `/v1/fazenda`            | Cria uma nova fazenda                 |
| PUT    | `/v1/fazenda/{id}`       | Atualiza uma fazenda                  |
| DELETE | `/v1/fazenda/{id}`       | Deleta uma fazenda                    |

### ✅ Exemplo de Request: `POST /v1/fazenda`

```json
{
  "name": "Fazenda Santiago",
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

| Método | Rota                | Descrição                                                              | Status esperado |
|--------|---------------------|------------------------------------------------------------------------|-----------------|
| POST   | `/v1/fazenda/:farmID/lote`             | Cria um novo lote agrícola                          | `201 Created`   |
| GET    | `/v1/fazenda/:farmID/lote`             | Lista todos os lotes agrícolas                      | `200 OK`        |
| GET    | `/v1/fazenda/:farmID/lote/:batchID`    | Busca um lote agrícola pelo ID                      | `200 OK`        |
| PUT    | `/v1/fazenda/:farmID/lote/:batchID`    | Atualiza os dados de um lote agrícola pelo ID       | `200 OK`        |
| DELETE | `/v1/fazenda/:farmID/lote/:batchID`    | Deleta um lote agrícola pelo ID                     | `204 No Content`|

---

### 📤 Exemplo de Request (POST / PUT  /v1/fazenda/1/lote)

```json
{
  "name": "Lote Norte",
  "area": 12.5,
  "unit": "hectare"
}
```
---

### 📥 Exemplo de Response (GET /v1/fazenda/1/lote/1)

```json
{
  "id": 1,
  "name": "Lote Norte",
  "area": 12.5,
  "unit": "hectare"
}
```
---

### 📥 Exemplo de Response (GET /v1/fazenda/1/lote)

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
| GET    | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/custos`                             | Lista todos os custos de produção         |
| GET    | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/custos/:costID`              | Detalha um custo de produção específico   |
| POST   | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/custos`                      | Cria um novo custo de produção            |
| PUT    | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/custos/:costID`              | Atualiza um custo de produção             |
| DELETE | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/custos/costID`               | Deleta um custo de produção               |

---

### ✅ Exemplo de Request: `POST /v1/fazenda/1/lote/2/plantacoes/1/custos`

```json
{
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
| GET    | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/vendas`                  | Lista todos as vendas de plantações       |
| GET    | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/vendas/:salePlantingID`               | Detalha uma venda de uma plantação        |
| POST   | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/vendas`                  | Cria uma nova venda de plantação          |
| PUT    | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/vendas/:salePlantingID`               | Atualiza uma venda de plantação           |
| DELETE | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/vendas/:salePlantingID`               | Deleta uma venda de uma plantação         |

---

### ✅ Exemplo de Request: `POST /v1/fazenda/1/lote/1/plantacoes/3/vendas`

```json
{
  "value_sale": 150.75
}

```

---

## 💰🌱 Lucro

| Método | Rota                                     | Descrição                                     |
|--------|------------------------------------------|-----------------------------------------------|
| GET    | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/lucros`                          | Calcula o lucro de uma plantação em específico|

### ✅ Exemplo de Response: `GET /v1/fazenda/1/lote/1/plantacoes/1/lucros`

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
| GET    | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/performances`             | Lista todas as performances de plantações          |
| GET    | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/performances/:performanceID`         | Detalha a performance de uma plantação específica  |
| POST   | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/performances`             | Cria um registro de performance de plantação       |
| PUT    | `/v1/fazendad/:farmID/lote/:batchID/plantacoes/:plantingID/performances/:performanceID`         | Atualiza uma performance de plantação              |
| DELETE | `/v1/fazenda/:farmID/lote/:batchID/plantacoes/:plantingID/performances/:performanceID`         | Remove uma performance de plantação                |

---

### ✅ Exemplo de Response: `GET /v1/fazenda/1/lote/1/plantacoes/2/performances/2`

```json
{
  "planting": {
    "batch_name": "Lote 04",
    "is_planting": false,
    "agriculture_culture_name": "Manga",
    "start_date_planting": "2025-04-23T12:21:53.399681Z"
  },
  "production_obtained": 1500.5,
  "production_obtained_formated": "1500.5kg",
  "harvested_area": 2.5,
  "harvested_area_formated": "2.5ha",
  "harvested_date": "2025-04-23T12:21:53.399681Z"
}
```

### ✅ Exemplo de Request: `POST /v1/fazenda/1/lote/1/plantacoes/1/performances`

```json
{
  "production_obtained": 1500.5,
  "unit_production_obtained": "kg",
  "harvested_area": 2.5,
  "unit_harvested_area": "ha",
  "harvested_date": 2.5
}

```

---

# 📊🌾 Identificacao de Pragas


| Método | Rota                                         | Descrição                                          |
|--------|----------------------------------------------|----------------------------------------------------|
| POST   | `/v1/pragas/reconhecimentos`                | Envia imagem para identificar qual a praga         |

### ✅ Exemplo de Request: `POST /v1/pragas/reconhecimentos`


```bash
curl -X POST /v1/performances-das-plantacoes/upload-imagem/ \
  -H "Content-Type: multipart/form-data" \
  -H "Authorization: Bearer SEU_TOKEN_AQUI" \
  -F "file=@/caminho/para/sua/imagem.jpg"
```

### ✅ Exemplo de Response: `/v1/pragas/reconhecimentos`

```json
{
    "detections": [
        {
            "pest": "Cicadellidae",
            "confidence": 0.8987361,
            "hit_percentage": 90,
            "hit_percentage_formated": "90%"
        }
    ]
}
```

# 📊🌾 Identificacao de Doencas


| Método | Rota                                         | Descrição                                          |
|--------|----------------------------------------------|----------------------------------------------------|
| POST   | `/v1/doencas/reconhecimentos`                | Envia imagem para identificar qual a doenca         |

### ✅ Exemplo de Request: `POST /v1/doencas/reconhecimentos`


```bash
curl -X POST /v1/doencas/reconhecimentos \
  -H "Content-Type: multipart/form-data" \
  -H "Authorization: Bearer SEU_TOKEN_AQUI" \
  -F "file=@/caminho/para/sua/imagem.jpg"
```

### ✅ Exemplo de Response: `/v1/doencas/reconhecimentos`

```json
{
    "disease": "Strawberry Leaf scorch"
}
```

---

# ☀️🌧️ Dados Climaticos

| Método | Rota              | Descrição                                                     |
|--------|-------------------|---------------------------------------------------------------|
| GET    | v1/weather-current?lat=?&long=?| Retorna as condições climáticas atuais da cidade |

### ✅ Exemplo de Request: `GET v1/weather-current?lat=-38.379&long=-89.2343`

### ✅ Exemplo de Response: `GET /weather-current`

```json
{
  "main": {
    "temp": 28.27,
    "temp_max": 28.27,
    "temperature_min": 28.27,
    "feels_like": 28.66,
    "pressure": 1016,
    "humidity": 49
  },
  "rain": {
    "1h": 0,
    "3h": 0
  },
  "wind": {
    "deg": 137,
    "speed": 3.79
  },
  "city": "Nome da cidade"
}
```

---

# 💧 Recomendacao de Irrigacao por IA Baseado nos Dados Climaticos 
| Método | Rota                             |                               Descrição                               |
|--------|----------------------------------|-----------------------------------------------------------------------|
| `GET`  | `/v1/fazenda/:farmID/irrigacao?lat=?&long=?`| calcula a quantidade correta de irrigacao que a planta necessita |

### ✅ Exemplo de Response: `GET /v1/fazenda/1/irrigacao?lat=-10.685&long=-38.2885`

```json
{
  "lote": "Lote 01",
  "estagio_fenologico_atual": "germinacao/emergencia",
  "decisao": false,
  "motivo": "Umidade do solo (90%) adequada para o estagio atual (80-90%)",
  "etc": 1.23,
  "lamina_de_irrigacao_em_mm": 0,
  "volume_por_planta_em_litros": 0
}
```

---

### 🚁 Monitoramento via Drones
| Método | Rota                             | Descrição |
|--------|----------------------------------|-----------|
| `POST` | `/v1/drones/monitoramento/ndvi` | Envia imagem para análise de saúde das plantas (NDVI) |
| `POST` | `/v1/drones/monitoramento/pragas`| Envia imagem para detecção de pragas via visão computacional |
| `GET`  | `/v1/drones/monitoramento/irrigacao`| Obtém dados de monitoramento para otimização de irrigação |


# Provavel Arquitetura do Sistema
![alt text](<./docs/images/ARQUITETURA DO SISTEMA IA.excalidraw(2).png>)