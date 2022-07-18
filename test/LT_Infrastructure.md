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
- **go**
```console
sudo apt install golang
```
- **postgres**
```console
sudo apt install postgresql postgresql-contrib
```
- **prometheus**

### 4. Настройте СУБД
1. **postgresql.conf** и **pg_hba.conf**

Найдите директорию **main**, в ней хранятся конфиг файлы постгреса, опычно путь до ней такой:
```console
ls /etc/postgresql/14/main
```

Отредактируйте файл postgresql.conf
```console
sudo nano /etc/postgresql/14/main/postgresql.conf
```
Вам нужно изменить поле listen_addresses как показано на скриншоте

![image](https://user-images.githubusercontent.com/67442103/179586118-38a83072-a811-4faf-8a1f-ab6bfc0ad9c9.png)

Отредактируйте файл pg_hba.conf
```console
sudo nano /etc/postgresql/14/main/pg_hba.conf
```
Вам нужно в конце файла добавить следующую строку

![image](https://user-images.githubusercontent.com/67442103/179594575-53865a80-f7be-4df8-aa0c-b9853ce45a12.png)

2. Перезагрузите Postgres чтобы применились все изменения

Статус работы сервиса можете посмотреть с помощью следующей команды
```console
systemctl status postgres
```
Перезагрузите следующим образом
```console
sudo systemctl restart postgresql
```

3. Измените пароль для пользователя **postgres**

Подключиться к СУБД через терминал можно с помощью следующей команды

```console
sudo -u postgres psql
```

Поменяйте пароль с помощью следующего **SQL** запроса

![image](https://user-images.githubusercontent.com/67442103/179604894-3163dde7-41e8-40d3-8270-7cf91a0e9a07.png)

### 5. Создайте Новую Базу данных

Для этого подключитесь к СУБД с помощью **DBeaver**

![image](https://user-images.githubusercontent.com/67442103/179607910-8ede41c1-3973-42c8-b0b8-ebbfe0463e07.png)
