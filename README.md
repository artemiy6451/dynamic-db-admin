# Dynamic DB Admin

**Dynamic DB Admin** — desktop-приложение для визуального проектирования, изменения структуры и управления содержимым баз данных SQLite.

Проект создаётся на Go и Wails v3 с frontend на Vue 3 + TypeScript. Приложение ориентировано на работу с локальными файлами SQLite: пользователь открывает базу данных и управляет её схемой и записями через графический интерфейс.

> **Status:** Early Development. Описанные ниже возможности являются целями проекта, если явно не отмечены как реализованные.

## Product scope

Dynamic DB Admin работает с файлами SQLite (`.db`, `.sqlite`, `.sqlite3`). Это не клиент для PostgreSQL/MySQL и не менеджер удалённых серверных подключений.

Основные сценарии:

- Открытие, создание и закрытие нескольких SQLite-файлов.
- Исследование структуры базы: таблицы, представления, индексы, триггеры, столбцы и ограничения.
- Изменение схемы базы данных через визуальный редактор.
- Просмотр и редактирование записей в табличном редакторе и автоматически генерируемых формах.
- Просмотр связей между таблицами в ER-диаграмме.
- Безопасное применение структурных изменений с резервным копированием и проверкой целостности.

### Planned features

**Database management**
- Создание новой SQLite-базы и открытие существующего файла.
- Несколько открытых баз и список недавно открытых файлов.
- Информация о файле, версии SQLite и состоянии базы.

**Database explorer**
- Дерево таблиц, views, indexes и triggers.
- Информация о столбцах, типах, первичных и внешних ключах, ограничениях.
- Повторное чтение структуры после изменений.

**Schema designer**
- Создание, переименование и удаление таблиц и столбцов.
- Настройка типов, PRIMARY KEY, FOREIGN KEY, UNIQUE, NOT NULL, DEFAULT и CHECK.
- Создание индексов.
- Безопасные миграции для операций, которые SQLite не поддерживает напрямую.

**Relationship designer**
- ER-диаграмма на основе реальных внешних ключей.
- Визуальное расположение таблиц и отображение связей.
- Сохранение координат и настроек диаграммы во внутреннем хранилище приложения.

**Data editor**
- Табличный просмотр записей и динамические формы.
- Добавление, редактирование и удаление записей.
- Сортировка, фильтрация, поиск и пагинация.
- Выбор связанных записей с учётом внешних ключей.

**Safety**
- Подтверждение разрушительных действий.
- Резервная копия перед рискованными изменениями схемы.
- Транзакции там, где они применимы.
- Проверка внешних ключей и целостности базы.
- Аккуратная работа с таблицами без первичного ключа и составными ключами.

SQL-редактор, расширенные инструменты импорта/экспорта и дополнительные средства работы с views/triggers могут быть добавлены после первой версии.

## Technology stack

### Backend
- [Go](https://go.dev/) — backend и предметная логика.
- [Wails v3](https://v3alpha.wails.io/) — desktop runtime и Go bindings.
- `database/sql` — стандартный интерфейс доступа к данным.
- SQLite driver — будет выбран и зафиксирован после проверки совместимости с целевыми ОС и необходимыми SQLite API.
- `log/slog` — структурированное логирование.

### Frontend
- [Vue 3](https://vuejs.org/) — Composition API и `<script setup>`.
- [TypeScript](https://www.typescriptlang.org/).
- [Vite](https://vite.dev/).
- [Pinia](https://pinia.vuejs.org/).
- [Vue Router](https://router.vuejs.org/).
- [Tailwind CSS v4](https://tailwindcss.com/).
- [Lucide](https://lucide.dev/) — иконки.

## Architecture

Приложение строится как модульный монолит с разделением на domain, application, infrastructure и adapters.

- **Domain** содержит предметные модели и правила, не зависит от Wails и SQLite-драйвера.
- **Application** реализует сценарии использования и координирует операции.
- **Infrastructure** реализует работу с SQLite-файлами, миграции, внутреннее хранилище, конфигурацию и логирование.
- **Wails adapters** предоставляют типизированный API frontend.
- **Frontend** организован по функциональным областям (features).

В приложении существуют два разных вида SQLite-хранилищ:

1. **Пользовательские базы** — выбранные пользователем файлы, содержимое и схема которых редактируются.
2. **Внутреннее хранилище приложения** — настройки, список недавних файлов, состояние интерфейса, расположение объектов ER-диаграмм и другие метаданные.

Эти хранилища не должны смешиваться. Закрытие приложения не должно приводить к изменению или перемещению пользовательских файлов.

Frontend передаёт backend типизированные описания операций, а не произвольные SQL-фрагменты. Backend проверяет входные данные, безопасно формирует SQL и использует параметры для значений.

Подробности и правила зависимостей описаны в [docs/architecture.md](docs/architecture.md).

## Project structure

Структура будет расширяться по мере появления реальных модулей. Целевое направление:

```text
dynamic-db-admin/
├── main.go
├── internal/
│   ├── app/                  # composition root и запуск Wails
│   ├── domain/               # предметные модели и правила
│   ├── application/          # use cases
│   ├── adapters/
│   │   └── wails/            # frontend-facing services
│   └── infrastructure/
│       ├── sqlite/            # пользовательские SQLite-файлы
│       ├── storage/            # внутреннее хранилище приложения
│       ├── config/
│       └── logging/
├── frontend/
│   ├── bindings/              # generated Wails bindings
│   └── src/
│       ├── app/
│       ├── components/
│       ├── layouts/
│       ├── pages/
│       ├── features/
│       ├── services/
│       ├── stores/
│       └── styles/
├── build/
├── docs/
└── go.mod
```

Не нужно создавать пустые пакеты заранее. Конкретные директории добавляются вместе с реализуемой функциональностью.

## Development

### Requirements

- Go version compatible with the version declared in `go.mod`.
- Node.js и npm.
- Wails v3 CLI.
- Системные зависимости Wails для целевой ОС (на Linux — GTK4 и WebKitGTK 6.0).
- Для проверки Windows-сборки — Windows environment или поддерживаемый Windows CI runner.

### Install dependencies

```bash
go mod download
cd frontend
npm install
cd ..
```

### Development mode

```bash
wails3 dev
```

### Generate Wails bindings

После изменения экспортируемых методов Wails-сервисов:

```bash
wails3 generate bindings
```

Generated файлы в `frontend/bindings/` не редактируются вручную.

### Build

```bash
wails3 build
```

Точные параметры сборки зависят от целевой платформы и конфигурации Wails v3 в репозитории.

## Git workflow

Используется простая двухветочная модель:

| Branch | Назначение |
|---|---|
| `dev` | Основная разработка, спринты и интеграция изменений |
| `main` | Проверенные стабильные версии |

Обычная работа ведётся в `dev`. Перенос в `main` выполняется только после проверки стабильной версии. Не выполнять force-push и не переписывать историю без отдельного решения.

### Commit convention

Используются Conventional Commits:

```text
feat(databases): add SQLite file opening
fix(sqlite): preserve foreign keys during migration
refactor(frontend): extract database explorer feature
test(migration): cover composite primary keys
docs: update architecture overview
build: configure Linux packaging
chore: update dependencies
```

Один коммит должен описывать одно логически завершённое изменение.

## Roadmap

1. Project Foundation
2. SQLite Database Management
3. Database Explorer
4. Schema Designer and Migration Engine
5. Relationship Designer
6. Data Editor
7. Advanced Data Management
8. Safety and Data Integrity
9. Testing and Stabilization
10. Release and Distribution

Подробный план находится в [docs/roadmap.md](docs/roadmap.md). Архитектурные решения описаны в [docs/architecture.md](docs/architecture.md). Рабочие заметки: [Sprint 1](docs/sprints/sprint-01.md) и [Sprint 2](docs/sprints/sprint-02.md).

## License

This project is licensed under the Apache License 2.0. See [LICENSE](LICENSE) for details.
