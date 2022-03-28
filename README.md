# Artifact Generator tg-Bot

> ### Генератор магических и не очень предметов, предположительно для игр вроде D&D — добро пожаловать в ридми.

### Предисловие
В качестве предисловия, хочу сказать, что не являюсь фанатом D&D-подобных игр ~~хотя и не без этого~~, но интересуюсь обработкой естесственного языка. Изучив некоторые языковые модели, я обратил внимание, что большинство из них «слишком серьёзны» для маленьких задач, которые проще выполнить «механическим» кодом без ML.
Этот пет-проект написан мной для демонстрации моей идеи по облегчению работы с русским языком для разработчиков, локализаторов, геймдизайнеров и т.д. Поэтому больше внимания здесь уделяется обработке прилагательных.

### Об идее
Если конкретнее, то сама идея заключается в разработке паттерна, который позволяет создавать огромный список уникальных предметов на русском языке, как в Diablo. Основная сложность этой задачи в том, что в русском языке для словообразования активно используются аффиксы. К сожалению, я не нашёл интересного способа создания/локализации слов с аффиксами в этой или подобных играх, кроме как варианта с большими таблицами со всеми формами слова. Учитывая количество ошибок с несогласованными словосочетаниями в генерируемых названиях, в этой области есть над чем поработать.


### О проекте 
#### ЗАДАЧА: написать телеграм-бот, который будет генерировать рандомные предметы, соединяя элементы заранее подготовленного текста в трёх гугл-таблицах 
Бот должен уметь генерировать 3 типа предметов: простые, необычные и редкие. Также он должен уметь выдавать список из 3 случайных предметов любого типа.
- Простые предметы имеют только название предмета и краткое описание.
- Необычные имеют магический атрибут, стоящий после названия предмета, и дополнительное предложение с описанием предмета, которое соответствует этому атрибуту.
- И, наконец, редкие предметы. Помимо предыдущего атрибута, имеют ещё один — прилагательное, стоящее перед названием.

#### РЕШЕНИЕ
Предмет генерируется на основе нескольких таблиц, отношения которых похожи на реляционные. Таблицы были незначительно модифицированы и представлены csv файлами — их легко выгружать из гугл-таблиц, которыми так часто пользуются геймдизайнеры (по крайней мере те, что мне попадались). Всего 4 таблицы:
1. Основная таблица с колонками для названий предметов и указанием грамматического рода. У каждого предмета может быть любое количество описаний, как правило их три.
2. Таблица с атрибутами, стоящими после названия предмета и описаниями атрибута.
3. Таблица атрибутов-прилагательных, что стоят перед названиями. Все атрибуты лемматизированы (оставлена только неизменяемая часть слова, н-р: волшебн), чтобы не хранить все возможные склонения по родам и не подбирать подходящие. Также есть колонки с указанием на тип склонения прилагательного и на описание атрибута.
4. Последняя таблица имеет в себе только окончания для прилагательных. На основе грамматического рода предмета из первой таблицы и на склонении из третьй подбирается правильное окончание и присоединяется к лемме атрибута. Работает безотказно, но лучше себя покажет при больших объёмах данных :) 

Выглядит это так:
![](https://github.com/hexhowk/Artifact-Generator-tg-Bot/blob/master/screen.png)


### Заключение 
Паттерн хорошо себя показывает, но его поведение ещё не фиксировано. Не исключено, что подобный или более продвинутый паттерн уже где-то используется (и является охраняемой тайной). Я продолжаю развивать свою идею за рамками этого проекта и возможно со временем мне попадётся информация о реализации чего-то подобного, но пока что встречалось не совсем то же самое.
