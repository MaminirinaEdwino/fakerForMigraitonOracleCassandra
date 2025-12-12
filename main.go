package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/gocql/gocql"
)

// type Article struct {
//     ID           string    `faker:"uuid_hyphenated"` // Génère un UUID
//     Title        string    `faker:"sentence"`        // Génère une phrase
//     Content      string    `faker:"paragraph"`       // Génère un paragraphe long
//     AuthorEmail  string    `faker:"email"`           // Génère un email
//     ViewsCount   int       `faker:"boundary_start=100, boundary_end=5000"` // Nombre entre 100 et 5000
// }

const keyspace_name = "migration_base"

type Specialite struct {
	Id_specialite  string `faker:"uuid_hyphenated"`
	Nom_specialite string `faker:"name"`
}

type Domaine struct {
	Id_domaine  string `faker:"uuid_hyphenated"`
	Nom_domaine string `faker:"name"`
}

type Enseignant struct {
	Id_enseignant string `faker:"uuid_hyphenated"`
	Nom           string `faker:"name"`
	Prenom        string `faker:"firstname"`
	Grade         string `faker:"name"`
	Domaine       string `faker:"uuid_hyphenated"`
	Niveau        string `faker:"name"`
	Specialite    string `faker:"uuid_hyphenated"`
}

type Etudiant struct {
	Id_etudiant string `faker:"uuid_hyphenated"`
	Nom         string `faker:"name"`
	Prenom      string `faker:"fisrtname"`
	Statut      string `faker:"bloodgrpoup"`
	Specialite  string `faker:"uuid_hyphenated"`
	Niveau      string `faker:"bloodtype"`
}

type Pfe struct {
	Id_pfe    string `faker:"uuid_hyphenated"`
	Titre_pfe string `faker:"sentence"`
}

type Cours struct {
	Id_cours    string `faker:"uuid_hyphenated"`
	Titre_cours string `faker:"sentence"`
	Salle       string `faker:"name"`
}

type Enseigne struct {
	Id_enseigne   string `faker:"uuid_hyphenated"`
	Id_enseignant string `faker:"uuid_hyphenated"`
	Id_cours      string `faker:"uuid_hyphenated"`
}

func ErrorLogger(err error) {
	if err != nil {
		log.Fatal(err)
	}
} 
func main() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "migration_base"
	cluster.Consistency = gocql.Quorum
	cluster.Port = 9042
	cluster.Timeout = 5 * time.Second

	session, errr := cluster.CreateSession()
	ErrorLogger(errr)

	defer session.Close()

	var spec Specialite
	var domaine Domaine
	var enseignant Enseignant
	var etudiant Etudiant
	var pfe Pfe
	var cours Cours
	var ensigne Enseigne


	err := faker.FakeData(&spec)
	ErrorLogger(err)

	fmt.Println(spec.Id_specialite)
	fmt.Println(spec.Nom_specialite)

	err = faker.FakeData(&domaine)
	ErrorLogger(err)

	err = faker.FakeData(&enseignant)
	ErrorLogger(err)

	err = faker.FakeData(&etudiant)
	ErrorLogger(err)

	err = faker.FakeData(&pfe)
	ErrorLogger(err)

	err = faker.FakeData(&cours)
	ErrorLogger(err)

	err = faker.FakeData(&ensigne)
	ErrorLogger(err)

	query := session.Query("INSERT INTO specialite(id_specialite, nom_specilaite) VALUES(?,?)", spec.Id_specialite, spec.Nom_specialite)

	errr = query.Exec()
	ErrorLogger(errr)

	

}
