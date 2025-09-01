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

//go:embed addCreditsCoaching.sparql
var AddCreditsCoaching string

//go:embed addCreditsProgramme.sparql
var AddCreditsProgramme string

var (
	userId        string
	creditNbValue string
	isValidFrom   string
	isValidUntil  string
	bookingType   string
)

func init() {
	flag.StringVar(&userId, "i", "http://senforsce.com/o8/brain/Some-ID-123", "The User Id created by `createUser` command")
	flag.StringVar(&creditNbValue, "nb", "3", "The Number of Credits to give")
	flag.StringVar(&isValidFrom, "d", "2025-31-12T10:00", "The Validity start date of the credits")
	flag.StringVar(&isValidUntil, "u", "2025-31-12T10:00", "The Validity end date of the credits")
	flag.StringVar(&bookingType, "t", "Coaching", "The booking type (Coaching|Programme)")

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
	fmt.Printf("the creditNbValue is %s\n", creditNbValue)
	fmt.Printf("the isValidFrom is %s\n", isValidFrom)
	fmt.Printf("the isValidUntil is %s\n", isValidUntil)
	fmt.Printf("the bookingType is %s\n", bookingType)

	userCreditId, err := gonanoid.New()

	if err != nil {
		log.Fatalf("Unable to Generate A userCreditId")
	}

	userPayableId, err := gonanoid.New()

	if err != nil {
		log.Fatalf("Unable to Generate A userPayableId")
	}

	infoIdentifier := fmt.Sprintf("http://senforsce.com/o8/brain/Credits-%s", userCreditId)
	payableIdentifier := fmt.Sprintf("http://senforsce.com/o8/brain/Payable-%s", userPayableId)

	fmt.Printf("the infoIdentifier  is %s\n", infoIdentifier)
	userConfig := userconfig.NewUserConfig()

	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	var formattedQuery string
	switch bookingType {
	case "Coaching":
		{
			formattedQuery = fmt.Sprintf(AddCreditsCoaching,
				payableIdentifier,
				userId,
				infoIdentifier,
				payableIdentifier,
				userId,
				creditNbValue,
				isValidFrom,
				isValidUntil,
			)
		}
	case "Programme":
		{
			formattedQuery = fmt.Sprintf(AddCreditsProgramme,
				payableIdentifier,
				userId,
				infoIdentifier,
				payableIdentifier,
				userId,
				creditNbValue,
				isValidFrom,
				isValidUntil,
			)
		}
	}

	log.Println(formattedQuery)

	r, err := repo.Update(formattedQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)

	fmt.Printf("\n SUCCESSFULLY CREATED : %s\n", infoIdentifier)
	fmt.Printf("\n SUCCESSFULLY CREATED : %s\n", payableIdentifier)

}
