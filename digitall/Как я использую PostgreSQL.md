Решил, что буду юзать постгрес через контейнер. 
Dockerfile:
```Dockerfile
FROM postgres:16.3

# Устанавливаем необходимые пакеты для поддержки локалей
RUN apt-get update && \
    apt-get install -y locales && \
    rm -rf /var/lib/apt/lists/*

# Генерируем локаль ru_RU.UTF-8
RUN localedef -i ru_RU -c -f UTF-8 -A /usr/share/locale/locale.alias ru_RU.UTF-8

# Устанавливаем переменные окружения для локали
ENV LANG ru_RU.utf8
```
Билд:
```bash
docker build -t my_postgres:latest .
```
Запуск:
```bash
docker run --name postgresql -e POSTGRES_PASSWORD=1234 -p 5432:5432 -d my_postgres:latest
```
Чтобы войти в контейнер:
```bash
docker exec -ti postgresql psql -h localhost -U postgres
```
Далее делаем там, что хотим. Например создаем юзера и БД
```SQL
CREATE ROLE user WITH LOGIN PASSWORD '1234';
CREATE DATABASE db WITH TEMPLATE=template0 ENCODING='UTF8' LC_COLLATE='ru_RU.UTF-8' LC_CTYPE='ru_RU.UTF-8' owner user;
```
И теперь мы можем сразу подключаться к этой БД:
```bash
docker exec -ti postgresql psql -h localhost -U user -d db
```
## GUI
В качестве удобного интерфейса мне понравился DBeaver. Чтобы подключиться к БД через него надо тыкнуть на розетку сверху и ввести все данные: пароль, юзера, порт, БД и хост. 
![[Pasted image 20240724230012.png]]