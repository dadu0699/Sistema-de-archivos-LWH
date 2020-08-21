package menu

import (
	"Sistema-de-archivos-LWH/analisis/lexico"
	"Sistema-de-archivos-LWH/analisis/sintactico"
	"Sistema-de-archivos-LWH/util"
	"Sistema-de-archivos-LWH/util/archivo"
	"fmt"
	"strings"
)

// Interfaz de línea de comandos
func Interfaz() {
	fmt.Println("╔══════════════════════╗")
	fmt.Println("║      Bienvenido      ║")
	fmt.Println("╚══════════════════════╝")
	fmt.Print(">> ")
	str := util.LecturaTeclado()

	for !strings.EqualFold(str, "exit") {
		if strings.EqualFold(str, "pause") {
			str = util.LecturaTeclado()
		} else {
			listaTokens, listaErrores := lexico.Scanner(str)
			if len(listaErrores) > 0 {
				fmt.Println(">> 'La entrada contiene errores lexicos'")
				fmt.Println(">> LISTADO DE ERRORES:", listaErrores)
				fmt.Println(">> LISTADO DE TOKENS:", listaTokens)
			} else if len(listaTokens) > 0 {
				if listaTokens[0].GetTipo() == "EXEC" {
					if len(listaTokens) == 4 && listaTokens[1].GetTipo() == "-PATH" &&
						listaTokens[2].GetTipo() == "ASIGNACION" &&
						(listaTokens[3].GetTipo() == "RUTA" ||
							listaTokens[3].GetTipo() == "CADENA") {
						archivo.Leer(listaTokens[3].GetValor())
					} else {
						fmt.Println(">> 'ERROR DE INSTRUCCION'")
					}
				} else {
					// INICIO ANALISIS SINTACTICO
					sintactico.Analizar(listaTokens)
				}
			}
		}
		fmt.Print(">> ")
		str = util.LecturaTeclado()
	}
}
