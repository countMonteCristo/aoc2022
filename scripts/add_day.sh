#!/bin/bash

set -e
#set -x

DAY=$1

TEMPLATE_SRC="dayXX.go"
DAYXX="day${DAY}"
DAY_DIR="days/${DAYXX}"
SRC="${DAY_DIR}/main.go"
COMMON="common.go"
MAKEFILE="Makefile"
INPUTS="inputs/${DAYXX}"
INPUT_FILE="${INPUTS}/input.txt"
TEST_FILE="${INPUTS}/test1.txt"

# If template source file does not exists - this is an error
[ -e "${TEMPLATE_SRC}" ] || { echo "[ERROR] Template ${TEMPLATE_SRC} does not exist"; exit 1; }

# Check if new target is not present in Makefile
[ -e "${MAKEFILE}" ] || { echo "[ERROR] ${MAKEFILE} does not exist"; exit 1; }
output=$(grep -c "${TARGET}: ${SRC}" "${MAKEFILE}"  || true)
if [[ "$output" != "0" ]]; then
    echo "[ERROR] Target ${TARGET} is present"
    exit 1
fi

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

# Add new targets to Makefile
{
    echo ""
    echo "${DAYXX}: ${SRC}"
    echo "	\${GO} build -o bin/\$@ ${SRC} ${DAY_DIR}/${COMMON}"
} >> ${MAKEFILE}
{
	echo ""
	echo "run_${DAYXX}: ${DAYXX}"
	echo "	./bin/${DAYXX}"
} >> ${MAKEFILE}
echo "[INFO] Succesfully add targets [${DAYXX}, run_${DAYXX}] to ${MAKEFILE}"

cd "${DAY_DIR}"
ln -s ../../${COMMON} ${COMMON}
