package flon

import (
	"fmt"

	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

var zlog, tracer = logging.PackageLogger("eos-go", "github.com/fullon-labs/flon-go")

type logStringerFunc func() string

func (f logStringerFunc) String() string { return f() }

func typeField(field string, v interface{}) zap.Field {
	return zap.Stringer(field, logStringerFunc(func() string {
		return fmt.Sprintf("%T", v)
	}))
}
