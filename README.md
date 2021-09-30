# rdbms-playground

## How to run

1. Create `.env`

   ```sh
   cp .env.example .env
   ```

2. Build Docker image and start development server

   ```sh
   docker-compose up --build
   ```

3. Open the following link in your browser

   <http://localhost:3000>

## Trouble Shooting

- Due to the upgrade of Docker Compose, the `docker-compose` command may have been replaced by `docker compose`.
