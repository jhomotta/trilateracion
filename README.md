# Trilateración
## _desarrollo realizado con _  - [Trilateracion_](https://es.wikipedia.org/wiki/Trilateraci%C3%B3n)



# DB

- Enter the database information in .env y las cordenadas de los satelites (en este caso no es neceario configurar la db)
```sh
 user="root"
 pass="root"
 host="localhost"
 port="33060"
 db_name="quasar"
 
 kenobi_x = 3
 kenobi_y = 5
 skywalker_x = -2
 skywalker_y = -3
 sato_x = 4
 sato_y = -4
```


# ejecutat la app
```sh
 go mod init meradolibre.com/q/v2
 go mod tidy
 rm q & go build
 ./q
 
 o simplemente con > go run .
```
# para consumir el servicio 
http://localhost:9000/api/satellite
```json
{
	"satellite": [{
			"name": "kenobi",
			"distance": 7.62,
			"message": [
				"este",
				"",
				"",
				"mensaje",
				""
			]
		},
		{
			"name": "skywalker",
			"distance": 5.385,
			"message": [
				"",
				"es",
				"",
				"",
				"secreto"
			]
		},
		{
			"name": "sato",
			"distance": 10,
			"message": [
				"este",
				"",
				"un",
				"",
				""
			]
		}
	]
}
```

