# Multiple Sensor System
📌 Overview
This project consists of two services:

Service A

Periodically generates sensor data.

Sends data to Service B via gRPC.

Provides REST endpoints to configure the sending interval.

Service B

Stores sensor data in MySQL.

Provides REST APIs to query, update, and delete data.

Runs a gRPC server to receive data from Service A.

# 🚀 How to Run
1. Clone Repository
- git clone https://github.com/yourusername/multiple-sensor.git
- cd multiple-sensor

2. Start with Docker Compose
- docker compose up --build

3. Services
- MySQL → localhost:3306 (DB: sensors, user: root, password: password)

- Service A → REST API http://localhost:8081

- Service B → REST API http://localhost:8083, gRPC localhost:50051

# 🧪 Testing the APIs
- Import Postman Collection
- File: postman_collection.json (included in the repo)

**Base URLs:**

{{serviceAUrl}} = http://localhost:8081

{{serviceBUrl}} = http://localhost:8083

# Service A Endpoints
```bash
- GET /health → Health check.
- GET /frequency → Get the current sending interval.
- POST /frequency → Update the sending interval.
```

# Service B Endpoints
```bash
GET /health → Health check.
GET /data → Query sensor data.
PUT /data/:id → Update specific sensor data.
DELETE /data/:id → Delete specific sensor data.
```


# Requirements
```bash
Docker & Docker Compose
Go 1.24+ (if you want to run locally without Docker)
```

# ⚙️ Notes
```bash
Service A → Periodically sends data every N seconds to Service B via gRPC.
Service B → Stores sensor data in MySQL and serves it via a REST API.
```
**The database schema is automatically created by mysql/init.sql when the MySQL container starts**
