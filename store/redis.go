package store

type RedisInterface interface {
	Get()
	Set()
	Del()
	Purge()
}

type Options struct {
	Addr     string
	Password string
	Database string
}

type Redis struct {
	Client  RedisInterface
	options *Options
}

func setUpRedis() {

}
