# Go + REACT docker boiler template

## Tech
- Docker for creating containers for UI, Server and DB
- React for user interface and visualization
- Go for backend APIs
- Postgre for DB
- _<will evolve as it goes :) >_

## Basic commands to run the containers

### UI application
```sh
cd ui
sudo docker compose -f docker-compose.dev.yml up
in browser -> http://localhost:3000
```

### Server application
```sh
cd ui
sudo docker compose up
/ or make dev
in browser -> http://localhost:3000
```
