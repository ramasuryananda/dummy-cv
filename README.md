# Invoice App

### 1. Requirement

| No. | Name | Version | Notes |
| --- | --- | --- | --- |
| 1 | Go | 1.20 | - |
| 2 | MySql | 8.0 | - |

### 2. Installation & Setup

- Clone project from the repository using http
```bash
git clone https://github.com/ramasuryananda/dummy-cv.git
```

- Clone project from the repository if using ssh

```bash
git clone git@github.com:ramasuryananda/dummy-cv.git
```

- Create *.env* file from *.env.example* file

```bash
cp .env.example .env
```

- Fill the *.env* file using your own environment

- Install vendor dependencies
```bash
go mod vendor
```

- Install Golang Migrator (for migrate case)
```bash
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

- Migrate database (installed golang migrator needed)
```bash
make migrate
```
or using migrate command if not have installed make
```bash
migrate -database ${DB_MIGRATION_CONNECTION} -path database/migrations up
```

- Seed database (installed golang migrator needed)
```bash
make seed
```
or using seed command if not have installed make
```bash
migrate -database ${DB_SEEDER_CONNECTION} -path database/seeders up
```

- Running the application
```bash
make run
```