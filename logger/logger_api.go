// logger api package
package logger

type Logger interface{

	Info(v ...interface{})
	Error(v ...interface{})
	Warn(v ...interface{})
	Debug(v ...interface{})
	Trace(v ...interface{})

}


