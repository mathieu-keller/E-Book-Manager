# E-Book Manager

To store EPUB ebooks on a server, sort and filter them.

Docker image can be found here: https://hub.docker.com/r/afrima/e-book-manager

To start, the following environment variables are needed:

| variables  | for what?                                 |
|------------|-------------------------------------------|
| GIN_MODE   | set to release                            |
| dbUser     | dbUser                                    |
| dbName     | name of the database                      |
| dbPassword | dbPassword                                |
| dbAddress  | postgres DB url or ip with port           |
| dbPort     | port from the db                          |
| user       | Optional if you want to have a basic login |
| password   | Optional if you want to have a basic login |

# **tl;dr**
```
services:
  server:
    image: afrima/e-book-manager:latest
    restart: on-failure
    ports:
      - "443:8080"
    environment:
      GIN_MODE: "release"
      dbPassword: "super-secret"
      dbUser: "postgres"
      dbAddress: "db"
      dbPort: "5432"
      dbName: "ebooks"
    volumes:
      - book-data:/home/appuser/upload/
    depends_on:
      - "db"
  db:
    image: postgres:latest
    restart: on-failure
    environment:
      POSTGRES_PASSWORD: "super-secret"
      POSTGRES_USER: "postgres"
      POSTGRES_DB: "ebooks"
    ports:
      - "5432:5432"
    volumes:
      - book-db:/var/lib/postgresql/data

volumes:
  book-data:
    external: false
  book-db:
    external: false
```
