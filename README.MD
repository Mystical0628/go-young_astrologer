Описание директорий
1. ./nginx - конфиг для nginx
2. ./package-core - базовый пакет с конфигами
3. ./service-api - сервис http api
4. ./service-migrations - Миграции и докерфайл который собирает образ с приложением миграций
5. ./service-nasa - сервис который качает apod через апи

Для запуска предварительно собираем пакеты service-nasa и service-api (в идеале сборка должна происходить в контейнере, но для экономии времени решил просто примонтировать готовые бинарники) и стартуем docker-compose, сразу создатся бд и запустяться миграции, также единожды выполниться service-nasa

Примечания
1. Nginx в качетсве прокси сервера
2. В service-api реализовал подогрев данных:
   1. Из db-repo вытягиваются все данные и загружаются в mem-repo, и затем mem-repo используется как основная репа для полeчения данных. 
   2. Обновление происходит каждые 10сек. 
   3. При желании можно в качестве основной репы использовать db-repo:
      1. Убрать вызов PreloadService
      2. В Core.Repositories.Apod закинуть инстанс infrastructure.db.NasaApodRepository
3. В контейнере бд пробросил порт, чтобы тестировать с локалки, по хорошему он должен быть закрыт
4. Не стал писать воркер для service-nasa, т.к. это можно делать либо через крон, либо через брокер сообщений
5. Подрузамеваеться что каждая отдельная папка в корне данного репозитория, должны быть отдельными репозиториями (за исколючением папки nginx, там чисто конфиг). А соответственно service-nasa и service-api должны подключать пакет package-core как репу. Env как мне кажеться должен деплоиться для каждого сервиса отдельно черзе менеджер конфигураций типа ansible, но тут для простоты я просто вынес в общий пакет package-core (Там 2 енва - для запуска с локалки и для запуска в контейнере)