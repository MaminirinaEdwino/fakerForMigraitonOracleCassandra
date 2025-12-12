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
	id_domaine  gocql.UUID
	nom_domaine string
}

type enseignant struct {
	id_enseignant gocql.UUID
	nom           string
	prenom        string
	grade         string
	domaine       gocql.UUID
	niveau        string
	specialite    gocql.UUID
}

type etudiant struct {
	id_etudiant gocql.UUID
	nom         string
	prenom      string
	statut      string
	specialite  gocql.UUID
	niveau      string
}

type pfe struct {
	id_pfe    gocql.UUID
	titre_pfe string
}

type cours struct {
	id_cours    gocql.UUID
	titre_cours string
	salle       string
}

type enseigne struct {
	id_enseigne   gocql.UUID
	id_enseignant gocql.UUID
	id_cours      gocql.UUID
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

	query := session.Query("INSERT INTO specialite(id_specialite, nom_specilaite) VALUES(?,?)",spec.Id_specialite,spec.Nom_specialite)
	errr = query.Exec()
	if errr != nil{
		log.Fatal(errr)
	}
	
}
