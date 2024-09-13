# Приветствие

Здравствуйте! Меня зовут Дмитрий Давиденко Николаевич. Я подготовил данный проект в рамках стажировки на Avito. Прошу обратить внимание, что из-за ограниченных сроков работы код может показаться немного неопрятным. В некоторых местах требуется декомпозиция, вынос констант и дополнительная оптимизация. Тем не менее, техническое задание полностью выполнено.

# Документация

### Пинг (Проверка доступности сервера)

#### Проверка доступности сервера
- **Эндпоинт:** GET /ping
- **Описание:** Этот эндпоинт используется для проверки доступности сервера.
- **Ожидаемый результат:** Статус код 200 и ответ "ok".

```yaml
GET /api/ping

Response:

  200 OK

  Body: "ok"
``` 



### Пользователи

#### Создание нового пользователя
- **Эндпоинт:** POST /users/new
- **Описание:** Создает нового пользователя в системе.
- **Ожидаемый результат:** Статус код 201 и информация о созданном пользователе.

```yaml
POST /api/users/new

Request Body:
{
  "username": "new_user",
  "firstName": "John",
  "lastName": "Doe"
}

Response:

  201 Created

  Body:
  {
    "id": 1,
    "username": "new_user",
    "firstName": "John",
    "lastName": "Doe",
    "createdAt": "2023-09-01T12:34:56Z",
    "updatedAt": "2023-09-01T12:34:56Z"
  }
```

#### Получение списка пользователей
- **Эндпоинт:** GET /users/
- **Описание:** Возвращает список всех пользователей системы.
- **Ожидаемый результат:** Статус код 200 и список пользователей.

```yaml
GET /api/users/

Response:

  200 OK

  Body:
  [
    {
      "id": 1,
      "username": "new_user",
      "firstName": "John",
      "lastName": "Doe",
      "createdAt": "2023-09-01T12:34:56Z",
      "updatedAt": "2023-09-01T12:34:56Z"
    },
    {
      "id": 2,
      "username": "another_user",
      "firstName": "Jane",
      "lastName": "Doe",
      "createdAt": "2023-09-01T12:35:56Z",
      "updatedAt": "2023-09-01T12:35:56Z"
    }
  ]
```

#### Получение информации о пользователе по ID
- **Эндпоинт:** GET /users/{user_id}
- **Описание:** Возвращает информацию о пользователе по указанному ID.
- **Ожидаемый результат:** Статус код 200 и данные пользователя.

```yaml
GET /api/users/1

Response:

  200 OK

  Body:
  {
    "id": 1,
    "username": "new_user",
    "firstName": "John",
    "lastName": "Doe",
    "createdAt": "2023-09-01T12:34:56Z",
    "updatedAt": "2023-09-01T12:34:56Z"
  }
```

#### Обновление информации о пользователе
- **Эндпоинт:** PATCH /users/{user_id}/edit
- **Описание:** Обновляет информацию о пользователе по указанному ID.
- **Ожидаемый результат:** Статус код 200 и обновленные данные пользователя.

```yaml
PATCH /api/users/1/edit

Request Body:
{
  "username": "updated_user",
  "firstName": "John",
  "lastName": "Doe"
}

Response:

  200 OK

  Body:
  {
    "id": 1,
    "username": "updated_user",
    "firstName": "John",
    "lastName": "Doe",
    "createdAt": "2023-09-01T12:34:56Z",
    "updatedAt": "2023-09-01T13:00:00Z"
  }
```

#### Удаление пользователя
- **Эндпоинт:** DELETE /users/{user_id}/delete
- **Описание:** Удаляет пользователя по указанному ID.
- **Ожидаемый результат:** Статус код 204, пользователь успешно удален.

```yaml
DELETE /api/users/1/delete

Response:

  204 No Content
```



### Организации

#### Создание новой организации
- **Эндпоинт:** POST /organizations/new
- **Описание:** Создает новую организацию.
- **Ожидаемый результат:** Статус код 201 и информация о созданной организации.

```yaml
POST /api/organizations/new

Request Body:
{
  "name": "New Organization",
  "description": "Organization Description",
  "type": "Corporate"
}

Response:

  201 Created

  Body:
  {
    "id": 1,
    "name": "New Organization",
    "description": "Organization Description",
    "type": "Corporate",
    "createdAt": "2023-09-01T12:34:56Z",
    "updatedAt": "2023-09-01T12:34:56Z"
  }
```

#### Получение списка организаций
- **Эндпоинт:** GET /organizations/
- **Описание:** Возвращает список всех организаций.
- **Ожидаемый результат:** Статус код 200 и список организаций.

```yaml
GET /api/organizations/

Response:

  200 OK

  Body:
  [
    {
      "id": 1,
      "name": "New Organization",
      "description": "Organization Description",
      "type": "Corporate",
      "createdAt": "2023-09-01T12:34:56Z",
      "updatedAt": "2023-09-01T12:34:56Z"
    },
    {
      "id": 2,
      "name": "Another Organization",
      "description": "Another Description",
      "type": "Non-Profit",
      "createdAt": "2023-09-02T14:35:00Z",
      "updatedAt": "2023-09-02T14:35:00Z"
    }
  ]
```

#### Получение информации об организации по ID
- **Эндпоинт:** GET /organizations/{org_id}
- **Описание:** Возвращает информацию об организации по указанному ID.
- **Ожидаемый результат:** Статус код 200 и данные организации.

```yaml
GET /api/organizations/1

Response:

  200 OK

  Body:
  {
    "id": 1,
    "name": "New Organization",
    "description": "Organization Description",
    "type": "Corporate",
    "createdAt": "2023-09-01T12:34:56Z",
    "updatedAt": "2023-09-01T12:34:56Z"
  }
```

#### Обновление информации об организации
- **Эндпоинт:** PATCH /organizations/{org_id}/edit
- **Описание:** Обновляет информацию об организации по указанному ID.
- **Ожидаемый результат:** Статус код 200 и обновленные данные организации.

```yaml
PATCH /api/organizations/1/edit

Request Body:
{
  "name": "Updated Organization",
  "description": "Updated Description",
  "type": "Corporate"
}

Response:

  200 OK

  Body:
  {
    "id": 1,
    "name": "Updated Organization",
    "description": "Updated Description",
    "type": "Corporate",
    "createdAt": "2023-09-01T12:34:56Z",
    "updatedAt": "2023-09-01T13:00:00Z"
  }
```

#### Удаление организации
- **Эндпоинт:** DELETE /organizations/{org_id}/delete
- **Описание:** Удаляет организацию по указанному ID.
- **Ожидаемый результат:** Статус код 204, организация успешно удалена.

```yaml
DELETE /api/organizations/1/delete

Response:

  204 No Content
```

### Ответственные

#### Добавление ответственного за организацию
- **Эндпоинт:** POST /organizations/{org_id}/responsibles/{user_id}/new
- **Описание:** Добавляет пользователя как ответственного за организацию.
- **Ожидаемый результат:** Статус код 204, пользователь успешно добавлен.

```yaml
POST /api/organizations/1/responsibles/1/new

Response:

  204 No Content
```

#### Удаление ответственного за организацию
- **Эндпоинт:** DELETE /organizations/{org_id}/responsibles/{user_id}/delete
- **Описание:** Удаляет пользователя из ответственных за организацию.
- **Ожидаемый результат:** Статус код 204, пользователь успешно удален.

```yaml
DELETE /api/organizations/1/responsibles/1/delete

Response:

  204 No Content
```

#### Получение списка ответственных за организацию
- **Эндпоинт:** GET /organizations/{org_id}/responsibles
- **Описание:** Возвращает список ответственных пользователей за указанную организацию.
- **Ожидаемый результат:** Статус код 200 и список ответственных пользователей.

```yaml
GET /api/organizations/1/responsibles

Response:

  200 OK

  Body:
  [
    {
      "id": 1,
      "username": "user1",
      "firstName": "John",
      "lastName": "Doe"
    },
    {
      "id": 2,
      "username": "user2",
      "firstName": "Jane",
      "lastName": "Doe"
    }
  ]
```

#### Получение ответственного по ID
- **Эндпоинт:** GET /organizations/{org_id}/responsibles/{user_id}
- **Описание:** Возвращает информацию о конкретном ответственном пользователе по ID организации и ID пользователя.
- **Ожидаемый результат:** Статус код 200 и данные ответственного пользователя.

```yaml
GET /api/organizations/1/responsibles/1

Response:

  200 OK

  Body:
  {
    "id": 1,
    "username": "user1",
    "firstName": "John",
    "lastName": "Doe"
  }
```



### Тендеры

#### Создание тендера
- **Эндпоинт:** POST /api/tenders/new
- **Описание:** Создание нового тендера от имени определенного пользователя.
- **Ожидаемый результат:** Статус код 200 и информация о созданном тендере.

```yaml
POST /api/tenders/new

Request Body:
{
  "name": "New Tender",
  "description": "Description of the tender",
  "serviceType": "construction",
  "status": "open",
  "organizationID": 1,
  "creatorUsername": "john_doe"
}

Response:

  200 OK

  Body:
  {
    "id": 1,
    "name": "New Tender",
    "description": "Description of the tender",
    "serviceType": "construction",
    "status": "open",
    "organizationID": 1,
    "createdAt": "2024-09-13T10:00:00Z",
    "version": 1
  }
```

#### Получение списка тендеров
- **Эндпоинт:** GET /api/tenders/
- **Описание:** Получение списка всех тендеров с возможностью фильтрации по типу сервиса.
- **Ожидаемый результат:** Статус код 200 и список тендеров.

```yaml
GET /api/tenders/?serviceType=construction

Response:

  200 OK

  Body: [
    {
      "id": 1,
      "name": "Tender 1",
      "description": "Description of Tender 1",
      "serviceType": "construction",
      "status": "open",
      "organizationID": 1,
      "createdAt": "2024-09-13T10:00:00Z",
      "version": 1
    },
    {
      "id": 2,
      "name": "Tender 2",
      "description": "Description of Tender 2",
      "serviceType": "construction",
      "status": "closed",
      "organizationID": 2,
      "createdAt": "2024-09-10T08:00:00Z",
      "version": 1
    }
  ]
```

#### Получение тендера по ID
- **Эндпоинт:** GET /api/tenders/{tenderId}
- **Описание:** Получение информации о конкретном тендере по его ID.
- **Ожидаемый результат:** Статус код 200 и информация о тендере.

```yaml
GET /api/tenders/1

Response:

  200 OK

  Body:
  {
    "id": 1,
    "name": "Tender 1",
    "description": "Description of Tender 1",
    "serviceType": "construction",
    "status": "open",
    "organizationID": 1,
    "createdAt": "2024-09-13T10:00:00Z",
    "version": 1
  }
```

#### Получение тендеров по имени пользователя
- **Эндпоинт:** GET /api/tenders/my/{username}
- **Описание:** Получение списка тендеров, созданных конкретным пользователем.
- **Ожидаемый результат:** Статус код 200 и список тендеров.

```yaml
GET /api/tenders/my/john_doe

Response:

  200 OK

  Body: [
    {
      "id": 1,
      "name": "Tender 1",
      "description": "Description of Tender 1",
      "serviceType": "construction",
      "status": "open",
      "organizationID": 1,
      "createdAt": "2024-09-13T10:00:00Z",
      "version": 1
    }
  ]
```

#### Обновление тендера
- **Эндпоинт:** PATCH /api/tenders/{tenderId}/edit
- **Описание:** Обновление информации о тендере.
- **Ожидаемый результат:** Статус код 200 и обновленная информация о тендере.

```yaml
PATCH /api/tenders/1/edit

Request Body:
{
  "name": "Updated Tender Name",
  "description": "Updated description of the tender"
}

Response:

  200 OK

  Body:
  {
    "id": 1,
    "name": "Updated Tender Name",
    "description": "Updated description of the tender",
    "serviceType": "construction",
    "status": "open",
    "organizationID": 1,
    "createdAt": "2024-09-13T10:00:00Z",
    "version": 2
  }
```

#### Публикация тендера
- **Эндпоинт:** POST /api/tenders/{tenderId}/publish
- **Описание:** Публикация тендера.
- **Ожидаемый результат:** Статус код 200, тендер успешно опубликован.

```yaml
POST /api/tenders/1/publish

Response:

  200 OK
```

#### Закрытие тендера
- **Эндпоинт:** POST /api/tenders/{tenderId}/close
- **Описание:** Закрытие тендера.
- **Ожидаемый результат:** Статус код 200, тендер успешно закрыт.

```yaml
POST /api/tenders/1/close

Response:

  200 OK
```

#### Откат версии тендера
- **Эндпоинт:** PUT /api/tenders/{tenderId}/rollback/{version}
- **Описание:** Откат тендера к предыдущей версии.
- **Ожидаемый результат:** Статус код 200 и информация о восстановленном тендере.

```yaml
PUT /api/tenders/1/rollback/1

Response:

  200 OK

  Body:
  {
    "id": 1,
    "name": "Tender 1",
    "description": "Description of Tender 1",
    "serviceType": "construction",
    "status": "open",
    "organizationID": 1,
    "createdAt": "2024-09-13T10:00:00Z",
    "version": 1
  }
```

#### Удаление тендера
- **Эндпоинт:** DELETE /api/tenders/{tenderId}/delete
- **Описание:** Удаление тендера по его ID.
- **Ожидаемый результат:** Статус код 200, тендер успешно удалён.

```yaml
DELETE /api/tenders/1/delete

Response:

  200 OK
```



Вот документация для роутов, которые ты предоставил:

### Ставки (Bids)

#### Создание новой ставки
- **Эндпоинт:** POST /api/bids/new
- **Описание:** Пользователь может создать новую ставку.
- **Ожидаемый результат:** Статус код 201, ставка успешно создана.

```yaml
POST /api/bids/new

Request Body:
{
  "name": "New Bid",
  "description": "Description of the bid",
  "tenderId": 1,
  "organizationId": 1,
  "creatorId": 123,
  "status": "pending"
}

Response:

  201 Created

  Body:
  {
    "id": 1,
    "name": "New Bid",
    "description": "Description of the bid",
    "tenderId": 1,
    "organizationId": 1,
    "creatorId": 123,
    "status": "pending",
    "createdAt": "2023-09-13T12:00:00Z",
    "version": 1
  }
```

#### Получение ставки по ID
- **Эндпоинт:** GET /api/bids/{bidId}
- **Описание:** Возвращает информацию о ставке по указанному ID.
- **Ожидаемый результат:** Статус код 200, информация о ставке.

```yaml
GET /api/bids/1

Response:

  200 OK

  Body:
  {
    "id": 1,
    "name": "Bid Name",
    "description": "Bid Description",
    "tenderId": 1,
    "organizationId": 1,
    "creatorId": 123,
    "status": "approved",
    "createdAt": "2023-09-13T12:00:00Z",
    "version": 1
  }
```

#### Получение всех ставок по ID тендера
- **Эндпоинт:** GET /api/bids/tender/{tenderId}
- **Описание:** Возвращает список ставок, связанных с указанным тендером.
- **Ожидаемый результат:** Статус код 200, список ставок.

```yaml
GET /api/bids/tender/1

Response:

  200 OK

  Body: [
    {
      "id": 1,
      "name": "Bid 1",
      "description": "Description 1",
      "tenderId": 1,
      "organizationId": 1,
      "creatorId": 123,
      "status": "approved",
      "createdAt": "2023-09-13T12:00:00Z",
      "version": 1
    },
    {
      "id": 2,
      "name": "Bid 2",
      "description": "Description 2",
      "tenderId": 1,
      "organizationId": 1,
      "creatorId": 124,
      "status": "pending",
      "createdAt": "2023-09-14T12:00:00Z",
      "version": 1
    }
  ]
```

#### Просмотр ставок пользователя
- **Эндпоинт:** GET /api/bids/my/{username}
- **Описание:** Возвращает список ставок, созданных указанным пользователем.
- **Ожидаемый результат:** Статус код 200, список ставок пользователя.

```yaml
GET /api/bids/my/user1

Response:

  200 OK

  Body: [
    {
      "id": 1,
      "name": "User1 Bid",
      "description": "Description",
      "tenderId": 1,
      "organizationId": 1,
      "creatorId": 123,
      "status": "approved",
      "createdAt": "2023-09-13T12:00:00Z",
      "version": 1
    }
  ]
```

#### Обновление ставки
- **Эндпоинт:** PATCH /api/bids/{bidId}/edit
- **Описание:** Пользователь может обновить данные существующей ставки.
- **Ожидаемый результат:** Статус код 200, обновлённая ставка.

```yaml
PATCH /api/bids/1/edit

Request Body:
{
  "name": "Updated Bid Name",
  "description": "Updated Description"
}

Response:

  200 OK

  Body:
  {
    "id": 1,
    "name": "Updated Bid Name",
    "description": "Updated Description",
    "tenderId": 1,
    "organizationId": 1,
    "creatorId": 123,
    "status": "approved",
    "createdAt": "2023-09-13T12:00:00Z",
    "version": 2
  }
```



### Комментарии

#### Создание нового комментария
- **Эндпоинт:** POST /api/comments
- **Описание:** Пользователь может добавить новый комментарий.
- **Ожидаемый результат:** Статус код 201, комментарий успешно создан.

```yaml
POST /api/comments

Request Body:
{
  "content": "This is a comment",
  "userId": 123,
  "bidId": 1
}

Response:

  201 Created

  Body:
  {
    "id": 1,
    "content": "This is a comment",
    "userId": 123,
    "bidId": 1,
    "createdAt": "2023-09-13T12:00:00Z"
  }
```

#### Просмотр отзывов на прошлые предложения
- **Эндпоинт:** GET /bids/{tenderId}/reviews
- **Описание:** Ответственный за организацию может посмотреть прошлые отзывы на предложения автора, который создал предложение для его тендера.
- **Ожидаемый результат:** Статус код 200 и список отзывов на предложения указанного автора.

```yaml
GET /api/bids/1/reviews?authorUsername=user2&organizationId=1

Response:

  200 OK

  Body: [
    {
      "id": 1,
      "content": "Great bid!",
      "userId": 123,
      "bidId": 1,
      "createdAt": "2023-09-13T12:00:00Z"
    },
    {
      "id": 2,
      "content": "Not bad",
      "userId": 124,
      "bidId": 1,
      "createdAt": "2023-09-14T12:00:00Z"
    }
  ]
```

#### Удаление комментария
- **Эндпоинт:** DELETE /api/comments/{commentId}
- **Описание:** Удаляет комментарий по его ID.
- **Ожидаемый результат:** Статус код 200, комментарий успешно удалён.

```yaml
DELETE /api/comments/1

Response:

  200 OK
```

# Заключение

Благодарю за внимание и за возможность участия в этом этапе отбора. Желаю вам приятной проверки кода, и надеюсь на положительный результат!
