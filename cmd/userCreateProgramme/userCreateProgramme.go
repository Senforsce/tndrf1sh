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

//go:embed createProgram.sparql
var CreateProgrammeQuery string

var (
	userId         string
	trainingMethod string
	persoProgDate  string
	persoProgName  string
)

func init() {
	flag.StringVar(&userId, "i", "http://senforsce.com/o8/brain/Some-ID-123", "The User Id created by `createUser` command")
	flag.StringVar(&trainingMethod, "m", "HIIT", "The Programme training Type (HIIT|AMRAP|MANUEL|...)")
	flag.StringVar(&persoProgDate, "d", "2025-10-20T10:10", "The Programme date and time")
	flag.StringVar(&persoProgName, "n", "Mon Programme", "The Name of the programme")

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
	fmt.Printf("the trainingMethod is %s\n", trainingMethod)
	fmt.Printf("the persoProgDate is %s\n", persoProgDate)
	fmt.Printf("the persoProgName is %s\n", persoProgName)

	userProgrammeId, err := gonanoid.New()

	if err != nil {
		log.Fatalf("Unable to Generate A userProgrammeId")
	}

	infoIdentifier := fmt.Sprintf("http://senforsce.com/o8/brain/userProgramme-%s", userProgrammeId)

	fmt.Printf("the infoIdentifier  is %s\n", infoIdentifier)
	userConfig := userconfig.NewUserConfig()

	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	formattedQuery := fmt.Sprintf(CreateProgrammeQuery,
		infoIdentifier,
		userId,
		fmt.Sprintf("http://senforsce.com/o8/brain/%s", trainingMethod),
		persoProgDate,
		persoProgName,
	)

	log.Println(formattedQuery)

	r, err := repo.Update(formattedQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)

	fmt.Printf("\n SUCCESSFULLY CREATED : %s\n", infoIdentifier)

}
