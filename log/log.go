package log

import (
	"log"
	"os"
)

var G_generalLogger *log.Logger = init_generalLogger()

func init_generalLogger() *log.Logger {
	// Flags: Create if needed, write only, remove contents. 0666 is read/write permission for everyone.
	file, err := os.OpenFile("general.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	var logger = log.New(file, "", 0)
	logger.SetFlags(log.Ltime | log.Lmicroseconds)
	return logger
}

func Init_positionLogger() *log.Logger {
	// Flags: Create if needed, write only, remove contents. 0666 is read/write permission for everyone.
	file, err := os.OpenFile("positions.csv", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	var logger = log.New(file, "", 0)
	logger.Println("time,id,x[cm],y[cm],theta[degrees] (the delimiter is a space)")
	logger.SetFlags(log.Ltime | log.Lmicroseconds)
	return logger
}