package customerror

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

func GetCustomStackTrace(err error) errors.StackTrace {
	var fs errors.StackTrace
	st, ok := err.(interface{ StackTrace() errors.StackTrace })
	if !ok {
		return fs
	}

	frames := st.StackTrace()

	for _, frame := range frames {
		pc := uintptr(frame)
		fun := runtime.FuncForPC(pc)
		f, _ := fun.FileLine(pc)

		// 不要なStackトレースが多すぎるのでフィルタリング
		if strings.Contains(f, "memoir-backend") || strings.Contains(f, "pkg/error") {
			if !strings.Contains(f, "usecase/customerror") {
				fs = append(fs, frame)
			}
		}
	}

	if os.Getenv("APP_ENV") == "local" {
		// エラー出力
		fmt.Print("-----------------------\n")
		fmt.Printf("■ error: %+v\n", err.Error())
		fmt.Printf("■ stack trace: %+v\n", fs)
		fmt.Print("-----------------------\n")
	}

	return fs
}

type RequestError struct {
	Code    string
	Message string
}

func (re RequestError) Error() string {
	return re.Message
}
