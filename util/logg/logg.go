package logg

import (
	"fmt"
	"runtime"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/converse"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/datetime"
)

func Printlogger_JsonMarshalIndent(header, prefix string, data interface{}) {
	if header == "" {
		header = "\t\t ********** DEBUGGER Struct JSON Indent **********"
	}
	if prefix == "" {
		prefix = "Result"
	}
	Printlogger(header, prefix, converse.JsonMarshalIndent(data))
}

func Printlogger_Variadic_JsonMarshalIndent(header, prefix string, data ...interface{}) {
	if header == "" {
		header = "\t\t ********** DEBUGGER Variadic JSON Indent **********"
	}
	if prefix == "" {
		prefix = "Result"
	}
	Printlogger(header, prefix, converse.JsonMarshalIndent(data))
}

func Printlogger(header, prefix string, data interface{}) {
	if header == "" {
		header = "\t\t ********** DEBUGGER Printlogger **********"
	}
	if prefix == "" {
		prefix = "Result"
	}
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("\t\t\t %v \n", header)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("\t", datetime.GetCurrentDateTimeNano())
	fmt.Printf("\t %v :: %v \n", prefix, fmt.Sprintf("%v", data))
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")
	fmt.Println()

}

func Printlogger_Variadic(header, prefix string, data ...interface{}) {
	if header == "" {
		header = "\t\t ********** DEBUGGER Printlogger_Variadic **********"
	}
	if prefix == "" {
		prefix = "Result"
	}
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("\t\t\t %v \n", header)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("\t", datetime.GetCurrentDateTimeNano())
	fmt.Printf("\t %v :: %v \n", prefix, data)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")
	fmt.Println()

}

func GetCallerPathNameFileNameLineNumber() string { //(pathName, fileName string, lineNumber int) {
	pc, fileName, lineNumber, _ := runtime.Caller(1)

	return fmt.Sprintf("Path in %s( Click this to go to %s:%d )", runtime.FuncForPC(pc).Name(), fileName, lineNumber)

}
