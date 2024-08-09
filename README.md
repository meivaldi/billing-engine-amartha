# Billing Engine Service

The **Billing Engine Service** is a microservice responsible for handling all billing-related operations within the application. It can make loan, apply interest, and scheduling for the loans.

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Installation](#installation)
- [API Endpoints](#api-endpoints)

## Features

- **Make Payment**: Handles customer's payment for the installment.
- **Repayment**: Handles customer's payment for the user that has deliquent.
- **Get Deliquent**: Get all users who has deliquent more than twice.
- **Make Loan**: Handles customer's loan request.
- **Get Outstanding**: Get total outstanding (remaining balance user should pay) for certain user.

## Architecture

The Billing Engine Service is built with a microservices architecture and build with clean architecture. It contains two services, the first one is REST API for billing engine itself, and another one is a cron/scheduler that running every minute to check the users who has deliquent more than twice.

- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **API**: RESTful API
- **Dockerized**: Fully containerized using Docker for easy deployment and scaling.

## Installation

To set up the Billing Engine Service locally:

### Prerequisites

- [Go](https://golang.org/doc/install) 1.18 or higher
- [PostgreSQL](https://www.postgresql.org/download/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Steps

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/meivaldi/billing-engine.git
   cd billing-engine

2. **Install Dependencies**:

   ```bash
   go mod tidy

3. **Run Docker Compose to Setup DB and Migration**

    ```bash
    docker compose up -d

4. **Run the service**

    ```bash
    go run main.go

5. **Run the scheduler (cron) in another terminal**

    ```bash
    cd delivery/cron
    go run main.go


### API Endpoints

Below are some key API endpoints exposed by the Billing Engine Service:

- **POST /loan/make**: To make a new loan.
- **POST /loan/pay**: To pay installment each week.
- **POST /loan/repay**: To repay deliquent.

- **GET /loan/outstanding/:userId**: Get total outstanding of certain user.
- **GET /loan/deliquent-users**: Get all the users who has delinquent more than twice.

You can also using postman collection that already provided in this repo.
- **Billing-Engine.postman_collection.json**
- **billing-engine-local.postman_environment.json**
