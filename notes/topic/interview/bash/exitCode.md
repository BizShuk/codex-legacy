# bash exit Code

```
($(exit 1); echo $?>&3| echo 5) 3>&1 | (read xs; echo xs:$xs)
# result: xs:1

(($(exit 1); echo $?>&3)| echo 5) 3>&1 | (read xs; echo xs:$xs)
# result: xs:5
```
