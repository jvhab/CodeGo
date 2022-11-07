Дан массив a, состоящий из n положительных целых чисел. Найдите непустое подмножество его элементов с чётной (т.е. делящейся на 2) суммой или определите, что такого подмножества нет.

И заданный массив и искомое подмножетсво могут содержать равные значения элементов.

Входные данные
В первой строке задано целое число t (1≤t≤100) — количество наборов входных данных, для которых требуется решить задачу. Затем следуют описания t наборов входных данных.

Описание каждого набора входных данных состоит из двух строк. В первой строке задано одно целое число n (1≤n≤100) — количество элементов в массиве a. Во второй строке заданы n целых чисел a1,a2,…,an (1≤ai≤100) — элементы массива a. Массив a может содержать одинаковые (равные) значения элементов.

Выходные данные
Для каждого набора входных данных выведите −1, если не существует требуемого подмножества элементов. Иначе выведите целое число k — количество элементов в подмножестве. Затем выведите k различных чисел (1≤pi≤n) — индексы элементов найденного подмножества. Если есть несколько подходящих подмножеств, выведите любое из них.