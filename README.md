# go-starter-kit

> Bu projeyi go diline yeni başlayanlar için hazırladım. Starter kit konusunda yıllarca çalıştığım için Go dilinde uzman kişiler içinde güzel bir boilerplate'dir.

## Projede neler var?

* Yüksek performanslı ve express.js'in Go diline uyarlanması olan [Fiber](https://gofiber.io) çatısını kullandım.
* Environment varible configurasyonu için popüler [Viper](https://github.com/spf13/viper) eklentisini kullandım.
* Cache kullanımı için [redis](https://github.com/go-redis/redis)
* Veritabanı yönetimi için popüler orm eklentisi [Gorm](https://gorm.io/)
* Rate limit için [Fiber Limiter](https://github.com/gofiber/limiter)
* Gzip, Brotli sıkıştırma için [Fiber Compression](https://github.com/gofiber/compression)
* Cors enable/disable için [Fiber Cors](https://github.com/gofiber/cors)
* config.json dosyası üzerinden proje configurasyonu
* Favicon disable/cache için [Fiber Favicon Middleware](https://github.com/gofiber/fiber/middleware). Performans artışı için kullanmanızı öneririm.
* Public klasörü ile html ve statik dosyalarınızı sunma imkanı.

## Kurulum için gerekli araçlar

* [golang](https://golang.org/)
* Çalışan bir [redis](https://redis.io/) server. Eğer config dosyanızdan redis cache'yi disable ediyorsanız kurmanıza gerek yok.

## Proje altyapısı

##### Folder Structure
```
config
└── database
    └── database.go
└── middlewares
    └── redis.go
└── providers
    └── redis
        └── redis.go 
└── routes
    └── routes.go 
└── config.json
controllers
└── todos.go
└── users.go
models
└── todo.go
└── user.go
public
└── index.html
└── favicon.go
.gitignore
go.mod
go.sum
main.go
README.md
```

## Kurulum

```
$ git clone git@github.com:ferdigokdemir/go-starter-kit.git
$ cd go-starter-kit
```

## Ayarlar
config/config.json dosyasını açın. Veritabanı ve redis host ve port configurasyonlarını ayarlayın. Eğer redisCache enabled: false ise redis ayarları yapmanız gerek yok.
```javascript
{
  "app": {
    "name": "go-starter-kit",
    "domain": "localhost",
    "port": 4141,
    "dev": true
  },
  "database": {
    "enabled": true,
    "mongodb": {
      "database": "icindenevar",
      "uri": "mongodb://lucy:mJCAbT3Vdc6I7bj0@icindenevar-shard-00-00.iby7b.mongodb.net:27017,icindenevar-shard-00-01.iby7b.mongodb.net:27017,icindenevar-shard-00-02.iby7b.mongodb.net:27017/icindenevar?ssl=true&replicaSet=atlas-6y38ji-shard-0&authSource=admin&retryWrites=true&w=majority"
    },
    "mysql": {
      "host": "localhost:3306",
      "user": "",
      "password": "",
      "database": ""
    },
    "redis": {
      "uri": "redis://h:pe6481c1f5f9a56da6cdd147661fd344189e2f81cc756d640f2a7a4ce6bc769a6@ec2-54-158-192-49.compute-1.amazonaws.com:22079"
    }
  },
  "compression": {
    "enabled": false // gzip compress default olarak kapalıdır açmak için true yapın
  },
  "cors": {
    "enabled": false // cors güvenlik için default kapalıdır. Production ortamında false yapın.
  },
  "prefork": {
    "enabled": false // Yüksek performans için açabilirsiniz.
  },
  "helmet": {
    "enabled": true // Http güvenliği için default değeri açıktır.
  },
  "logger": {
    "enabled": false // Api'ye gelen requestleri console'da görmek için true yapın
  },
  "ratelimit": {
    "enabled": false // Brute force ataklarına karşı enable yapabilirsiniz.
  },
  "redisCache": {
    "enabled": true // Veritabanından çektiğiniz verileri daha hızlı sunmak için redis'e kaydedin.
  },
  "favicon": {
    "ignore": false, // Tarayıcıdan apinize gelen requestlerde, extra favicon isteğide atılır. Buda gereksiz disk okuması demektir. Ignore true yaparsanız performans artışı sağlar. Veya cache true yapın.
    "cache": true
  }
}

```

## Çalıştır
``` 
go run main.go
```

## Build

```
go build main.go
```

## License
MIT
