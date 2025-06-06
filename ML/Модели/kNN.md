## Гиперпараметры
- `n_neighbors` - кол-во соседей для присвоения нужного класса
- `metric` - метрика расстояния
- `p` - множитель для расстояния
## Сводка
kNN - это алгоритм обучения с учителем, который используется для классификации и регрессии. Сильно подвержен переобучению. 
**Преимущества**:
- Непараметрический
- Простой
- Достаточно точный
- Используется для классификации и регрессии
**Недостатки**:
- Неэффективный по памяти
- Вычислительно дорогой
- Чувствителен к масштабу данных
- Для применения алгоритма необходимо, чтобы метрическая близость признаков совпадала с их семантикой
**Применение**:
- Рекомендательные системы. Как какое-то начало, так как есть алгоритмы и лучше
- Поиск семантически похожих документов
- Поиск аномалий и выбросов
- Задача кредитного скоринга
## Обучение в scikit-learn
```python
knn = KNeighborsClassifier(n_neighbors=5, p=2, metric="minkowski")
knn.fit(features_train, labels_train)
```