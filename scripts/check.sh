#! /bin/bash

#set -x
set -e

CFAIL='\033[0;31m'    # red for fail
COK='\033[0;32m'  # green for success
CVERB='\033[0;35m'  # magenta for verbose text
NC='\033[0m'        # No Color

verbose="0"
while getopts v flag
do
    case "${flag}" in
        v) verbose="1";;
        *) exit 1;;
    esac
done

for f in days/day[0-2][0-9]/main.go; do
    target=$(echo "${f}" | grep -oP 'day\d\d')
    day=$(echo "${f}" | grep -oP '\d\d')
    [ "$verbose" == "1" ] && echo -e "${CVERB}Check Day${day}:${NC}"

	run_target="run_${target}"
    output=$(make -B "${run_target}")
    result="$?"
    [ "$verbose" == "1" ] && echo "${output}"
    [ "${result}" != "0" ] && echo -e "${CFAIL}Day${day} failed!${NC}" || echo -e "${COK}Day${day} succeded!${NC}"
	[ "$verbose" == "1" ] && echo "" || :
done
