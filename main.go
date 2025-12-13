package main

import (
	"log"
	"math/rand"
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
	Prenom        string `faker:"name"`
	Grade         string `faker:"name"`
	Domaine       string `faker:"uuid_hyphenated"`
	Niveau        string `faker:"sentence"`
	Specialite    string `faker:"uuid_hyphenated"`
}

type Etudiant struct {
	Id_etudiant string `faker:"uuid_hyphenated"`
	Nom         string `faker:"name"`
	Prenom      string `faker:"name"`
	Statut      string `faker:"sentence"`
	Specialite  string `faker:"uuid_hyphenated"`
	Niveau      string `faker:"sentence"`
}

type Pfe struct {
	Id_pfe      string `faker:"uuid_hyphenated"`
	Titre_pfe   string `faker:"sentence"`
	Id_etudiant string `faker:"uuid_hyphenated"`
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

func RandomNumber(max int) int {
	return rand.Intn(max)
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
	var liste_specialite []Specialite
	var liste_domaine []Domaine
	var liste_enseignant []Enseignant
	var liste_etudiant []Etudiant
	var liste_cours []Cours
	// var liste_enseigne []Enseigne

	for range 10 {
		ErrorLogger(faker.FakeData(&spec))
		liste_specialite = append(liste_specialite, spec)
		ErrorLogger(session.Query("INSERT INTO specialite (id_specialite, nom_specilaite) VALUES(?,?)", spec.Id_specialite, spec.Nom_specialite).Exec())
	}


	for range 10 {
		ErrorLogger(faker.FakeData(&domaine))
		liste_domaine = append(liste_domaine, domaine)
		ErrorLogger(session.Query("INSERT INTO domaine (id_domaine, nom_domaine) VALUES(?,?)", domaine.Id_domaine, domaine.Nom_domaine).Exec())
	}

	for range 20{
		ErrorLogger(faker.FakeData(&enseignant))

		enseignant.Domaine = liste_domaine[RandomNumber(10)].Id_domaine
		enseignant.Specialite = liste_specialite[RandomNumber(10)].Id_specialite
		liste_enseignant = append(liste_enseignant, enseignant)

		ErrorLogger(session.Query("INSERT INTO enseignants (id_enseignant,  nom, prenom, grade, domaine, niveau, specialite) VALUES(?,?,?,?,?,?,?)", enseignant.Id_enseignant, enseignant.Nom, enseignant.Prenom, enseignant.Grade, enseignant.Domaine, enseignant.Niveau, enseignant.Specialite).Exec())
	}

	for range 20 {
		ErrorLogger(faker.FakeData(&etudiant))

		etudiant.Specialite = liste_specialite[RandomNumber(10)].Id_specialite
		liste_etudiant = append(liste_etudiant, etudiant)
		ErrorLogger(session.Query("INSERT INTO etudiants (id_etudiants, nom, prenom, statut, specialite, niveau) VALUES(?,?,?,?,?,?)", etudiant.Id_etudiant, etudiant.Nom, etudiant.Prenom, etudiant.Statut, etudiant.Specialite, etudiant.Niveau).Exec())
	}

	for range 20 {
		ErrorLogger(faker.FakeData(&pfe))

		pfe.Id_etudiant = liste_etudiant[RandomNumber(20)].Id_etudiant

		ErrorLogger(session.Query("INSERT INTO pfe (id_pfe, titre_pfe, id_etudiant) VALUES(?,?,?)", pfe.Id_pfe, pfe.Titre_pfe, pfe.Id_etudiant).Exec())
	}

	for range 20 {
		ErrorLogger(faker.FakeData(&cours))

		liste_cours = append(liste_cours, cours)

		ErrorLogger(session.Query("INSERT INTO cours (id_cours, titre_cours, salle) VALUES(?,?,?)", cours.Id_cours, cours.Titre_cours, cours.Salle).Exec())
	}

	for range len(liste_cours) {
		ErrorLogger(faker.FakeData(&ensigne))
		ensigne.Id_enseignant = liste_enseignant[RandomNumber(20)].Id_enseignant
		ensigne.Id_cours = liste_cours[RandomNumber(20)].Id_cours
		// liste_enseigne = append(liste_enseigne, ensigne)

		ErrorLogger(session.Query("INSERT INTO enseigne (id_enseigne, id_enseignant, id_cours) VALUES(?,?,?)", ensigne.Id_enseigne, ensigne.Id_enseignant, ensigne.Id_cours).Exec())
	}
}
