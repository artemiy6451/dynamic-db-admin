# Dynamic DB Admin

**Dynamic DB Admin** — desktop-приложение для управления PostgreSQL базами данных с графическим интерфейсом. Проект разрабатывается на Go и Vue 3 с использованием Wails.

Приложение предназначено для удобного просмотра, анализа и управления структурой баз данных и их содержимым без необходимости постоянно писать SQL-запросы вручную.

> **Status:** Early Development

## Overview

Dynamic DB Admin позволит подключаться к нескольким PostgreSQL серверам, просматривать структуру баз данных, управлять таблицами и работать с записями через динамический графический интерфейс.

Основная идея проекта — предоставить универсальный инструмент администрирования PostgreSQL, в котором интерфейс автоматически адаптируется к структуре выбранной таблицы.

## Features

Планируемые возможности:

### Connection Management
- Управление несколькими PostgreSQL подключениями.
- Проверка доступности сервера.
- Настройка параметров подключения.
- Поддержка SSL/TLS.
- Управление состоянием подключений.

### Database Explorer
- Просмотр доступных баз данных.
- Навигация по схемам.
- Просмотр таблиц и представлений.
- Просмотр структуры объектов базы данных.

### Schema Management
- Создание, удаление и переименование таблиц.
- Добавление, удаление и переименование столбцов.
- Изменение типов данных.
- Управление DEFAULT и NOT NULL.
- Управление первичными и внешними ключами.
- Просмотр и управление индексами и ограничениями.

### Record Management
- Просмотр записей в таблицах.
- Добавление, редактирование и удаление записей.
- Динамическое построение таблиц и форм на основе схемы PostgreSQL.
- Сортировка и фильтрация данных.
- Постраничная загрузка записей.

### Security & Audit
- Безопасное хранение параметров подключений.
- Контроль выполнения операций.
- Подтверждение потенциально опасных изменений.
- Журналирование операций с базами данных.

## Tech Stack

### Backend
- [Go](https://go.dev/) — основной язык backend.
- [Wails](https://wails.io/) — создание desktop-приложения и взаимодействие Go с frontend.
- [pgx](https://github.com/jackc/pgx) — работа с PostgreSQL.
- [SQLite](https://sqlite.org/) — локальное хранилище настроек и метаданных.

### Frontend
- [Vue 3](https://vuejs.org/) — пользовательский интерфейс.
- [TypeScript](https://www.typescriptlang.org/) — типизация.
- [Vite](https://vite.dev/) — сборка frontend.
- [Pinia](https://pinia.vuejs.org/) — управление состоянием.
- [Vue Router](https://router.vuejs.org/) — маршрутизация.
- [Tailwind CSS](https://tailwindcss.com/) — стилизация.
- [TanStack Table](https://tanstack.com/table/latest) — работа с таблицами.

## Architecture

Проект использует архитектуру модульного монолита с принципами Clean Architecture и Hexagonal Architecture.

Основные уровни приложения:

| Layer | Responsibility |
|---|---|
| Domain | Бизнес-сущности и правила |
| Application | Сценарии использования приложения |
| Infrastructure | PostgreSQL, SQLite, файловая система |
| Adapters | Взаимодействие с Wails и frontend |
| Platform | Конфигурация, системные сервисы и общие механизмы |

Бизнес-логика не должна зависеть от Wails или Vue. Это позволит в дальнейшем расширить приложение, например, добавить REST API или серверный режим работы.

## Project Structure

```text
dynamic-db-admin/
├── main.go
├── app.go
├── internal/
│   ├── app/
│   ├── domain/
│   ├── application/
│   ├── infrastructure/
│   │   ├── postgres/
│   │   ├── sqlite/
│   │   ├── security/
│   │   └── logging/
│   ├── adapters/
│   │   ├── wails/
│   │   └── events/
│   └── platform/
├── frontend/
│   └── src/
│       ├── app/
│       ├── layouts/
│       ├── pages/
│       ├── features/
│       ├── components/
│       ├── stores/
│       ├── composables/
│       ├── services/
│       ├── types/
│       └── styles/
├── build/
├── docs/
│   └── sprints/
├── tests/
├── go.mod
├── go.sum
├── wails.json
└── README.md
```

Структура будет постепенно расширяться по мере реализации функциональности.

## Development

### Requirements

Для разработки потребуются:

- Go — актуальная поддерживаемая версия.
- Node.js и npm.
- Wails CLI.
- Системные зависимости Wails для целевой ОС.
- PostgreSQL — для интеграционного тестирования будущих модулей.

### Clone repository

```bash
git clone https://github.com/USERNAME/dynamic-db-admin.git
cd dynamic-db-admin
```

### Install dependencies

```bash
go mod download

cd frontend
npm install
cd ..
```

### Development mode

```bash
wails dev
```

### Build

```bash
wails build
```

Для Windows:
```bash
wails build GOOS=windows
```

Готовая сборка будет находиться в каталоге `build/bin/`.

> Команды и инструкции будут уточняться по мере настройки проекта и закрепления версий инструментов.

## Development Workflow

Разработка ведётся с использованием упрощённой двухветочной модели:

| Branch | Purpose |
|---|---|
| `main` | Стабильные версии приложения |
| `dev` | Основная разработка и интеграция изменений |

Работа организована по спринтам.

- Каждый спринт имеет отдельное описание в `docs/sprints/`.
- Конкретные задачи отслеживаются через GitHub Issues.
- Завершённые и проверенные изменения интегрируются в `main`.
- Стабильные версии отмечаются Git-тегами.

### Commit Convention

Для сообщений коммитов используется Conventional Commits:

```text
feat: add PostgreSQL connection manager
fix: handle connection timeout
refactor: simplify schema introspection
docs: update installation instructions
test: add database integration tests
build: configure Windows packaging
chore: update dependencies
```

## Roadmap

Проект развивается поэтапно:

- [ ] Sprint 01 — Project Foundation
- [ ] Sprint 02 — Local Storage & Configuration
- [ ] Sprint 03 — PostgreSQL Connection Manager
- [ ] Sprint 04 — Database Schema Introspection
- [ ] Sprint 05 — Schema Management
- [ ] Sprint 06 — Record Management
- [ ] Sprint 07 — Frontend & User Experience
- [ ] Sprint 08 — Security & Audit
- [ ] Sprint 09 — Testing & Stabilization
- [ ] Sprint 10 — Release & Distribution

Подробный план разработки находится в каталоге `docs/sprints/`.

## Security

Безопасность является одним из основных требований проекта.

Основные принципы:

- Использование параметризованных SQL-запросов.
- Безопасная обработка SQL identifiers.
- Проверка операций на стороне backend.
- Минимально необходимые права PostgreSQL.
- Защита сохранённых учётных данных.
- Маскирование конфиденциальных данных в логах.
- Подтверждение потенциально опасных операций.

Не используйте учётные записи PostgreSQL с избыточными привилегиями для повседневной работы.

## License

This project is licensed under the Apache License 2.0.

See the [LICENSE](LICENSE) file for details.

## Author

Developed by **Lokach**.

