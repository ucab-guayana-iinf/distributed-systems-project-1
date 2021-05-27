# Proyecto #1 Sistemas Distribuidos

## TODO

- Servicio de Cuentas
Servicio de Cuentas, mantiene las cuenta y los diferentes servidores se
comunican con él vía una cola de mensajes y el accede, vía un semáforo,
la cuenta en memoria compartida.

- Servidor UDP
- Servidores TCP (procesos e hilos)
- Cliente UDP – Con interfaz a usuario
- Cliente TCP – Con Interfaz a usuario
- Consola local – Con interfaz a usuario
Consola local, provee una interfaz de usuario en una terminal local que
permite manipular la cuenta. Para lo cual accede, vía un semáforo, la
cuenta en memoria compartida, de la misma forma provee información
del servicio, la cola de mensaje y servidores (UDP y TCP). El servicio, así
como los servidores deben provee la capacidad de monitorear su
actividad.

- Consola remota – Con interfaz a usuario
Consola remota; ofrece las mismas funcionalidades de la consola local, y
para su implementación se debe hacer uso de RPC (llamado a
procedimientos remotos) y la función remota en el servidor debe accede,
vía un semáforo, la cuenta en memoria compartida.

## Configuración

TODO: documentar software requerido para correr el proyecto

## Comandos

- Instalar dependencias
  1. `go get -u github.com/manifoldco/promptui`
  2. `go get -u gorm.io/gorm`
  3. `go get -u gorm.io/driver/sqlite`
- Correr la aplicación principal 
  - `go install proyecto1`
  - `go run src/main.go`

## Integrantes 

- Juan Vera  CI 27375479
- Angel Rodríguez CI 27015036
- Brenda Ramos CI 27308627
- Miguel Valdez No es persona natural
