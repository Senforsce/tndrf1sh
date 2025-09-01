package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
	"github.com/senforsce/userconfig"

	"github.com/senforsce/sparql"
)

//go:embed listMoves.sparql
var ListMovesQuery string

//go:embed listUsers.sparql
var ListUsersQuery string

//go:embed listTrainingMethods.sparql
var ListTrainingMethodsQuery string

//go:embed listUserProgrammes.sparql
var ListUserProgrammeQuery string

//go:embed showUserProgramme.sparql
var ShowUserProgrammeQuery string

var (
	listType  string
	userId    string
	programId string
)

func init() {
	flag.StringVar(&listType, "lt", "Moves", "The Type of Elements to list (Moves|TrainingMethods|Users|MainInfo|Objectives|Params|Booking|PR|Programmes)")
	flag.StringVar(&userId, "i", "http://senforsce.com/o8/brain/User-1234", "The UserId Parameter")
	flag.StringVar(&programId, "p", "http://senforsce.com/o8/brain/Programme-1234", "The ProgrammeId Parameter")

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

	fmt.Printf("the listType is %s\n", listType)

	userConfig := userconfig.NewUserConfig()

	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	var formattedQuery string

	switch listType {
	case "Moves":
		{
			formattedQuery = (ListMovesQuery)
			log.Println(formattedQuery)

			r, err := repo.Query(formattedQuery)
			if err != nil {
				log.Fatal(err)
			}
			liste := sparql.ListOfSubjects(r.Results.Bindings)
			for elem := range liste {
				view := NewMoveInfo(liste[elem])
				fmt.Printf("name: %s \t id: %s\n", view.MoveExerciceName, view.Subject)
			}
		}
	case "Users":
		{
			formattedQuery = (ListUsersQuery)
			log.Println(formattedQuery)

			r, err := repo.Query(formattedQuery)
			if err != nil {
				log.Fatal(err)
			}

			liste := sparql.ListOfSubjects(r.Results.Bindings)
			for elem := range liste {
				view := NewUserInfo(liste[elem])
				fmt.Printf("email: %s \t id: %s\n", view.Email, view.Subject)
			}
		}
	case "Types":
		{
			formattedQuery = (ListTrainingMethodsQuery)
			log.Println(formattedQuery)

			r, err := repo.Query(formattedQuery)
			if err != nil {
				log.Fatal(err)
			}

			liste := sparql.ListOfSubjects(r.Results.Bindings)
			for elem := range liste {
				view := NewProgramTypeInfo(liste[elem])
				fmt.Printf("name: %s \t id: %s\n", view.ProgramTypeShortName, view.Subject)
			}
		}
	case "UserProgramme":
		{
			formattedQuery = fmt.Sprintf(ListUserProgrammeQuery, userId)
			log.Println(formattedQuery)

			r, err := repo.Query(formattedQuery)
			if err != nil {
				log.Fatal(err)
			}

			liste := sparql.ListOfSubjects(r.Results.Bindings)
			for elem := range liste {
				view := NewProgramListInfo(liste[elem])
				fmt.Printf("name: %s \t id: %s\n", view.PersoProgName, view.Subject)
			}
		}

	case "Programme":
		{
			formattedQuery = fmt.Sprintf(ShowUserProgrammeQuery, programId)
			log.Println(formattedQuery)

			r, err := repo.Query(formattedQuery)
			if err != nil {
				log.Fatal(err)
			}

			view := NewProgramInfo(r.Results.Bindings)
			fmt.Printf("name: %s \t id: %s\n", view.PersoProgName, programId)
			fmt.Println("Exercices:")
			fmt.Println("\tType\t|\t\t\t\tSubject\t\t\t\t|\t\tName\t\t\t|\tMoveName\t|")
			fmt.Println("__________________________________________________________________________________________________________________________________________________________")

			for _, Ex := range view.HasEchauffement {
				fmt.Printf("%s\t| %s\t|\t%s\t|\t%s\t|\n", "Echauffement:", Ex.Subject, Ex.ExerciseName, Ex.Move.MoveName)
			}
			fmt.Println("__________________________________________________________________________________________________________________________________________________________")

			for _, Ex := range view.HasCorpsDeSeance {
				fmt.Printf("%s| %s\t|\t%s\t|\t%s\t|\n", "Corps de Séance:", Ex.Subject, Ex.ExerciseName, Ex.Move.MoveName)
			}
			fmt.Println("__________________________________________________________________________________________________________________________________________________________")

			for _, Ex := range view.HasRetourAuCalme {
				fmt.Printf("%s| %s\t|\t%s\t|\t%s\t|\n", "Retour Au Calme:", Ex.Subject, Ex.ExerciseName, Ex.Move.MoveName)
			}

			fmt.Println("Valeurs:")

			for _, Ex := range view.HasEchauffement {
				fmt.Printf("%s  >>  %s  |  %s%s  |  %s sets  |  %s reps  |  %s%s  |\n", Ex.Subject, Ex.Move.MoveName,
					Ex.MoveExerciceWeight, Ex.MoveExerciceWeightUnit,
					Ex.MoveExerciceSets, Ex.MoveExerciceReps,
					Ex.MoveExerciceTime, Ex.MoveExerciceTimeUnit,
				)
			}

			for _, Ex := range view.HasCorpsDeSeance {
				fmt.Printf("%s  >>  %s  |  %s%s  |  %s sets  |  %s reps  |  %s%s  |\n", Ex.Subject, Ex.Move.MoveName,
					Ex.MoveExerciceWeight, Ex.MoveExerciceWeightUnit,
					Ex.MoveExerciceSets, Ex.MoveExerciceReps,
					Ex.MoveExerciceTime, Ex.MoveExerciceTimeUnit,
				)
			}

			for _, Ex := range view.HasRetourAuCalme {
				fmt.Printf("%s  >>  %s  |  %s%s  |  %s sets  |  %s reps  |  %s%s  |\n", Ex.Subject, Ex.Move.MoveName,
					Ex.MoveExerciceWeight, Ex.MoveExerciceWeightUnit,
					Ex.MoveExerciceSets, Ex.MoveExerciceReps,
					Ex.MoveExerciceTime, Ex.MoveExerciceTimeUnit,
				)
			}

		}
	}

}

type ShowMove struct {
	MoveExerciceName string
	Subject          string
}

func NewMoveInfo(results []map[string]sparql.Binding) ShowMove {
	return ShowMove{
		MoveExerciceName: sparql.GetValue("MoveExerciceName", results),
		Subject:          sparql.GetValue("s", results),
	}
}

type ShowUser struct {
	Email   string
	Subject string
}

func NewUserInfo(results []map[string]sparql.Binding) ShowUser {
	return ShowUser{
		Email:   sparql.GetValue("UserEmail", results),
		Subject: sparql.GetValue("s", results),
	}
}

type ShowProgramType struct {
	ProgramTypeName      string
	ProgramTypeShortName string
	Subject              string
}

func NewProgramTypeInfo(results []map[string]sparql.Binding) ShowProgramType {
	return ShowProgramType{
		ProgramTypeName:      sparql.GetValue("ProgramTypeName", results),
		ProgramTypeShortName: sparql.GetValue("ProgramTypeShortName", results),
		Subject:              sparql.GetValue("s", results),
	}
}

type ShowProgramList struct {
	TrainingMethod string
	PersoProgDate  string
	PersoProgName  string
	Subject        string
}

func NewProgramListInfo(results []map[string]sparql.Binding) ShowProgramList {
	return ShowProgramList{
		TrainingMethod: sparql.GetValue("TrainingMethod", results),
		PersoProgDate:  sparql.GetValue("PersoProgDate", results),
		PersoProgName:  sparql.GetValue("PersoProgName", results),

		Subject: sparql.GetValue("s", results),
	}
}

type ShowProgram struct {
	TrainingMethod   string
	PersoProgDate    string
	PersoProgName    string
	Subject          string
	HasCorpsDeSeance []ProgramExercice
	HasRetourAuCalme []ProgramExercice
	HasEchauffement  []ProgramExercice
}

func NewProgramInfo(results []map[string]sparql.Binding) ShowProgram {
	return ShowProgram{
		TrainingMethod:   sparql.GetValue("tProgramTypeName", results),
		PersoProgDate:    sparql.GetValue("sPersoProgDate", results),
		PersoProgName:    sparql.GetValue("sPersoProgName", results),
		HasCorpsDeSeance: NewProgramExerciceList("c", results),
		HasRetourAuCalme: NewProgramExerciceList("r", results),
		HasEchauffement:  NewProgramExerciceList("e", results),
		Subject:          sparql.GetValue("s", results),
	}
}

type ExerciceMovement struct {
	MoveName             string
	MoveTips             []byte
	MoveGifCredits       string
	MoveGifMedia         string
	MoveMusclesGroupList string
	MoveDescription      string
	MovePoster           string
}

type ProgramExercice struct {
	Move                     ExerciceMovement
	MoveExerciceWeightUnit   string
	MoveExercicePRMin        string
	MoveExercicePRUnit       string
	MoveExerciceReps         string
	ExerciseName             string
	MoveExerciceSets         string
	MoveExerciceTimeUnit     string
	MoveExercicePRLast       string
	MoveExercicePRBest       string
	MoveExerciceLastTimeDate string
	MoveExerciceWeight       string
	MoveExerciceTime         string
	Subject                  string
}

func NewProgramExerciceList(prefix string, results []map[string]sparql.Binding) []ProgramExercice {
	liste := sparql.ListOf(results, prefix)
	userProgs := []ProgramExercice{}
	for i := range liste {
		prog := NewProgramExercice(prefix, liste[i])
		userProgs = append(userProgs, prog)

	}
	return userProgs
}

func NewProgramExercice(prefix string, results []map[string]sparql.Binding) ProgramExercice {
	return ProgramExercice{
		Move:                     NewExerciceMovement(prefix, results),
		MoveExerciceWeightUnit:   sparql.GetValue(prefix+"MoveExerciceWeightUnit", results),
		MoveExercicePRMin:        sparql.GetValue(prefix+"MoveExercicePRMin", results),
		MoveExercicePRUnit:       sparql.GetValue(prefix+"MoveExercicePRUnit", results),
		MoveExerciceReps:         sparql.GetValue(prefix+"MoveExerciceReps", results),
		MoveExerciceSets:         sparql.GetValue(prefix+"MoveExerciceSets", results),
		ExerciseName:             sparql.GetValue(prefix+"ExerciseName", results),
		MoveExerciceTimeUnit:     sparql.GetValue(prefix+"MoveExerciceTimeUnit", results),
		MoveExercicePRLast:       sparql.GetValue(prefix+"MoveExercicePRLast", results),
		MoveExercicePRBest:       sparql.GetValue(prefix+"MoveExercicePRBest", results),
		MoveExerciceLastTimeDate: sparql.GetValue(prefix+"MoveExerciceLastTimeDate", results),
		MoveExerciceWeight:       sparql.GetValue(prefix+"MoveExerciceWeight", results),
		MoveExerciceTime:         sparql.GetValue(prefix+"MoveExerciseTime", results),
		Subject:                  sparql.GetValue(prefix, results),
	}
}

func NewExerciceMovement(prefix string, results []map[string]sparql.Binding) ExerciceMovement {
	maybeUnsafeHTML := markdown.ToHTML([]byte(sparql.GetValue(prefix+"Tips", results)), nil, nil)
	html := bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)

	return ExerciceMovement{
		MoveName:             sparql.GetValue(prefix+"MoveName", results),
		MoveTips:             html,
		MoveGifCredits:       sparql.GetValue(prefix+"Credits", results),
		MovePoster:           sparql.GetValue(prefix+"Poster", results),
		MoveGifMedia:         sparql.GetValue(prefix+"GifMedia", results),
		MoveMusclesGroupList: sparql.GetValue(prefix+"Muscles", results),
		MoveDescription:      sparql.GetValue(prefix+"MoveDescription", results),
	}
}
