#!/bin/bash
function do_make()
{
    if [ "$1" == "clean" ]; then
        rm -rf *.o main
    else
        set -x
        g++ -std=c++11 -I${problem_name} -c -o main.o main.cpp && g++ main.o -o main
        set +x
    fi
}

function main()
{
    if [ "$1" == "clean" ]; then
        # make clean
        do_make clean
    else
        problem_name="${1%/}" # remove tailing /
        if [ "${problem_name}" == "" ]; then
            problem_name=42-trapping-rain-water
        elif [ "${problem_name}" != "$(tail -n 4 $0 | awk -F'=' '{print $2}' | awk -F'#' '{print $1}')" ]; then
            # make clean
            do_make clean
        fi
        if [ ! -d "${problem_name}" ]; then
            echo "warn: problem \"${problem_name}\" not exists"
            exit 1
        fi
        # make problem_name=42-trapping-rain-water${problem_name}
        do_make
    fi
    sed -i "s#problem_name=42-trapping-rain-water#problem_name=${problem_name}#" $0
}

main $1