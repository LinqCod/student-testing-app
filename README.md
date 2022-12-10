# Тестирование Студентов
## РТУ МИРЭА - курсовая работа
### Описание:

Изначальная цель реализации проекта - курсовая работа по предмету "Разработка серверных частей интернет ресурсов". Представляет из себя веб-приложение, специализирующееся на тестировании студентов и отслеживании их успеваемости. Студентам доступен список предметов курса, в каждом из которых доступны задания по различным разделам этого предмета. В качестве мотивации у каждого студента имеется показатель успеваемости. Все преподаватели имеют доступ к их предмету и ведомым группам, включая возможность модификации заданий. Легкий и понятный интерфейс, в котором способен разобраться и начать развиваться каждый учащийся.
> Приложение в состоянии разработки. В дальнейшем доведется до ума в качестве pet проекта
***
### Стек технологий:

- `Golang` - бекенд

  + `gin` - фреймворк для разработки API
  + `viper` - библиотека для парсинга конфиг файлов
  + `jwt-go` - библиотека для реализации аутентификации
  + `pq` - postgres драйвер
  + `bcrypt` - библиотека для хеширования
  
- `JS/HTML/CSS` - фронтенд
- `Postgres` - базы данных
- `Docker` - контейнеризация
- `Docker Compose` - развертывание приложения
- `Swagger` - документация api

***
### UML диаграмма моделей:

![Screenshot from 2022-12-10 00-26-30](https://user-images.githubusercontent.com/58244765/206798916-20f4d514-f71e-41e2-b8b2-b8c2448c101e.png)
***
### API:

#### Students:
```
 - /api/v1/students/ | POST | регистрация студента
 - /api/v1/students/login | POST | авторизация студента
 - /api/v1/students/:student_id | GET | получение данных студента
 - /api/v1/students/:student_id | PUT | изменение данных студента
```
#### Subjects:
```
 #Students
 - /api/v1/subjects/:group | GET | получение списка предметов группы
```
#### Tasks:
```
 #Students
 - /api/v1/tasks/:subject_id | GET | получение списка разделов задач по предмету
 - /api/v1/tasks/categories/:category_id | GET | получение списка задач по id раздела
```
