# Sensor Data Processing System

## Overview

**Sensor Data Processing System** is a backend system for processing sensor data using a microservice architecture. Built as a coding assignment for Jr. Backend Developer, the system is designed with scalability, clean architecture, and best practices in mind.

The system consists of two microservices:

- **Microservice A**: Generates simulated sensor data streams and sends them to Microservice B via gRPC. Multiple instances can run, each tied to a specific sensor type.
- **Microservice B**: Receives sensor data, stores it in MySQL, and exposes REST APIs for retrieval, editing, and deletion. Supports filtering, pagination, authentication, and authorization.

Technologies used: **Go (Echo Framework), MySQL, GORM, gRPC, Docker, Swagger, Postman**.

---