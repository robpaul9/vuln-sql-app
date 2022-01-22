# vuln-sql-app

!! USE FOR TESTING PURPOSES ONLY !!

A web app with a SQL injection vulnerability and modsec waf proxy for testing and research 

## Prerequisites
1. An AWS param store secret used for the postgres password
2. An AWS service account with a role or key access to read/decrypt param
2. Golang >= `v1.16`

## Installation
1. Clone repo and create `.env` and `.secrets/postgress_db_password` files in root dir.
2. Add AWS postgres secret param to file `.secrets/postgress_db_password`
3. Fill our .env file with the following information:
```bash
SERVICE_NAME=
SERVICE_PORT=
DATABASE_NAME=
DATABASE_PORT=5432
DATABASE_HOST=
DATABASE_USER=
DB_PASSWORD_PARAM
AWS_SDK_LOAD_CONFIG=true
```
4. Auth to awscli using AWS service account credentials or [not recommended] add `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, and `AWS_DEFAULT_REGION` env vars to `.env`. 
5. Run `docker-compose up -d`

## Usage

SQL Injection vulnerable web app URL: `http://localhost:8080/`

Vulnerable app with a WAF in front of it URL: `http://localhost:9085/`