# E-Book Manager

To store EPUB ebooks on a server, sort and filter them.

Docker image can be found here: https://hub.docker.com/r/afrima/e-book-manager

To start, the following environment variables are needed:

| variables  | for what?                                  |
|------------|--------------------------------------------|
| GIN_MODE   | set to release                             |
| dbUser     | dbUser                                     |
| dbName     | name of the database                       |
| dbPassword | dbPassword                                 |
| dbAddress  | maria DB url or ip with port               |
| dbPort     | port from the db                           |
| user       | Optional if you want to have a basic login |
| password   | Optional if you want to have a basic login |

If you want to start and test on your own machine, you need go version 1.18 and nodejs.
You can change the database connection to a sqlite connection in the backend/db/connector.go file.
