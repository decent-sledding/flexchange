# FlexChange

> Provides actual information for `UAH`/`USD` exchange rate 
> Allows subsribtion to the exchange rate updates

** Unfortunately this project is not finished yet


# Technical Requirements

Implement a service with API, which allows:
- acquire current exchange rate USD to UAH (the API can be extended to allow other currencies also)
- email subscription for alternations in the exchange rate


Additional Requirements:
1. Service must be implemented in appliance with described API
2. Emails have to be dispatched once a day
3. All the application data must be stored inside database. Additionally "migrations" mechanism has to be provided/implemented.
4. A `Dockerfile` is a requirement as well as `docker-compose.yml`.
5. Comments are encouraged. Description of the program logic is encouraged.
6. Properly described logic maybe an advantage in project estimation in case not all features are implemented.
7. Implementation can be provided in one of the next languages: `PHP`, `Go`, `NodeJS`. This is not a limitation, but there is no advantage in the choice of a language.
8. Finished project has to be published to a github repository.


# Project


## Requirements
### For containerized application
- `Docker` v26.0.2
- `Docker Compose` v26.0.2

### For local development
- `Go` >= v1.20
- `PostgreSQL` >= 14
- `nodejs` == 20.0.*
- `npm` >= 9.8.0

---
## Simple Frontend
> This is simple frontend app composed using `Next.js` and `Typescript`
>
> It provides bare simple interface to interact with the backend through UI in the browser

- #### Proceed with [frontend documentation](frontend/)


# Backend
> This app is written in `Go` and provides backend API service.
>
> As well as a CLI Crontab app to perform routine tasks

- #### Proceed with [backend documentation](backend/)
