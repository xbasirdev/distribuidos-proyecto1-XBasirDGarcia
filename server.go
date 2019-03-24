package main
////////////////////////////////////////////////  IMPORTANDO LIBRERIAS   //////////////////////////////////////////
//importando librerias necesarias
import (
    "net"                              //servidores tcp udp y http
    "fmt"                              //Escribir en consolo
    "bufio"                            //Enviar y recivir mensajes tcp
    "database/sql"                     //Interactuar con bases de datos
    _ "github.com/go-sql-driver/mysql" //conectar a MySQL                    
    "strconv"                          //convertir cadena a entero y viceversa
    "strings"                          //manipular strings
    "log"                              //Mostrar mensajes de errores

)

/////////////////////////////////////////////  CONEXION CON LA BASE DE DATOS   ///////////////////////////////////////
//funcion para la conexion con la bse de datos
// retorna la base de datos o un error si la conexion no existe
func obtenerBaseDeDatos() (db *sql.DB, e error) {
  //configuracion de la base de datos
  usuario := "root"
  pass := "root"
  host := "tcp(127.0.0.1:3306)"
  nombreBaseDeDatos := "inventario"
  // Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
  db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
  if err != nil {
    return nil, err
  }
  return db, nil
}

///////////////////////////////////////// CREANDO LA ESTRUCTURA (MODELOS) /////////////////////////////////////////////
//definimos nuestro modelo
type Person struct{
  // los datos a recibir seran en formato json
  // el ID lo enviara con el nombre de id (minuscula) y que no este vacio, para poder almacenarlo
  ID int `json:"id,omitempty"`
  Nombre string `json:"nombre,omitempty"`
  Direccion string `json:"direccion,omitempty"`
}

type Producto struct{
  id int
  nombre string
  marca string
  precio int
  cantidad int
}

type Cajero struct{
  id int
  usuario string
  password string

}


////////////////////////////////// /  DEFINIENDO ARRAY (BASE DE DATOS)    /////////////////////////////
//Arreglo de persona para manejar los datos, hacemos referencia a la estructura Person
var ArrPersonasBD []Person


//////////////////////////////////////////////// FUNCIONES CRUD CON SERVIDOR TCP y UDP ////////////////////////////////////

// funcion para eliminar una persona
func EliminarProductoTCPyUDP(i int) bool{

  //abrimos nuestra base de datos
  db, err := obtenerBaseDeDatos()
  //comprobamos que no exista errores
  if err != nil {
    fmt.Printf("Error obteniendo base de datos: %v", err)
    return false
  }
  // Terminar conexión al terminar función
  defer db.Close()


  var producto Producto
  //verificamos si existe el id
  db.QueryRow("SELECT id FROM producto where id = ?", i).Scan(&producto.id)

  //si existe la persona la eliminados
  //de lo contrario se retorna false
  if producto.id == i {
       //preparamos nuestro Query
    sentenciaPreparada, err := db.Prepare("DELETE FROM producto WHERE id = ?")
    if err != nil {
      fmt.Printf("Error de sentencia eliminar: %v", err)
      return false
    }
    defer sentenciaPreparada.Close()

    //Eliminamos el dato
    _, err = sentenciaPreparada.Exec(i)
    if err != nil {
      fmt.Printf("Error al Eliminar dato: %v", err)
      return false
    }
    return true
  }else{
    return false
  }

}

// funcion para obtener una persona
func ObtenerUnProductoTCPyUDP(i int) (string) {
  
  //abrimos nuestra base de datos
  db, err := obtenerBaseDeDatos()
  //comprobamos que no exista errores
  if err != nil {
    fmt.Printf("Error obteniendo base de datos: %v", err)
  }
  // Terminar conexión al terminar función
  defer db.Close()

  // creamos una variable de tipo Person para almacenar el dato
  var product Producto

  //lanzamos la consulta a nuestra base de datos y guardamos los datos en persona con Scan
  db.QueryRow("SELECT id, nombre, marca, precio, cantidad FROM producto where id = ?", i).Scan(&product.id, &product.nombre, &product.marca, &product.precio, &product.cantidad)

  //si encuentra el id retorna los datos de la persona

  if product.id == i {

    idCadena := strconv.Itoa(product.id)
    precioCadena := strconv.Itoa(product.precio)
    cantidadCadena := strconv.Itoa(product.cantidad)
    return "ID: "+idCadena+", Nombre: "+product.nombre+", Marca: "+product.marca+", Precio: "+precioCadena+", Cantidad: "+cantidadCadena

  }else{
    
    return ""

  }

}

// funcion para crear una persona
func InsertarProductoTCPyUDP(id int, nombre string, marca string, precio int, cantidad int) (string) {
  
  //abrimos nuestra base de datos
  db, err := obtenerBaseDeDatos()
  //comprobamos que no exista errores
  if err != nil {
    fmt.Printf("Error obteniendo base de datos: %v", err)
    return ""
  }
  // Terminar conexión al terminar función
  defer db.Close()

  //crearemos una persona de tipo Producto
  var producto Producto

  db.QueryRow("SELECT id, nombre, direccion FROM persona where id = ?", id).Scan(&producto.id, &producto.nombre, &producto.marca, &producto.precio, &producto.cantidad)

  if producto.id == id {

    return ""

  }else{

    //preparamos nuestro Query
      sentenciaPreparada, err := db.Prepare("INSERT INTO producto (id, nombre, marca, precio, cantidad) VALUES (?,?,?,?,?)")
      if err != nil {
        fmt.Printf("Error en la sentencia insertar: %v", err)
        return ""
      }
      defer sentenciaPreparada.Close()

      // Ejecutar sentencia, un valor por cada '?'
      _, err = sentenciaPreparada.Exec(id, nombre, marca, precio, cantidad)
      if err != nil {
        fmt.Printf("Error al insertar: %v", err)
        return ""
      }

      return "Insertado con exito"
  }
}

func InsertarCajeroTCPyUDP(id int, user string, pass string) (string) {
  
 //abrimos nuestra base de datos
  db, err := obtenerBaseDeDatos()
  //comprobamos que no exista errores
  if err != nil {
    fmt.Printf("Error obteniendo base de datos: %v", err)
    return ""
  }
  // Terminar conexión al terminar función
  defer db.Close()

  var cajero Cajero

  db.QueryRow("SELECT id, usuario FROM cajero where id = ?", id).Scan(&cajero.id, &cajero.usuario)

  if cajero.id == id || cajero.usuario == user{

    return ""

  }else{

    //preparamos nuestro Query
      sentenciaPreparada, err := db.Prepare("INSERT INTO cajero (id, usuario, pass) VALUES (?,?,?)")
      if err != nil {
        fmt.Printf("Error en la sentencia insertar: %v", err)
        return ""
      }
      defer sentenciaPreparada.Close()

      // Ejecutar sentencia, un valor por cada '?'
      _, err = sentenciaPreparada.Exec(id, user, pass)
      if err != nil {
        fmt.Printf("Error al insertar: %v", err)
        return ""
      }

      return "Insertado con exito"
  }
}

func EliminarCajeroTCPyUDP(i int) bool {
    //abrimos nuestra base de datos
  db, err := obtenerBaseDeDatos()
  //comprobamos que no exista errores
  if err != nil {
    fmt.Printf("Error obteniendo base de datos: %v", err)
    return false
  }
  // Terminar conexión al terminar función
  defer db.Close()


  var cajero Cajero
  //verificamos si existe el id
  db.QueryRow("SELECT id FROM cajero where id = ?", i).Scan(&cajero.id)

  //si existe la persona la eliminados
  //de lo contrario se retorna false
  if cajero.id == i {
       //preparamos nuestro Query
    sentenciaPreparada, err := db.Prepare("DELETE FROM cajero WHERE id = ?")
    if err != nil {
      fmt.Printf("Error de sentencia eliminar: %v", err)
      return false
    }
    defer sentenciaPreparada.Close()

    //Eliminamos el dato
    _, err = sentenciaPreparada.Exec(i)
    if err != nil {
      fmt.Printf("Error al Eliminar dato: %v", err)
      return false
    }
    return true
  }else{
    return false
  }


}

func VenderProductoUDP(id int, cantidad int) bool {
  
  //abrimos nuestra base de datos
  db, err := obtenerBaseDeDatos()
  //comprobamos que no exista errores
  if err != nil {
    fmt.Printf("Error obteniendo base de datos: %v", err)
    return false
  }
  // Terminar conexión al terminar función
  defer db.Close()

  var producto Producto
  //verificamos si existe el id
  db.QueryRow("SELECT id, cantidad FROM producto where id = ?", id).Scan(&producto.id,&producto.cantidad)

 
 if producto.id == id {

    if cantidad <= producto.cantidad{

        cantidadNueva := producto.cantidad - cantidad

        if cantidadNueva == 0 {

          sentenciaPreparada, err := db.Prepare("DELETE FROM producto WHERE id = ?")
          if err != nil {
            fmt.Printf("Error de sentencia eliminar: %v", err)
            return false
          }
          defer sentenciaPreparada.Close()

          //Eliminamos el dato
          _, err = sentenciaPreparada.Exec(id)
          if err != nil {
            fmt.Printf("Error al Eliminar dato: %v", err)
            return false
          }
          return true

        }else{

          sentenciaPreparada, err := db.Prepare("UPDATE producto SET cantidad = ? WHERE id = ?")
          if err != nil {
            fmt.Printf("Error en la sentencia update: %v", err)
            return false
          }
          defer sentenciaPreparada.Close()

          // Pasar argumentos en el mismo orden que la consulta
          _, err = sentenciaPreparada.Exec(cantidadNueva, producto.id)
          if err != nil {
            fmt.Printf("Error al Editar: %v", err)
            return false
          }
          return true

       }
      
    }else{

      return false

    }

  }else{

    return false

  }


}

func IniciarSesion(user string, pass string) (string) {
  
  //abrimos nuestra base de datos
  db, err := obtenerBaseDeDatos()
  //comprobamos que no exista errores
  if err != nil {
    fmt.Printf("Error obteniendo base de datos: %v", err)
    return ""
  }
  // Terminar conexión al terminar función
  defer db.Close()

  var evaluarUser string
  var evaluarpass string

  db.QueryRow("SELECT usuario, pass FROM supervisor where usuario = ?", user).Scan(&evaluarUser, &evaluarpass)

  if user == evaluarUser && pass == evaluarpass {
    return ""
  }else{
    return "no existe"
  }

}

func IniciarSesionUDP(user string, pass string) (string){
  
  //abrimos nuestra base de datos
  db, err := obtenerBaseDeDatos()
  //comprobamos que no exista errores
  if err != nil {
    fmt.Printf("Error obteniendo base de datos: %v", err)
    return ""
  }
  // Terminar conexión al terminar función
  defer db.Close()

  var evaluarUser string
  var evaluarPass string

  db.QueryRow("SELECT usuario, pass FROM cajero where usuario = ?", user).Scan(&evaluarUser, &evaluarPass)

  if user == evaluarUser && pass == evaluarPass {
    return ""
  }else{
    return "no existe"
  }

}

///////////////////////////////////////////////////// FUNCION RECORTAR CADENAS /////////////////////////////////////////////
func between(value string, a string, b string) string {
    
    posFirst := strings.Index(value, a)
    if posFirst == -1 {
        return ""
    }
    posLast := strings.Index(value, b)
    if posLast == -1 {
        return ""
    }
    posFirstAdjusted := posFirst + len(a)
    if posFirstAdjusted >= posLast {
        return ""
    }
    return value[posFirstAdjusted:posLast]
}


////////////////////////////////////////////////////// BUCLE DE EJECUCION UDP//////////////////////////////////////////////////////  
 func UDPConexion(conn *net.UDPConn) {
    //bucle de peticiones udp
    for{
         //buffer para los mensajes entrantes
         message := make([]byte, 1024)
         //leer mensajes
         n, addr, err := conn.ReadFromUDP(message)
         if err != nil {
              log.Println(err)
         }


         longitudCadena := between(string(message[:n]), "<", ">")

         longitudEntero, error := strconv.Atoi(longitudCadena)
         if error != nil {
             fmt.Println("Error al convertir: ", error)
          }

         cadena := between(string(message[:n]), "+", "-")

         cadenaS:= len(cadena)

         fmt.Println("No se ha perdido ningun dato importante")

         if  longitudEntero != cadenaS {
           
            //Enviar mensaje al cliente a la direccion del cliente
            message := []byte("Se ha perdido un dato imporatante del mensaje, Reenvie el mensaje\n")
            _, err = conn.WriteToUDP(message, addr)
            if err != nil {
                   log.Println(err)
            }

         }else{

           opcionCadena := between(string(message), "?", "@")


           switch opcionCadena{

              case "Mostrar":
                //cortamos el mensaje para obtener solo el id
                idCadena := between(string(message[:n]), ":", "\r")
                //pasamos la cadena a entero
                idEntero, error := strconv.Atoi(idCadena)
                if error != nil {
                   fmt.Println("Error al convertir: ", error)
                }

                //algunos mensajes en el servidor
                fmt.Println("UDP cliente: ", addr)
                fmt.Println("mensaje de cliente UDP recivido: ", opcionCadena+": "+idCadena+"\n")


                datos:=ObtenerUnProductoTCPyUDP(idEntero)

                if datos == ""{

                   //Enviar mensaje al cliente a la direccion del cliente
                    message := []byte("ID de producto no encontrado\n")
                    _, err = conn.WriteToUDP(message, addr)
                    if err != nil {
                           log.Println(err)
                    }

                }else{

                  message := []byte(datos+"\n")
                    _, err = conn.WriteToUDP(message, addr)
                    if err != nil {
                           log.Println(err)
                  }

                }

                break

              case "Vender":

                 //cortamos el mensaje para obtener solo el id
                idCadena := between(string(message[:n]), ":", "\r")
                cantidadCadena := between(string(message[:n]), "{", "}")
                //pasamos la cadena a entero
                idEntero, error := strconv.Atoi(idCadena)
                if error != nil {
                   fmt.Println("Error al convertir: ", error)
                }

                cantidadEntero, error := strconv.Atoi(cantidadCadena)
                if error != nil {
                   fmt.Println("Error al convertir: ", error)
                }

                //algunos mensajes en el servidor
                fmt.Println("UDP cliente: ", addr)
                fmt.Println("mensaje de cliente UDP recivido: ", opcionCadena+": "+idCadena+"\n")

              
                 if VenderProductoUDP(idEntero, cantidadEntero) == true {
                  //Enviar mensaje al cliente a la direccion del cliente
                  message := []byte("Producto "+idCadena+" vendido con exito!!!\n")
                  _, err = conn.WriteToUDP(message, addr)
                  if err != nil {
                         log.Println(err)
                   }

                 }else{
                    //Enviar mensaje al cliente a la direccion del cliente
                    message := []byte("ID de producto no encontrado o cantidad mayor al inventario\n")
                    _, err = conn.WriteToUDP(message, addr)
                    if err != nil {
                           log.Println(err)
                    }
                 }

                break

               case "iniciar sesion":
                  //cortamos el mensaje para obtener solo el id
                  userCadena := between(string(message[:n]), "{", "}")
                  passCadena := between(string(message[:n]), "[", "]")
                
                  //algunos mensajes en el servidor
                  fmt.Println("UDP cliente: ", addr)
                  fmt.Println("mensaje de cliente UDP recivido: ", opcionCadena+": "+userCadena+"\n")


                  datos:=IniciarSesionUDP(userCadena, passCadena)

                  if datos == ""{

                    message := []byte("Inicio de sesion exitoso!!!\n")
                    _, err = conn.WriteToUDP(message, addr)
                    if err != nil {
                           log.Println(err)
                    }

                  }else{

                  //Enviar mensaje al cliente a la direccion del cliente
                    message := []byte("Cajero no registrado\n")
                    _, err = conn.WriteToUDP(message, addr)
                    if err != nil {
                           log.Println(err)
                    }
                  }

                break


              default:
                conn.Write([]byte("opcion invalida\n"))
                break
          }

         }

         
    }
 }

////////////////////////////////////////////////// BUCLE DE EJECUCION TCP////////////////////////////////////////////////////
 func TcpConexion(ln net.Listener) {
   
    conn, _ := ln.Accept()
    // escuchará el mensaje para procesar que termina en nueva línea (\ n)
    message, _ := bufio.NewReader(conn).ReadString('&')

    opcionCadena := between(string(message), "?", "@")



    switch opcionCadena{

        case "Mostrar":
          numeroCadena := between(string(message), ":", "\r")

          fmt.Print("mensaje de cliente TCP recivido: ",opcionCadena+": "+numeroCadena+"\n")
    
          numeroEntero, error := strconv.Atoi(string(numeroCadena))
          if error != nil {
            fmt.Println("Error al convertir: ", error)
          }


          datos:=ObtenerUnProductoTCPyUDP(numeroEntero)

          if datos == ""{

            conn.Write([]byte("ID no encontrado\n"))

          }else{

            conn.Write([]byte(datos+"\n"))

          }

          break

        case "Eliminar":

          numeroCadena := between(string(message), ":", "\r")

          fmt.Print("mensaje de cliente TCP recivido: ",opcionCadena+": "+numeroCadena+"\n")
    
          numeroEntero, error := strconv.Atoi(string(numeroCadena))
          if error != nil {
            fmt.Println("Error al convertir: ", error)
          }

          CadenaFinal := strconv.Itoa(numeroEntero)


           if EliminarProductoTCPyUDP(numeroEntero) == true {

            conn.Write([]byte("Producto "+CadenaFinal+" eliminado correctamente\n"))

          }else{

            conn.Write([]byte("ID de producto no encontrado\n"))

          }
          break

        case "Insertar":

            idCadena := between(string(message), ":", "\r")
            precioCadena := between(string(message), "{", "}")
            cantidadCadena := between(string(message), "(", ")")

            fmt.Print("mensaje de cliente TCP recivido: ",opcionCadena+": "+idCadena+"\n")
      
            idEntero, error := strconv.Atoi(string(idCadena))
            if error != nil {
              fmt.Println("Error al convertir: ", error)
            }

            precioEntero, error := strconv.Atoi(string(precioCadena))
            if error != nil {
              fmt.Println("Error al convertir: ", error)
            }

            cantidadEntero, error := strconv.Atoi(string(cantidadCadena))
            if error != nil {
              fmt.Println("Error al convertir: ", error)
            }


            nombreCadena := between(string(message), "$", "#")
            marcaCadena := between(string(message), "[", "]")


            evaluar:=InsertarProductoTCPyUDP(idEntero, nombreCadena, marcaCadena, precioEntero, cantidadEntero)

            if evaluar == ""{

              conn.Write([]byte("ID Existente\n"))

            }else{

              conn.Write([]byte("Producto Insertado Correctamente\n"))

            }
          break

        case "iniciar sesion":

            user:= between(string(message), "{", "}")
            pass:= between(string(message), "[", "]")

            fmt.Print("mensaje de cliente TCP recivido: ",opcionCadena+"\n")

            evaluar:=IniciarSesion(user,pass)

            if evaluar == ""{

              conn.Write([]byte("Inicio de sesion exitoso!!!\n"))

            }else{

              conn.Write([]byte("Supervisor no registrado\n"))

            }

          break

        case "Registrar cajero":

            idCadena := between(string(message), ":", "\r")
            userCadena := between(string(message), "[", "]")
            passCadena := between(string(message), "{", "}")

            fmt.Print("mensaje de cliente TCP recivido: ",opcionCadena+": "+idCadena+"\n")
      
            idEntero, error := strconv.Atoi(string(idCadena))
            if error != nil {
              fmt.Println("Error al convertir: ", error)
            }

            evaluar:=InsertarCajeroTCPyUDP(idEntero, userCadena, passCadena)

            if evaluar == ""{

              conn.Write([]byte("Cajero Existente\n"))

            }else{

              conn.Write([]byte("Cajero Registrado Correctamente\n"))

            }
          break

        case "Eliminar Cajero":

            numeroCadena := between(string(message), ":", "\r")

            fmt.Print("mensaje de cliente TCP recivido: ",opcionCadena+": "+numeroCadena+"\n")
      
            numeroEntero, error := strconv.Atoi(string(numeroCadena))
            if error != nil {
              fmt.Println("Error al convertir: ", error)
            }

            CadenaFinal := strconv.Itoa(numeroEntero)


             if EliminarCajeroTCPyUDP(numeroEntero) == true {

              conn.Write([]byte("Cajero "+CadenaFinal+" eliminado correctamente\n"))

            }else{

              conn.Write([]byte("ID de Cajero no encontrado\n"))

            }
          break

        default:
          conn.Write([]byte("opcion invalida\n"))
          break
    }

 }

///////////////////////////////////////////////////// BUCLE DE EJECUCION MULTICAST ////////////////////////////////////////////////////
 type Dir struct{
  direccion *net.UDPAddr
  nick string
  
 }

 var Arr []Dir

 func  MulticastConexion(l *net.UDPConn) {
   
  for {

    b := make([]byte, 1024)
    n, src, err := l.ReadFromUDP(b)
    if err != nil {
      log.Fatal("ReadFromUDP failed:", err)
    }

    opcion := between(string(b[:n]), "{", "}")


    if opcion == "union al grupo" {

      nick := between(string(b[:n]), ":", "?")

      var direc Dir

      direc.direccion = src
      direc.nick = nick

      Arr = append(Arr, direc)

      fmt.Print("Multicast cliente: ", nick+ " unido correctamente al grupo"+"\n")

      message := []byte(nick+" unido correctamente al grupo\n")
        _, err = l.WriteToUDP(message, src)
        if err != nil {
               log.Println(err)
       }

    }else{

      nick := between(string(b[:n]), ":", "?")
      mensaje := between(string(b[:n]), "[", "]")

      fmt.Print("\n")
      fmt.Print("Multicast cliente: ", nick+"\n")
      fmt.Print("mensaje de cliente recivido: ", mensaje+"\n")

       for _, item := range Arr {
          if item.nick != nick {
              //Enviar mensaje al cliente a la direccion del cliente
            message := []byte(nick+": "+mensaje)
            _, err = l.WriteToUDP(message, item.direccion)
            if err != nil {
                   log.Println(err)
            }
          }
       }
    }
    
  }

 }


func main() {

  //////////////////////////////////////////////////// CONEXION A LA BASE DE DATOS ///////////////////////////////////////////
  // Verificamos la conexion a la base de datos
  db, err := obtenerBaseDeDatos()
  if err != nil {
    fmt.Printf("Error obteniendo base de datos: %v", err)
    return
  }
  // Terminar conexión al terminar función
  defer db.Close()

  // Ahora vemos si tenemos conexión (opcional)
  err = db.Ping()
  if err != nil {
    fmt.Printf("Error conectando: %v", err)
    return
  }
  // Listo, aquí ya podemos usar a db!
  fmt.Printf("Conectado correctamente a la BD \n")


//////////////////////////////////////////////////////// SERVIDOR TCP  /////////////////////////////////////////////////////////
  // escuchar las interfaces TCP
  go net.Listen("tcp", ":8080")
  ln, _ := net.Listen("tcp", ":8080")
  defer ln.Close()

//////////////////////////////////////////////////////// SERVIDOR UDP  /////////////////////////////////////////////////////////
  //escuchar las interfaces UDP
  hostName := "localhost"
  portNum := "6000"
  service := hostName + ":" + portNum
  udpAddr, err := net.ResolveUDPAddr("udp4", service)
  if err != nil {
         log.Fatal(err)
  }
  go net.ListenUDP("udp", udpAddr)
  sudp, err := net.ListenUDP("udp", udpAddr)

  if err != nil {
         log.Fatal(err)
  }
  defer sudp.Close()

//////////////////////////////////////////////////////// SERVIDOR MULTICAST  /////////////////////////////////////////////////////////
  srvAddr := "224.0.0.1:9999"

  addr, err := net.ResolveUDPAddr("udp", srvAddr)
  if err != nil {
    log.Fatal(err)
  }
  l, err := net.ListenMulticastUDP("udp", nil, addr)

  fmt.Println("Servidores: Multicast pueto:9999, tcp puerto:8080, udp pueto: 6000 preparados...\n")

////////////////////////////////////////////////  CICLO DE EJECUCION TCP UDP MULTICAST /////////////////////////////////////////////////////////

  // bucle de ejecución para escuchar peticiones TCP y UDP para siempre 
  for {

    // udp
    // blucle de ejecucion para UDP en una gorutine en un hilo aparte
    go UDPConexion(sudp)

    //Multicast
    go MulticastConexion(l)

    //tcp
    TcpConexion(ln)

  }

}