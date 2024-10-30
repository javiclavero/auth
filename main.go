package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	//"github.com/javiclavero/go-auth-service/ent" // Reemplaza con tu módulo real
	"github.com/javiclavero/go-auth-service/controllers"
)

func main() {

	http.HandleFunc("/login", controllers.LoginHandler)
	fmt.Sprintf("Servidor iniciado en http://localhost:8080")
	log.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	/*
		// Cargar variables de entorno desde .env
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		// Crear el DSN usando las variables de entorno
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)

		// Conectar a la base de datos
		client, err := ent.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("failed to open connection to mysql: %v", err)
		}
		defer client.Close()

		// Ejecutar la migración
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		// Crear un usuario de prueba
		createUser(client)
	*/
}

/*
func createUser(client *ent.Client) {
	ctx := context.Background()
	user, err := client.User.
		Create().
		SetUseUsername("example_user").
		SetUseEmail("example_user@example.com").
		SetUsePwd("securepassword").
		SetUseType(1).
		SetUseCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Printf("user created: %+v\n", user)
}
*/
