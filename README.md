# Database Migration Tool

A CLI tool to manage database migrations, seeding, and schema management for chat and analytics services.

## Prerequisites

- Go 1.23+
- PostgreSQL database
- Environment configuration (`.env` file)

## Setup

1. **Clone and navigate to the repository:**
   ```bash
   cd db-migration
   ```

2. **Copy environment configuration:**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

3. **Build the binary:**
   ```bash
   go build -o db-migration .
   ```

## Available Commands

### 1. **`droptables`** - Drop All Database Tables
Removes all existing tables from the database.

```bash
# Using Go run
go run main.go droptables

# Using compiled binary
./db-migration droptables
```

**⚠️ Warning:** This will permanently delete all data in the database.

### 2. **`migrate`** - Run Database Migrations
Creates all database tables with their schemas.

```bash
# Using Go run
go run main.go migrate

# Using compiled binary
./db-migration migrate
```

**Tables Created:**
- `users` - User accounts and profiles
- `chat_rooms` - Chat room definitions
- `chat_room_members` - Room membership tracking
- `chat_messages` - Chat messages
- `chat_message_reactions` - Message reactions
- `chat_message_attachments` - File attachments
- `analytics_events` - User activity tracking
- `analytics_user_sessions` - Session tracking
- `analytics_daily_metrics` - Daily statistics
- `analytics_room_metrics` - Room-specific metrics
- `analytics_user_metrics` - User-specific metrics

### 3. **`seed`** - Populate Database with Seed Data
Inserts test/sample data into the database tables.

```bash
# Using Go run
go run main.go seed

# Using compiled binary
./db-migration seed
```

**Note:** Seeding will fail if data already exists due to unique constraints.

### 4. **`alter`** - Alter Existing Database Tables
Modifies existing table structures (add/remove columns, indexes, etc.).

```bash
# Using Go run
go run main.go alter

# Using compiled binary
./db-migration alter
```

### 5. **`server`** - Run Test HTTP Server
Starts a test HTTP server for database connectivity testing.

```bash
# Using Go run
go run main.go server

# Using compiled binary
./db-migration server
```

### 6. **`help`** - Get Help Information
Display help information for any command.

```bash
# General help
go run main.go help

# Command-specific help
go run main.go help migrate
go run main.go help seed
```

## Common Workflows

### Fresh Database Setup
For a completely fresh database setup:

```bash
# 1. Drop all existing tables (if any)
go run main.go droptables

# 2. Create all tables
go run main.go migrate

# 3. Insert seed data
go run main.go seed
```

### Reset Database with Fresh Data
To reset the database and reload seed data:

```bash
# 1. Drop all tables
go run main.go droptables

# 2. Recreate tables
go run main.go migrate

# 3. Load fresh seed data
go run main.go seed
```

### Update Existing Schema
To modify existing tables without losing data:

```bash
# Run schema alterations
go run main.go alter
```

## Environment Configuration

The tool reads configuration from a `.env` file. Key variables:

```bash
# Database Connection
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_SSLMODE=disable

# Connection Pool Settings
DB_MAX_OPEN_CONNS=100
DB_MAX_IDLE_CONNS=10

# Application Settings
ENVIRONMENT=local
PORT=8080
LOG_LEVEL=info
```

## Troubleshooting

### Connection Issues
- Verify PostgreSQL is running
- Check database credentials in `.env`
- Ensure database exists (or use an existing one like `postgres`)

### Seeding Failures
- Seeding fails if data already exists due to unique constraints
- Run `droptables` first to clean the database
- Check for conflicting data in the target database

### Build Issues
- Ensure Go 1.23+ is installed
- Run `go mod tidy` to resolve dependencies
- Check for syntax errors in Go files

## Development

### Adding New Tables
1. Create table definition in `tables/` directory
2. Add migration logic in `database/migrator.go`
3. Add seed data in `seeds/` directory
4. Update this README with new table information

### Modifying Existing Tables
1. Add alteration logic in `database/alter.go`
2. Test with `go run main.go alter`
3. Update seed data if necessary

## Project Structure

```
db-migration/
├── commands/          # CLI command implementations
├── configs/           # Configuration management
├── database/          # Database connection and operations
├── tables/            # Table schema definitions
├── seeds/             # Seed data generators
├── .env              # Environment configuration
├── .env.example      # Environment template
├── main.go           # CLI entry point
└── README.md         # This file
```