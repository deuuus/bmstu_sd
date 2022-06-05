package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kirktriplefive/labsMed"
	"github.com/kirktriplefive/labsMed/pkg/handler"
	"github.com/kirktriplefive/labsMed/pkg/repository"
	"github.com/kirktriplefive/labsMed/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	repos_d := repository.NewDoctorRepository(db)
	repos_r := repository.NewRecordRepository(db)

	services := service.NewServicePatient(repos, repos_d, repos_r)

	handler := handler.NewHandlerPatient(services)

	mux := initRouter(handler)           // инициализируем роутер
	errChan := make(chan error)          // канал для получения ошибок работы сервера
	if err := initConfig(); err != nil { // инициализируем конфиг для того, чтобы узнать порт
		logrus.Fatalf("error init config: %v", err)
	}
	t:=testing.T{}
	TestErr(&t)
	TestForErr(&t)
	TestForRecord(&t)

	go func() {
		errChan <- http.ListenAndServe(":8080", mux) // запускаем сервер
	}()
	select {
	case err = <-errChan: // проверяем нет ли ошибок в работе сервера
		if err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
	}

}

func initRouter(h *handler.Handler) *http.ServeMux {
	mux := &http.ServeMux{}
	mux.HandleFunc("/doctors", http.HandlerFunc(h.GetDoctors)) // хэндлеры
	mux.HandleFunc("/create_doctor", http.HandlerFunc(h.CreateDoctor))
	mux.HandleFunc("/create_patient", http.HandlerFunc(h.CreatePatient))
	mux.HandleFunc("/records", http.HandlerFunc(h.GetRecords))
	mux.HandleFunc("/create_record", http.HandlerFunc(h.CreatePatientRecord))
	mux.HandleFunc("/patient_records", http.HandlerFunc(h.GetRecordOfPatient))

	return mux
}

type server struct {
	httpServer *http.Server
}

func TestErr(t *testing.T) {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	db, _ := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	repos := repository.NewRepository(db)
	repos_d := repository.NewDoctorRepository(db)
	repos_r := repository.NewRecordRepository(db)

	services := service.NewServicePatient(repos, repos_d, repos_r)
	var want error
	want = nil
	_, testErrStr := services.GetDoctors()
	if want != testErrStr {
		t.Errorf("err = %q, want %q", testErrStr, want)
	}
	fmt.Println("TESTS: OK")
}

func TestForErr(t *testing.T) {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	db, _ := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	repos := repository.NewRepository(db)
	repos_d := repository.NewDoctorRepository(db)
	repos_r := repository.NewRecordRepository(db)

	services := service.NewServicePatient(repos, repos_d, repos_r)
	var want error
	want = nil
	_, testErrStr := services.GetRecords()
	if want != testErrStr {
		t.Errorf("err = %q, want %q", testErrStr, want)
	}
	fmt.Println("TESTS: OK")
}

func TestForRecord(t *testing.T) {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	db, _ := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	repos := repository.NewRepository(db)
	repos_d := repository.NewDoctorRepository(db)
	repos_r := repository.NewRecordRepository(db)

	services := service.NewServicePatient(repos, repos_d, repos_r)
	want := []labsMed.Record{
		{
			RecordId:  2,
			Date:      "2022-05-30",
			DoctorId:  1,
			PacientId: 1,
			Diagnosis: "Шизофрения",
		},
	}
	got, _ := services.GetRecordOfPatient(1)
	for i := range got {
		if want[i] != got[i] {
			t.Errorf("err = %q, want %q", got, want)
		}
	}
	fmt.Println("TESTS: OK")
}
