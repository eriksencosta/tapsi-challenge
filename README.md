# Tapsi-challenge

Proyecto que soluciona el desafio propuesto por tapsi.

## Problema
Dado un tablero de damas y el turno de un jugador, encontrar un movimiento que permita comer una ficha, si no existe tal, encontrar
un movimiento sencillo (es decir un movimiento valido donde no se come una ficha)

## Dependencias

* Lenguaje : Go lang 1.8.3
* Librerias externas : github.com/gorilla/mux

## Requisitos para la solucion
* Se asume que se encuentra un bloque negro en la parte inferior izquierda de cada jugador
* No se tiene un limite de fichas por cada jugador
* Se asume que si no hay fichas de un jugador, el otro jugador ya ha ganado y no se pude hacer una jugada
* Se asume que si no hay fichas, el estado del tablero no es correcto
* Se asume que las fichas blancas empiezan en la parte de abajo y las negras arriba
* Se asume que no hay reyes, es decir fichas que se encuentran al otro extremo de donde empezaron.
* Si no existe una jugada posible se informara en la respuesta
* El tablero debe ser de 8 *8

Tablero de ejemplo

![picture alt](https://skillgamesboard.com/GamesResources/checkers/Images/is_checkers.jpg "Title is optional")

Si el request tiene inconsistencias se retorna un Bad Request

## Solucion
Se busca en todo el tablero las fichas del jugador en turno, cada que se encuentra una de estas, se verifica que a un cuadro de
distancia en la diagonal de su posible movimiento, se encuentre una ficha del jugador contrario, si esta condicion se da, se 
verifica que a dos cuadros de distancia en la diagonal exista un cuadro vacio, de cumplirse ambas condiciones, se retorna la solucion,
de no ser asi se siguen comprobando ficha por ficha.

De no encontrar ninguna ficha que cumpla con las condiciones anteriores, se deduce que no hay un movimiento que permita comer una ficha,
lo que conduce a que se busque hacer un movimiento simple, este movimiento se puede dar solo si a un cuadro de distancia en la
diagonal de posible movimiento este se encuentra vacio.

## Request

Se espera un request que contenga los siguientes campos:

board : un string con 64 caracteres los cuales pueden ser w, b, 0
player : un string que representa que jugador tiene el turno de juego, los valores pueden ser w o b

w = white
b = black
0 = casillas vacias

## Respuesta

La respuesta estara dada por un Json con la siguiente estructura:
{"play" : valor, "play2d" : valor, "error" : valor}

donde el valor de play indica la posicion inicial en el string de entrada donde empezo la ficha y la posicion final despues de 
hacer el movimiento

play2d indica la posicion inicial y la posicion final en formato de matriz donde la esquina superior izquierda es representada por
el valor (0,0) y la esquina inferior derecha es representada por el valor (7,7)

error indica si no se pudo realizar un movimiento con la configuracion dada

## Ejemplo

Si se corre el proyecto localmente y usando CURL se podria usar como request
curl -H "Content-Type: application/json" -d '{"board":"00000000000000000000000000000000000b00000000w0000000000000000000", "player" : "w"}' http://localhost:7100/move

Obtenemos como response
{"play":"44-26","play2d":"( 5,4 ) -> ( 3,2 )","error":""}



