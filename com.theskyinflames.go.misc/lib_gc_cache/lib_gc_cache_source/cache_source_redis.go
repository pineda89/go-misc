package lib_gc_cache_source

import (
	EVENT "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event"
	FIFO "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_fifo"
	LOG "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_log"
	POOL "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_pool"
	POOL_SHARDER "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_pool/lib_gc_pool_sharder_adapter"
	RANDOMIZER "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_randomizer"
	UTIL "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_util"

	REDIS "github.com/fzzy/radix/redis"

	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

func init() {

}

var start_pipeline_loop_chan chan int

var ICacheRedis ICacheSource
var MAX_IDLE_CONNECTIONS int
var NETWORK string
var ADDRESS string
var PASSWD string
var PIPELINE_LIMIT int
var PIPELINE_LIMIT_BY_INTERVAL_OF_TIME int
var pipeline_chan chan *[]interface{}
var pipeline_finalized_chan chan int
var pipeline_ticker time.Ticker

var redismutex3 *sync.Mutex = &sync.Mutex{}

// Pool
var conPools []POOL.Pooler
var conPoolSharder POOL.PoolSharder_I

var IN_USE_CONNECTIONS int = 0
var AVAILABLE_CONNECTIONS int = 0
var shutdown bool = false

type redis_source struct{}

var Pipeline PipelineFlusher
var pipelineConnections map[string]*REDIS.Client

type PipelineAppener interface {
	AppendToPipeline() (map[string]int, error)
	GetKeys() *[][]byte
}

/*
	PIPELINED ADD DATA
*/
type PipelinedActionAddData struct {
	key   *[]byte
	value *[]byte
}

func (pAddData *PipelinedActionAddData) AppendToPipeline() (map[string]int, error) {
	if sharded_pool, err := conPoolSharder.GetPoolPerKey(pAddData.key); err != nil {
		return nil, err
	} else {
		c := pipelineConnections[sharded_pool.Name]
		if c == nil || pAddData == nil || pAddData.key == nil || pAddData.value == nil {
			msg, _ := EVENT.NotifyEvent("014-004", "", &[]string{"ADD", fmt.Sprint(c), fmt.Sprint("%v", *pAddData)})
			return nil, errors.New(msg)
		} else {
			c.Append("SET", *pAddData.key, *pAddData.value)
			z := make(map[string]int)
			z[sharded_pool.Name] = 1
			return z, nil
		}
	}
}
func (pAddData *PipelinedActionAddData) GetKeys() *[][]byte {
	return &[][]byte{*pAddData.key}
}

/*
	PIPELINED REMOVE DATA
*/
type PipelinedActionRemoveData struct {
	key *[]byte
}

func (pRemoveData *PipelinedActionRemoveData) AppendToPipeline() (map[string]int, error) {
	if sharded_pool, err := conPoolSharder.GetPoolPerKey(pRemoveData.key); err != nil {
		return nil, err
	} else {
		c := pipelineConnections[sharded_pool.Name]
		if c == nil || pRemoveData == nil || pRemoveData.key == nil {
			msg, _ := EVENT.NotifyEvent("014-004", "", &[]string{"REMOVE", fmt.Sprint(c), fmt.Sprint("%v", *pRemoveData)})
			return nil, errors.New(msg)
		} else {
			c.Append("DEL", *pRemoveData.key)
			z := make(map[string]int)
			z[sharded_pool.Name] = 1
			return z, nil
		}
	}
}
func (pRemoveData *PipelinedActionRemoveData) GetKeys() *[][]byte {
	return &[][]byte{*pRemoveData.key}
}

/*
	PIPELINED MSET
*/
type PipelinedActionAddMData struct {
	Data *[][]byte
}

func (pAddMData *PipelinedActionAddMData) AppendToPipeline() (map[string]int, error) {
	keys := make([][]byte, len(*pAddMData.Data)/2)
	z := 0
	_keys_values := make(map[string][]byte)
	zm := make(map[string]int)
	for i, _b := range *pAddMData.Data {
		if i%2 != 0 {
			keys[z] = _b
			_keys_values[string(_b)] = (*pAddMData.Data)[i+1]
			z += 1
		}
	}

	if pools_per_keys, err := conPoolSharder.GetPoolsPerKeys(&keys); err != nil {
		return nil, err
	} else {
		for _, v := range pools_per_keys {
			c := pipelineConnections[v.Name]
			_keys := make([][]byte, len(*v.Keys)*2)
			n := 0
			for _, _key := range *v.Keys {
				_keys[n] = _key
				_keys[n+1] = _keys_values[string(_key)]
				n += 2
			}
			zm[v.Name] = 1
			c.Append("MSET", _keys)
		}
	}

	return zm, nil
}

func (pAddMData *PipelinedActionAddMData) GetKeys() *[][]byte {
	z := 0
	keys := make([][]byte, len(*pAddMData.Data)/2)
	for i, _b := range *pAddMData.Data {
		if i%2 != 0 {
			keys[z] = _b
			z += 1
		}
	}
	return &keys
}

/*
	PIPELINED ADD TO HASH
*/
type PipelinedActionAddToHash struct {
	Hash  *string
	Key   *[]byte
	Value *[]byte
}

func (pAddToHash *PipelinedActionAddToHash) AppendToPipeline() (map[string]int, error) {
	_h := []byte(*pAddToHash.Hash)
	if sharded_pool, err := conPoolSharder.GetPoolPerKey(&_h); err != nil {
		return nil, err
	} else {
		c := pipelineConnections[sharded_pool.Name]
		if c == nil || pAddToHash == nil || pAddToHash.Hash == nil || pAddToHash.Key == nil || pAddToHash.Value == nil {
			msg, _ := EVENT.NotifyEvent("014-004", "", &[]string{"HSET", fmt.Sprint(c), fmt.Sprint("%v", *pAddToHash)})
			return nil, errors.New(msg)
		} else {
			c.Append("HSET", *pAddToHash.Hash, *pAddToHash.Key, *pAddToHash.Value)
			z := make(map[string]int)
			z[sharded_pool.Name] = 1
			return z, nil
		}
	}
}
func (pAddToHash *PipelinedActionAddToHash) GetKeys() *[][]byte {
	return &[][]byte{[]byte(*pAddToHash.Hash)}
}

/*
	PIPELINED REMOVE FROM HASH
*/
type PipelinedActionRemoveFromHash struct {
	Hash *string
	Key  *[]byte
}

func (pRemoveFromHash *PipelinedActionRemoveFromHash) AppendToPipeline() (map[string]int, error) {
	_h := []byte(*pRemoveFromHash.Hash)
	if sharded_pool, err := conPoolSharder.GetPoolPerKey(&_h); err != nil {
		return nil, err
	} else {
		c := pipelineConnections[sharded_pool.Name]
		if c == nil || pRemoveFromHash == nil || pRemoveFromHash.Hash == nil || pRemoveFromHash.Key == nil {
			msg, _ := EVENT.NotifyEvent("014-004", "", &[]string{"HSET", fmt.Sprint(c), fmt.Sprint("%v", *pRemoveFromHash)})
			return nil, errors.New(msg)
		} else {
			c.Append("HDEL", *pRemoveFromHash.Hash, *pRemoveFromHash.Key)
			z := make(map[string]int)
			z[sharded_pool.Name] = 1
			return z, nil
		}
	}
}
func (pRemoveFromHash *PipelinedActionRemoveFromHash) GetKeys() *[][]byte {
	return &[][]byte{[]byte(*pRemoveFromHash.Hash)}
}

/*
	PIPELINE FLUSHER
*/
type PipelineFlusher interface {
	AppendToPipeLine(appener PipelineAppener) error
	ManagePipeLineBySize() error
	ManagePipeLineByIntervalOfTime() error
	FlushPipeLine() error
}

/*
	PIPELINE CHANGES POOL
*/
type PipelinedChanges struct {
	PendingChanges_chan chan PipelineAppener
	PendingChanges      FIFO.Fifo
}

func (pc *PipelinedChanges) AppendToPipeLine(appener PipelineAppener) error {
	pc.PendingChanges_chan <- appener
	return nil
}
func (pc *PipelinedChanges) ManagePipeLineBySize() error {
	redismutex3.Lock()
	redismutex3.Unlock()

	ok := true
	for ok {
		select {
		case appener, ok := <-pc.PendingChanges_chan:
			if !ok || pc.PendingChanges.Len() >= int32(PIPELINE_LIMIT) {

				//println("*jas* execute pipeline by size, putting")
				pc.executePipeLine()

				if !ok {
					break
				}
			}
			pc.PendingChanges.Put(appener)
		}
	}

	return nil
}
func (pc *PipelinedChanges) ManagePipeLineByIntervalOfTime() error {
	pipeline_ticker := time.NewTicker(time.Duration(PIPELINE_LIMIT_BY_INTERVAL_OF_TIME) * time.Microsecond)
	for t := range pipeline_ticker.C {
		_ = t
		//println("*jas* execute pipeline by timing, putting")
		redismutex3.Lock()
		pc.executePipeLine()
		redismutex3.Unlock()
	}
	return nil
}
func (pc *PipelinedChanges) executePipeLine() []error {

	if !shutdown {
		pc_fifo := pc.PendingChanges
		pc.PendingChanges = FIFO.Fifo{}

		_f := &[]interface{}{func(pending_changes *FIFO.Fifo) []error {
			z := make(map[string]int)
			for k, _ := range pipelineConnections {
				z[k] = 0
			}
			for !pending_changes.Empty() {
				appener := pending_changes.Pop()
				if _z, err := appener.(PipelineAppener).AppendToPipeline(); err != nil {
					return []error{err}
				} else {
					for k, v := range _z {
						z[k] += v
					}
				}
			}

			var _err []error = nil
			for k, v := range z {
				for _z := 0; _z < v; _z++ {
					reply := pipelineConnections[k].GetReply()
					if reply.Err != nil {
						if _err == nil {
							_err = make([]error, 0)
						}
						_err = append(_err, reply.Err)
					}
				}
			}
			return _err
		}, &pc_fifo}
		pipeline_chan <- _f
	}
	return nil
}
func (pc *PipelinedChanges) FlushPipeLine() error {
	close(pc.PendingChanges_chan)
	pipeline_ticker.Stop()
	close(pipeline_chan)
	<-pipeline_finalized_chan
	return nil
}

/*
	lib_cache.CacheSource interface implementation
*/
func (s *redis_source) InitSource() error {

	// Take environment variables
	_MAX_IDLE_CONNECTIONS := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_CACHE_SOURCE__MAX_IDLE_CONNECTIONS", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_NETWORK := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_CACHE_SOURCE__NETWORK", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_STRING}
	_ADDRESS := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_CACHE_SOURCE__ADDRESS", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_STRING}
	_PASSWD := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_CACHE_SOURCE__PASSWD", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_STRING}
	_PIPELINE_LIMIT := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_CACHE_SOURCE__PIPELINE_LIMIT", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_PIPELINE_LIMIT_BY_INTERVAL_OF_TIME := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_CACHE_SOURCE__PIPELINE_LIMIT_BY_INTERVAL_OF_TIME", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	evs := []*UTIL.EnvironmentVariable{_MAX_IDLE_CONNECTIONS, _NETWORK, _ADDRESS, _PASSWD, _PIPELINE_LIMIT, _PIPELINE_LIMIT_BY_INTERVAL_OF_TIME}
	if _, err := UTIL.GetEnvironmentVariables(evs); err != nil {
		panic(err)
	} else {
		MAX_IDLE_CONNECTIONS = _MAX_IDLE_CONNECTIONS.Var_value.(int)
		NETWORK = _NETWORK.Var_value.(string)
		ADDRESS = _ADDRESS.Var_value.(string)
		PASSWD = _PASSWD.Var_value.(string)
		PIPELINE_LIMIT = _PIPELINE_LIMIT.Var_value.(int)
		PIPELINE_LIMIT_BY_INTERVAL_OF_TIME = _PIPELINE_LIMIT_BY_INTERVAL_OF_TIME.Var_value.(int)
	}

	// Instance the Pileline connections map
	pipelineConnections = make(map[string]*REDIS.Client)

	// Start Pipeline
	Pipeline = &PipelinedChanges{PendingChanges: FIFO.Fifo{}, PendingChanges_chan: make(chan PipelineAppener)}
	pipeline_chan = make(chan *[]interface{}, 2)
	pipeline_finalized_chan = make(chan int)
	start_pipeline_loop_chan = make(chan int)
	go func(start_pipeline_loop_chan *chan int) {
		for {
			<-*start_pipeline_loop_chan
			startPipelineLoop()
		}
		close(*start_pipeline_loop_chan)
	}(&start_pipeline_loop_chan)
	start_pipeline_loop_chan <- 0

	// Register the cache source
	ICacheRedis = &redis_source{}
	CacheSourceContainer.AddICacheSoure("REDIS", ICacheRedis)

	// Initialize sharded pools
	var err error
	if err = initShardedPools(); err != nil {
		return err
	} else {

		// Start
		go func(PipeLine PipelineFlusher) {
			Pipeline.ManagePipeLineBySize()
		}(Pipeline)
		go func(PipeLine PipelineFlusher) {
			Pipeline.ManagePipeLineByIntervalOfTime()
		}(Pipeline)
	}

	LOG.Info.Println("lib_cache_source (REDIS adapter) PIPELINE_LIMIT : ", PIPELINE_LIMIT)
	LOG.Info.Println("lib_cache_source (REDIS adapter) PIPELINE_LIMIT_BY_INTERVAL_OF_TIME : ", PIPELINE_LIMIT_BY_INTERVAL_OF_TIME)
	LOG.Info.Println("lib_cache_source (REDIS adapter) MAX_IDLE_CONNECTIONS : ", MAX_IDLE_CONNECTIONS)
	LOG.Info.Println("lib_cache_source (REDIS adapter) NETWORK : ", NETWORK)
	LOG.Info.Println("lib_cache_source (REDIS adapter) ADDRESS : ", ADDRESS)
	LOG.Info.Println("lib_cache_source (REDIS adapter) PASSWD : ", PASSWD)

	LOG.Info.Println("lib_cache_source (REDIS adapter) package initialized. ")
	fmt.Println("Init finish lib_cache_source_redis")

	return nil
}

func (s *redis_source) CloseSource() error {
	redismutex3.Lock()
	defer redismutex3.Unlock()

	if !shutdown {
		// Set shutdown flag
		shutdown = true

		// Finalize started pipelines
		if err := Pipeline.FlushPipeLine(); err != nil {
			return err
		} else {
			// Close the connections.
			if err := closeConnections(); err != nil {
				return err
			}
			return nil
		}
	} else {
		return nil
	}
}

func (source *redis_source) DeleteData(key *[]byte) (bool, error) {

	if err := checkConPool(); err != nil {
		return false, err
	} else {
		if !shutdown {
			appener := PipelinedActionRemoveData{key}
			addConnectionToPipelinedConnections(appener.GetKeys())
			Pipeline.AppendToPipeLine(&appener)
		}
	}
	return true, nil
}

func (source *redis_source) GetData(key *[]byte) (*[]byte, error) {
	if c, err := getCon(key); err != nil {
		return nil, err
	} else {
		defer releaseCon(c, key)

		if reply := c.Cmd("GET", *key); reply.Err != nil {
			return nil, reply.Err
		} else {
			l, _ := reply.Str()
			if len(l) == 0 {
				msg, _ := EVENT.NotifyEvent("013-013", "", &[]string{string(*key)})
				return nil, errors.New(msg)
			} else {
				if n, err := reply.Bytes(); err == nil {
					return &n, err
				} else {
					return nil, err
				}
			}
		}
	}
}

func (source *redis_source) MGetData(keys *[][]byte) (*[][]byte, error) {

	// Take the keys grouped by its pool
	if poolsPerKeys, err := conPoolSharder.GetPoolsPerKeys(keys); err != nil {
		return nil, err
	} else {
		wg := &sync.WaitGroup{}
		res := make([][]byte, 0)
		mutex := &sync.Mutex{}

		// Launch the mget against each pool
		errs_chan := make(chan error, 2)
		for _, poolPerKeys := range poolsPerKeys {

			wg.Add(1)
			go func(poolPerKeys *POOL.PoolPerKeys, res *[][]byte, wg *sync.WaitGroup, mutex *sync.Mutex) {
				defer wg.Done()

				// Take the connection from the pool
				poolable := <-poolPerKeys.PKPool.GetPool()

				// At ends, returns the connection to the pool.
				defer func(poolable *POOL.Poolable, poolPerKeys *POOL.PoolPerKeys) {
					poolPerKeys.PKPool.GetPool() <- poolable
				}(poolable, poolPerKeys)

				// Execute the mget
				c := poolable.Item.(*REDIS.Client)
				if reply := c.Cmd("MGET", *poolPerKeys.Keys); reply.Err != nil {
					errs_chan <- reply.Err
					return
				} else {
					l, _ := reply.List()
					if len(l) == 0 {
						msg, _ := EVENT.NotifyEvent("013-015", "", &[]string{""})
						errs_chan <- errors.New(msg)
						return
					} else {
						_res := make([][]byte, len(l))
						for i, _ := range l {
							_res[i] = []byte(l[i])
						}
						mutex.Lock()
						*res = append(*res, _res...)
						mutex.Unlock()
					}
				}

				// fires up the defer function execution.
				return
			}(poolPerKeys, &res, wg, mutex)
		}

		// Wait for all mget have finished
		wg.Wait()
		var err error = nil
		select {
		case err = <-errs_chan:
		default:
		}

		// Return the result.
		return &res, err
	}
}
func (s *redis_source) AddData(key, value *[]byte) error {
	if err := checkConPool(); err != nil {
		return err
	} else {
		if !shutdown {
			appener := PipelinedActionAddData{key, value}
			addConnectionToPipelinedConnections(appener.GetKeys())
			Pipeline.AppendToPipeLine(&appener)
		}
	}
	return nil
}

func (mp *redis_source) AddMData(data *[][]byte) error {

	if err := checkConPool(); err != nil {
		return err
	} else {
		if !shutdown {
			appener := PipelinedActionAddMData{data}
			addConnectionToPipelinedConnections(appener.GetKeys())
			Pipeline.AppendToPipeLine(&appener)
		}
		return nil
	}
}

func (s *redis_source) AddToHash(hash *string, key, value *[]byte) error {

	if err := checkConPool(); err != nil {
		return err
	} else {
		if !shutdown {
			appener := PipelinedActionAddToHash{hash, key, value}
			addConnectionToPipelinedConnections(appener.GetKeys())
			Pipeline.AppendToPipeLine(&appener)
		}
		return nil
	}
}

func (s *redis_source) GetKeysFromHash(hash *string) (*[][]byte, error) {
	_hash := []byte(*hash)
	if c, err := getCon(&_hash); err != nil {
		return nil, err
	} else {
		defer releaseCon(c, &_hash)

		if reply := c.Cmd("HKEYS", *hash); reply.Err != nil {
			return nil, reply.Err
		} else {
			_keys, err := reply.ListBytes()
			if len(_keys) == 0 {
				msg, _ := EVENT.NotifyEvent("014-004", "", &[]string{*hash, "n/a"})
				return nil, errors.New(msg)
			} else {
				return &_keys, err
			}
		}
	}
}

func (s *redis_source) RemoveKeyFromHash(hash *string, key *[]byte) error {

	if err := checkConPool(); err != nil {
		return err
	} else {
		if !shutdown {
			appener := PipelinedActionRemoveFromHash{hash, key}
			addConnectionToPipelinedConnections(appener.GetKeys())
			Pipeline.AppendToPipeLine(&appener)
		}
		return nil
	}
}

func (s *redis_source) RemoveHash(hash *string) error {
	//	c := pool.Get()
	//	defer c.Close()

	//	if _, err := c.Do("DEL", hash); err != nil {
	//		return err
	//	} else {
	//		return nil
	//	}
	return nil
}

//
//
//
func initShardedPools() error {
	// Building the pool of connections
	fmt.Println("Creating REDIS sharded pool")

	// Initializing the defined pools
	addresses := strings.Split(ADDRESS, ",")
	conPools := make([]*POOL.ShardedPool, len(addresses))
	var poolAlias string
	for i, address := range addresses {
		connections_pool := POOL.POOLMaker.GetPool(int32(MAX_IDLE_CONNECTIONS))

		if strings.Contains(address, ">>") {
			parts := strings.Split(address, ">>")
			poolAlias = parts[0]
			address = parts[1]
		} else {
			poolAlias = address
		}

		fmt.Printf("Will create %v REDIS connection for address %s with alias %v\n", MAX_IDLE_CONNECTIONS, address, poolAlias)
		for i := 0; i < MAX_IDLE_CONNECTIONS; i++ {
			if client, err := REDIS.Dial(NETWORK, address); err != nil {
				fmt.Printf("Error creating pool: %v\n", err)
				return err
			} else {
				connections_pool.GetPool() <- &POOL.Poolable{client}
				AVAILABLE_CONNECTIONS += 1
			}
		}
		conPools[i] = &POOL.ShardedPool{Name: poolAlias, PKPool: connections_pool, Weight: 1}
	}

	// Initalizes the pools sharder
	conPoolSharder, _ = POOL.PoolSharderFactory.GetPoolSharder(POOL_SHARDER.HASHRING_ADAPTER_NAME)
	if err := conPoolSharder.InitializeWithWeightedPools(conPools); err != nil {
		return err
	} else {
		fmt.Printf("Finally created %v connections on %v sharded pools of connections to REDIS\n", AVAILABLE_CONNECTIONS, len(addresses))
		return nil
	}
}

func checkConPool() error {
	if conPoolSharder == nil {
		msg, _ := EVENT.NotifyEvent("014-015", "", &[]string{})
		return errors.New(msg)
	} else {
		return nil
	}
}

func getCon(key *[]byte) (*REDIS.Client, error) {
	if err := checkConPool(); err != nil {
		return nil, err
	} else {
		if conPool, err := conPoolSharder.GetPoolPerKey(key); err != nil {
			return nil, err
		} else {
			poolable := <-conPool.PKPool.GetPool()
			c := poolable.Item.(*REDIS.Client)
			IN_USE_CONNECTIONS += 1
			return c, nil
		}
	}
}

func releaseCon(con *REDIS.Client, key *[]byte) error {

	if err := checkConPool(); err != nil {
		return err
	} else {
		if conPool, err := conPoolSharder.GetPoolPerKey(key); err != nil {
			return err
		} else {
			conPool.PKPool.GetPool() <- &POOL.Poolable{con}
			IN_USE_CONNECTIONS += -1
			return nil
		}
	}
}

func closeConnections() error {
	for _, pool := range conPools {

		// Close connection from pools
		var z int32
		cancel := false
		for z = 0; z < pool.GetSize() && cancel; z++ {
			select {
			case con := <-pool.GetPool():
				con.Item.(*REDIS.Client).Close()
			default:
				cancel = true
				break
			}
		}

		// close connections from pipeline
		for _, v := range pipelineConnections {
			v.Close()
		}
	}
	return nil
}
func addConnectionToPipelinedConnections(keys *[][]byte) error {
	if poolPerKeys, err := conPoolSharder.GetPoolsPerKeys(keys); err != nil {
		return err
	} else {
		for _, v := range poolPerKeys {
			if _, ok := pipelineConnections[v.Name]; !ok {
				poolable := <-v.PKPool.GetPool()
				pipelineConnections[v.Name] = poolable.Item.(*REDIS.Client)
			}
		}
		return nil
	}
}

func startPipelineLoop() {
	go func() {
		defer func() {

			if r := recover(); r != nil {

				// Restart the pipeline loop
				defer func(start_pipeline_loop_chan *chan int) {
					*start_pipeline_loop_chan <- 0
				}(&start_pipeline_loop_chan)

				t := time.Now()
				id := RANDOMIZER.GetRandom()
				msg := fmt.Sprintf("[%d], %s PANIC at %s : PANIC Defered recover: %v.\n", id, t, "cache_source_redis", r)
				fmt.Println(msg)

				// Capture the stack trace --
				buf := make([]byte, 10000)
				runtime.Stack(buf, false)

				msg = fmt.Sprintf("[%d], PANIC Stack Trace at %s : %s\n", id, "cache_source_redis", string(buf))
				fmt.Println(msg)
			}
		}()

		ok := true
		var _f *[]interface{}
		for ok {
			select {
			case _f, ok = <-pipeline_chan:
				if !ok {
					break
				}
				errs := (*_f)[0].(func(*FIFO.Fifo) []error)((*_f)[1].(*FIFO.Fifo))
				if errs != nil && len(errs) > 0 {
					for _, e := range errs {
						if e != nil {
							fmt.Println("Something web wrong on execute pipeline: ", e.Error())
						}
					}
				}
			}
		}
		pipeline_finalized_chan <- 0
	}()
}
