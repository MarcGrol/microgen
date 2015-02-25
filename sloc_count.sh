#!/bin/sh

printLine() {
    printf "  %-20s: %s\n"  ${1} ${2}
}
countLinesInDir() {
    _DIR=${1}
    _LINE_COUNT=$(find ${_DIR} -name "*.go" -exec cat {} \; | wc -l | tr -d '[[:space:]]' )
    printLine ${_DIR} ${_LINE_COUNT} 
}

countGeneratedInDir() {
    _DIR=${1}
    _LINE_COUNT=$(find ${_DIR} -name "interface.go" -exec cat {} \; | wc -l | tr -d '[[:space:]]')
    printLine ${_DIR} ${_LINE_COUNT} 
}

countLogicInDir() {
    _DIR=${1}
    _LINE_COUNT=$(find ${_DIR} -name "logic.go" -exec cat {} \; | wc -l | tr -d '[[:space:]]')
    printLine ${_DIR} ${_LINE_COUNT} 
}

countTestInDir() {
    _DIR=${1}
    _LINE_COUNT=$(find ${_DIR} -name "logic_test.go" -exec cat {} \; | wc -l | tr -d '[[:space:]]')
    printLine ${_DIR} ${_LINE_COUNT} 
}

countSpecInDir() {
    _DIR=${1}
    _LINE_COUNT=$(find ${_DIR} -name "application.go" -exec cat {} \; | wc -l | tr -d '[[:space:]]')
    printLine ${_DIR}/application.go ${_LINE_COUNT} 
}

echo "\nTotal:"
countLinesInDir .

echo "\nLibraries:"
countLinesInDir ./dsl
countLinesInDir ./bus
countLinesInDir ./store
countLinesInDir ./myerrors
countLinesInDir ./tourApp/http
countLinesInDir ./tourApp/test
countLinesInDir ./tourApp/collector

echo "\nGenerated:"
countLinesInDir ./tourApp/events
countGeneratedInDir ./tourApp/tour
countGeneratedInDir ./tourApp/gambler
countGeneratedInDir ./tourApp/score

echo "\nBusiness logic:"
countSpecInDir . 

echo "\n  Tour:"
countLogicInDir ./tourApp/tour
countTestInDir ./tourApp/tour

echo "\n  Gambler:"
countLogicInDir ./tourApp/gambler
countTestInDir ./tourApp/gambler

echo "\n  Score:"
countLogicInDir ./tourApp/score/
countTestInDir ./tourApp/score/

echo ""
