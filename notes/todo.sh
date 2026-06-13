#!/bin/bash

function genTODO_auto {
        echo "# TODO" >&3
        echo >&3
        egrep -R TODO . \
                --exclude-dir=tmp \
                --exclude-dir=.git \
                --exclude=todo.sh \
                --exclude=TODO*.md | \
        sed -E 's|(\./[^:]*):.*\[TODO\] *(.*)|- [LINK](\1)  \2|g' >&3
        echo >&3
}

function genASK_auto {
        echo "# ASK" >&3
        echo >&3
        egrep -R ASK . \
                --exclude-dir=tmp \
                --exclude-dir=.git \
                --exclude=todo.sh \
                --exclude=TODO*.md | \
        sed -E 's|(\./[^:]*):.*\[ASK\] *(.*)|- [LINK](\1)  \2|g' >&3
        echo >&3
}

exec 3> TODO_auto.md
genTODO_auto
genASK_auto
exec 3>&-