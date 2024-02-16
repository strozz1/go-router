

# Router

Framework de un router en Go con fines de estudio.

## Objetivos del proyecto
El desarrollo del proyecto es para aprender y mejorar sobre enrutamiento y middleware.

Al acabar el proyecto, este debe tener:

1. Sistema de rutas mendiante el router.
2. Sistema dinamico de rutas (con parametros y valores dinamicos)
3. Un buen sistema de middlewares
4. Un buen sistema de lectura de la ruta otorgada por el cliente
5. buena gestion de errores e informes al cliente
6. buenos test para comprobar el funcionamiento correcto.


## TODO LIST
- [ ] Rutas dinamicas
- [ ] Sistema bueno de middlewares
- [ ] Sanitizacion y limpieza de la ruta
- [ ] Tests
- [x] Sistema de enrutamiento basico


## Basic usage

```go
func main() {

    //Example usage
    addr := ":8000"   
  	router := routes.NewRouter()

	router.Endpoint("/", handlers.HandleIndex)
    router.Endpoint("/inicio", handlers.HandleHome)      
	router.Endpoint("/rutas", handlers.HandleRutas) 

    log.Println("Listening on ",addr)
	http.ListenAndServe(addr, &router)
}
```

## Estructura basica del proyecto

La estructura principal de este proyecto es el `router`. Este es el que se va a encargar de realizar todas las tareas relacionadas con enrutamiento y
gestion de las peticiones entrantes. 

### Endpoints

Para crear nuevos endpoints, se usa el metodo Endpoint de router. Se le pasa una ruta y una funcion que se encargara de gestionar la ruta en cuestion.
Esta funcion debe ser de tipo `http.HandlerFunc`  de la libreria std de Go.

```go
// With function already defined
router.Endpoint("/home", homeHandler)

//or lambda function
router.Endpoint("/home", func (w http.ResponseWriter, r *http.Request){
    //do logic here
    }) 

```
Puedes definir tus propios handlers.

Es importante destacar que el programa sacara un error si la ruta esta mal formada o si es vacia.
Por ejemplo, si es `home` salta un error ya que debe preceder con una `/`.

### Peticiones entrantes

Para escuchar las peticiones entrantes se usa el metodo http.ListenAndServe de la libreria http de Go, pasandole la direccion de nuestra server y el router.
```go
http.ListenAndServe(":80", &router)

```

### Middlewares y Auth

El framework trae algunos middlewares por defecto, como GET, POST, LOG y algunos otros que podemos usar para agregar funcionalidad y limitar el uso de la ruta.

Por ejemplo si queremos que una ruta sea solo accesible con el metodo http GET, podemos usar el middleware GET.
```go
//todo
```
Es importante saber que los middleware se aplican en orden inverso al que son puestos. Es decir, si el ultimo middleware que pones es el LOG, es el primero que se ejecuta.

Lo que hace principalmente un middleware es envolver a el controlador anterior de la ruta y meterle funcionalidad extra por encima, como un envoltorio.
De esta manera se pueden poner tantos middleware como se quiera.

Tambien puedes crear middlewares propios. Simplemente estos tienen que ser de tipo Middleware del paquete middlewares


### Funcionamiento interno

Vamos a explicar brevemente como funciona el framework por dentro.

El componente principal es el `router`, el cual tiene unos atributos importantes.

Uno de ellos es un arbol que contiene todas las rutas que hemos definido con el metodo `Endpoint`. El arbol contiene en cada Nodo un valor de tipo Segment.

Este struct `Segment` representa un trozo de la ruta(ej: "/home/profile", los segments serian "/home" y "/profile"). De esta forma cuando llega una peticion podemos recorrer el arbol
para encontrar la ruta final si es que existe.

No todos los segmentos tienen porque tener una ruta final. Puede que quiera tener "/books/author/3" y que la ruta "/books" no sea una ruta valida. Para ello se encuentran los handlers que estan almacenados en los segmentos.
Si queremos que una ruta no sea final, tan solo debemos dejar el controlador vacio. O si definimos una ruta con subrutas que no estaban definidas, se guardaran sin un controlador(handler).

De esta manera podemos manejar de forma sencilla que rutas tienen final y cuales no.

El segmento tambien tiene una propiedad llamada `SegmentType` que define que clase de segmento es.

Hay dos tipos de segmentos:

- *Static*: es el por defecto y el mas basico. Es el que la ruta es estatica y no cambia (ej: "/home")
- *Dynamic*: es la que puede variar segun su uso y no ser siempre igual (ej: "/user/juan" siendo /juan la ruta dinamica).


