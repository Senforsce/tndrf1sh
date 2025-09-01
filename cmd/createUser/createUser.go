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

//go:embed createuser.sparql
var CreateUserQuery string

var (
	userEmail   string
	userComment string
)

func init() {
	flag.StringVar(&userEmail, "e", "test+user@senforsce.com", "The Email Of the user to create")
	flag.StringVar(&userComment, "c", "Created by Tndr OWL Generator", "A Comment The User Carries")

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

	fmt.Printf("the email  is %s\n", userEmail)
	fmt.Printf("the comment is %s\n", userComment)

	userid, err := gonanoid.New()

	if err != nil {
		log.Fatalf("Unable to Generate A UserID")
	}

	userIdentifier := fmt.Sprintf("http://senforsce.com/o8/brain/User-%s", userid)

	fmt.Printf("the userIdentifier  is %s\n", userIdentifier)
	userConfig := userconfig.NewUserConfig()

	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	formattedQuery := fmt.Sprintf(CreateUserQuery,
		userIdentifier,
		userComment,
		userEmail,
	)

	r, err := repo.Update(formattedQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)

	fmt.Printf("\n SUCCESSFULLY CREATED : %s\n", userIdentifier)

}
