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

//go:embed addAddress.sparql
var addAddressQuery string

var (
	userId       string
	city         string
	country      string
	addressLine1 string
	addressLine2 string
	postcode     string
)

func init() {
	flag.StringVar(&userId, "i", "http://senforsce.com/o8/brain/Some-ID-123", "The User Id created by `createUser` command")
	flag.StringVar(&city, "ci", "Paris", "The City of the user")
	flag.StringVar(&country, "cy", "France", "The Country of the user")
	flag.StringVar(&addressLine1, "a1", "9 rue des addresses", "the address (line 1) of the user")
	flag.StringVar(&addressLine2, "a2", "Flat32", "the address (line 2) of the user")

	flag.StringVar(&postcode, "zip", "75001", "The PostCode/ZipCode of the user")

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
	fmt.Printf("the city is %s\n", city)
	fmt.Printf("the country is %s\n", country)
	fmt.Printf("the addressLine1 is %s\n", addressLine1)
	fmt.Printf("the addressLine2 is %s\n", addressLine2)
	fmt.Printf("the postcode is %s\n", postcode)

	addressId, err := gonanoid.New()

	if err != nil {
		log.Fatalf("Unable to Generate An addressId")
	}

	infoIdentifier := fmt.Sprintf("http://senforsce.com/o8/brain/addr-%s", addressId)

	fmt.Printf("the infoIdentifier  is %s\n", infoIdentifier)
	userConfig := userconfig.NewUserConfig()

	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	formattedQuery := fmt.Sprintf(addAddressQuery,
		infoIdentifier,
		city,
		country,
		addressLine1,
		addressLine2,
		postcode,
		userId,
	)

	log.Println(formattedQuery)

	r, err := repo.Update(formattedQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)

	fmt.Printf("\n SUCCESSFULLY CREATED : %s\n", infoIdentifier)

}
