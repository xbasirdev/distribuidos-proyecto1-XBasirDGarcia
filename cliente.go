package main

import (
	"log"
	"net"
	"fmt"
	"bufio"
    "os"
    "strings"
    "strconv"
    "time"

)

func main() {

	for{
		// leer en la entrada de entrada estándar
    	reader := bufio.NewReader(os.Stdin)
    	fmt.Print("\n")
		fmt.Print("APLICACION FARMACIA\n")
		fmt.Print("ELIJA UNA OPCION\n")
	    fmt.Print("1 --> Supervisor TCP \n")
	    fmt.Print("2 --> Cajero UDP \n")
	    fmt.Print("3 --> Chat Multidifucion \n")
	  	fmt.Print("opcion: ")
        opc, _ := reader.ReadString('\n')

        switch opc {

            case "1\r\n":

            	  var cerrar string
            	  cerrar=""

                  for{
	                  	if cerrar=="salir"{
					       break
					     }
				    // conectarse a este zócalo
					ts := time.Second * 5
				    conn, err:= net.DialTimeout("tcp", "localhost:8080", ts)
				     if err != nil {
				        log.Println(err)
				     }
                   //cerramos la conexion
				    defer conn.Close()

				    timeoutDuration := time.Second * 5
					
				    // leer en la entrada de entrada estándar
				    reader := bufio.NewReader(os.Stdin)

				    //validar sesion
				    fmt.Print("\n")
				    fmt.Print("INICIAR SESION\n")
				    fmt.Print("Usuario: ")
				    user, _ := reader.ReadString('\n')
				    fmt.Print("Password: ")
				    pass, _ := reader.ReadString('\n')

				    userFinal := strings.TrimRight(user, "\r\n")
				    passFinal := strings.TrimRight(pass, "\r\n")
				    // enviar a socket
				    conn.SetWriteDeadline(time.Now().Add(timeoutDuration))
				    conn.Write([]byte("{"+userFinal+"}"+"["+passFinal+"]"+"?"+"iniciar sesion"+"@"+"&"))
				    
				    conn.SetReadDeadline(time.Now().Add(timeoutDuration))
				    message, _ := bufio.NewReader(conn).ReadString('\n')
				    fmt.Print("mensaje del servidor TCP: "+message+"\n")

				    if message == "Inicio de sesion exitoso!!!\n"{
				       
				        for {

					        if cerrar=="salir"{
					        	break
					        }
				        // conectarse a este zócalo
				        conn, _ := net.Dial("tcp", "127.0.0.1:8080")
				        //cerramos la conexion
				        defer conn.Close()
				       	
				       	fmt.Print("\n")
				        fmt.Print("ELIJA UNA OPCION\n")
				        fmt.Print("1 --> Mostrar Producto \n")
				        fmt.Print("2 --> Eliminar Producto \n")
				        fmt.Print("3 --> Insertar Producto \n")
				        fmt.Print("4 --> Registrar Cajero \n")
				        fmt.Print("5 --> Eliminar Cajero \n")
				        fmt.Print("6 --> salir \n")
				        fmt.Print("opcion: ")
				        opc, _ := reader.ReadString('\n')


				        switch opc {

				            case "1\r\n":

				                fmt.Print("\n")
				                fmt.Print("MOSTRAR PRODUCTO\n")
				                fmt.Print("Mostra producto con ID: ")
				                text, _ := reader.ReadString('\n')
				                
				                // enviar a socket
				                conn.Write([]byte(":"+text+"?"+"Mostrar"+"@"+"&"))
				                
				                // escucha la respuesta
				                message, _ := bufio.NewReader(conn).ReadString('\n')
				                fmt.Print("mensaje del servidor TCP: "+message+"\n")

				              break

				            case "2\r\n":

				                fmt.Print("\n")
				                fmt.Print("ELIMINAR PRODUCTO\n")
				                fmt.Print("Desea eliminar el producto con ID: ")
				                text, _ := reader.ReadString('\n')
				                
				                // enviar a socket
				                conn.Write([]byte(":"+text+"?"+"Eliminar"+"@"+"&"))
				                
				                // escucha la respuesta
				                message, _ := bufio.NewReader(conn).ReadString('\n')
				                fmt.Print("mensaje del servidor TCP: "+message+"\n")

				              break

				            case "3\r\n":

				                fmt.Print("\n")
				                fmt.Print("INSERTAR PRODUCTO\n")
				                fmt.Print("Ingrese ID: ")
				                id, _ := reader.ReadString('\n')
				                fmt.Print("Ingrese Nombre: ")
				                nombre, _ := reader.ReadString('\n')
				                fmt.Print("Ingrese Marca: ")
				                marca, _ := reader.ReadString('\n')
				                fmt.Print("Ingrese Precio: ")
				                precio, _ := reader.ReadString('\n')
				                fmt.Print("Ingrese Cantidad: ")
				                cantidad, _ := reader.ReadString('\n')

				                nombreFinal := strings.TrimRight(nombre, "\r\n")
				                marcaFinal := strings.TrimRight(marca, "\r\n")
				                precioFinal := strings.TrimRight(precio, "\r\n")
				                cantidadFinal := strings.TrimRight(cantidad, "\r\n")
				                
				                // enviar a socket
				                conn.Write([]byte(":"+id+"?"+"Insertar"+"@"+"$"+nombreFinal+"#"+"["+marcaFinal+"]"+"{"+precioFinal+"}"+"("+cantidadFinal+")"+"&"))
				                
				                // escucha la respuesta
				                message, _ := bufio.NewReader(conn).ReadString('\n')
				                fmt.Print("mensaje del servidor TCP: "+message+"\n")

				              break

				             case "4\r\n":

				                fmt.Print("\n")
				                fmt.Print("REGISTRAR CAJERO\n")
				                fmt.Print("Ingrese id: ")
				                id, _ := reader.ReadString('\n')
				                fmt.Print("Ingrese Usuario: ")
				                user, _ := reader.ReadString('\n')
				                fmt.Print("Ingrese Password: ")
				                pass, _ := reader.ReadString('\n')

				                userFinal := strings.TrimRight(user, "\r\n")
				                passFinal := strings.TrimRight(pass, "\r\n")
				                // enviar a socket
				                conn.Write([]byte(":"+id+"?"+"Registrar cajero"+"@"+"["+userFinal+"]"+"{"+passFinal+"}"+"&"))
				                
				                // escucha la respuesta
				                message, _ := bufio.NewReader(conn).ReadString('\n')
				                fmt.Print("mensaje del servidor TCP: "+message+"\n")

				              break

				             case "5\r\n":

				                fmt.Print("\n")
				                fmt.Print("ELIMINAR CAJERO\n")
				                fmt.Print("Desea eliminar el Cajero con ID: ")
				                text, _ := reader.ReadString('\n')
				                
				                // enviar a socket
				                conn.Write([]byte(":"+text+"?"+"Eliminar Cajero"+"@"+"&"))
				                
				                // escucha la respuesta
				                message, _ := bufio.NewReader(conn).ReadString('\n')
				                fmt.Print("mensaje del servidor TCP: "+message+"\n")

				              break

				             case "6\r\n":
				             	cerrar = "salir"
				              break

				            default:
				              fmt.Print("Elija una opcion correcta")
				              break
				            }

				        }

				    }else{
				       fmt.Print("Supervisor no registrado\n")
				       break
				    }

				  }

              break

            case "2\r\n":

            	var cerrar string
            	cerrar=""

                for{

        		if cerrar=="salir"{
				     break
				}

		        hostName := "localhost"
		        portNum := "6000"
		        service := hostName + ":" + portNum
		        RemoteAddr, err := net.ResolveUDPAddr("udp", service)
		        // conectarse a este zócalo
		        conn, err := net.DialUDP("udp", nil, RemoteAddr)
		        if err != nil {
		                 log.Fatal(err)
		        }
		        // cerramos la conexion
		        defer conn.Close()

		        // leer en la entrada de entrada estándar
		        reader := bufio.NewReader(os.Stdin)

		        var cadenaFinal string

		        for{

		        	fmt.Print("\n")
			        fmt.Print("INICIAR SESION\n")
			        fmt.Print("Usuario: ")
			        user, _ := reader.ReadString('\n')
			        fmt.Print("Password: ")
			        pass, _ := reader.ReadString('\n')

			        userFinal := strings.TrimRight(user, "\r\n")
			        passFinal := strings.TrimRight(pass, "\r\n")

			        cadena := "{"+userFinal+"}"+"["+passFinal+"]"+"?"+"iniciar sesion"+"@"

			        i := len(cadena)

			        iCadena := strconv.Itoa(i)


			        cadenaFinal = "+"+"{"+userFinal+"}"+"["+passFinal+"]"+"?"+"iniciar sesion"+"@"+"-"+"<"+iCadena+">"

			        if len(cadenaFinal) <= 1024 {
			        	break
			        }else{
			        	fmt.Println("Datos superan el tamaño del buffer asignados en el sistema")
			        }
		        }

		        conn.Write([]byte(cadenaFinal))

		        // escucha la respuesta
		        buffer := make([]byte, 1024)
		        n, addr, err := conn.ReadFromUDP(buffer)
		        if err != nil {
		                 log.Println(err)
		        }
		        fmt.Println("UDP Server : ", addr)
		        fmt.Println("mensaje del servidor UDP: ", string(buffer[:n]))

		        if string(buffer[:n]) == "Inicio de sesion exitoso!!!\n"{

		            for {
		            	if cerrar=="salir"{
				        	break
				        }
		                hostName := "localhost"
		                portNum := "6000"
		                service := hostName + ":" + portNum
		                RemoteAddr, err := net.ResolveUDPAddr("udp", service)
		                // conectarse a este zócalo
		                conn, err := net.DialUDP("udp", nil, RemoteAddr)
		                if err != nil {
		                         log.Fatal(err)
		                }
		                // cerramos la conexion
		                defer conn.Close()


		                fmt.Print("\n")
		                fmt.Print("ELIJA UNA OPCION\n")
		                fmt.Print("1 --> Consultar Producto \n")
		                fmt.Print("2 --> Vender Producto \n")
		                fmt.Print("3 --> salir \n")
		                fmt.Print("opcion: ")
		                opc, _ := reader.ReadString('\n')

		                switch opc {

		                    case "1\r\n":

		                    	var cadenaFinal string
		                    	for{

		                    		fmt.Print("\n")
			                        fmt.Print("CONSULTAR PRODUCTO\n")
			                        fmt.Print("Consultar producto con id: ")
			                        text, _ := reader.ReadString('\n')

		                    		cadena := ":"+text+"?"+"Mostrar"+"@"

		                    		i := len(cadena)

			                        iCadena := strconv.Itoa(i)

						            cadenaFinal = "+"+":"+text+"?"+"Mostrar"+"@"+"-"+"<"+iCadena+">"

							        if len(cadenaFinal) <= 1024 {
							        	break
							        }else{
							        	fmt.Println("Datos superan el tamaño del buffer asignados en el sistema")
							        }
		                    	}
		                        
		                        // enviar a socket
		                        _, err = conn.Write([]byte(cadenaFinal))
		                        if err != nil {
		                                 log.Println(err)
		                        }
		                        
		                        // escucha la respuesta
		                        buffer := make([]byte, 1024)
		                        n, addr, err := conn.ReadFromUDP(buffer)
		                        if err != nil {
		                                 log.Println(err)
		                        }
		                        fmt.Println("UDP Server : ", addr)
		                        fmt.Println("mensaje del servidor UDP: ", string(buffer[:n]))


		                      break

		                    case "2\r\n":
		                    	var cadenaFinal string
		                    	for{
		                    		fmt.Print("\n")
			                        fmt.Print("VENDER PRODUCTO\n")
			                        fmt.Print("Ingrese el id: ")
			                        id, _ := reader.ReadString('\n')
			                        fmt.Print("Ingrese la cantidad a vender: ")
			                        cantidad, _ := reader.ReadString('\n')

			                        cantidadFinal := strings.TrimRight(cantidad, "\r\n")

			                        cadena := ":"+id+"?"+"Vender"+"@"+"{"+cantidadFinal+"}"

			                        i := len(cadena)

			                        iCadena := strconv.Itoa(i)

			                        cadenaFinal = "+"+":"+id+"?"+"Vender"+"@"+"{"+cantidadFinal+"}"+"-"+"<"+iCadena+">"

							        if len(cadenaFinal) <= 1024 {
							        	break
							        }else{
							        	fmt.Println("Datos superan el tamaño del buffer asignados en el sistema")
							        }
		                    	}
		                        // enviar a socket
		                        _, err = conn.Write([]byte(cadenaFinal))
		                        if err != nil {
		                                 log.Println(err)
		                        }

		                        // escucha la respuesta
		                        buffer := make([]byte, 1024)
		                        n, addr, err := conn.ReadFromUDP(buffer)
		                        if err != nil {
		                                 log.Println(err)
		                        }
		                        fmt.Println("UDP Server : ", addr)
		                        fmt.Println("mensaje del servidor UDP: ", string(buffer[:n]))

		                      break

		                    case "3\r\n":
				             	cerrar = "salir"
				              break

		                    default:
		                      fmt.Print("Elija una opcion correcta")
		                      break
		                }

		            }

		        }else{

		            fmt.Print("Cajero no registrado\n") 
		        }

		    }

              break

            case "3\r\n":

         		hostName := "localhost"
			    portNum := "9999"
			    service := hostName + ":" + portNum
			    RemoteAddr, err := net.ResolveUDPAddr("udp", service)
			    // conectarse a este zócalo
			    conn, err := net.DialUDP("udp", nil, RemoteAddr)
			    if err != nil {
			             log.Fatal(err)
			    }
			    // cerramos la conexion
			    defer conn.Close()
			    reader := bufio.NewReader(os.Stdin)

			    fmt.Print("\n")
			    fmt.Print("ingresa tu nick para unirte al chat: ")
			    nick, _ := reader.ReadString('\n')

			    nickFinal := strings.TrimRight(nick, "\r\n")

			    _, err = conn.Write([]byte(":"+nickFinal+"?"+"{union al grupo}"))
			    if err != nil {
			             log.Println(err)
			    }

			    // escucha la respuesta
			    buffer := make([]byte, 1024)
			    n, addr, err := conn.ReadFromUDP(buffer)
			    if err != nil {
			             log.Println(err)
			    }
			    fmt.Print("\n")
			    fmt.Print("Multicast Server : ", addr)
			    fmt.Print("\n")
			    fmt.Print("mensaje del servidor Multicast: ", string(buffer[:n]))
			    fmt.Print("\n")
			    // escucha la respuesta
	
		    	go escuchar(conn)

				escribir(conn, nickFinal)

              break

            default:
              fmt.Print("Elija una opcion correcta")
              break
            }

	}

}

func escuchar(conn *net.UDPConn){

	for{
		buffer := make([]byte, 1024)
	    n, _, err := conn.ReadFromUDP(buffer)
	    if err != nil {
	             log.Println(err)
	    }

	    fmt.Println(string(buffer[:n]))
	 }
}

func escribir(conn *net.UDPConn, nickFinal string){
	reader := bufio.NewReader(os.Stdin)
	for{
	
        text, _ := reader.ReadString('\n')

        textFinal := strings.TrimRight(text, "\r\n")
        
        // enviar a socket
        _, err := conn.Write([]byte(":"+nickFinal+"?"+"["+textFinal+"]"))
        if err != nil {
                 log.Println(err)
        }
	}
}
