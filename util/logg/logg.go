package logg

import (
	"fmt"
	"runtime"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/converse"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/datetime"
)

func PrintloggerJsonMarshalIndentHasHeader(header, prefix string, data interface{}) {
	if header == "" {
		header = "\t\t ********** DEBUGGER Struct JSON Indent **********"
	}
	if prefix == "" {
		prefix = "Result"
	}
	PrintloggerHasHeader(header, prefix, converse.JsonMarshalIndent(data))
}

func PrintloggerVariadicJsonMarshalIndentHasHeader(header, prefix string, data ...interface{}) {
	if header == "" {
		header = "\t\t ********** DEBUGGER Variadic JSON Indent **********"
	}
	if prefix == "" {
		prefix = "Result"
	}
	PrintloggerVariadicHasHeader(header, prefix, converse.JsonMarshalIndent(data))
}

func PrintloggerHasHeader(header, prefix string, data interface{}) {
	if header == "" {
		header = " ********** DEBUGGER Printlogger **********"
	}
	if prefix == "" {
		prefix = "Result"
	}
	fmt.Printf("%v | %v | %v :: %+v \n", header, datetime.GetCurrentDateTimeNano(), prefix, fmt.Sprintf("%s", data))

}

func PrintloggerVariadicHasHeader(header, prefix string, data ...interface{}) {
	if header == "" {
		header = " ********** DEBUGGER Printlogger_Variadic **********"
	}
	if prefix == "" {
		prefix = "Result"
	}
	fmt.Printf("%v | %v | %v :: %+v \n", header, datetime.GetCurrentDateTimeNano(), prefix, data)
}

func PrintloggerJsonMarshalIndent(prefix string, data interface{}) {
	if prefix == "" {
		prefix = "Result"
	}
	Printlogger(prefix, converse.JsonMarshalIndent(data))
}

func PrintloggerVariadicJsonMarshalIndent(prefix string, data ...interface{}) {
	if prefix == "" {
		prefix = "Result"
	}
	PrintloggerVariadic(prefix, converse.JsonMarshalIndent(data))
}

func PrintloggerHasDateTime(prefix string, data interface{}) {
	if prefix == "" {
		prefix = "Debugger logger"
	}
	fmt.Printf("Date Time %v :: %v %v \n", datetime.GetCurrentDateTimeNano(), prefix, fmt.Sprintf("%+v", data))
}

func PrintloggerVariadicHasDateTime(prefix string, data ...interface{}) {
	if prefix == "" {
		prefix = "Debugger logger"
	}
	fmt.Printf("Date Time %v :: %v %v \n", datetime.GetCurrentDateTimeNano(), prefix, fmt.Sprintf("%+v", data))

}

func Printlogger(prefix string, data interface{}) {
	if prefix == "" {
		prefix = "Debugger logger"
	}
	fmt.Printf("%v %v \n", prefix, fmt.Sprintf("%+s", data))
}

func PrintloggerVariadic(prefix string, data ...interface{}) {
	if prefix == "" {
		prefix = "Debugger logger"
	}

	fmt.Printf("%v %v \n", prefix, fmt.Sprintf("%+s", data))

}

func GetCallerPathNameFileNameLineNumber() string { //(pathName, fileName string, lineNumber int) {
	pc, fileName, lineNumber, _ := runtime.Caller(1)

	return fmt.Sprintf("Path in %s( Click this to go to %s:%d )", runtime.FuncForPC(pc).Name(), fileName, lineNumber)

}

func PrintloggerDebuggerHasHeader(header, prefix string, data interface{}) {
	if header == "" {
		header = "\t\t ********** DEBUGGER Printlogger **********"
	}
	if prefix == "" {
		prefix = "Result"
	}
	fmt.Println()
	fmt.Println("=============================================================================================================================")
	fmt.Printf("\t\t\t %v \n", header)
	fmt.Println("=============================================================================================================================")
	fmt.Println("\t", datetime.GetCurrentDateTimeNano())
	fmt.Printf("\t %v :: %+s \n", prefix, fmt.Sprintf("%s", data))
	fmt.Println("=============================================================================================================================")
	fmt.Println("=============================================================================================================================")
	fmt.Println()

}

func PrintloggerDebuggerVariadicHasHeader(header, prefix string, data ...interface{}) {
	if header == "" {
		header = "\t\t ********** DEBUGGER Printlogger_Variadic **********"
	}
	if prefix == "" {
		prefix = "Result"
	}
	fmt.Println()
	fmt.Println("=============================================================================================================================")
	fmt.Printf("\t\t\t %v \n", header)
	fmt.Println("=============================================================================================================================")
	fmt.Println("\t", datetime.GetCurrentDateTimeNano())
	fmt.Printf("\t %v :: %+v \n", prefix, data)
	fmt.Println("=============================================================================================================================")
	fmt.Println("=============================================================================================================================")
	fmt.Println()

}

func PrintloggerTopupLineDouble() {
	fmt.Println()
	fmt.Println("=============================================================================================================================")
}

func PrintloggerUnderLineDouble() {
	fmt.Println("=============================================================================================================================")
	fmt.Println()
}

func PrintloggerTopupLineSingle() {
	fmt.Println()
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")
}

func PrintloggerUnderLineSingle() {
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")
	fmt.Println()
}

func PrintloggerTopupLineStar() {
	fmt.Println()
	fmt.Println("*****************************************************************************************************************************")
}

func PrintloggerUnderLineStar() {
	fmt.Println("*****************************************************************************************************************************")
	fmt.Println()
}
