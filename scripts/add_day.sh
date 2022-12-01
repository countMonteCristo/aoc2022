#!/bin/bash

set -xe

DAY=$1

TEMPLATE_SRC="dayXX.go"
DAYXX="day${DAY}"
DAY_DIR="days/${DAYXX}"
SRC="${DAY_DIR}/main.go"
#MAKEFILE="Makefile"
INPUTS="inputs/${DAYXX}"
INPUT_FILE="${INPUTS}/input.txt"
TEST_FILE="${INPUTS}/test1.txt"

# If template source file does not exists - this is an error
[ -e "${TEMPLATE_SRC}" ] || { echo "[ERROR] Template ${TEMPLATE_SRC} does not exist"; exit 1; }

# Check if new target is not ipresent in Makefile
#[ -e "${MAKEFILE}" ] || { echo "[ERROR] ${MAKEFILE} does not exist"; exit 1; }
#output=$(grep -c "${TARGET}: ${SRC}" "${MAKEFILE}"  || true)
#if [[ "$output" != "0" ]]; then
#    echo "[ERROR] Target ${TARGET} is present"
#    exit 1
#fi

# Check if new source file does not exist
[ -e "${SRC}" ] && { echo "[ERROR] ${SRC} exists"; exit 1; } || echo "[INFO] Create source file ${SRC}..."

# Create empty input file for source
[ -d "${INPUTS}" ] || mkdir "${INPUTS}"
[ -f "${INPUT_FILE}" ] || { touch "${INPUT_FILE}"; echo "[INFO] Successfuly add input file ${INPUT_FILE}"; }
[ -f "${TEST_FILE}" ] || { touch "${TEST_FILE}"; echo "[INFO] Successfuly add input test file ${TEST_FILE}"; }

# Create source file
mkdir -p "${DAY_DIR}"
cp "${TEMPLATE_SRC}" "${SRC}"
sed -i -- "s/@@dayXX@@/${DAYXX}/g" "${SRC}"
echo "[INFO] Successfuly add new source file ${SRC}"

# Add new target to Makefile
#{
#    echo ""
#    echo "${TARGET}: ${SRC}"
#    echo "	\${CXX} \${CXXFLAGS} ${SRC} -o ${TARGET}"
#} >> ${MAKEFILE}
#echo "[INFO] Succesfully add target ${TARGET} to ${MAKEFILE}"
