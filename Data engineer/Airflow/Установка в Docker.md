[[Airflow]]
1. Скачиваем docker compose - https://airflow.apache.org/docs/apache-airflow/stable/howto/docker-compose/index.html
2. Создаем в проекте папки logs, plugins, dags, так как в контейнеры примонтированы эти папки
3. В корне прописываем `sudo docker compose up airflow-init`, чтобы создать юзеров, базы и накатить миграции
4. Выполнить `sudo docker-compose up -d`
