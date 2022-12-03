package tool

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"server/conf"
)

//
// RabbitMQ
//  @Description: RabbitMQ结构体
//
type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string //  队列名称
	Exchange  string //  交换机
	Key       string //  Key
	url       string //  连接信息
}

//
// NewRabbitMQ
//  @Description: 创建RabbitMQ实例
//  @param queueName 队列名称
//  @param exchange 交换机
//  @param key
//  @return *RabbitMQ
//
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	host := conf.Cfg.Section("RABBITMQ").Key("host").String()
	port := conf.Cfg.Section("RABBITMQ").Key("port").String()
	user := conf.Cfg.Section("RABBITMQ").Key("user").String()
	password := conf.Cfg.Section("RABBITMQ").Key("password").String()
	url := "amqp://" + user + ":" + password + "@" + host + ":" + port + "/"

	rabbitMQ := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		url:       url,
	}

	var err error
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.url)
	rabbitMQ.failOnErr(err, "创建RabbitMQ连接错误")
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	rabbitMQ.failOnErr(err, "获取Channel失败")
	return rabbitMQ
}

//
// failOnErr
//  @Description: 错误处理
//  @receiver r RabbitMQ
//  @param err 错误
//  @param message 信息
//
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

//
// Close
//  @Description: 关闭
//  @receiver r RabbitMQ
//
func (r *RabbitMQ) Close() {
	err := r.channel.Close()
	if err != nil {
		Logger.Error(err.Error())
		return
	}
	err = r.conn.Close()
	if err != nil {
		Logger.Error(err.Error())
		return
	}
}

var RabbitMQSimpleLogger *RabbitMQ

//
// init
//  @Description: 初始化
//
func init() {
	RabbitMQSimpleLogger = NewRabbitMQSimpleLogger("Logger")
	RabbitMQSimpleLogger.ConsumeSimpleLogger()
}

//
// NewRabbitMQSimpleLogger
//  @Description: 创建简单模式RabbitMQ实例来处理异步写日志
//  @param queueName
//  @return *RabbitMQ
//
func NewRabbitMQSimpleLogger(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

//
// PublishSimpleLogger
//  @Description: 简单模式下生产日志信息
//  @receiver r RabbitMQ
//  @param message 信息
//
func (r *RabbitMQ) PublishSimpleLogger(level LogLevel, message string) {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, //  是否持久化
		false, //  是否自动删除
		false, //  是否具有排他性
		false, //  是否阻塞
		nil,   //  额外属性
	)
	if err != nil {
		Logger.Error(err.Error())
	}

	data := LogTemp{
		level,
		message,
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		Logger.Error(err.Error())
	}

	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, //  如果为true 根据exchange类型和routeKey规则 如果无法找到符合条件的队列 那么会把发送的消息返回给发送者
		false, //  如果为true 当exchange发送消息到队列侯发现队列上没有绑定消费者 则会把消息发还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        dataBytes,
		},
	)
	if err != nil {
		Logger.Error(err.Error())
		return
	}
}

//
// ConsumeSimpleLogger
//  @Description: 简单模式下消费日志信息
//  @receiver r RabbitMQ
//
func (r *RabbitMQ) ConsumeSimpleLogger() {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, //  是否持久化
		false, //  是否自动删除
		false, //  是否具有排他性
		false, //  是否阻塞
		nil,   //  额外属性
	)
	if err != nil {
		Logger.Error(err.Error())
	}

	message, err := r.channel.Consume(
		r.QueueName,
		"",    //  用来区分多个消费者
		true,  //  是否自动应答
		false, //  是否具有排他性
		false, //  如果设置为true 表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false, //  是否阻塞
		nil,
	)
	if err != nil {
		Logger.Error(err.Error())
	}

	go func() {
		for m := range message {
			//  实现逻辑函数
			var logTemp LogTemp
			err := json.Unmarshal(m.Body, &logTemp)
			if err != nil {
				Logger.Warning(err.Error())
				continue
			}

			Logger.Log(logTemp.Level, logTemp.Message)
		}
	}()
}
