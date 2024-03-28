# Bokkatalog-Applikation i Go

## Instruktioner
För att köra denna exempel-applikation i Go behöver du ha följande verktyg installerade:
- Go
- PostgreSQL

Klona detta Git-repositorium till din lokala maskin.

Använd följande kommando i din PostgreSQL CLI för att skapa en PostgreSQL-databas med namnet bookCatalogue:

```CREATE DATABASE bookCatalogue;```

Ange dina PostgreSQL-användaruppgifter och andra konfigurationsinställningar i filen bookCatalogue.go under variabeln connStr.

```connStr := "postgres://postgres:password@localhost:5432/bookCatalogue?sslmode=disable"```

Kompilera och kör Go-applikationen genom att köra följande kommando från terminalen:

```go run bookCatalogue.go```

Följ anvisningarna på terminalen för att interagera med applikationen.
Du kan lägga till, ta bort, redigera och visa alla böcker i katalogen.

## Kort Beskrivning
Denna Go-applikation är en enkel bokkatalogapplikation som låter användare interagera med en databas av böcker. 
Applikationen är baserad på de grundläggande principerna för CRUD-operationer (Create, Read, Update, Delete). 
Användare kan lägga till nya böcker, ta bort befintliga böcker, redigera bokinformation och visa alla tillgängliga böcker i katalogen.
Alla ändringar användarna gör kommer att återspeglas i databasen. Programmet är skrivit med programspråket Go och använder sig av databasen PostgreSQL.

## Tidsåtgång
Projektet tog ungefär 5 dagar att genomföra. Eftersom jag inte hade någon tidigare erfarenhet med Go eller PostgreSQL var det särskilt tidskrävande att lära mig verktygen.
