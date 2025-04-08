<h2>Quick Start</h2>

Из исходников:
```
1. go mod tidy
2. docker-compose up
3. make mig-u
4. make build || go run ./...
```

<h3>Сервер должен реализовывать следующую бизнес-логику:</h3>
1. [x] регистрация, аутентификация и авторизация пользователей;
2. [x] хранение приватных данных;
3. [x] синхронизация данных между несколькими авторизованными клиентами одного владельца;
4. [x] передача приватных данных владельцу по запросу.
<h3>Клиент должен реализовывать следующую бизнес-логику:</h3>
5. [x] аутентификация и авторизация пользователей на удалённом сервере;
6. [x] доступ к приватным данным по запросу.
<h3>Функции, реализация которых остаётся на усмотрение исполнителя:</h3>
- создание, редактирование и удаление данных на стороне сервера или клиента;
- формат регистрации нового пользователя; (на вход email, пароль) - на выходе JWT-token
- выбор хранилища и формат хранения данных; (мета-информация в БД, сами файлы на диске)
- обеспечение безопасности передачи и хранения данных; (Все данные шифруются на сервере, с помощью секретного слова и userID)
- протокол взаимодействия клиента и сервера; (gRPC)
- механизмы аутентификации пользователя и авторизации доступа к информации.
<h3>Дополнительные требования:</h3>
1. [x] клиент должен распространяться в виде CLI-приложения с возможностью запуска на платформах Windows, Linux и Mac OS;
2. [x] клиент должен давать пользователю возможность получить информацию о версии и дате сборки бинарного файла клиента.
<h3>Типы хранимой информации</h3>
3. [x] пары логин/пароль;
4. [x] произвольные текстовые данные;
5. [x] произвольные бинарные данные;
6. [x] данные банковских карт.
7. [x] Для любых данных должна быть возможность хранения произвольной текстовой метаинформации (принадлежность данных к веб-сайту, личности или банку, списки одноразовых кодов активации и прочее).