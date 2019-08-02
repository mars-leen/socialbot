package logger

import (
	"fmt"
	"os"
)

func LoadLog() error {
	return nil
}

func Info(v ...interface{})  {
	fmt.Printf("[INFO] %+v \n", v)
}

func Infof(f string, v ...interface{})  {
	fmt.Printf("[INFO] " + f + "\n", v...)
}


func Warn(v ...interface{})  {
	fmt.Printf("[WARN] %+v \n", v)
}

func Warnf(f string, v ...interface{})  {
	fmt.Printf("[WARN] " + f  + "\n", v...)
}

func Error(v ...interface{})  {
	fmt.Printf("[ERROR] %+v \n", v)
}

func Errorf(f string, v ...interface{})  {
	fmt.Printf("[ERROR] " + f + "\n", v...)
}


func Fatal(v ...interface{}) {
	fmt.Printf("[FATAL] %+v \n", v)
	os.Exit(1)
}


