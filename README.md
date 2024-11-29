# Graphql

Цель проекта познакомиться с элементами graphql, изучение которых вышло за рамки пары тестовых файлов и чтения доков. Изначально мне было интересно
как интегрировать graphql и зачем он нужен, впоследствии написал на скорую руку сервис с http ендпоинтами и затем для них же написал graphql реализацию. 
В рамках репы была использована библиотека gqlgen от 99designs и http сервер fiber. Также распишу какие элементы изучил и какие мне 
не поддались по разным причинам.

## Для запуска проекта необходимо

- написать .env по подобию .env.example
- запустить проект в докере `docker-compose up`
- воспользоваться postman коллекцией `./docs/Graphql_Project_Http.postman_collection`

## О сервисе

Имитация работы магазина со следующими модулями и эндпоинтами:

#### Модуль user
- post `/user/register` - занесение пользователя в бд
- get `/user` - получение пользователей по фильтру

#### Модуль products
- post `/products` - занесение типизированных сведений в postgres и его нетипизированных атрибутов в mongodb
- get `/products` - получение сведений о продуктах по фильтру

#### Модуль order
- post `/order/:product_id` - создание сведений о заказе в бд
- get `/order` - получение сведений о заказах по фильтру

Также на каждую ручку была сделана зеркальная копия graphql формата, которые можно получить после запуска проекта с помощью интроспекции по адресам

#### Модуль user
- `/user/graphql` - query + mutation

#### Модуль products
- `/products/graphql` - query + mutation

#### Модуль order
- `/orders/graphql` - query + mutation

Изначально планировалось добавить subscription, но в go еще нет реализации библиотеки gqlgen, которая сможет адаптировать вебсокет для fasthttp (Поэтому изучение subscription делалось отдельно в тестовых файлах) 
- https://github.com/99designs/gqlgen/issues/1664 - открытый issue с этой проблемой
- https://github.com/gofiber/fiber/issues/403 - issue, в которой обсуждается проблемы Graphql support для фреймворков

## Какие элементы graphql были изучены

Большинство элементов были показаны в [официальной репе gqlgen](https://github.com/99designs/gqlgen/tree/master/_examples), остальные же пришлось познавать самому:
- Schema, Query, Mutation, Subscription, Resolvers 
- Fields, Arguments, Variables, Fragments, Directives 
- Scalars, Enums, Objects, Input, Interfaces, Unions, Lists
- Federation, Code Generator
- GraphiQL/Playground, Apollo Client, Introspection 
