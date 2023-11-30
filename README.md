# Proyecto

La carpeta `cmd` contiene el archivo main, que es el punto de entrada del programa.
En la carpeta `models` están definidos las estructuras de datos que usé para modelar la información, además de funciones para crear y leer registros.
En la carpeta `db` está la configuración para conectarse a la base de datos.

Los requisitos para correr el programa en local son:
- tener un compilador de Go instalado [https://go.dev/](https://go.dev/).
- instalar la base de datos: [https://surrealdb.com/docs/installation](https://surrealdb.com/docs/installation)

Cuando se tengan esas dos herramientas, se procede a correr el programa. La base de datos se puede iniciar en una terminal aparte, ubicado en la carpeta raíz con el comando `surreal start -u root -p root file://./my.db`.

Y el programa, también en una terminal ubicada en la carpeta raíz del programa, se corre con el comando `go run cmd/main.go`
