# AvitoTest

Это минимальная реализация тестового задания для Avito на Backend

### Запуск приложения
Для работы веб-приложения необходим установленный Docker, для запуска нужно ввести в терминале: `docker compose up`

### Endpoint'ы

`GET http://localhost:3000/segments` - Получить все сегменты
`GET http://localhost:3000/user_segments/:id` - Получить все сегменты юзера с указанным id
`GET http://localhost:3000/users` - Получить всех юзеров
`POST http://localhost:3000/users/:name` - Создать юзера с именем name
`DELETE http://localhost:3000/users/:id` - удалить юзера по id
`PUT http://localhost:3000/users/:SegmentsToAdd/:SegmentsToDelete/:UserId` - Изменить сегменты юзера по id (Сегменты, которые нужно добавить, удалить и юзер)
`POST http://localhost:3000/segments/:name` - Создать сегмент с именем name
`DELETE http://localhost:3000/segments/DeleteById/:id` - Удалить сегмент по id
`DELETE http://localhost:3000/segments/:name` - Удалить сегмент по имени
