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

//go:embed addMainInfo.sparql
var AddMainInfoQuery string

var (
	userId       string
	firstname    string
	lastname     string
	birthdayDate string
	userEmail    string
	telephone    string
	joinedSince  string
)

func init() {
	flag.StringVar(&userId, "i", "http://senforsce.com/o8/brain/Some-ID-123", "The User Id created by `createUser` command")
	flag.StringVar(&userEmail, "e", "test+user@senforsce.com", "The Email Of the user to create")
	flag.StringVar(&firstname, "fn", "Abdoul", "The First Name of the User")
	flag.StringVar(&birthdayDate, "dob", "2025-12-31T00:00", "the date of birth of the user")
	flag.StringVar(&lastname, "ln", "Sy", "The Last Name of the User")
	flag.StringVar(&telephone, "tel", "+336666666", "The Téléphone of the User")

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

	joinedSince = time.Now().Format(time.RFC3339)

	fmt.Printf("the email  is %s\n", userEmail)
	fmt.Printf("the userId is %s\n", userId)
	fmt.Printf("the firstname is %s\n", firstname)
	fmt.Printf("the lastname is %s\n", lastname)
	fmt.Printf("the birthdayDate is %s\n", birthdayDate)
	fmt.Printf("the telephone is %s\n", telephone)
	fmt.Printf("joined since is %s\n", joinedSince)

	mainInfoId, err := gonanoid.New()

	if err != nil {
		log.Fatalf("Unable to Generate A mainInfoId")
	}

	infoIdentifier := fmt.Sprintf("http://senforsce.com/o8/brain/mainInfo-%s", mainInfoId)

	fmt.Printf("the infoIdentifier  is %s\n", infoIdentifier)
	userConfig := userconfig.NewUserConfig()

	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	formattedQuery := fmt.Sprintf(AddMainInfoQuery,
		infoIdentifier,
		userId,
		firstname,
		birthdayDate,
		userEmail,
		lastname,
		telephone,
		joinedSince,
	)

	log.Println(formattedQuery)

	r, err := repo.Update(formattedQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)

	fmt.Printf("\n SUCCESSFULLY CREATED : %s\n", infoIdentifier)

}
