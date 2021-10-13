# Trilateración
## _desarrollo realizado con _  - [Trilateracion_](https://es.wikipedia.org/wiki/Trilateraci%C3%B3n)



# DB

- Ingrese la informacion de la base de datos en .env y las cordenadas de los satelites (en este caso no es neceario configurar la db)
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


# ejecuta la app
```sh
 go mod init meradolibre.com/q/v2
 go mod tidy
 rm q & go build
 ./q
 
 o simplemente con > go run .
```
# para consumir el servicio por POST
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

# debo mencionar que para realizar el ejercicio me base en este [proyecto](https://github.com/bryanpaluch/trilateration-go/blob/master/main.go)
