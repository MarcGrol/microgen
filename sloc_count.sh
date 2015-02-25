#!/bin/sh

printLineCount() {
    printf "  %-40s: %4s loc %s\n"  ${1} ${2} ${3}
}

printFileCount() {
    printf "  %-40s: %4s files %s\n"  ${1} ${2} ${3}
}

countFilesInDir() {
    _DIR=${1}
    _EXTRA=${2}
    _LINE_COUNT=$(find ${_DIR} -name "*" | wc -l | tr -d '[[:space:]]')
    printFileCount ${_DIR} ${_LINE_COUNT} "${_EXTRA}"
}

countGoFilesInDir() {
    _DIR=${1}
    _EXTRA=${2}
    _LINE_COUNT=$(find ${_DIR} -name "*.go" | wc -l | tr -d '[[:space:]]')
    printFileCount ${_DIR} ${_LINE_COUNT} "${_EXTRA}"
}

countLinesInDir() {
    _DIR=${1}
    _EXTRA=${2}
    _LINE_COUNT=$(find ${_DIR} -name "*.go" -exec cat {} \; | wc -l | tr -d '[[:space:]]' )
    printLineCount ${_DIR} ${_LINE_COUNT} "${_EXTRA}"
}

countGeneratedLinesInDir() {
    _DIR=${1}
    _EXTRA=${2}
    _LINE_COUNT=$(find ${_DIR} -name "interface.go" -exec cat {} \; | wc -l | tr -d '[[:space:]]')
    printLineCount ${_DIR}/interface.go ${_LINE_COUNT} "${_EXTRA}"
    _LINE_COUNT=$(find ${_DIR} -name "start.go" -exec cat {} \; | wc -l | tr -d '[[:space:]]')
    printLineCount ${_DIR}/start.go ${_LINE_COUNT} "${_EXTRA}"
}

countLogicLinesInDir() {
    _DIR=${1}
    _EXTRA=${2}
    _LINE_COUNT=$(find ${_DIR} -name "logic.go" -exec cat {} \; | wc -l | tr -d '[[:space:]]')
    printLineCount ${_DIR}/logic.go ${_LINE_COUNT} 
    _LINE_COUNT=$(find ${_DIR} -name "logic_test.go" -exec cat {} \; | wc -l | tr -d '[[:space:]]')
    printLineCount ${_DIR}/logic_test.go ${_LINE_COUNT} "${_EXTRA}"
}

countSpecLinesInDir() {
    _DIR=${1}
    _LINE_COUNT=$(find ${_DIR} -name "application.go" -exec cat {} \; | wc -l | tr -d '[[:space:]]')
    printLineCount ${_DIR}/application.go ${_LINE_COUNT} 
}

echo "\nTotal:"
countLinesInDir .
countGoFilesInDir .

echo "\nLibraries:"
countLinesInDir ./dsl
countLinesInDir ./bus
countLinesInDir ./store
countLinesInDir ./myerrors
countLinesInDir ./tourApp/http
countLinesInDir ./tourApp/test
countLinesInDir ./tourApp/infra
countLinesInDir ./tourApp/collector

echo "\nGenerated contracts:"
countLinesInDir ./tourApp/events "(gen)"
countFilesInDir ./tourApp/doc "(gen)"

echo "\nBusiness logic:"
countSpecLinesInDir .

echo "\n  Tour:"
countGeneratedLinesInDir ./tourApp/tour "(gen)"
countLogicLinesInDir ./tourApp/tour

echo "\n  Gambler:"
countGeneratedLinesInDir ./tourApp/gambler "(gen)"
countLogicLinesInDir ./tourApp/gambler

echo "\n  Score:"
countGeneratedLinesInDir ./tourApp/score "(gen)"
countLogicLinesInDir ./tourApp/score/

echo ""
