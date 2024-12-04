# Go Application Template

This repository is a **template** for building Go applications with Fiber. It includes:

- Configuration loading via environment variables.
- Database connection with migration support.
- Basic Auth middleware.
- Example handler and routing setup.

---

## Getting Started

1. **Clone the repository**:
   ```bash
   git clone https://github.com/zyeloni/golang-fiber-api-template.git
   cd golang-fiber-api-template
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**:
   Export required environment variables. Example:

   ```
   DB_URL=localhost
   DB_DRIVER_NAME=postgres
   DB_MIGRATION_PATH=file://database/migrations
   API_USER=admin
   API_PASSWORD=password
   ```

4. **Run the application**:
   ```bash
   go run main.go
   ```

---

## Configuration

The application uses environment variables for configuration. The following variables are required:

- `DB_URL`: The URL of the database (default: `localhost`).
- `DB_DRIVER_NAME`: The name of the database driver (default: `postgres`).
- `DB_MIGRATION_PATH`: The path to the database migration files (default: `file://database/migrations`).
- `API_USER`: The username for the API (default: `admin`).
- `API_PASSWORD`: The password for the API (default: `password`).

## Database Migrations

This template uses `golang-migrate` for database migrations.

### Creating a Migration

1. **Install the `golang-migrate` CLI**:
   ```bash
   brew install golang-migrate
   ```

2. **Create a new migration file**:
   ```bash
   migrate create -ext sql -dir database/migrations -seq create_example_table
   ```

3. **Edit the generated migration files** in the `migrations/` directory:
    - `database/migrations/000001_create_example_table.up.sql`:
      ```sql
      CREATE TABLE example (
          id SERIAL PRIMARY KEY,
          name VARCHAR(255) NOT NULL
      );
      ```
    - `migrations/000001_create_example_table.down.sql`:
      ```sql
      DROP TABLE example;
      ```

4. **Apply the migrations**:
   ```bash
   migrate -database $DATABASE_URL -path migrations up
   ```

5. **Rollback the last migration** (optional):
   ```bash
   migrate -database $DATABASE_URL -path migrations down
   ```

---

## License

This template is open-source and available under the [MIT License](LICENSE).
