# Trap

`trap <Command> <Signal Type>`
<https://www.linuxjournal.com/content/bash-trap-command>

### Run something important, no Ctrl-C allowed

trap "" SIGINT
important_command

### Less important stuff from here on out, Ctrl-C allowed

trap - SIGINT
not_so_important_command
