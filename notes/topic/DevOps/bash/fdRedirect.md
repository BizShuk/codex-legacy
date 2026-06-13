# FD redirect

<https://wiki.bash-hackers.org/howto/redirection_tutorial>

### open fd

`exec [fd]< <filename>`
`exec [fd]> <filename>`

### fd meaning

0, stdin
1, stdout
2, stderr
3~9, extra need to open first

### <>

`exec 3<> <filename>`, put file content to fd 3 and also open output to 3

### redirection before command

`echo output > tmp` == `>tmp echo output`

### truncates file to zero-length

`> <filename>` or `:> <filename>`

### Sequence matter

`badcommand >tmp 2>&1`, write stderr to tmp file != `badcommand 2>&1 >tmp` didn't write stderr to tmp file

### close fd

- `[fd]>&-` , close output fd
- `[fd]<&-` , close input fd

### socket

`exec 3 <> /dev/<protocol>/<host>/<port>`
