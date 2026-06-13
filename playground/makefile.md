# Makefile


[GNU make manual](https://www.gnu.org/software/make/manual/html_node/index.html)

### option
- `-f <makefile name>`, specific makefile


### all
default make progress target with sequence. If you want to add other make command , check [.PHONY](#.PHONY)

`all: clean Composer.o Database.o Database.out`



### target
```
<target>: <parameter1> <parameter2> ...
    operations ...
```

### variable
- `$@` , label
- `$+` , parameters


### .PHONY
use label as make command. i.e., `.PHONY clean, mrproper`


### clean

##### c
`rm -f *.o <excutable file>`


