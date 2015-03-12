#!/bin/sh

printLineCount() {
    printf "  %-40s: %4s loc %s\n"  ${1} ${2} ${3}
}

printFileCount() {
    printf "  %-40s: %4s files %s\n"  ${1} ${2} ${3}
}

printHighlighted() {
    printf "\e[1;34m"
    printf  "  %-40s: %4s loc %s\n"  ${1} ${2} ${3}
    printf "\e[0m"
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
     printHighlighted ${_DIR}/logic.go ${_LINE_COUNT} 
    _LINE_COUNT=$(find ${_DIR} -name "logic_test.go" -exec cat {} \; | wc -l | tr -d '[[:space:]]')
     printHighlighted ${_DIR}/logic_test.go ${_LINE_COUNT} "${_EXTRA}"
}

countSpecLinesInDir() {
    _DIR=${1}
    _LINE_COUNT=$(find ${_DIR} -name "application.go" -exec cat {} \; | wc -l | tr -d '[[:space:]]')
    printHighlighted ${_DIR}/application.go ${_LINE_COUNT} 
}

echo "\nTotal:"
countLinesInDir .
countGoFilesInDir .

echo "\nTools and libraries:"
countLinesInDir ./tool/dsl
countLinesInDir ./lib/myerrors
countLinesInDir ./lib/test
countLinesInDir ./lib/envelope
countLinesInDir ./infra/bus
countLinesInDir ./infra/store
countLinesInDir ./infra/myhttp

echo "\nApplication specification:"
countSpecLinesInDir .

echo "\n  Contracts:"
countLinesInDir ./tourApp/events "(gen)"
countFilesInDir ./tourApp/doc "(gen)"

echo "\n  Tour-service:"
countGeneratedLinesInDir ./tourApp/tour "(gen)"
countLogicLinesInDir ./tourApp/tour

echo "\n  Gambler-service:"
countGeneratedLinesInDir ./tourApp/gambler "(gen)"
countLogicLinesInDir ./tourApp/gambler

echo "\n  News-service:"
countGeneratedLinesInDir ./tourApp/news "(gen)"
countLogicLinesInDir ./tourApp/news

echo "\n  Notification-service:"
countGeneratedLinesInDir ./tourApp/notification "(gen)"
countLogicLinesInDir ./tourApp/notification

echo "\n  Collector-service:"
countGeneratedLinesInDir ./tourApp/collector
countLogicLinesInDir ./tourApp/collector

echo "\n  Proxy-service:"
countLinesInDir ./tourApp/proxy

echo "\n  Provisioning-tool:"
countLinesInDir ./tourApp/client
countLinesInDir ./tourApp/prov

echo "\n  Web ui:"
countFilesInDir ./tourApp/ui

echo ""
