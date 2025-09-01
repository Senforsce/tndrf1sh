package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/senforsce/userconfig"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/senforsce/sparql"
)

//go:embed addPR.sparql
var AddPRQuery string

var (
	userId     string
	comment    string
	date       string
	prMovement string
	unit       string
	value      string
)

func init() {
	flag.StringVar(&userId, "i", "http://senforsce.com/o8/brain/Some-ID-123", "The User Id created by `createUser` command")
	flag.StringVar(&comment, "c", "Un Commentaire", "The PR Comment")
	flag.StringVar(&date, "d", "2025-31-12T10:00", "The Last Name of the User")
	flag.StringVar(&prMovement, "m", "http://senforsce.com/o8/brain/coach-mj/Some-ID-123", "The Téléphone of the User")
	flag.StringVar(&unit, "c", "reps", "The PR unit (reps|kg|s|min)")
	flag.StringVar(&value, "c", "25", "The PR value")

}

func main() {
	projectRoot := os.Getenv("TNDR_GENERATE_ROOT")
	generateKey := os.Getenv("TNDR_GENERATE_KEY")

	flag.Parse()

	if projectRoot == "" {
		log.Fatal("Please Set the environment variable TNDR_GENERATE_ROOT")
	}

	if generateKey == "" {
		log.Fatal("Please Set the environment variable TNDR_GENERATE_KEY")
	}

	fmt.Printf("the userId is %s\n", userId)
	fmt.Printf("the comment is %s\n", comment)
	fmt.Printf("the date is %s\n", date)
	fmt.Printf("the prMovement is %s\n", prMovement)
	fmt.Printf("the unit is %s\n", unit)
	fmt.Printf("the value is %s\n", value)

	userPrId, err := gonanoid.New()

	if err != nil {
		log.Fatalf("Unable to Generate A userPrId")
	}

	infoIdentifier := fmt.Sprintf("http://senforsce.com/o8/brain/coach-mj/PR-%s", userPrId)

	fmt.Printf("the infoIdentifier  is %s\n", infoIdentifier)
	userConfig := userconfig.NewUserConfig()

	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	formattedQuery := fmt.Sprintf(AddPRQuery,
		infoIdentifier,
		"Created by terminal",
		date,
		prMovement,
		userId,
		unit,
		value,
	)

	log.Println(formattedQuery)

	r, err := repo.Update(formattedQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)

	fmt.Printf("\n SUCCESSFULLY CREATED : %s\n", infoIdentifier)

}
