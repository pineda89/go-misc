package lib_gc_kafka_client

import (
	"sync"

	UTIL "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_util"

	"strings"
	"time"

	"github.com/Shopify/sarama"
)

func init() {

	// Take environment variables
	_BROKER_LIST := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__BROKER_LIST", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_STRING}
	_CHANNEL_BUFFERSIZE := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__CHANNELBUFFERSIZE", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_MAXOPENREQUESTS := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__MAXOPENREQUESTS", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_DIALTIMEOUT := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__DIALTIMEOUT", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_READTIMEOUT := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__READTIMEOUT", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_WRITETIMEOUT := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__WRITETIMEOUT", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_PRODUCER_MAXMESSAGESBYTES := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__PRODUCER_MAXMESSAGESBYTES", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_PRODUCER_REQUIREDACKS := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__PRODUCER_REQUIREDACKS", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_PRODUCER_TIMEOUT := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__PRODUCER_TIMEOUT", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_PRODUCER_COMPRESSION := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__PRODUCER_COMPRESSION", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_PRODUCER_RETRY_MAX := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__PRODUCER_RETRY_MAX", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_PRODUCER_RETRY_BACKOFF := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__PRODUCER_RETRY_BACKOFF", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_CONSUMER_MAXWAITTIME := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__CONSUMER_MAXWAITTIME", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_CONSUMER_MAXPROCESSINGTIME := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__CONSUMER_MAXPROCESSINGTIME", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_CONSUMER_FETCH_MIN_ := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__CONSUMER_FETCH_MIN", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_CONSUMER_FETCH_DEFAULT_ := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__CONSUMER_FETCH_DEFAULT", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	_CONSUMER_FETCH_MAX_ := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_KAKFA_CLIENT__CONSUMER_FETCH_MAX", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_INT64}
	evs := []*UTIL.EnvironmentVariable{_BROKER_LIST, _CONSUMER_FETCH_MIN_, _CONSUMER_FETCH_DEFAULT_, _CONSUMER_FETCH_MAX_, _CHANNEL_BUFFERSIZE, _CONSUMER_MAXPROCESSINGTIME, _CONSUMER_MAXWAITTIME, _DIALTIMEOUT, _MAXOPENREQUESTS, _PRODUCER_COMPRESSION, _PRODUCER_MAXMESSAGESBYTES, _PRODUCER_REQUIREDACKS, _PRODUCER_TIMEOUT, _READTIMEOUT, _WRITETIMEOUT, _PRODUCER_RETRY_MAX, _PRODUCER_RETRY_BACKOFF}
	if _, err := UTIL.GetEnvironmentVariables(evs); err != nil {
		panic(err)
	} else {
		broker_list = strings.Split(_BROKER_LIST.Var_value.(string), ",")
		channelbuffersize = int(_CHANNEL_BUFFERSIZE.Var_value.(int64))
		maxopenrequests = int(_MAXOPENREQUESTS.Var_value.(int64))
		dialtimeout = time.Duration(_DIALTIMEOUT.Var_value.(int64)) * time.Millisecond
		readtimeout = time.Duration(_READTIMEOUT.Var_value.(int64)) * time.Millisecond
		writetimeout = time.Duration(_WRITETIMEOUT.Var_value.(int64)) * time.Millisecond
		producer_maxmessagesbytes = int(_PRODUCER_MAXMESSAGESBYTES.Var_value.(int64))
		producer_requiredacks = sarama.RequiredAcks(_PRODUCER_REQUIREDACKS.Var_value.(int64))
		producer_timeout = time.Duration(_PRODUCER_TIMEOUT.Var_value.(int64)) * time.Millisecond
		producer_retry_backoff = time.Duration(_PRODUCER_RETRY_BACKOFF.Var_value.(int64)) * time.Millisecond
		producer_retry_max = int(_PRODUCER_RETRY_MAX.Var_value.(int64))
		consumer_maxwaittime = time.Duration(_CONSUMER_MAXWAITTIME.Var_value.(int64)) * time.Millisecond
		consumer_maxprocessingtime = time.Duration(_CONSUMER_MAXPROCESSINGTIME.Var_value.(int64)) * time.Millisecond
		consumer_fetch_min = int32(_CONSUMER_FETCH_MIN_.Var_value.(int64))
		consumer_fetch_default = int32(_CONSUMER_FETCH_DEFAULT_.Var_value.(int64))
		consumer_fetch_max = int32(_CONSUMER_FETCH_MAX_.Var_value.(int64))
	}

	// Set kafka client factory
	KafkaClientFactory = &kafkaClientFactory{}
}

const FIRST_TOPIC_PARTITION = -1

var KafkaClientFactory KafkaClientFactory_I
var client KafkaClient_I

var broker_list []string
var channelbuffersize int
var maxopenrequests int
var dialtimeout time.Duration
var readtimeout time.Duration
var writetimeout time.Duration
var producer_maxmessagesbytes int
var producer_requiredacks sarama.RequiredAcks
var producer_timeout time.Duration
var producer_compression sarama.CompressionCodec
var producer_retry_max int
var producer_retry_backoff time.Duration
var consumer_maxwaittime time.Duration
var consumer_maxprocessingtime time.Duration
var consumer_fetch_min int32
var consumer_fetch_default int32
var consumer_fetch_max int32

type MessageProcessor interface {
	ProcessMessage(message []byte, partition int32, offset int64) error
}

type KafkaClientFactory_I interface {
	GetNewClient() (KafkaClient_I, error)
}

type kafkaClientFactory struct{}

func (kcf *kafkaClientFactory) GetNewClient() (KafkaClient_I, error) {
	if client == nil || client.IsClosed() {
		client = &kafkaClient{closed: false, checkInitialized: &sync.Once{}, client_initialized: false}
	}
	return client, nil
}

type KafkaClient_I interface {
	Close() error
	IsClosed() bool
	StartAsyncProducer(topic string) (chan<- []byte, <-chan *error, error)
	StartSyncProducer(topic string) (chan<- []byte, <-chan *error, error)
	StartConsumer(message_processor MessageProcessor, topic string, partition int32, offset int64) (<-chan *error, error)
	GetPartitionsFromATopic(string) ([]int32, error)
	GetOldestOffset(topic string, partiionID int32) (int64, error)
	GetNewestOffset(topic string, partiionID int32) (int64, error)
	GetOffsetByPartition(topic string) (map[int32]int64, map[int32]int64, error)
}

type kafkaClient struct {
	sarama.Client
	client_initialized bool
	config             *sarama.Config
	async_producer     sarama.AsyncProducer
	sync_producer      sarama.SyncProducer
	consumer           sarama.Consumer
	shutdown           chan struct{}
	closed             bool
	checkInitialized   *sync.Once
}

func (kc *kafkaClient) IsClosed() bool {
	return kc.closed
}

func (kc *kafkaClient) initializeClient() error {

	// Make the config
	kc.config = sarama.NewConfig()
	kc.config.ChannelBufferSize = channelbuffersize
	kc.config.Net.MaxOpenRequests = maxopenrequests
	kc.config.Net.DialTimeout = dialtimeout
	kc.config.Net.ReadTimeout = readtimeout
	kc.config.Net.WriteTimeout = writetimeout
	kc.config.Producer.MaxMessageBytes = producer_maxmessagesbytes
	kc.config.Producer.RequiredAcks = producer_requiredacks
	kc.config.Producer.Timeout = producer_timeout
	kc.config.Producer.Compression = producer_compression
	kc.config.Consumer.MaxWaitTime = consumer_maxwaittime
	kc.config.Consumer.MaxProcessingTime = consumer_maxprocessingtime
	kc.config.Consumer.Fetch.Min = consumer_fetch_min
	kc.config.Consumer.Fetch.Default = consumer_fetch_default
	kc.config.Consumer.Fetch.Max = consumer_fetch_max
	kc.config.Producer.Retry.Max = producer_retry_max
	kc.config.Producer.Retry.Backoff = producer_retry_backoff

	// Make the client
	var err error
	if kc.Client, err = sarama.NewClient(broker_list, kc.config); err != nil {
		return err
	} else {

		if kc.consumer, err = sarama.NewConsumerFromClient(kc.Client); err != nil {
			return err

		} else {

			// Start shutdown channel
			kc.shutdown = make(chan struct{})

			// Set client initialization status
			kc.client_initialized = true

			return nil
		}
	}
}

func (kc *kafkaClient) check() error {
	var err error
	f := func() {
		if _err := kc.initializeClient(); _err != nil {
			err = _err
		}
	}
	kc.checkInitialized.Do(f)
	return err
}

func (kc *kafkaClient) Close() error {
	if kc.client_initialized {
		close(kc.shutdown)
	}
	kc.closed = true
	return nil
}

func (kc *kafkaClient) StartAsyncProducer(topic string) (chan<- []byte, <-chan *error, error) {

	// Check for kafka client initialization
	if err := kc.check(); err != nil {
		return nil, nil, err
	}

	// Crate the asynchronous producer
	var err error
	if kc.async_producer, err = sarama.NewAsyncProducerFromClient(kc.Client); err != nil {
		return nil, nil, err

	} else {

		// Create the errors channel
		errors_chan := make(chan *error, channelbuffersize)
		go func(producer_errors_chan <-chan *sarama.ProducerError, errors_chan chan<- *error) {
			for err := range producer_errors_chan {
				errors_chan <- &err.Err
			}
		}(kc.async_producer.Errors(), errors_chan)

		// Create the messages queue
		messages_queue := make(chan []byte, channelbuffersize)

		// Start the asynchronous producer
		go func(topic string, async_producer sarama.AsyncProducer, messages_queue <-chan []byte, shutdown chan struct{}) {
			for {
				select {
				case message, ok := <-messages_queue:
					{
						if !ok {
							return
						} else {
							async_producer.Input() <- &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(message)}
						}
					}
				case <-shutdown:
					return
				}
			}
		}(topic, kc.async_producer, messages_queue, kc.shutdown)

		return messages_queue, errors_chan, nil
	}
}

func (kc *kafkaClient) StartSyncProducer(topic string) (chan<- []byte, <-chan *error, error) {

	// Check for kafka client initialization
	if err := kc.check(); err != nil {
		return nil, nil, err
	}

	// Crate the asynchronous producer
	var err error
	if kc.sync_producer, err = sarama.NewSyncProducerFromClient(kc.Client); err != nil {
		return nil, nil, err

	} else {

		// Create the errors chan
		errors_chan := make(chan *error, channelbuffersize)

		// Create the messages queue
		messages_queue := make(chan []byte, channelbuffersize)

		// Start the asynchronous producer
		go func(topic string, sync_producer sarama.SyncProducer, messages_queue <-chan []byte, shutdown chan struct{}) {
			for {
				select {
				case message, ok := <-messages_queue:
					{
						if !ok {
							return
						} else {
							if _, _, err := sync_producer.SendMessage(&sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(message)}); err != nil {
								errors_chan <- &err
							}
						}
					}
				case <-shutdown:
					return
				}
			}
		}(topic, kc.sync_producer, messages_queue, kc.shutdown)

		return messages_queue, errors_chan, nil
	}
}

func (kc *kafkaClient) StartConsumer(message_processor MessageProcessor, topic string, partition int32, offset int64) (<-chan *error, error) {

	// Check for kafka client initialization
	if err := kc.check(); err != nil {
		return nil, err
	}

	// Crate the consumer
	var err error

	// Make errors chan
	errors_chan := make(chan *error)

	// Consume messages
	if partition == int32(FIRST_TOPIC_PARTITION) {
		if partitions, perr := kc.consumer.Partitions(topic); perr != nil {
			return nil, err
		} else {
			partition = partitions[0]
		}
	}
	if partition_consumer, err := kc.consumer.ConsumePartition(topic, partition, offset); err != nil {
		return nil, err
	} else {
		go func(partition_consumer sarama.PartitionConsumer, message_processor MessageProcessor, errors_chan chan *error, shutdown chan struct{}) {
			for {
				select {
				case message, ok := <-partition_consumer.Messages():
					if !ok {
						return
					}
					if err := message_processor.ProcessMessage(message.Value, message.Partition, message.Offset); err != nil {
						errors_chan <- &err
					}
				case <-shutdown:
					return
				}
			}
		}(partition_consumer, message_processor, errors_chan, kc.shutdown)
	}

	return errors_chan, nil
}

func GetOffsetOldestMark() int64 {
	return sarama.OffsetOldest
}

func GetOffsetNewestMark() int64 {
	return sarama.OffsetNewest
}

func (kc *kafkaClient) GetPartitionsFromATopic(topic string) ([]int32, error) {

	// Check for kafka client initialization
	if err := kc.check(); err != nil {
		return nil, err
	}

	return kc.consumer.Partitions(topic)
}

func (kc *kafkaClient) GetOldestOffset(topic string, partiionID int32) (int64, error) {
	// Check for kafka client initialization
	if err := kc.check(); err != nil {
		return -1, err
	}
	return kc.Client.GetOffset(topic, partiionID, GetOffsetOldestMark())
}

func (kc *kafkaClient) GetNewestOffset(topic string, partiionID int32) (int64, error) {
	// Check for kafka client initialization
	if err := kc.check(); err != nil {
		return -1, err
	}
	if lastOffset, err := kc.Client.GetOffset(topic, partiionID, GetOffsetNewestMark()); err != nil {
		return -1, err
	} else {
		return (lastOffset - 1), nil
	}
}

// Take the progresssive loading topic current offset.
func (kc *kafkaClient) GetOffsetByPartition(topic string) (map[int32]int64, map[int32]int64, error) {

	// Check for kafka client initialization
	if err := kc.check(); err != nil {
		return nil, nil, err
	}

	// Create the maps to put the newest/oldest offsets
	oldest := make(map[int32]int64)
	newest := make(map[int32]int64)

	// Take the number of partition for the topic
	if partitions, err := client.GetPartitionsFromATopic(topic); err != nil {
		return nil, nil, err
	} else {
		// Take the newest offset for each partition of the topic
		for partition := 0; partition < len(partitions); partition += 1 {
			// Take the newest offset
			if offset, err := client.GetNewestOffset(topic, partitions[partition]); err != nil {
				return nil, nil, err
			} else {
				newest[partitions[partition]] = offset
			}
			// Take the oldest offset
			if offset, err := client.GetOldestOffset(topic, partitions[partition]); err != nil {
				return nil, nil, err
			} else {
				oldest[partitions[partition]] = offset
			}
		}
	}
	return oldest, newest, nil
}

func CloseCurrentClient() error {
	if client != nil {
		return client.Close()
	} else {
		return nil
	}
}
