package log

import "log"

func LogJson(logType,controllerName, functionName, description, errorDetail string)  {
	errorMessage:= `{
						"Type":` + logType + `,
						"Controller_Name":` + controllerName + `
						"Function_Name":` + functionName + `,
						"Description":` + description + `,
						"Error_Detail":` + errorDetail + `
					}`
	log.Fatalln(errorMessage)
}