# Как запустить проект 

В данный проект я не стал добавлять файл .gitignore и добавлять в него .env, так как проект тестовый, поэтому для запуска проекта достаточно склонировать его и ввести следующую команду:
```shell
docker compose up -d --build
```
После того как поднимется контейрнеры приложения, для полного ознакомления будет достпен Swagger по адресу 

http://localhost:8080/swagger/index.html#/
