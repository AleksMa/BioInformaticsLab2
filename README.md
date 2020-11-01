## Запуск
Как запускать:
```bash
go build -o lab2 *.go
./lab2 [--gap <gap_value>] [--mode {simple|dnafull|blosum62}] [--out <output_file] <input_file_1> [<input_file_2>]
```
`input_file_2` может быть опущен - в этом случае обе последовательности должны находиться в первом файле.

Параметр `mode` может иметь значение `simple`, `dnafull` либо `blosum62`, где `simple` - режим работы со скоринговыми параметрами "совпадение +1, различие -1", оставшиеся режимы соответствуют одноименным скоринговым таблицам. Таблицы были сгенерированы посредством скрипта `gen_table/gen_blosum.go`.

## Тесты
Тесты находятся в файле `engine_test`.

Запуск: 
```bash
go test .
```
Рассматриваются различные размеры последовательностей РНК, проверяется правильность постройки глобального выравнивания и скора. 
