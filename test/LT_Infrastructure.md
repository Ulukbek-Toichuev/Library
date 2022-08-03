# Как развернуть тестовую среду


## Тестовая среда состоит из следующих компонентов

1. Сам сервис реализован на **REST** архитектуре, написан на **Go**, присутствует **Swagger** документация
2. В качестве СУБД используется **Postgres DB**
3. Для мониторинга тестовой среды используется **Prometheus**
4. Проект работает в **Docker** через **Compose**


## Шаги для развертывания тестовой среды
-------------------------------------
### 1. Запустите на **VirtualBox/VMware** Linux дистрибутив **Ubuntu Server**
Материалы:

Где скачать: https://ubuntu.com/download/server/

Мануал №1: https://ubuntu.com/tutorials/install-ubuntu-server

Мануал №2: https://www.linuxtechi.com/install-ubuntu-server-22-04-step-by-step/

-------------------------------------
### 2. Проверьте на наличие следующих пакетов (при отсутствии установите их, команды указаны ниже)
- **net-tools**
```console
sudo apt install net-tools
```
- **git**
```console
sudo apt install git
```

- **docker**
1. Как устанавливать:

https://docs.docker.com/engine/install/ubuntu/

2. Что нужно сделать после установки, так называемый Post install

https://docs.docker.com/engine/install/linux-postinstall/

**ВНИМАНИЕ! №1**

Если вы хотите не заморачиваться, то склонируйте репозиторий (команда такая же как в пункте 4) и запустите скрипт docker.sh который лежит в папке **configs**

**ВНИМАНИЕ! №2**
В АРГУМЕНТАХ СКРИПТА ПЕРЕДАЙТЕ ИМЯ ПОЛЬЗОВАТЕЛЯ ПОД КОТОРЫМ ВЫ ЗАЛОГИНИЛИСЬ В LINUX


-------------------------------------
### 3. Для удобства работы используйте SSH
Для того чтобы узнать IP адрес виртуальной машины, используйте **ifconfig**
```console
ifconfig
```
![image](https://user-images.githubusercontent.com/67442103/179555998-2b4b85ea-bcfc-4376-b294-c749e4b1651c.png)


-------------------------------------
### 4. Склонируйте репозиторий Library

```console
git clone https://github.com/Ulukbek-Toychuev/Library.git
```


-------------------------------------
### 5. Перейдите в директорию Library и разверните среду с помощью Docker

```console
cd library && docker compose up -d
```


-------------------------------------
### 6. Проверьте работоспособность компонентов

#### 1. Откройте prometheus через http://host-ip:9090

Все инстансы должны быть отмечены зеленым цветом и словом UP

![image](https://user-images.githubusercontent.com/67442103/182324449-60940628-3310-451d-8f2f-bcfaa675aa80.png)


#### 2. Подключите grafana к prometheus

Рекомендуется установить **Grafana** на хост систему (то есть на ту где у вас работает виртуальная машина)

![image](https://user-images.githubusercontent.com/67442103/182325333-8574349b-56a7-4a31-93f9-eb16cb2dbc6c.png)

Так же для вас есть готовые дашборды для node exporter, postgres exporter

https://grafana.com/grafana/dashboards/1860

https://grafana.com/grafana/dashboards/9628

#### 3. Подключитесь к БД через dbeaver, у вас должны быть таблицы **users, books, users_books**

Пользователь: **postgres**

База данных: **postgres**

Пароль: **12345**

Адрес: **ip адрес виртуальный машины**

Порт: **5432**

![image](https://user-images.githubusercontent.com/67442103/182324776-81cb0b8f-6eb5-4443-b98e-57d786172f3e.png)


#### 4. Проверьте работоспособность Swagger
Пройдитесь по каждому эндпоинту через Swagger, расположен по адресу http://host-ip/swagger/index.html
![image](https://user-images.githubusercontent.com/67442103/182549806-59ee1056-c512-4470-89eb-a62bad2aae7f.png)


#### 5. Проверьте работоспособность самого сервиса
Пройдитесь по каждому эндпоинту через Postman, коллекцию можете импортировать, она расположена в папке **test**

![image](https://user-images.githubusercontent.com/67442103/182549725-41a5c035-3c32-4f10-b995-5e1a271c58c3.png)


-------------------------------------
### 7. Если у вас указанные пункты работают, поздравляем, вы развернули тестовую среду. Можете приступать к проведению нагрузочного тестирования.
