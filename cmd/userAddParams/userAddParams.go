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

//go:embed addUserParameters.sparql
var AddParametersQuery string

var (
	userId      string
	bgColor     string
	textColor   string
	colorScheme string
)

const (
	DARK  = "dark"
	LIGHT = "light"
)

func init() {
	flag.StringVar(&userId, "i", "http://senforsce.com/o8/brain/Some-ID-123", "The User Id created by `createUser` command")
	flag.StringVar(&bgColor, "bg", "#EEEEEE", "The First Name of the User")
	flag.StringVar(&textColor, "t", "#1FD2AC", "The Last Name of the User")
	flag.StringVar(&colorScheme, "cs", "light", "The Téléphone of the User")

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
	fmt.Printf("the bgColor is %s\n", bgColor)
	fmt.Printf("the textColor is %s\n", textColor)
	fmt.Printf("the colorScheme is %s\n", colorScheme)

	userParamsId, err := gonanoid.New()

	if err != nil {
		log.Fatalf("Unable to Generate A userParamsId")
	}

	infoIdentifier := fmt.Sprintf("http://senforsce.com/o8/brain/userParams-%s", userParamsId)

	fmt.Printf("the infoIdentifier  is %s\n", infoIdentifier)
	userConfig := userconfig.NewUserConfig()

	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	formattedQuery := fmt.Sprintf(AddParametersQuery,
		infoIdentifier,
		userId,
		bgColor,
		colorScheme,
		textColor,
	)

	log.Println(formattedQuery)

	r, err := repo.Update(formattedQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)

	fmt.Printf("\n SUCCESSFULLY CREATED : %s\n", infoIdentifier)

}
