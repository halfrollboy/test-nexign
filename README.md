# Тестовое задание для Nexign

## Описание
Данное решение выполняет поиск ошибок в получаемом тексте и выводит уже исправленный текст


### Чтобы сделать запрос к сервису требуется: 
Сделать **POST** запорос на базовый endpoint `localhost:8080\`
в качестве **body** он ожидает структуру типа **map[string][]string**
```
{
    "texts": [
        "Время лесных малышей",
        "Пришло теплое лето. ",
        "На лисной опушки распускаюца колоколчики, незабутки, шыповник."
    ],
}
```
Ответ будет сохранять последовательность и структуру полученного json но с исправленным текстом

