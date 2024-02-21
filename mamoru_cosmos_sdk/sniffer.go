package mamoru_cosmos_sdk

import (
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer/cosmos"
	"github.com/cometbft/cometbft/libs/log"
)

var snifferConnectFunc = cosmos.CosmosConnect

func InitConnectFunc(f func() (*cosmos.SnifferCosmos, error)) {
	snifferConnectFunc = f
}

//

func init() {

	mamoru_sniffer.InitLogger(func(entry mamoru_sniffer.LogEntry) {
		kvs := mapToInterfaceSlice(entry.Ctx)
		msg := "Mamoru core: " + entry.Message
		var tmLogger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
		switch entry.Level {
		case mamoru_sniffer.LogLevelDebug:
			tmLogger.Debug(msg, kvs...)
		case mamoru_sniffer.LogLevelInfo:
			tmLogger.Info(msg, kvs...)
		case mamoru_sniffer.LogLevelWarning:
			tmLogger.With("Warn").Error(msg, kvs...)
		case mamoru_sniffer.LogLevelError:
			tmLogger.Error(msg, kvs...)
		}
	})
}

func mapToInterfaceSlice(m map[string]string) []interface{} {
	var result []interface{}
	for key, value := range m {
		result = append(result, key, value)
	}

	return result
}

type SnifferI interface {
	CheckRequirements() bool
	Client() *cosmos.SnifferCosmos
}

var _ SnifferI = (*Sniffer)(nil)

type Sniffer struct {
	mu     sync.Mutex
	logger log.Logger
	client *cosmos.SnifferCosmos
}

func NewSniffer(logger log.Logger) *Sniffer {
	return &Sniffer{
		logger: logger,
	}
}

func (s *Sniffer) isSnifferEnable() bool {
	val, ok := os.LookupEnv("MAMORU_SNIFFER_ENABLE")
	if !ok {
		return false
	}

	isEnable, err := strconv.ParseBool(val)
	if err != nil {
		s.logger.Error("Mamoru Sniffer env parse error", "err", err)
		return false
	}

	return isEnable
}

func (s *Sniffer) CheckRequirements() bool {
	return s.isSnifferEnable() && s.connect()
}

func (s *Sniffer) Client() *cosmos.SnifferCosmos {
	return s.client
}

func (s *Sniffer) connect() bool {
	if s.client != nil {
		return true
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	var err error
	s.client, err = snifferConnectFunc()
	if err != nil {
		erst := strings.Replace(err.Error(), "\t", "", -1)
		erst = strings.Replace(erst, "\n", "", -1)
		s.logger.Error("Mamoru Sniffer connect", "err", erst)
		return false
	}

	return true
}
