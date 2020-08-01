# go-netflow2squid

## Tnaslator from netflow v5 format to default squid log format

### How To Build

If you do not have a golang, please use the [installation instruction](https://golang.org/doc/install).

and next

    git clone https://github.com/Rid-lin/go-netflow2squid.git
    cd go-netflow2squid
    make

### How to use

    flow-cat ft-v05.YYYY-MM-DD.HHMMSS+Z | flow-print -f 5 | go-netflow2squid > access_netflow.log

After that, you need to use any **analyzer of squid** logs.

    Extra options:
    -gmt - sets the time zone in the format "+0300", by default "+0500".
    -year - specifies which year to use for time conversion, by default 2020

### Thanks

- [Translation of Squid documentation](http://break-people.ru)
- [Very simple log format description](https://wiki.enchtex.info/doc/squidlogformat)

-------------------------------------------------

### Это транслятор из формата netflow v5 в формат логов Squid по-умолчанию

### Как установить

Если у Вас не установлен Golang, пожалуйста воспользуйтесь [инструкцией по установке](https://golang.org/doc/install)

и далее

    git clone https://github.com/Rid-lin/go-netflow2squid.git
    cd go-netflow2squid
    make

### Как использовать

    flow-cat ft-v05.YYYY-MM-DD.HHMMSS+Z | flow-print -f 5 | go-netflow2squid > access_netflow.log

После этого можно использовать любой **анализатор логов Squid-а**.

    Дополнительные параметры:
    -gmt - устанавливает часовой пояс в формате "+0300", по-умолчанию "+0500".
    -year - указывает какой год использовать для перевода времени, по-умолчанию 2020

### Благодарности

- [Перевод документации по Squid](http://break-people.ru/)
- [Очень простое описание формата логов](https://wiki.enchtex.info/doc/squidlogformat)
