# Moveinsync Assignment


## Architecture

![architecture](./screenshots/architecture.png)


## Instructions

1. Start all the containers using `docker compose up` command.
2. make clean ( put password if asks)
3. Use the following command to setup the mysql database and tables.

```bash
docker exec -i mysql mysql -u root -proot < init_database.sql
```

