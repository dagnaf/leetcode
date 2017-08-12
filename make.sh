if [ "$1" == "clean" ]; then
    make clean
else
    problem_name="${1%/}" # remove tailing /
    if [ "${problem_name}" == "" ]; then
        problem_name=44-wildcard-matching
    elif [ "${problem_name}" != "$(tail -n 1 $0 | awk -F'=' '{print $2}' | awk -F'/' '{print $1}')" ]; then
        make clean
    fi 
    if [ "$(find . -name "${problem_name}")" == "" ]; then
        echo "warn: problem \"${problem_name}\" not exists"
        exit 1
    fi
    make problem_name=${problem_name}
fi
sed -i "s/problem_name=44-wildcard-matching/problem_name=${problem_name}/" $0