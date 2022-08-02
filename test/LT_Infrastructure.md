# Как развернуть тестовую среду


## Тестовая среда состоит из следующих компонентов

1. Сам сервис реализован на **REST** архитектуре, написан на **Go**, присутствует **Swagger** документация
2. В качестве СУБД используется **Postgres DB**
3. Для мониторинга тестовой среды используется **Prometheus**


## Шаги для развертывания тестовой среды

### 1. Запустите на **VirtualBox** Linux дистрибутив **Ubuntu Server**
Материалы:

Где скачать: https://ubuntu.com/download/server/

Мануал №1: https://ubuntu.com/tutorials/install-ubuntu-server

Мануал №2: https://www.linuxtechi.com/install-ubuntu-server-22-04-step-by-step/

### 2. Для удобства работы используйте SSH

![image](https://user-images.githubusercontent.com/67442103/179555998-2b4b85ea-bcfc-4376-b294-c749e4b1651c.png)


### 3. Установите следующие пакеты
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

### 4. Склонируйте репозиторий Library

```console
git clone https://github.com/Ulukbek-Toychuev/Library.git
```

### 5. Перейдите в директорию Library и разверните среду с помощью Docker

```console
docker compose up -d
```

### 6. Проверьте работоспособность компонентов

#### 1. Откройте prometheus через http://host-ip:9090

Все инстансы должны быть отмечены зеленым цветом и словом UP

![image](https://user-images.githubusercontent.com/67442103/182324449-60940628-3310-451d-8f2f-bcfaa675aa80.png)


#### 2. Подключите grafana (дополнительно установите если у вас его нет) к prometheus

![image](https://user-images.githubusercontent.com/67442103/182325333-8574349b-56a7-4a31-93f9-eb16cb2dbc6c.png)

Так же для вас есть готовые дашборды для node exporter, postgres exporter

https://grafana.com/grafana/dashboards/1860

https://grafana.com/grafana/dashboards/9628

#### 3. Подключитесь к БД через dbeaver, у вас должны быть таблицы **users, books, users_books**

![image](https://user-images.githubusercontent.com/67442103/182324776-81cb0b8f-6eb5-4443-b98e-57d786172f3e.png)


#### 4. Отправьте запрос через swagger, расположен по адресу http://host-ip/swagger/index.html

#### 5. Отправьте запрос через Postman, коллекцию можете импортировать, она расположена в папке **test**
