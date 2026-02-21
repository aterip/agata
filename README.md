# AGATA

**Analytics Gateway And Transform Architecture** — a universal data collection gateway for ingesting data from any source and delivering it to modern analytical databases.

[![Go Report Card](https://goreportcard.com/badge/github.com/aterip/agata)](https://goreportcard.com/report/github.com/aterip/agata)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![CLA assistant](https://cla-assistant.io/readme/badge/aterip/agata)](https://cla-assistant.io/aterip/agata)

## What is AGATA?

AGATA is a lightweight, open-source, and extensible tool for building data pipelines. It addresses one of the most common tasks in modern infrastructure: collecting data from external APIs (pull model) or receiving events from clients (push model) and writing them to databases such as ClickHouse, PostgreSQL, TimescaleDB, and more.

The project is designed to free teams from writing and maintaining countless custom scripts, offering a ready‑to‑use solution with a user‑friendly web interface, built‑in scheduler, transformation support, and enterprise‑grade features (RBAC, audit, SSO, encryption).

## Key Features

- **Pull Model**: Automatically fetch data from external HTTP APIs on a cron schedule. Supports headers, basic authentication, tokens.
- **Push Model**: HTTP endpoint for receiving JSON/XML data with API key authentication.
- **Flexible Writing**: Pluggable adapters for sending data to ClickHouse, PostgreSQL, TimescaleDB, MongoDB, and other databases.
- **Web UI & API Management**: Configure sources, users, and view statistics and logs via a modern interface.
- **Task Scheduler**: Built‑in cron with distributed mode support (using PostgreSQL advisory locks).
- **Monitoring & Observability**: Prometheus metrics, structured logging, request tracing.
- **Enterprise‑Ready**: Role‑based access control (RBAC), audit trails, SSO integration (LDAP, OIDC, SAML), encryption of sensitive data.
- **Python Analytics API** (optional): A microservice for complex analytics — moving averages, forecasting, anomaly detection, CSV/JSON export.

## For Russian
## Что такое AGATA?

AGATA — это легковесный, открытый и расширяемый инструмент для построения пайплайнов данных. Он решает одну из самых частых задач в современной инфраструктуре: сбор данных из внешних API (pull-модель) или приём событий от клиентов (push-модель) и их запись в такие базы, как ClickHouse, PostgreSQL, TimescaleDB и другие.

Проект создан для того, чтобы избавить команды от необходимости писать и поддерживать множество кастомных скриптов, и предлагает готовое решение с удобным веб-интерфейсом, планировщиком, поддержкой трансформаций и enterprise-функциями (RBAC, аудит, SSO, шифрование).

## Основные возможности

- **Pull-модель**: автоматический сбор данных из внешних HTTP API по расписанию (cron). Поддержка заголовков, базовой авторизации, токенов.
- **Push-модель**: HTTP-эндпоинт для приёма данных в форматах JSON/XML с аутентификацией по API-ключу.
- **Гибкая запись**: подключаемые адаптеры для отправки данных в ClickHouse, PostgreSQL, TimescaleDB, MongoDB и другие СУБД.
- **Управление через веб-интерфейс и API**: настраивайте источники, пользователей, просматривайте статистику и логи.
- **Планировщик задач**: встроенный cron с поддержкой распределённого режима (через блокировки PostgreSQL).
- **Мониторинг и наблюдаемость**: метрики Prometheus, структурированное логирование, трейсинг запросов.
- **Enterprise‑готовность**: ролевая модель доступа (RBAC), аудит действий, интеграция с SSO (LDAP, OIDC, SAML), шифрование чувствительных данных.
- **Python Analytics API** (опционально): микросервис для сложных аналитических задач — скользящие средние, прогнозирование, обнаружение аномалий, экспорт в CSV/JSON.
