#! /bin/bash


function match_uppercase() {
    typeset -a data=("$@")
    typeset -i count=0
    for e in ${data[@]}; do
        if [[ $e =~ [A-Z] ]]; then
            count=$(($count + 1))
        fi
    done
    echo $count
}

aaa=(onE tWo tHree four)


match_uppercase "${aaa[@]}"


#if [[ "aBc" =~ [A-Z] ]]; then
#    echo a
#fi