# Online store
Простой проект онлайн-магазина.<br> Мой первый проект на Go

# Технологии
* Gin - веб-фреймворк
* GORM - ORM
* Postgresql - БД

# Установка
В директории `config` создайте файл .env по образу .env.template, 
наполните данный файл значениями

В корневой директории проекта создайте `docker-compose.yml` или `docker-compose.*.yml`.

После этого запустите команду 
```
docker-compose up --build
```
Или, если вы желаете запустить опредленный файл
```
docker-compose -f docker-compose.name.yml up --build
```
