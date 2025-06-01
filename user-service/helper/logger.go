package helper

import (
	"fmt"
	"log"
	"os"
	"time"
)


func LoggingError(message string) {
	lg := log.New(os.Stdout, "", 0)
	logMsg := fmt.Sprintf("[%s] \033[31m[ERROR]\033[0m  %s", time.Now().Format("2006-01-02 15:04:05"), message)

	// Output the log message using the standard logger
	lg.Output(2, logMsg)
}

func LoggingInfo(message string) {
	lg := log.New(os.Stdout, "", 0)
	logMsg := fmt.Sprintf("[%s] \033[34m[INFO]\033[0m  %s", time.Now().Format("2006-01-02 15:04:05"), message)

	// Output the log message using the standard logger
	lg.Output(2, logMsg)
}