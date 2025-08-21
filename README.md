# Multiple Sensor System

## 📌 Overview
This project consists of two services:

- **Service A**
  - Periodically generates sensor data.
  - Sends data to Service B via gRPC.
  - Provides REST endpoints to configure the sending interval.

- **Service B**
  - Stores sensor data in MySQL.
  - Provides REST APIs to query, update, and delete data.
  - Runs a gRPC server to receive data from Service A.

---

## 🚀 How to Run
1. Clone Repository
git clone [https://github.com/yourusername/multiple-sensor.git](https://github.com/yourusername/multiple-sensor.git)
cd multiple-sensor

2. Start with Docker Compose
docker compose up --build

3. Services
MySQL → localhost:3306 (DB: sensors, user: root, password: password)

Service A → REST API http://localhost:8081

Service B → REST API http://localhost:8083, gRPC localhost:50051

🧪 Testing the APIs
Import Postman Collection
File: postman_collection.json (included in the repo)

Base URLs:
{{serviceAUrl}} = http://localhost:8081

{{serviceBUrl}} = http://localhost:8083

Service A Endpoints
GET /health → Health check.
GET /frequency → Get the current sending interval.
POST /frequency → Update the sending interval.

Service B Endpoints
GET /health → Health check.
GET /data → Query sensor data.
PUT /data/:id → Query sensor data.
DELETE /data/:id → Query sensor data.

✅ Example Requests
Query Data (Service B)

Bash

curl "http://localhost:8083/data?id1=A&id2=1&limit=5&offset=0"
Set Frequency (Service A)

Bash

curl -X POST "http://localhost:8081/frequency" \
     -H "Content-Type: application/json" \
     -d '{"seconds": 5}'
📂 Requirements
Docker & Docker Compose

Go 1.24+ (if you want to run locally without Docker)

⚙️ Notes
Service A → Periodically sends data every N seconds to Service B via gRPC.

Service B → Stores sensor data in MySQL and serves it via a REST API.

The database schema is automatically created by mysql/init.sql when the MySQL container starts.