#### funtemps
Funtemps står for funfacts temprature og er en oppgave som ble publisert i januar på https://uia-worker.github.io/is105komp/testing.html. Hovedmålet med oppgaven var å prøve seg ut på test-drevet utvikling. Eleven ble tildelt en mal for repositorien funtemps gjennom github.com. I repositorien funtemps var det en gitt mappe struktur som skulle følge med undermappene conv og funfacts. 

#### Formål
Formålen med oppgaven var å sette sammen de ulike abstraksjonene fra repositorien og konfigurere det slik at det endrer som et program som kan konvertere ulike tempratur enheter og skrive ut funfacts. For å kunne gjøre testene fra pakke-mappene converter (conv) og funfacts operativ måtte flagg konfigureres. Blant annet Luna, Terra og Sun. For konvertering av tempratur fra en enhet til en annen måtte flagg variablene C, F og K konfigureres og -t for temprature. 


#### conv
Converter bruker test bruker funksjoner til å konvertere og teste/validere en float64 verdi til en annen. 

#### main.go
Main.go er et program som tar input av tempraturer i enhetene Kelvin, Celsius og Fahrenheit. Dette gjøres fra terminal ved bruk av kommandoer der flagg som C. F og K brukes. Den returnere tempraturen i output " -out x". Koden bruker init() funksjonen for å initialisere/påkalle flagg variablene, main fuksjonen brukes til å konvertere flaggene. Koden bruker også isFlagPassed() funksjoen for å validere at flagget er spesifisert i kommandolinjen i terminal

#### Funfacts
Mappe-pakken funfacts inneholder to filer funfacts.go inneholder en funksjone GetFunFacts og en struktur med tre felt. I disse feltene står det ulike Funfacts il flaggene Luna, Terra og Sun i form av String. Disse blir hentet opp av GetFunFact funksjonen når den påkalles.

  It creates a test struct which contains an input string and and the expected output. It then creates a slice of tests which contain different input strings and the corresponding output. Finally, it runs a for loop over the tests and compares the expected output with the output from the GetFunFacts function. If the outputs are not equal, it will return an error.

#### Refleksjon

Jeg samarbeidet med Tony Le på denne oppgaven. 

