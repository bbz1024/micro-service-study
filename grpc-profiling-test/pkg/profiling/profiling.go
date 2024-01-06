package profiling

import (
	"github.com/pyroscope-io/client/pyroscope"
	"log"
	"runtime"
)

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: profiling
 * @Date:
 * @Desc: 性能分析，
 *
 */

const (
	ServiceAddress = "http://192.168.40.129:4040"
)

func InitProfiling(serviceName string) {
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: serviceName,
		ServerAddress:   ServiceAddress,
		Tags:            map[string]string{"hostname": "2632141215@qq.com"},
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})

	if err != nil {
		log.Fatalf("error starting profiling: %v", err)
		return
	}
}
