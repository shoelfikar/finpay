package middleware

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func loggingError(message string, code string) {
	lg := log.New(os.Stdout, "", 0)
	logMsg := fmt.Sprintf("[%s] \033[31m[ERROR]\033[0m -- message : %s | code : %s", time.Now().Format("2006-01-02 15:04:05"), message, code)

	// Output the log message using the standard logger
	lg.Output(2, logMsg)
}

func loggingInfo(message any, code string) {
	lg := log.New(os.Stdout, "", 0)
	logMsg := fmt.Sprintf("[%s] \033[34m[INFO]\033[0m -- message : %+v | duration : %s", time.Now().Format("2006-01-02 15:04:05"), message, code)

	// Output the log message using the standard logger
	lg.Output(2, logMsg)
}


func UnaryLoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	
	duration := time.Since(start)
    if err != nil {
        st, _ := status.FromError(err)
		loggingError(st.Message(), st.Code().String())
        // log.Printf("[ERROR] <-- gRPC error: %s | Code: %s | Duration: %s", st.Message(), st.Code(), duration)
    } else {
		loggingInfo(resp, duration.String())
        // log.Printf("[INFO] <-- gRPC response: %+v | Duration: %s", resp, duration)
    }

	return resp, err
}