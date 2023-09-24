# Citix billboard Backend
This project was made for Citix Lab in a hackathon by InnovateX

[CITIX BILLBOARD FRONT API](https://github.com/Eager-coder/citix-billboard-frontend/tree/main)
# How to run
1. Setup .env, in .env will be stored way to your config file
    ```env
    CONFIG_FILE_PATH=
    CONFIG_FILE_NAME=
    CONFIG_FILE_FORMAT=
    ```
2. Setup config.yaml, this file will be contained all configs data that needed for app
    ```yaml
    app:
      mode:
      host:
      port:
      timeout:
    
    source:
      news_api:
      news_api_key:
      timeout:
    
    cloudinary:
      cloud_name:
      api_key:
      api_secret:
    
    wasabi:
      endpoint:
      access_key_id:
      secret_key:
      region:
      bucket:
      path:
      timeout:
    
    postgres:
      hostname:
      port:
      user:
      password:
      database:
      user_table:
      survey_table:
      photo_table:
      timeout:
    ```

3. Setup docker-compose.yaml empty field and run app
```yaml
version: "3.9"

networks:
  internal:
    driver: bridge

volumes:
  database_data:

services:
  postgres:
    image: postgres:alpine
    restart: always
    container_name: citix-db
    hostname: postgres
    volumes:
      - ./scripts/postgres/:/docker-entrypoint-initdb.d/
      - database_data:/var/lib/postgresql/data
    ports:
      - "5436:5432"
    environment:
      - POSTGRES_DB=via-db
      - POSTGRES_USER=via-user
      - POSTGRES_PASSWORD=via-qwerty
    networks:
      - internal
    healthcheck:
      test: [ "CMD-SHELL", "pg_isready -U onelab -d onelab_db" ]
      interval: 10s
      timeout: 5s
      retries: 5

  app:
    build:
      context: .
      dockerfile: ./build/Dockerfile
    restart: on-failure
    container_name: citix-app
    networks:
      - internal
    ports:
      - "8080:8080"
    volumes:
      - ./config/config.yaml:/root/config/config.yaml
      - .env:/root/.env
    environment:
      - CONFIG_FILE_PATH=config
      - CONFIG_FILE_NAME=config
      - CONFIG_FILE_FORMAT=yaml
    depends_on:
      postgres:
        condition: service_healthy
```

```shell
docker compose up --build
```

# Realized
- News list
- Survey module
- Real Time Photo Contest

# TODO
- Life Chat video streaming
- Admin dashboard for government institutes and Citix billboard managers
- etc
