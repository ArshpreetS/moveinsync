# Moveinsync Assignment


## Architecture

![architecture](./screenshots/architecture.png)


## Instructions

1. Start all the containers using `docker compose up` command.
2. Use the following command to setup the mysql database and tables.

```bash
docker exec -i mysql mysql -u root -proot < init_database.sql
```

