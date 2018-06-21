# handling-excel-files
An application on go that processes excel files

## Table of Contents
- [Install](#install)

## Install
Клонируем репозиторий проекта к себе на компьютер
```sh
$ git clone git@github.com:surrexi/parsing-excel-on-Go.git
```
Проходим в папку проекта
```sh
$ cd parsing-excel-on-Go
```
Далее копируем конфиг и настраиваем в нем подключение к БД
```sh
$ cp config.toml.example config.toml
```
Компилируем проект
```sh
$ go build
```
Запускаем проект
```sh
$ ./parsing-excel-on-Go
```