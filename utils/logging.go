package utils

import (
    "os"
    "io"
    "log"
)

func LoggingSettings(logfile string) {
    logFile, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln(err)
    }
    mutilLogFile := io.MultiWriter(os.Stdout, logFile)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.SetOutput(mutilLogFile)
}
