package main

import (
	CACHE "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_cache"
	CONF "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_conf"
	CONFIG "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_config"
	CONFIGURABLE_FROM_ENV "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_configurable_from_env"
	CONFIGURABLE_FROM_JSON "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_configurable_from_json"
	CONTAINER "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_container"
	EVENT "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event"
	FTS_HELPER "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_fts_helper"
	IDX "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_idx"
	KAFKA_CLIENT "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_kafka_client"
	LOG "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_log"
	METRICS "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_metrics_endpoint"
	PANIC_CATCHING "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_panic_catching"
	POOL "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_pool"
	RESTFUL "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_restful"
	SEMAPHORE "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_semaphore"
	STATUS "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_status"
	TIMEOUT_WRAPPER "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_timeout_wrapper"
	//UTIL "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_util"
	MESSAGES_COLLECTOR "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_messages_collector"

	"expvar"
	"time"
)

func init() {

	// expvar variables declaration
	counter = expvar.NewInt("myCounter")
}

// expvar varaibles
var counter *expvar.Int

func main() {
	defer PANIC_CATCHING.PanicCatching("main")

	_ = CONFIG.PtrConfigProvider
	_ = CONF.Dummy
	_ = CONFIGURABLE_FROM_JSON.Configurable{}
	_ = CONTAINER.GenericContainerFactory
	_ = EVENT.Event{}
	_ = POOL.Poolable{}
	_ = SEMAPHORE.SemaphoreFactory
	_ = TIMEOUT_WRAPPER.TimeoutWrapper
	_ = CACHE.DOCache
	_ = IDX.DOIndexesManager
	_ = RESTFUL.Dummy
	_ = METRICS.Dummy
	_ = KAFKA_CLIENT.KafkaClientFactory
	_ = FTS_HELPER.H_FtsMessageMaker
	_ = MESSAGES_COLLECTOR.Dummy
	_ = CONFIGURABLE_FROM_ENV.Dummy

	LOG.Info.Println("MDH Common library compiled !!")

	STATUS.StartStatusExposer()

	for {
		counter.Add(1)
		time.Sleep(500 * time.Millisecond)
	}
}
