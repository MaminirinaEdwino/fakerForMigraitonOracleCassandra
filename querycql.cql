--cqlsh : l'outil en ligne de commande pour se connecter a cassandra, utilise le langage cql, une forme de sql mais spécifique a cassandra
--creation de base de donnée --
CREATE KEYSPACE migration_base WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
--precise la base de donnée a utilisé--
USE migration_base;

--type uuid : identifiant auto generé et incrementé de cassandra
CREATE TABLE Specialite (
    id_specialite uuid PRIMARY KEY,
    nom_specilaite VARCHAR
);

CREATE TABLE Domaine (
    id_domaine uuid PRIMARY KEY,
    nom_domaine VARCHAR
);



CREATE TABLE Enseignants (
    id_enseignant uuid PRIMARY KEY,
    nom VARCHAR,
    prenom VARCHAR,
    grade VARCHAR,
    domaine uuid,
    niveau VARCHAR,
    specialite uuid
);



CREATE TABLE Etudiants (
    id_etudiants uuid PRIMARY KEY,
    nom VARCHAR,
    prenom VARCHAR,
    statut VARCHAR,
    specialite uuid,
    niveau VARCHAR
);

CREATE TABLE PFE(
    id_pfe uuid PRIMARY KEY,
    titre_pfe VARCHAR,
    id_etudiant uuid
);

CREATE TABLE Cours(
    id_cours uuid PRIMARY KEY,
    titre_cours VARCHAR,
    salle VARCHAR
);

CREATE TABLE Enseigne(
    id_enseigne uuid PRIMARY KEY,
    id_enseignant uuid,
    id_cours uuid
);

