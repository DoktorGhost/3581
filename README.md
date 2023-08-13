# Задание 35.8.1

Разработайте сетевую службу по аналогии с сервером времени (не RPC), которая бы каждому подключившемуся клиенту показывала раз в 3 секунды случайную Go-поговорку. Поговорки возьмите с [сайта](https://go-proverbs.github.io/)
Служба должна поддерживать множественные одновременные подключения. Служба не должна завершать соединение с клиентом.
Вы должны проверить работу приложения с помощью telnet.

## Решение
1. В отдельном пакете (pcg/parse/parse.go) парсер поговорок с [сайта](https://go-proverbs.github.io/)
2. Для завершения соединения используется канал, в который горутина отправит TRUE, если пользователь отправит что-либо на ввод.
