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

//go:embed addmedicalinfo.sparql
var addMedicalInfo string

var (
	userId            string
	doctorEmail       string
	doctorPhoneNumber string
	emergencyEmail    string
	emergencyFullname string
	gpName            string
	specialistName    string
	weight            string
	height            string
)

func init() {
	flag.StringVar(&userId, "i", "http://senforsce.com/o8/brain/Some-ID-123", "The User Id created by `createUser` command")
	flag.StringVar(&doctorEmail, "de", "test+doctor@senforsce.com", "The Email Of the user's doctor")
	flag.StringVar(&emergencyEmail, "ee", "test+doctor@senforsce.com", "The Email Of the user's emergency contact")
	flag.StringVar(&doctorPhoneNumber, "dpn", "+336666666", "The Téléphone of the User's Doctor")
	flag.StringVar(&emergencyFullname, "efn", "Abdoul Syouk", "The Full Name of the User emergency contact")
	flag.StringVar(&gpName, "gpn", "House MD", "the full name of the doctor/GP")
	flag.StringVar(&specialistName, "spn", "Dr Cuddy", "The full Name of the specialist Doctor")
	flag.StringVar(&weight, "wkg", "60", "the weight of the user in KG")
	flag.StringVar(&height, "hcm", "170", "the height of the user in CM")

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
	fmt.Printf("the doctorEmail  is %s\n", doctorEmail)
	fmt.Printf("the emergencyEmail is %s\n", emergencyEmail)
	fmt.Printf("the doctorPhoneNumber is %s\n", doctorPhoneNumber)
	fmt.Printf("the emergencyFullname is %s\n", emergencyFullname)
	fmt.Printf("the gpName is %s\n", gpName)
	fmt.Printf("the specialistName is %s\n", specialistName)
	fmt.Printf("the weight is %skg\n", weight)
	fmt.Printf("the height is %scm\n", height)

	mainInfoId, err := gonanoid.New()

	if err != nil {
		log.Fatalf("Unable to Generate A mainInfoId")
	}

	infoIdentifier := fmt.Sprintf("http://senforsce.com/o8/brain/medicalInfo-%s", mainInfoId)

	fmt.Printf("the infoIdentifier  is %s\n", infoIdentifier)
	userConfig := userconfig.NewUserConfig()

	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	formattedQuery := fmt.Sprintf(addMedicalInfo,
		infoIdentifier,
		userId,
		doctorEmail,
		doctorPhoneNumber,
		emergencyEmail,
		emergencyFullname,
		gpName,
		specialistName,
		height,
		weight,
	)

	log.Println(formattedQuery)

	r, err := repo.Update(formattedQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)

	fmt.Printf("\n SUCCESSFULLY CREATED : %s\n", infoIdentifier)

}
