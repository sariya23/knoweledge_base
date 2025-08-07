[[Airflow]]
Переменные в Airflow, то же самое, что и переменные в Gitlab CI/CD - пара ключ-значение. В них можно хранить какую-то конфиг информацию, а потом использовать ее в рабочем процессе. 
Создать переменную можно через UI:
1. Admin
2. Variables
3. Add a new record
Либо через API
```python
import requests

url = "http://localhost:8080/api/v1/variables"
headers = {"Content-Type": "application/json"}
data = {
    "key": "aboba",
    "value": "True"
}

response = requests.post(url, headers=headers, json=data, auth=("airflow", "airflow"))

if response.status_code == 200:
    print("Variables set")
else:
    print(response)
```
