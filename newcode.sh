#!/bin/bash
trap 'echo "interrupted ... "; do_exit' INT

function do_exit()
{
    rm -rf "${problem_dir}"
    exit 1
}

function check_empty()
{
    for var in $1; do
        if [ "$(eval echo '$'$var)" == "" ]; then
            echo "warn: $var is empty"
            do_exit
        fi
    done
}

function main()
{
    cd $(dirname $0)
    if [ "$1" == "--help" -o "$1" = "" ]; then
        echo `usage: $0 "problem name"`
        exit 0
    fi

    problem_name=$(echo $1 | tr ' ' '-' | tr '[:upper:]' '[:lower:]')
    if [ "$(find . -name "*${problem_name}")" != "" ]; then
        echo "warn: problem ${problem_name} exists"
        exit 1
    fi

    echo "creating ${problem_name} ..."

    # cd "${problem_dir}"
    # ln -s ../Makefile Makefile

    # description contains ' as &#39; and " as &quot;
    echo "curling ..."
    problem="$1"
    url="https://leetcode.com/problems/${problem_name}/description/" # without qid
    html=$(curl -L "$url")
    code=$(echo "${html}" | sed -nr "/codeDefinition/ {s/.*'C\+\+', 'defaultCode': '([^']+).*/\1/;p;q}")
    description=$(echo "${html}" | sed -nr "/\"description\"/,/\" \/>/ {s/(.*content=\")|(\" \/>)//;s/&quot;/\"/g;s/&#39;/'/g;p}")
    difficulty=$(echo "${html}" | sed -nr '/Difficulty:/ {s/.*label-([a-z,A-Z]+).*/\1/;p;q}')
    qid=$(echo "${html}" | grep questionId | grep -oP "\d+")
    problem_name="${qid}-${problem_name}" # rename with qid
    problem_dir="$(pwd)/${problem_name}"

    check_empty "html code description difficulty qid"

    mkdir ${problem_name} || exit 1

    # write readme.md
    while read line
    do
        line=${line//#/\\#}
        line=${line//\`/\\\`}
        eval echo -e "$line"
    done < "tpl/readme.md" > "${problem_name}/readme.md"

    # write default solution code solution.cpp, 4-space indent by default
    indent="    "
    echo -ne "${code}" | sed "/^};$/ i\\\n${indent}void main() {\n${indent}${indent}\n${indent}}" > "${problem_name}/solution.cpp"

    # always make new problem
    export problem_name
    ./make.sh clean

    echo "finish"
}

main $1