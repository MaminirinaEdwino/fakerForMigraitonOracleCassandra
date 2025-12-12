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

type domaine struct {
	id_domaine  string `faker:"uuid_hyphenated"`
	nom_domaine string `faker:"name"`
}

type enseignant struct {
	id_enseignant string `faker:"uuid_hyphenated"`
	nom           string `faker:"name"`
	prenom        string `faker:"firstname"`
	grade         string `faker:"name"`
	domaine       string `faker:"uuid_hyphenated"`
	niveau        string `faker:"name"`
	specialite    string `faker:"uuid_hyphenated"`
}

type etudiant struct {
	id_etudiant string `faker:"uuid_hyphenated"`
	nom         string `faker:"name"`
	prenom      string `faker:"fisrtname"`
	statut      string `faker:"bloodgrpoup"`
	specialite  string `faker:"uuid_hyphenated"`
	niveau      string `faker:"bloodtype"`
}

type pfe struct {
	id_pfe    string `faker:"uuid_hyphenated"`
	titre_pfe string `faker:"sentence"`
}

type cours struct {
	id_cours    string `faker:"uuid_hyphenated"`
	titre_cours string `faker:"sentence"`
	salle       string `faker:"name"`
}

type enseigne struct {
	id_enseigne   string `faker:"uuid_hyphenated"`
	id_enseignant string `faker:"uuid_hyphenated"`
	id_cours      string `faker:"uuid_hyphenated"`
}

func main() {
	// 1. Déclarer une variable de type Article
	// var article Article

	// // 2. Remplir la structure avec des données factices
	// err := faker.FakeData(&article)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // 3. Afficher les données générées
	// fmt.Println("--- Article Généré ---")
	// fmt.Printf("ID: %s\n", article.ID)
	// fmt.Printf("Titre: %s\n", article.Title)
	// // On affiche juste les 100 premiers caractères du contenu pour la console
	// fmt.Printf("Contenu (Extrait): %s...\n", article.Content)
	// fmt.Printf("Auteur: %s\n", article.AuthorEmail)
	// // fmt.Printf("Date: %s\n", article.CreatedAt.Format("2006-01-02"))
	// fmt.Printf("Vues: %d\n", article.ViewsCount)
	// // fmt.Printf("Publié: %t\n", article.IsPublished)

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "migration_base"
	cluster.Consistency = gocql.Quorum
	cluster.Port = 9042
	cluster.Timeout = 5 * time.Second

	session, errr := cluster.CreateSession()
	if errr != nil {
		log.Fatal(errr)
	}

	defer session.Close()

	var spec Specialite

	err := faker.FakeData(&spec)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(spec.Id_specialite)
	fmt.Println(spec.Nom_specialite)

	query := session.Query("INSERT INTO specialite(id_specialite, nom_specilaite) VALUES(?,?)", spec.Id_specialite, spec.Nom_specialite)
	errr = query.Exec()
	if errr != nil {
		log.Fatal(errr)
	}

}
