package main

/*
=== HTTP server ===

Сделать HTTP  сервер (без применения фреймворков), умеющий создавать, изменять и удалять записи двух типов:
    • Информация о покупателе
        ◦ Фамилия
        ◦ Имя
        ◦ Отчество
        ◦ Возраст (необязательное поле)
        ◦ Дата регистрации
    • Информацию о магазине
        ◦ Название
        ◦ Адрес
        ◦ Работающий или нет
        ◦ Владелец (необязательное поле)

Непосредственную работу с записями необходимо осуществлять с использованием одной функции, умеющей принимать в качестве входного значения оба типа записи. Хранить и накапливать информацию можно по выбору: в рантайм, СУБД, файлах. Входными и выходными параметрами в  HTTP  запросах являются данные в формате  JSON.
На выходе сервис должен уметь возвращать всю запись, либо одно поле из записи в зависимости от запроса пользователя, осуществлять поиск по Фамилии и Названию  для соответствующих записей.



Можно немного усложнить ее, разбив на два микросервиса:

Сделать два микросервиса работающих через  GRPC:
    1. HTTP  сервер для общения с пользователем
    2. Сервис хранения и обработки данных, в который ходит HTTP  сервер за данными.
Остальные вводные те же.
*/

func main() {

}
