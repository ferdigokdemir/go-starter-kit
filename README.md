# go-starter-kit

> Bu projeyi go diline yeni başlayanlar için hazırladım. Starter kit konusunda yıllarca çalıştığım için Go dilinde uzman kişiler içinde güzel bir boilerplate'dir.

## Projede neler var?

* Yüksek performanslı ve express.js'in Go diline uyarlanması olan [Fiber](https://gofiber.io) çatısını kullandım.
* Environment varible configurasyonu için popüler [Viper](https://github.com/spf13/viper) eklentisini kullandım.
* Cache kullanımı için [redis](github.com/go-redis/redis)
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

Repoyu Klonlayın:

```
$ git clone git@github.com:ferdigokdemir/go-starter-kit.git
$ cd go-starter-kit
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
