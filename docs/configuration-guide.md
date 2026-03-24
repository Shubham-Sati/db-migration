# Configuration Guide

## Overview
This document explains the configuration setup for the chat-analytics database migration tool.

## Configuration Structure

### App Configuration (`configs/app.go`)
- **Purpose**: General application settings that control migration tool behavior
- **Key Settings**:
  - `Env`: Environment mode (development/staging/production)
  - `Port`: Port for test server when running migration tool as service
  - `LogLevel`: Controls verbosity of logging output
  - `Debug`: Enables/disables debug mode for troubleshooting

### Database Configuration (`configs/db.go`)
- **Purpose**: Database connection and performance tuning settings
- **Key Settings**:
  - `Host/Port/Username/Password/Database`: Basic PostgreSQL connection parameters
  - `SSLMode`: Security mode for database connections
    - `disable`: No encryption (development)
    - `require`: SSL required (production)
  - `MaxOpenConns`: Maximum simultaneous database connections (default: 100)
  - `MaxIdleConns`: Number of warm connections to maintain (default: 10)

### Configuration Loading (`configs/configs.go`)
- **Purpose**: Centralized configuration initialization
- **Process**:
  1. Locates `.env` file relative to application root
  2. Loads environment variables using `godotenv`
  3. Maps environment variables to Go structs using `envconfig`
  4. Provides logging for debugging configuration issues

## Environment Variables
```bash
# Database settings
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=your_password
DB_NAME=chatapp
DB_SSLMODE=disable
DB_MAX_OPEN_CONNS=100
DB_MAX_IDLE_CONNS=10

# Application settings
ENVIRONMENT=development
PORT=8080
LOG_LEVEL=info
DEBUG=false
```

## Load Testing Impact

### **MaxOpenConns** during load tests:
```
Scenario: 1000 concurrent users hitting your API
- Without proper MaxOpenConns: Database crashes or refuses connections
- With MaxOpenConns=100: Only 100 users get DB access, others wait
- This prevents database overload but creates bottlenecks
```

### **MaxIdleConns** for sustained load:
```
Scenario: Continuous traffic for hours
- Without idle connections: Every request pays connection overhead
- With MaxIdleConns=10: Fast response times for sustained traffic
- Critical for maintaining low latency under load
```

## Real Load Testing Scenarios

### **Scenario 1: Chat Message Burst**
```
Event: 500 users send messages simultaneously
Problem: Each message = 1 DB connection needed
Solution: MaxOpenConns=100 handles burst, queues the rest
```

### **Scenario 2: Analytics Processing**
```
Event: Processing 10,000 events per minute
Problem: Each analytics write needs DB connection
Solution: Connection pooling prevents connection exhaustion
```

## Load Testing Strategy

When you use **k6** or similar tools later, you'll test:

1. **Connection Pool Exhaustion**:
   ```bash
   k6 run --vus 200 --duration 5m  # 200 virtual users for 5 minutes
   ```

2. **Response Time Degradation**:
   - Monitor how response times increase as connections max out
   - Find optimal MaxOpenConns vs response time balance

3. **Database Resource Usage**:
   - Track PostgreSQL connection count: `SELECT count(*) FROM pg_stat_activity;`
   - Monitor memory usage vs connection pool size

## Production Tuning Based on Load Tests

Your load tests will reveal:
- **Sweet spot** for MaxOpenConns (performance vs resources)
- **Connection timeout** settings needed
- **Database server limits** you'll hit

This is why these seemingly small config values become **make-or-break** for scalability!