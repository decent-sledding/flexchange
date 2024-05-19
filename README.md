# UnKEx

> Provides actual information for `UAH`/`USD` exchange rate 
> Allows subsribtion to the exchange rate updates


# Requirements
## For containerized application
- `Docker` v26.0.2
- `Docker Compose` v26.0.2

## For local development
- `Go` >= v1.20
- `PostgreSQL` >= 14
- `nodejs` == 20.0.*
- `npm` >= 9.8.0


# Specification
Потрібно реалізувати сервіс з АРІ, який дозволить:

- дізнатись поточний курс долара (USD) у гривні (UAH)
- підписати емейл на отримання інформації по зміні курсу.

#### Додаткові вимоги:
1. Сервіс має відповідати описаному ниже АРІ. NB Закривати рішення аутентифікацією не потрібно.
2. Eмейли з актуальний курсом мають відправлятись юзерам 1 раз на добу.
3. Всі дані, для роботи додатку повинні зберігатися в базі даних. Також потрібно реалізувати виконання міграції структури БД при піднятті сервісу.
4. В репозиторії повинен бути Dockerfile, та Docker Compose який дозволяє запустити систему в Docker. З матеріалом по Docker вам необхідно ознайомитись самостійно.
5. Також ти можеш додавати коментарі чи опис логіки виконання роботи в README.md документі. Правильна логіка може стати перевагою при оцінюванні, якщо ти не повністю виконаєш завдання.
6. Очікувані мови виконання завдання: PHP, --Go--, NodeJS. Виконувати завдання іншими мовами можна, проте, це не буде перевагою.
7. Виконане завдання необхідно завантажити на GitHub (публічний репозиторій) та сабмітнути виконане завдання в гугл-форму.


---
# Simple Frontend
> This is simple frontend app composed using `Next.js` and `Typescript`
>
> It provides bare simple interface to interact with the backend through UI in the browser

- #### Proceed with [frontend documentation](frontend/)


# Backend
> This app is written in `Go` and provides backend API service.
>
> As well as a CLI Crontab app to perform routine tasks

- #### Proceed with [backend documentation](backend/)



# Notes
1. It is usually a bad practice to commit `.env` files into repository. Though...:
> Though in this case we only pushed a basic template to ease subsequent project configuration (for other comrad devs)
>
> Only a template is commited, next commits are ignored because file is removed from tracking

2. You need to set proper permissions within database
> Set proper permissions for a database that is being used for local development/testing

> This should be automatically configured in case `Docker Compose` is used