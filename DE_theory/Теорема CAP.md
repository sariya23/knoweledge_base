[[DE]]
Теорема CAP - это фундаментальная теорема распределенных систем. И это не аксиома, а есть доказательство.
## C
C - сonsistency (согласованность) - каждый узел системы видит одинаковые данные. То есть нет такого, что для одного узла последний комментарий `n`, а для другого `m`. Согласованность на уровне БД можно достичичь синхронной репликацией, когда мы ждем подтверждение транзакции от всех узлов
## A
A - availability (доступность) - система работает даже если часть узлов отвалилась. И тут имеется ввиду не обычная доступность, а доступность в контексте теоремы. Доступность - любая неупавшая нода должна ответить. То есть если всего нод 100, упало 99, то работает одна и отвечает супер долго. В привычном понимании система не очень доступна, но с точки зрения CAP доступнсоть есть, так как нода не лежит и отвечает. А во время, за которое она отвечает теорему не волнует
## P
P - partiotion tolerance - толерантность к разделению сети. Это значит, что в системе из-за нестабильной связи между узлами, все равно должно идти все ок. Но это вроде как миф и P это про то, что в системе это расщепление возможно

Суть теоремы в том, что нельзя добиться исполнения каждого из пунктов, нужно чем-то жертвовать. В большинстве случаев P идет по дефолту, поэтому у нас остаются CP и AP системы. Но, системы AC существуют - например соверменные процессоры. Теорема CAP это не только про веб системы, это вообще, про любые системы

Источники:
- https://habr.com/ru/articles/322276/
- http://blog.thislongrun.com/2015/04/the-unclear-cp-vs-ca-case-in-cap.html