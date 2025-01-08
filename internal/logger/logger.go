package logger

type Field struct {
	Key   string
	Value interface{}
}

func NewField(key string, value interface{}) Field {
	return Field{
		Key:   key,
		Value: value,
	}
}

type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
}

// func Get() *log.Logger {
// 	f, err := os.OpenFile("log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
// 	if err != nil {
// 		log.Println("can't set logger")
// 		return nil
// 	}
// 	return log.New(f, "", log.Ldate|log.Ltime)
// }

// func Debug(v ...any) {
// 	Get().Println(v...)
// }
