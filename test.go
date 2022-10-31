package main

import (
		"time"
	    "strings"
	    "strconv"
	    "bufio"
	    "os"
	    "fmt"
)

var almacenar_tops_en_archivo [60000]string /*        TOPS_lines          */

func escribir_tops_a_base_de_datos() {
   var i, j int
   var n int = 10000 /*        TOPS_lines/6         */
   var cuenta int = 0
   for  i = 0; i < n; i++ {
      for j = 0; j < 6; j++ {
         almacenar_tops_en_archivo[cuenta] = tops[i][j]
         cuenta++
      }
   }
}

func writeLines() error {
    file, err := os.Create("database/TOPS")
    if err != nil {
        return err
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    for _, line := range almacenar_tops_en_archivo {
        fmt.Fprintln(w, line)
    }
    return w.Flush()
}

var tops [10000][6]string = [10000][6]string{
	{"1", "Username", "125.00 (wpm)", "2022 (fecha)", "0", "123456 (User.ID)"},
	{"2", "0", "0", "0", "0", "0"},
	{"3", "0", "0", "0", "0", "0"},
	{"4", "0", "0", "0", "0", "0"},
	{"5", "0", "0", "0", "0", "0"},

	{"1", "0", "0", "0", "0", "0"},
	{"2", "0", "0", "0", "0", "0"},
	{"3", "0", "0", "0", "0", "0"},
	{"4", "0", "0", "0", "0", "0"},
	{"5", "0", "0", "0", "0", "0"},
}

var levels [500][2]string = [500][2]string{
	{"12345678910...", "0"},
	{"87652345654...", "0"},

}
const INVISIBLE string = "​"

func how_many_texts() int {
	var cuenta int
	for i := 1; i <= len(textos); i++ {
	    if textos[i] != "" {
	    	cuenta++
	    } else { return cuenta+1 }
	}
	return 0
}

func split_curr() string {
	arrayed := strings.Split(current_text, " ")
	return arrayed[0] + " " + arrayed[1] + " " + arrayed[2] + " " + arrayed[3] + " " + arrayed[4] + " " + arrayed[5] + "..."
}

func is_illegal(s string) bool{
	x := strings.Contains(s, INVISIBLE)
	return x
}

func split(tosplit string, sep rune) []string {
	var fields []string

	last := 0
	for i, c := range tosplit {
		if c == sep {
			fields = append(fields, string(tosplit[last:i]))
			last = i + 1
		}
	}
	fields = append(fields, string(tosplit[last:]))

	return fields
}

func CountWords(s string) int {
	return len(strings.Fields(s))
}

func average_word_length(s string) float64 {
	arrayed := strings.Split(s, " ")
	var total = 0
	for i := 0; i < len(arrayed); i++ {
		total += len([]rune(arrayed[i]))
	}
	var x float64 = float64(total)/float64(len(arrayed))
	return x
}

var random int
var current_text = textos[0]
var is_started bool = false
var time_elapsed float64 = 0

var time_soon bool = true

var wpm float64 = 0

var start_author string

var result_splited[]string

var conseguiste_alguna_marca bool = false
var conseguiste_superar_alguna_marca bool = true

func start() {
	time_elapsed = 0
	time_soon = true
	if is_started == false {
		is_started = true
		time.Sleep(500 * time.Millisecond)
		for is_started {
			time.Sleep(1 * time.Millisecond)
			time_elapsed = time_elapsed + 1
			if time_elapsed == 6000 {
				time_soon = false
			}
			if time_elapsed == 180000 {
				is_started = false
			}
		}
	}

}

func calculate_wpm() {
	var length = len([]rune(current_text))/* - (CountWords(current_text)) */
	var length_as_a_float float64 = float64(length)
	wpm = length_as_a_float / 5 / time_elapsed * 60000
}

var errores int = 0
var errores_s string

var lista_errores string

func calculate_errors(sent string, current string) {
	
	/* reseting */
	errores = 0
	lista_errores = ""

	A := sent
	sent_arrayed := strings.Split(A, " ")

	B := current
	text_arrayed := strings.Split(B, " ")

	if len(text_arrayed) == len(sent_arrayed) {
		for i := 0; i < len(text_arrayed); i++ {
				if text_arrayed[i] != sent_arrayed[i] {
					if lista_errores != "" {
						lista_errores = lista_errores + ", " + sent_arrayed[i]
						errores++
					} else {
						lista_errores = sent_arrayed[i]
						errores++
					}
				}
			}
	}

	if  len(text_arrayed) > len(sent_arrayed) {

	}


	errores_s = strconv.FormatInt(int64(errores), 10)
}

 func printslice(slice []string) {
 	fmt.Println("slice = ", slice)
 }

 func dup_count(list []string) map[string]int {

 	duplicate_frequency := make(map[string]int)

 	for _, item := range list {
 		// check if the item/element exist in the duplicate_frequency map

 		_, exist := duplicate_frequency[item]

 		if exist {
 			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
 		} else {
 			duplicate_frequency[item] = 1 // else start counting from 1
 		}
 	}
 	return duplicate_frequency
 }

var palabras []string = []string{"la","ser","y","a","de","a","en","yo","tú","Glevion","tener","a","sorpresa","por","hacer","tengo","con","en","Este","nuevo","nosotros","que","no","pero","ellos","decir","a","What","su","desde","Vamos","o","por","obtener","ella","mi","perro","as","saber","si","yo","su","allá","quien","sobre","sus","voluntad","sur","haría","hacer","sólo","arriba","pensar","tiempo","allá","ver","su","as","afuera","una","comer","gente","tomar","año","aél","aellos","alguno","desear","cómo","cuando","cual","ahora","megusta","otro","pudo","nuestro","dentro","aquí","después","que","Mira","camino","más","estosestos","No","cosa","bien","porque","además","dos","usar","decir","bueno","primero","hombre","día","encontrar","dar","más","nuevo","una","anosotros","ningún","aquellos","very","su","necesitar","espalda","allá","debería","incluso","solamente","muchos","DeVerdad","trabajar","vida","por","Correcto","abajo","en","probar","dejar","algunacosa","también","llamar","mujer","mayo","todavía","mediante","significar","After","nunca","No","mundo","en","sentir","sí","estupendo","ultimo","niño","Oh","sobre","pedir","cuando","as","escuela","estado","mucho","hablar","afuera","mantener","abandonar","poner","megusta","ayuda","grande","dónde","mismo","allá","propio","tiempo","comienzo","Tres","alto","cada","otro","convertirseen","lamayoría","Entre","suceder","Familia","sobre","presidente","antiguo","es","casa","Espectáculo","otravez","estudiantes","SUDOESTE","parecer","puedeque","parte","oír","su","lugar","problema","dónde","creer","país","siempre","semana","punto","mano","apagado","desempeñar","giro","pocos","grupo","tal","contra","correr","tipo","sobre","caso","pregunta","trabajar","noche","Vivo","juego","número","escribe","","sinque","dinero","lote","lamayoría","libro","sistema","gobierno","Siguiente","ciudad","empresa","historia","EsteDia","trabajo","Muevete","deber","malo","amigo","durante","empezar","love","cada","mantener","diferente","americano","poco","antesde","algunavez","palabra","hecho","Correcto","leer","Anything","nada","Porsupuesto","pequeña","mes","programa","quizás","Correcto","pordebajo","negocio","hogar","tipo","deténgase","pagar","estudiar","yaque","tema","batata","ocurrencia","habitación","porcentaje","lejos","lejos","ley","Realmente","largo","aunque","proveer","losé","energía","niños","guerra","comprender","cabeza","madre","real","mejor","equipo","ojo","largo","largo","lado","inodoro","joven","Espere","OK","ambascosas","aún","After","reunir","Servicio","área","importante","persona","Oye","gracias","mucho","alguien","final","cambio","sinembargo","solamente","alrededor","hora","todo","nacional","cuatro","línea","muchachas","alrededor","reloj","Hastaque","padre","sentar","crear","información","coche","aprender","elmenos","ya","matar","minuto","fiesta","incluir","puesto","juntos","espalda","seguir","salud","recuerda","confrecuencia","razón","hablar","ago","establecer","negro","miembro","comunidad","once","Social","noticias","permitir","victoria","mono","Plomo","Seguir","yasea","suficiente","gastar","nivel","poder","política","casi","chico","Universidad","antesde","Quédate","agregar","luego","cambio","cinco","probablemente","centro","entre","cara","público","morir","alimento","más","historia","comprar","resultado","Mañana","apagado","padre","oficina","curso","enviar","investigar","caminar","puerta","blanco","varios","tribunal","hogar","crecer","mejor","abierto","momento","incluido","considerar","ambascosas","tal","poco","dentrode","segundo","latidos","calle","libre","mejor","todos","políticas","mesa","perdón","cuidado","bajo","hombre","porfavor","esperar","CIERTO","proceso","profesor","datos","oferta","muerte","entero","experiencia","plan","fácil","educación","construir","suponer","fallar","élmismo","años","difícil","sentido","atravésde","Espectáculo","temprano","Universidad","música","Aparecer","mente","clase","policía","usar","efecto","temporada","impuesto","corazón","son","Arte","Posible","atender","descanso","apesardeque","final","mercado","incluso","aire","fuerza","requerir","pies","arriba","escuchar","aceptar","respectivamente","cualquiera","bebé","equivocado","love","Corte","decidir","republicano","completo","detrás","pasar","interesante","Algunasveces","laseguridad","comer","reporte","control","Velocidad","local","sugerir","reporte","nación","vender","acción","apoyo","esposa","decisión","recibir","valor","base","elegir","teléfono","Gracias","evento","conducir","fuerte","alcanzar","permanecer","explique","situ","pegar","jalar","iglesia","modelo","quizás","relaciones","seis","multa","películas","campo","elevar","menos","jugador","pareja","millón","ellos","registro","especialmente","diferencia","Luz","desarrollo","federal","anterior","role","bonita","mímismo","vista","precio","esfuerzo","bonito","retirar","alolargode","voz","finalmente","Departamento","o","hacia","líder","porque","fotos","tenerpuesto","espacio","proyectos","devolver","posición","especial","millón","película","necesitar","importante","escribe","pueblo","artículos","carretera","forma","oportunidad","","economía","situación","elegir","práctica","causa","contento","Ciencias","unirse","enseñar","temprano","desarrollar","compartir","túmismo","llevar","claro","hermano","asunto","muerto","imagen","estrella","costo","simplemente","correo","sociedad","fotos","pieza","papel","energía","personal","edificio","militar","abierto","médico","actividad","exactamente","americano","mitad","señorita","evidencia","producto","darsecuenta","ahorrar","brazo","tecnología","captura","comentario","Mira","término","color","cubrir","describe","Adivinar","elección","fuente","mamá","pronto","director","internacional","normas","Campaña","terrestre","elecciones","cara","oh","controlar","página","lucha","símismo","prueba","paciente","produce","Porsupuesto","loquesea","mitad","video","apoyo","lanzar","tercera","cuidado","descansar","reciente","disponible","pasos","Listo","oportunidad","oficial","aceite","llamar","organización","caracteres","único","Actual","probable","condado","futuro","papá","cuyo","menos","disparar","industria","segundo","lista","general","cosas","figura","atención","olvidar","riesgo","No","enfoque","pantalonescortos","fuego","perro","red","pelo","punto","condición","pared","hija","antesde","acuerdo","autor","verdad","al","esposo","período","serie","ordenar","oficial","cerca","tierra","Nota","ordenadores","pensamiento","economía","meta","banco","comportamiento","sonido","acuerdo","seguramente","porpoco","aumentar","Actuar","norte","bien","sangre","cultura","medicamento","okey","todoelmundo","partesuperior","difícil","cerca","idioma","ventanas","respuesta","población","mentir","árbol","parque","trabajadores","dibujar","plan","soltar","empujar","tierra","causa","pero","privado","estanoche","laraza","que","carta","otro","pistola","simple","curso","preguntarse","involucrar","infierno","pobre","cada","responder","naturaleza","administración","común","No","difícil","mensaje","canción","disfrutar","Similarmente","congreso","ataque","pasado","caliente","buscar","Monto","análisis","Tienda","defensa","Factura","megusta","célula","lejos","actuación","hospital","cama","junta","proteger","siglo","elverano","material","individual","recientemente","ejemplo","representar","llenar","estado","lugar","animal","fallar","factor","natural","señor","agencia","normalmente","importante","ayuda","capacidad","mil","declaración","entero","democrático","piso","enserio","carreraprofesional","dolares","votar","sexy","comparar","sur","adelante"}

var current_text_10ff string

var textos [1000]string = [1000]string{
	"Escucha, las reglas propias... se tratan de decidir conseguir algo usando medios y maneras propias para conseguirlo. Por eso decimos que son nuestras reglas. Precisamente por eso podemos afrontar sinceramente los desafíos y darlo todo. Y si fracasamos, hay que retomar la práctica y soportar duros entrenamientos para lograrlo. Y así, dedicándote a ello, creas tus propias reglas.",
	"Pocos ven lo que somos, pero todos ven lo que aparentamos. Nunca intentes ganar por la fuerza lo que puede ser ganado por la mentira. Nunca se logró nada grandioso sin peligro.",
	"Tal vez solo nos espere un camino oscuro por delante. Pero aún así necesitas creer y seguir adelante. Cree en que las estrellas iluminarán tu camino, incluso un poco. Vamos, ¡emprendamos una aventura!",
	"Sólo soy verdaderamente libre cuando todos los seres humanos que me rodean, hombres y mujeres, son igualmente libres, de manera que cuanto más numerosos son los hombres libres que me rodean y más profunda y más amplia es su libertad, más extensa, más profunda y más amplia viene a ser mi libertad.",
	"Cada tipo que crece podría ser un repetidor, un pequeño maestro, el desencadenante de una reacción en cadena que en sí misma es capaz de cambiar el mundo.",
	"Es natural que estas cosas se produzcan necesariamente así a partir de tales hombres. Y el que así no lo acepta, pretende que la higuera no produzca su zumo.",
	"Su cabello empezó a caer, sus ojos a derretirse, su boca se abría y veía solamente mierda. El castaño desaparece, el blanco se arruga; una dulce niña sin más. Y la veía, y no miraba; y en ella pensaba, se escondía. En ese instante se fue la furia, el dolor cambió a escalofríos, el desaliento en calma; y ahora moriré solo como siempre lo soñé y como siempre lo esperé.",
	"También sabía que era uno de los músicos del coro, y aunque nunca se había atrevido a levantar la vista para comprobarlo durante la misa, un domingo tuvo la revelación de que mientras los otros instrumentos tocaban para todos, el violín tocaba solo para ella.",
	"No me mires que nos miran, nos miran que nos miramos. Miremos que no nos miren, y cuando no nos miren, nos miraremos. Porque si nos miran que nos miramos, miran que nos amamos.",
	"En el símbolo yin-yang hay un punto blanco sobre la parte negra y un punto negro sobre la parte blanca. Esto sirve para ilustrar el equilibrio en la vida, porque nada puede sobrevivir mucho tiempo yendo a cualquiera de los extremos, tanto si es el yin puro (negativismo) como el yang puro (positivismo). El calor extremo mata, como también lo hace el frío extremo. Ningún extremo violento perdura. Nada permanece, excepto la sobria moderación.",
	"Hace mucho tiempo que me hubiera suicidado de no haber leído en alguna parte que es un pecado quitarse voluntariamente la vida mientras pueda hacerse todavía una buena acción. La vida es hermosa, pero la mía está envenenada para siempre.",
	"En un lugar de la Mancha, de cuyo nombre no quiero acordarme, no hace mucho tiempo que vivía un hidalgo de los de lanza en astillero, adarga antigua, rocín flaco y galgo corredor. Una olla de algo más vaca que carnero, salpicón las más noches, duelos y quebrantos los sábados, lentejas los viernes, algún palomino de añadidura los domingos, consumían las tres partes de su hacienda.",
	"Tú dices que amas la lluvia, sin embargo usas un paraguas cuando llueve. Tú dices que amas el sol, pero siempre buscas una sombra cuando el sol brilla. Tú dices que amas al viento, pero cierras las ventanas cuando el viento sopla. Por eso es que tengo miedo cuando dices que me amas.",
	"El miedo es como un fuego. Te va quemando por dentro. Si lo controlas, entras en calor. Pero si llega a dominarte, te quemará a ti y a todo lo que te rodea.",
	"El sapo sapote no sabe decir \"power\", el sapo sapote creo que sabe decir \"power\", el sapo sapote sí sabe decir \"power\". Además, el sapo sapote no come camote.",
	"Raras veces nos topamos con un individuo capaz de revisar la idea que tiene de su propia inteligencia y sus propios límites bajo la presión de los acontecimientos, hasta el punto de abrirse a nuevas perspectivas sobre lo que aún puede aprender.",
	"Yo aconsejo a mis pacientes que evalúen cuidadosamente el momento adecuado para la confrontación. No se trata de disparar sin haber apuntado, pero tampoco de postergar indefinidamente la confrontación.",
	"El día de pascua nevó mucho, pero al siguiente sopló de improviso un viento cálido, amontonándose las nubes, y por espacio de tres días con sus noches no dejó de caer una lluvia tibia. El viento se calmó el jueves y entonces se extendió sobre la tierra una espesa bruma de color gris, como para ocultar los misterios que se producían en la naturaleza.",
	"Las personas con alta tolerancia a la frustración pueden hacer frente a los contratiempos con éxito. Las personas con baja tolerancia a la frustración pueden sentirse frustradas por inconvenientes cotidianos aparentemente menores, como atascos de tráfico o esperar en la cola del supermercado. Las personas con baja tolerancia a la frustración pueden renunciar a tareas difíciles de inmediato. La idea de tener que esperar en la cola o trabajar en una tarea que no entienden puede ser intolerable.",
	"Te sientas frente a un tablero y repentinamente tu corazón brinca. Tu mano tiembla al tomar una pieza y moverla. Pero lo que el ajedrez te enseña es que tú debes permanecer ahí con calma y pensar si realmente es una buena idea o si hay otras ideas mejores.",
	"Sabes, Cristo me contó una vez el secreto de la resurrección. Pero estábamos en una boda en Canaán y me emborraché y se me olvidó. ¿Que si lo conocía? El tipo me debe doce dólares.",
	"Debía de tener unos diez u once años cuando desapareció su madre. Era una mujer alta, elegante, más bien callada, de movimientos lentos y precioso cabello rubio. A su padre lo recordaba de manera más vaga como un hombre moreno y delgado, vestido siempre pulcramente de negro.",
	"Leyendo novelas de detectives suelo sentir una sensación de alivio ante el asesinato de un personaje que al pasar la página no me molestará con su existencia. De todos modos, las historias de detectives siempre tienen demasiados personajes. Muchos mencionados al principio pero nunca vistos, que permanecen fuera de la historia, adquiriendo una calidad portentosa horrible. Es mejor que se hayan ido.",
	"La forma en que actuó muestra que le falta la madurez que es necesaria para este trabajo. Necesito a alguien que tenga, pues primero, un mínimo de cuidado para este trabajo, pero también, alguien que quiera aprender y ampliar sus conocimientos. Y por lo que puedo decir, él no tiene ninguna de esas cualidades. Entiendo que él necesita dinero y trabajo, pero yo no soy caridad para los que viven en los sótanos de sus padres. Dile que necesita salir de su trasero y trabajar.",
	"Me giré sobre mis pasos y a medida que me acercaba a la habitación de mi compañero de piso, el olor se hacía más espeso. Vacilé antes de abrir la puerta. Las cortinas flotaban sin rumbo, despavoridas y sobre su lecho, mi amigo yacía con los ojos vidriosos y el cuerpo inerte, sin vida, acompañado por ella.",
	"No te obstines, pues, en mantener como única opinión la tuya creyéndola la única razonable. Todos los que creen que ellos solos poseen una inteligencia, una elocuencia o un genio superior a los de los demás, cuando se penetra dentro muestran solo la desnudez de su alma.",
	"¿Por qué salimos corriendo cuando notamos la menor conexión? ¿Por qué luchamos contra lo que más nos atrae? Quizá es porque cuando encontramos algo o alguien a quién aferrarnos lo ansiamos como al aire. Y nos aterra perderlo... Y creedme, uno aprende a estar solo. Pero la mayoría de las cosas son mejores si las compartes.",
	"Pues si vemos lo presente cómo en un punto se es ido y acabado, si juzgamos sabiamente, daremos lo no venido por pasado. No se engañe nadie, no, pensando que ha de durar lo que espera más que duró lo que vio, pues que todo ha de pasar por tal manera.",
	"Había una casa abajo, junto al estruendo de las olas desbaratándose contra los cantiles, donde el amor era más intenso porque tenía algo de naufragio.",
	"El hombre intentó mover la cabeza en vano. Echó una mirada de reojo a la empuñadura del machete, húmeda aún del sudor de su mano. Apreció mentalmente la extensión y la trayectoria del machete dentro de su vientre, y adquirió fría, matemática e inexorable, la seguridad de que acababa de llegar al término de su existencia. La muerte.",
	"Todos mis instintos son una forma, y todos los hechos son otras, y me temo mucho que los jurados británicos todavía no han alcanzado ese tono de inteligencia cuando van a dar la preferencia a mis teorías.",
	"No somos más que unos hombres elaborados para actuar en nombre de la venganza que consideramos justicia. Pero cuando llamamos a nuestra venganza justicia, solo engendramos más venganza, forjando el primer eslabón de las cadenas de odio.",
	"La oscuridad, la algarabía y sobre todo el acto ilegal que estaba a punto de cometer no formaban parte de mi vida cotidiana, sin embargo, proseguí y llegué a la habitación. Respiré hondamente. Abrí la puerta con el temblor de la mano como testigo y reuniendo el poco valor que me quedaba crucé el umbral prohibido.",
	"Existe una conexión fundamental entre lo que uno parece y lo que uno es. Todos nos contamos una historia sobre nosotros mismos. Siempre. Continuamente. Esa historia es la que nos convierte en lo que somos. Nos construimos a nosotros mismos a partir de esa historia.",
	"No hay razón para sufrir. La única razón por la que sufres es porque así tú lo exiges. Si observas tu vida encontrarás muchas excusas para sufrir, pero ninguna razón válida. Lo mismo es aplicable a la felicidad. La única razón por la que eres feliz es porque tú decides ser feliz. La felicidad es una elección, como también lo es el sufrimiento.",
	"Si aún respiras, eres de los que más suerte tienen, porque la mayoría de nosotros jadeamos con pulmones corruptos y nos desatendemos a nosotros mismos en nuestro afán de recolectar los nombres de los amantes con quienes no fue bien.",
	"Si hay algo más importante que mi ego en los alrededores, lo quiero capturado y matado cuanto antes.- Por un momento, no pasó nada. Luego, después de un segundo más o menos, nada continuó sucediendo.",
	"Mi madre me dice que no mastico mi comida lo suficiente. Dice que estoy haciendo que sea más difícil para mi cuerpo obtener los nutrientes esenciales que necesita. Si ella estuviera aquí, le recordaría que estoy comiendo una tarta de arándanos.",
	"El libro es una criatura frágil. Sufre el paso del tiempo, el acoso de los roedores y las manos torpes, así que el bibliotecario protege los libros no sólo contra el género humano sino también contra la naturaleza, dedicando su vida a esta guerra contra las fuerzas del olvido.",
	"No es de extrañar que nos pasemos la vida entera soñando. Estar despierto y verlo todo tal y como es en realidad... nadie podría soportarlo mucho tiempo.",
	"Despertó empapado en sudor, con el corazón desbocado y casi sin aire. Al abrir los ojos, todo era oscuridad. Intentó incorporarse en la cama, estirar los brazos, pero el techo y las paredes de madera se lo impedían.",
	"Sabía que esto enviudaría a tu madre y te dejaría huérfano. Pero había encontrado mi destino. Así que abandoné a mi hijo. Ahora tengo trabajo que hacer, una cantidad infinita. Debo encontrar vida inteligente.",
	"La gente que más me gusta es la que ha fallado, ha sido lastimada, ha llorado, ha visto cosas terribles, y sin embargo, no ha perdido su capacidad para seguir amando.",
	"Generalmente, esa sensación de estar solo en el mundo aparece mezclada con un orgulloso sentimiento de superioridad: desprecio a los hombres, los veo sucios, feos, incapaces, ávidos, groseros, mezquinos; mi soledad no me asusta, es casi olímpica.",
	"como pasa mismo nos oh creo mejor nada tenemos papas qué sea acuerdo ellos un vas tiene fue nadie estamos va parece mi tus sin ahí mira alguien",
	"A primera vista, la llave y la cerradura en la que encaja pueden parecer muy distintas. Diferentes en su forma, diferentes en su función, diferentes en su diseño. El hombre que las mira sin conocer su verdadera naturaleza puede pensar que son opuestas, pues una sirve para abrir y la otra para mantener cerrado. Sin embargo, examinándolas con atención, se ve que sin una la otra no sirve para nada. El hombre sabio ve que la cerradura y la llave fueron creadas para el mismo propósito.",
	"No conocerás el miedo. El miedo mata la mente. El miedo es la pequeña muerte que conduce a la destrucción total. Afrontaré mi miedo. Permitiré que pase sobre mí y a través de mí. Y cuando haya pasado giraré mi ojo interior para escrutar su camino. Allá donde haya pasado el miedo ya no habrá nada. Solo estaré yo.",
	"No tengo hambre, tengo ansiedad. Ver tanta gente acá reunida me dan ganas de fumar. Lo que pasa es que en casa muy solo se está. Ahí va otro inadaptado intentando encajar. No sé muy bien en dónde... Encajar, no sé ni para qué... Si no hay botón de pausa, no hay rebobinar. Por eso es que no me pierdo ningún evento social. Hace tiempo estoy pensando tengo que parar, mientras tanto por las dudas sigo a toda velocidad.",
	"Amar es prolongar el breve instante de angustia de ansiedad y de tormento en que, mientras espero, te presiento en la sombra suspenso y delirante. Yo quisiera anular de tu cambiante y fugitivo ser el movimiento, y cautivarte con el pensamiento y por él sólo ser tu solo amante. Pues si no quiero ver, mientras avanza el tiempo indiferente, a quien más quiero, para soñar despierto en tu tardanza. La sola posesión de lo que espero, es porque cuando llega mi esperanza es cuando ya sin esperanza muero.",
	"En las noches de luna llena, la lívida luz se refleja en los adoquines del pavimento, que refulgen con un aura fantasmal, alimentando la imaginación de los cuentos de abuelas generación tras generación. Ese brillo mortecino abraza la soledad de sus calles en las noches eternas y se une al tenue ulular del viento evocando un lamento por la grandeza perdida de los tiempos pasados.",
	"El muro de tu olvido absorbe la caída de mis pupilas en la quieta mirada del retiro... La noche exhala sus tules. Estás y no... Con el ojo estrangulado en el vano intento de acariciar tu imagen verifico tu ausencia. Estás y no... Te figuro en un rayo de la luna, parece que me observas, levanto la cabeza, para asirme en tus cabellos. Estás y no... En el muro, en la morada, en mis ojos y en la luna en el hoy y en el ayer. Estás y no...",
	"Estás preocupado por lo que fue y por lo que va a ser... Hay un dicho: \"El ayer es historia, el mañana es un misterio, pero el hoy es un obsequio. Por eso se llama presente.\"",
	"Mi vida comenzó un otoño del 96, y acabó mucho antes de los 16, mi habla fue precoz, pero qué poquito le duró la voz, cuando el tiempo la cercenó con su hoz.",
	"Cerca de quinientas almas viven en esta pequeña villa amurallada de origen medieval que se eleva orgullosa sobre la cima de una colina, rodeada de verdes pastos y tierras de labranza. A lo largo de los siglos sus vecinos se han distinguido como gente sencilla y trabajadora.",
	"De Tarde, arrimaba a la puerta una de las sillas y mateaba con seriedad, puestos los ojos en la enredadera del muro de la inmediata casa de altos. Años de soledad le habían enseñado que los días, en la memoria, tienden a ser iguales, pero que no hay un día, ni siquiera de cárcel o de hospital, que no traiga sorpresas, que no sea al trasluz una red de mínimas sorpresas.",
	"No sabía en qué estaba pensando, no comprendía el significado de su aroma. Estos aromas eran brillantes, cinéticos, me quemaban la nariz. Pero por debajo de ellos, estaba Joe. Era como humo y tierra y lluvia. Eran los aromas que siempre había asociado a él, pero intensificados por miles de veces más. Quería enterrarme en ellos, enroscándome hasta que su esencia me cubriera por completo.",
	"Me hubiera gustado que viese usted mi nombre en un libro, aunque no pudiese leerlo. Me hubiera gustado que estuviese aquí, conmigo, para ver que su hijo conseguía abrirse camino y llegaba a hacer algunas de las cosas que a usted nunca le dejaron. Me hubiera gustado conocerle, padre, y que usted me hubiera conocido a mí. Le convertí a usted en un extraño para olvidarle y ahora el extraño soy yo.",
	"La vida de un crítico es sencilla en muchos aspectos. Arriesgamos poco y tenemos poder sobre aquellos que ofrecen su trabajo y su servicio a nuestro juicio. Preferimos las críticas negativas, que es divertida de escribir y de leer. Pero la triste verdad que debemos afrontar es que en el gran orden de las cosas, cualquier basura tiene más significado que lo que deja ver nuestra crítica.",
	"Cualquiera puede enfadarse, eso es algo muy sencillo. Pero enfadarse con la persona adecuada, en el grado exacto, en el momento oportuno, con el propósito justo y del modo correcto, eso, ciertamente, no resulta tan sencillo.",
	"Nuestro conocimiento nos ha hecho cínicos. Nuestra inteligencia, duros y secos. Pensamos demasiado y sentimos muy poco. Más que máquinas, necesitamos humanidad. Más que inteligencia, tener bondad y dulzura. Sin estas cualidades, la vida será violenta. Se perderá todo... A los que puedan oírme, les digo; no desesperéis. La desdicha que padecemos no es más que la pasajera codicia y la amargura de hombres que temen seguir el camino del progreso humano. El odio de los hombres pasará.",
	"Si escuchas mi voz y al girar tu cabeza no me encuentro ahí, el que te hable será mi espíritu que viaja a través del tiempo y del espacio en tus recuerdos.",
	"La guerra es lo más importante para el estado, el terreno de la vida y de la muerte, el camino a la supervivencia o la desaparición. No puede ser ignorada.",
	"No me gustan los términos de buena o mala persona porque es imposible ser totalmente bueno con todo el mundo, o totalmente malo con todo el mundo. Para algunos, eres una buena persona, mientras que para otros eres una mala persona.",
	"Las palabras son pálidas sombras de nombres olvidados. Los nombres tienen poder, y las palabras también. Las palabras pueden hacer prender el fuego en la mente de los hombres. Las palabras pueden arrancarles lágrimas a los corazones más duros. Existen siete palabras que harán que una persona te ame. Existen diez palabras que minarán la más poderosa voluntad de un hombre. Pero una palabra no es más que la representación de un fuego. Un nombre es el fuego en sí.",
	"En el camino hacia aquí se sientan cómodos y disfrutan el trayecto. A veces nos paramos y vemos los pájaros donde no los hay, y miramos atardeceres cuando está lloviendo. Nos lo pasamos fantásticamente y recibo muy buena propina.",
	"Si la felicidad tuviera una forma, tendría forma de cristal, porque puede estar a tu alrededor sin que la notes. Pero si cambias de perspectiva, puede reflejar una luz capaz de iluminarlo todo.",
	"Había una mujer, fue la primera vez que encontré a alguien que estuviera verdaderamente vivo. Al menos, eso fue lo que pensé. Ella era... la parte de mí que perdí en algún lugar del camino, la parte que faltaba, la que deseaba.",
	"Las pesadillas son muy raras. Tu cerebro es el autor, espectador y escenario de una película de terror cuyo guión está siendo escrito mientras lo ves.",
	"Ya no jugamos tres en raya porque es un juego aburrido y siempre acaba en empate; no hay forma de ganar, el juego en sí no tiene sentido. Pero en la sala de guerra creen que puedes ganar una guerra nuclear, que puede haber pérdidas aceptables.",
	"Súbitamente desapareció en el fondo de la bahía, sin duda para embestirme desde abajo. Primero fue un terrible costalazo. Luego se asomó para mostrarme sus nueve hileras de colmillos blancos, sus ojos apagados, sin odio ni crueldad, como los de quien ejecuta una rutina.",
	"Qué extraño. Esto es irracional. Sin embargo, fueron los seres humanos quienes prefirieron las emociones en lugar de las ganancias o las pérdidas. Yo solo quería que mi hermano mayor me reconociera, pero me faltaba algo. Lo siento, no pude cambiar hasta el final.",
	"Eso es porque todos queremos no tener problemas, arreglarnos a nosotros mismos. Buscamos una solución mágica para mejorarnos, pero ninguno de nosotros sabe realmente lo que está haciendo. Eso es todo lo que los humanos podemos hacer: adivinar, intentar, esperar.",
	"Ahora que hemos establecido el dilema de nuestro protagonista, pasemos a nuestro antagonista. A muchos kilómetros de distancia, a través de las llanuras abiertas, otra bella bestia salvaje se abre paso hasta un abrevadero.",
	"Sabes es increíble tener la misma rutina desde hace tres años, levantarte, escuchar unas rolitas, imaginar cosas para ser feliz, solo te haces daño a ti mismo, pero- por lo menos eres feliz en ese mundo, en tu imaginación, eso es lo único que te mantiene feliz, ver esa sonrisa en su rostro, imaginarte un beso, dormir con esa persona, una salida al parque y que por primera vez te dé una señal de que le gustas...",
	"Crearía un perfume que no sólo fuera humano, sino sobrehumano. Un aroma de ángel, tan indescriptiblemente bueno y pletórico de vigor que quien lo oliera quedaría hechizado y no tendría más remedio que amar a la persona que lo llevara, o sea, amarle a él, Grenouille, con todo su corazón.",
	"No podía moverme, no podía hablar. Solo podía escuchar y ver. Mi grado de desesperación se acercaba a dimensiones infinitas mientras la silueta seguía ahí, muy bajito, desde su sucio rincón, alimentando mi nuevo estado de locura.",
	"Tenía una mezcla de miedo y resaca porque había escuchado desde chaval miles de historias sobre heroína, putas y problemas. Comencé a andar más rápido para salir de allí cuanto antes y a unos metros de mí escuché una gran carcajada seguido de una voz que se acercaba, pero no acerté a entender lo que decía.",
	"¿Realmente cree que el enemigo atacaría sin provocación, usando tantos misiles, bombarderos y submarinos, para que no tengamos más remedio que aniquilarlos por completo? General, está escuchando a una máquina, hágale un favor al mundo y no actúe como una.",
	"Pero luego el tráfico, la estampida, los autos, los camiones, las bocinas... Esta pobre mujer gritaba y me di cuenta de que no era un renacimiento, era una especie de ilusión vertiginosa de renovación que ocurre en el momento final antes de la muerte.",
	"Deseaba ver el infinito, pero en realidad nunca salía de los límites de su propia cabeza. Decirle la verdad sería como dar una patada a un perro de aguas.",
	"Abrí la puerta del coche para, como un día cualquiera, recorrer el tramo que separa mi casa del trabajo. El tráfico era el habitual, como habitual es también el cabreo que te produce.",
	"La niebla, impregnada de aquella luminosidad extraña, a ras de suelo, empezó a arremolinarse y parecía cobrar vida, ascendiendo alrededor de él, que seguía arrodillado en el suelo, con los brazos cruzados y la cabeza baja.",
	"En segundo lugar, opino que no es tan importante cuantas medidas tomen las autoridades: no basta que estén asociados a campañas nacionales de concientización y enseñanza, si no pueden transmitir eficazmente esta conciencia. Estoy seguro de que un gran porcentaje de la población desconoce las consecuencias del uso del plástico o que este se obtiene del petróleo, un recurso natural no renovable. Cuando se desconoce, no se entiende. Y, si no se entiende, no se hace.",
	"La vida es un viaje y no nos damos cuenta, nos anclamos a lo que sea que permita que nuestra existencia tenga algún sentido, somos adictos a creer en algo pero muchos son incapaces de creer en ellos mismos.",
	"Hay mucho más. Hay todo lo que va más allá, todo lo que está afuera. Y todo lo de atrás, desde hace muchísimo, muchísimo tiempo. Yo recibí todas esas cosas cuando me seleccionaron. Y aquí en esta habitación, yo solo, las vuelvo a experimentar una y otra vez. Es así como viene la sabiduría. Y como configuramos nuestro futuro.",
	"Ahora se sorprendía a menudo enfadado: irracionalmente enfadado con sus compañeros de grupo, porque estaban satisfechos con unas vidas que no tenían nada de la vibración que la suya estaba adquiriendo. Y enfadado consigo mismo, por no poder cambiar esa situación.",
	"Un amor por ocultar. Aunque en cueros no hay donde esconderlo, lo disfrazan de amistad cuando sale a pasear por la ciudad. Una opina que aquello no está bien. La otra opina que qué se le va a hacer. Y lo que opinen los demás está de más. Quien detiene palomas al vuelo volando a ras de suelo. Mujer contra mujer.",
	"Entrará el mar lentamente en tus venas, oh nadador que esperaste la noche y la soledad para medir tus fuerzas con la tormenta, digo con tu propio destino, desde el principio con algo que sabías superior a ti mismo.",
	"Fue cuestión de segundos, bebió más de la cuenta, estaba alegre, aún oigo su carcajada. Para él no tiene remedio y para mí ya nada tiene sentido. Bajo del metro y recorro lentamente los trescientos sesenta y cinco escalones que me conducen a él.",
	"Al recuperarme, abrí los ojos y vi a mi mujer a mi lado. Estaba tumbado en la cama de ese hospital que había visto desde la ventana. Un sentimiento de alivio recorrió todo mi cuerpo. Había vuelto de un viaje que nunca recomendaré a nadie.",
	"El dolor desaparece con el tiempo. Pero no deseo ser curado por el tiempo, porque cuando huyes del dolor, con el anhelo de olvidar, lo único que logras es quedarte atascado. Te vuelves incapaz de seguir adelante.",
	"en sabe estoy papas tipo espero fuera mejor ahora quiere favor necesito ser podría nada ella sé papas desde estado era mejor noche amigo creo puedes",
	"tipo sabe nadie no nadie ya usted has sea dos espero de sobre día desde antes sin algo papas todos así porque con te nuevo ni por otro cosa hablar",
	"Manos a la obra, señor Freeman, manos a la obra. No quiero decir que se haya dormido en los laureles, en el trabajo. Nadie se merece más un alto en el camino. Todos los esfuerzos realizados habrían sido en vano hasta que... Bueno, digamos que otra vez ha llegado su hora. El hombre adecuado en el sitio equivocado puede cambiar el rumbo del mundo. Despierte, señor Freeman, despierte y mire a su alrededor.",
	"Nos sentimos más apasionados cuando besamos, más humildes cuando nos arrodillamos y más enojados cuando agitamos los puños. Es decir, que el beso no solamente expresa pasión sino que la origina. Los papeles llevan consigo al mismo tiempo ciertas acciones y las emociones y actitudes que corresponden a esas acciones. El profesor que pone en escena un acto que finge sabiduría, llega a sentirse sabio. El predicador llega a creer en lo que predica.",
	"En realidad, a uno no le gobierna en modo alguno la mente. Lo que la mente revela es una corriente interminable de opciones, todas ellas disfrazadas en la forma de recuerdos, fantasías, temores, conceptos, etc. Para liberarse del dominio de la mente, solo hay que darse cuenta de que su desfile de temas no más que un autoservicio arbitrario de opciones que pasan a través de la pantalla de la mente.",
	"Mientras sacaba las llaves del bolso, oí de repente unos pasos que se acercaban. Al volverme, vi la figura de un hombre joven que se dirigía hacia mí con una amplia sonrisa en la cara. Según avanzaba comenzó a hablar...",
	"Quisiera que me vieras y descubrieras que solo un murmullo me haría feliz, quisiera que me dieras una sonrisa y poder sentir, que estoy cerca de ti. Quisiera, que quisieras quererme, olvidarme para siempre y de tu mano andar; quisiera que quisieras quereme, y el primer beso me dieras. Eso quiera.",
	"Cuando algo sobre una persona resuena profundamente dentro de ti, es el sentimiento más maravilloso. Tu corazón se libera, tu mente se abre y te das cuenta de que hay más en el mundo de lo que nunca imaginaste.",
	"Se movía como un depredador, sus anchos hombros se balanceaban siguiendo los vaivenes de su andar y la cabeza giraba de un lado a otro, explorando. Ella tuvo la incómoda sensación de que si él quisiera, podía matar a todos los presentes sin usar más arma que sus manos.",
	"Solo podía apreciar el típico ruido que hacen esos scanners. De repente, bajo mis pies se abrió una puerta negra. Caí en picado por un túnel. Las imágenes pasaban rápidamente a mi alrededor. Eran todos los acontecimientos de mi vida. Sentí un fuerte golpe y perdí la consciencia.",
	"Te escribí cuando mi tren salía de su andén. Fue una madrugada bendecida por el beso de Diciembre. Mi mano temblaba de confesiones de cobardía; pero no pude dar marcha atrás. Así que comencé a escribirte un poema mientras mi tren salía. Y al hacerlo, el peso dentro de mí disminuyó, y solo lo hacía en viajes como estos.",
	"La fantasía es una bicicleta estática para la mente. Es posible que no te lleve a ninguna parte, pero tonifica los músculos que pueden. Por supuesto, podría estar equivocado.",
	"Rendirse es lo que destruye a la gente, cuando te niegas con todo tu corazón a rendirte entonces trasciendes tu humanidad, incluso ante la muerte nunca te rindas.",
	"Yo tengo miedo de todo. Y estoy loca. Como si tal vez piensas que estoy un poco loca, pero la gente sólo ve la punta del iceberg de la locura. Por debajo de esta apariencia de ligeramente loca y socialmente inepta, soy un completo desastre.",
	"En la vida se tienen que tomar muchas decisiones; si esas decisiones son correctas o no, nadie lo sabe. Por eso las personas suelen elegir lo que ellos creen que es correcto.",
	"Estaba permanentemente impresionado por las banalidades más irrelevantes, pero era imposible de impresionar con una verdadera novedad, significado o conflicto. Y era demasiado estúpido como para odiarse a sí mismo, así que era mi deber hacerlo por él.",
	"Caminó rápidamente a través de la oscuridad con el paso franco de alguien que al menos estaba seguro de que el bosque, en esta noche húmeda y ventosa, contenía cosas extrañas y terribles y ella las era.",
	"Pues, no, hermano, no la quiero, que es historia muy cansada ver que al pasar del arroyo te llegue a la boca el agua. La mujer que ha de ser propia ha de estar en una caja como el gusano de seda, hasta ser paloma blanca. Si fuiste abeja en su rosa, que buen provecho te haga; que lo que no fue posible olvidar con la mudanza de su traje, ni acabaron sus desdenes y desgracias, con lo que me has dicho sólo, hoy para siempre se acaba.",
	"Indicaciones terapéuticas del ibuprofeno: artritis reumatoide (incluyendo artritis reumatoide juvenil), espondilitis anquilopoyética, artrosis y otros procesos reumáticos agudos o crónicos. Alteraciones musculoesqueléticas y traumáticas con dolor e inflamación. Tto. sintomático del dolor leve o moderado (dolor de origen dental, dolor posquirúrgico, dolor de cabeza, migraña). Dismenorrea primaria. Cuadros febriles.",
	"Es una verdad mundialmente reconocida que un hombre soltero, poseedor de una gran fortuna, necesita una esposa. Sin embargo, poco se sabe de los sentimientos u opiniones de un hombre de tales condiciones cuando entra a formar parte de un vecindario. Esta verdad está tan arraigada en las mentes de algunas de las familias que lo rodean, que algunas le consideran de su legítima propiedad y otras de la de sus hijas.",
	"Los hombres se equivocan al creerse libres, opinión que obedece al solo hecho de que son conscientes de sus acciones e ignorantes de las causas que las determinan. Y, por tanto, su idea de \"libertad\" se reduce al desconocimiento de las causas de sus acciones, pues todo eso que dicen de que las acciones humanas dependen de la voluntad son palabras, sin idea alguna que les corresponda.",
	"Hay una fuerza motriz más poderosa que el vapor, la electricidad y la energía atómica: la voluntad. Albert Einstein. 14 de marzo de 1879 - 18 de abril de 1955",
	"Hora crepuscular. Un guardillón con ventano angosto, lleno de sol. Retratos, grabados, autógrafos repartidos por las paredes, sujetos con chinches de dibujante. Conversación lánguida de un hombre ciego y una mujer pelirrubia, triste y fatigada. El hombre ciego es un hiperbólico andaluz, poeta de odas y madrigales, Máximo Estrella. A la pelirrubia, por ser francesa, le dicen en la vecindad Madama Collet.",
	"Es un gran placer para mí darles la bienvenida a la 19 Conferencia de Plenipotenciarios (PP-14), el evento de mayor trascendencia en el calendario de la UIT, que congrega a casi tres mil delegados de los Estados Miembros junto con otros Miembros de la UIT, entre ellos empresas del sector privado, organizaciones internacionales, la sociedad civil e Instituciones Académicas. En la Conferencia se abordarán cuestiones importantes relativas al futuro de la Unión, desde la integración digital a la implantación de la banda ancha.",
	"La manera en la que este bot calcula el WPM, es mediante la siguiente función: func calculate_wpm() {var length = len([]rune(current_text)); var length_as_a_float float64 = float64(length); wpm = length_as_a_float / 5 / time_elapsed * 60000}. Como puedes ver, solo consta de aplicar una pequeña fórmula a la longitud del texto y los milisegundos transcurridos.",
	"Estoy formado por los recuerdos de mis padres y mis abuelos y todos mis antepasados. Están en mi aspecto, en el color de mi cabello. Estoy hecho de todos los que he conocido y han cambiado mi forma de pensar.",
	"Aprecio mi pasado en detrimento de mi futuro; aunque encuentro excelentes ambos, no puedo otorgar primacía a ninguno, y solo debo censurar la injusticia de la providencia que tanto me favorece.",
	"No más que un soplo de viento es el rumor de la aprobación humana, que tan pronto viene de un extremo como del opuesto y cambia de nombre al cambiar de dirección.",
	"Deberías ser más desafiante, porque solo vives una vez, deberías enfrentar, lo que sea que quieres hacer cuando eres joven, cuando seas viejo y pienses en el pasado, no será algo malo de recordar incluso si fallaste.",
	"Así es como me gusta tenerlo. Bajo mi poder. Entre mis manos. No por ahí conspirando y tramando planes, hablando con vampiros. Ahora te tengo, pienso, finalmente te tengo donde quería tenerte.",
	"Muchos de los que viven merecen morir y algunos de los que mueren merecen la vida. ¿Puedes devolver la vida? Entonces no te apresures a dispensar la muerte, pues ni el más sabio conoce el fin de todos los caminos.",
	"Eres grande Ox. Más grande que cualquier lobo que haya visto antes. Más grande que yo, que mi padre, pero tiene sentido, ¿sabes? Porque siempre has sido así para mí. Más grande que cualquier cosa. El día que te vi, supe que las cosas jamás volverían a ser como antes. Lo abarcas todo, empequeñeces todo lo demás y cuando te veo, Ox, eres todo lo que puedo ver.",
	"La llegada de su amigo y compañero que por un contratiempo permanente la puerta nunca abrirá, aunque me digan que nunca llegará, lo esperaré y cuando cruce la puerta lo recibiré. O tal vez, solo tal vez me reciba él.",
	"La guerra deberá existir mientras defendamos nuestras vidas contra un destructor que lo devoraría todo. Pero no amo la espada por su filo, ni la flecha por su rapidez, ni al guerrero por su gloria. Solo amo lo que defienden.",
	"Significa poder enfrentar los problemas, aguantar el dolor y los golpes que te da la vida, no significa que no puedas sentir dolor ni ponerte triste, pero alguien fuerte no deja que su vida sea afectada por cualquier desdicha.",
	"La gente tiene diferentes formas de pensar, incluso cuando se comete un error... Si la persona se da cuenta de su error es posible que lo enmiende, si mantienes tu visión clara verás el futuro, es de lo que se trata esta vida...",
	"El problema es, Leslie, que el mundo no se va a acabar mañana. El sol se va a levantar por ese lado, será un viernes común y corriente, y tú seguirás en la misma posición en la que estás ahora.",
	"Al lado, se encontraba la enorme escultura de un payaso realizada en un descolorido material que Cristina tampoco pudo reconocer. Sorprendida, reparó en la falta de nariz tan significativa en todos los payasos.",
	"No es nada, hija, es el viento. Las noches aquí son muy distintas a las de la ciudad. Ahí se oyen los coches, los tranvías... Aquí las casas son viejas, gruñen, hasta parece que hablan.",
	"Lo que realmente significa el poder del hombre sobre la naturaleza es el poder ejercido por algunas personas sobre otras, con el conocimiento de la naturaleza como instrumento.",
	"La vida no es más que una sombra en marcha; un mal actor que se pavonea y se agita una hora en el escenario y después no vuelve a saberse de él: es un cuento contado por un idiota, lleno de ruido y de furia, que no significa nada.",
	"Este es mi hogar. Este siempre fue un viaje de ida, hijo mío. Nunca hubo nada para mí allí, nunca me importasteis ni tú ni tu madre ni ninguna de vuestras pequeñas ideas. Durante treinta años he estado respirando este aire, comiendo esta comida, soportando estas dificultades y nunca pensé en casa.",
	"En mi experiencia, algunas de las relaciones más exitosas se basan en mentiras y engaños. Como de todos modos es ahí donde generalmente acaban, me parece un lugar muy lógico para comenzar.",
	"Hay momentos en la vida en los que las personas deben saber cuándo no dejar que se nos escapen ciertas cosas. Los globos están diseñados para enseñar esto a los niños pequeños.",
	"Sobresaltada, corrió hasta la ventana para cerrarla, pero no pudo llegar hasta ella. Las sombras de la noche, el alma de su madre, a la que abandonó hacía tantos años, la arrastró definitivamente a un sueño del que nunca volvería a despertar.",
	"Primero fuimos formas difusas en el mar, luego peces, después lagartos y ratas, monos y cientos de cosas en medio. Esta mano fue una vez una aleta, también tuvo garras. En mi boca humana tengo los dientes puntiagudos de un lobo, los de cincel de un conejo y los molientes de una vaca.",
	"Nunca volveré a ser así. No volveré a sentirme tan alta como el cielo, tan vieja como las colinas y tan fuerte como el mar. Me han concedido algo durante un rato, y el precio es que tengo que devolverlo.",
	"Los artistas dedican más tiempo a crear obras malas de práctica que a sus obras maestras, sobre todo al principio. E incluso cuando un artista se hace maestro, hay trabajos que no le salen bien. Y otros que ya están mal desde el primer al último trazo. Se aprende más del mal arte que del bueno, ya que tus errores importan más que tus éxitos.",
	"Piensa en lo importante que eres para la vida de aquellos a quienes encuentras; cuán importante podrías ser para las personas en las que ni siquiera pensarías. Hay algo de ti que dejas en cada interacción con otra persona.",
	"Una sombra pasó junto al coche, indiferente a nuestro melodrama en la acera. Esta fue mi segunda vez en peligro en un vehículo estacionado en el espacio de tres horas. Me preguntaba qué espectáculos tontos había pasado por alto en mi propia carrera como caminante de aceras.",
	"Nuestra sangre es tan salada como el mar en el que vivíamos. Cuando estamos asustados, el vello de nuestra piel se eriza, al igual que cuando teníamos pelaje. Somos historia. Todo lo que hemos sido mientras nos convertíamos en nosotros, todavía lo somos.",
	"En Egipto se llamaban las bibliotecas el tesoro de los remedios del alma. En efecto, curábase en ellas de la ignorancia, la más peligrosa de las enfermedades y el origen de todas las demás.",
	"La imaginación es una cualidad que se le ha concedido al hombre para compensarlo por lo que no es, mientras que el sentido del humor le ha sido dado para consolarlo por lo que es.",
	"Todo lo que hay en esta sala es comestible. Hasta yo lo soy. Pero eso sería canibalismo, queridos niños, y está mal visto en la mayoría de las sociedades.",
	"Aprendí que el coraje no era la ausencia de miedo, sino el triunfo sobre él. El valiente no es quien no siente miedo, sino aquel que conquista ese miedo.",
	"Todo comienza en alguna parte, aunque muchos físicos no están de acuerdo. Pero las personas siempre han sido poco conscientes del problema con el comienzo de las cosas. Se preguntan cómo el conductor de la quitanieves va al trabajo, o cómo los fabricantes de diccionarios buscan la ortografía de las palabras.",
	"Tras una pausa de un microsegundo y una micromodulación magníficamente calculada de tono y timbre, algo que no podría considerarse insultante, Marvin logró transmitir su absoluto desprecio y horror por todas las cosas humanas.",
	"Siempre había creído que los asesinos tendrían otro aspecto. Tal vez pensaba que por el simple hecho de matar a alguien le cambiaría la mirada, le aturdirían los remordimientos o el nerviosismo entumecería sus movimientos. Se miró al espejo, pero su aspecto era el mismo.",
	"Sus cejas eran tan rubias que eran casi invisibles, no contrastaban con el resto de su rostro, lo cual le dificultaba mucho parecer enojada, arrepentida o burlona.",
	"Es muy curioso. Él está muerto bajo tierra, yo paralizado sin piernas, y tú perdido sin dónde ir. Supongo que siempre es así. La guerra, quiero decir.",
	"La muerte únicamente es una preocupación de aquellos que están vivos. Los muertos como yo solo nos preocupamos de la descomposición y de los necrófilos.",
	"Ante los dilemas morales nunca he elegido la opción más noble. Pero es justamente lo que recomiendo a otros que hagan porque así tengo ventaja sobre ellos.",
	"Si hay algo que he aprendido, es que la piedad es más inteligente que el odio, que la misericordia es preferible aún a la justicia misma, que si uno va por el mundo con mirada amistosa, uno hace buenos amigos.",
	"La mala noticia es que nunca serás el mismo, nunca vas a estar completo de nuevo. Perdiste a tu hija y nada va a reemplazar eso. Ahora la buena noticia es: tan pronto como aceptes eso y te permitas sufrir, te permitirás visitarla en tu mente. Recordarás todo el amor que ella dio, toda la alegría que conoció.",
	"Tienes que recuperar el control de tus pensamientos, imágenes mentales, sueños y ensoñaciones, en fin, de tu comportamiento. A partir de ahora, todos tus pensamientos, palabras y acciones tienen que ser voluntarios y puestos al servicio de tus intereses, valores y objetivos.",
	"Todos los seres humanos tienen un número de oportunidades para salir del caparazón y crecer. Pero no pueden aprovecharlo solo, necesitan adversarios poderosos que les exijan toda su concentración y aliados con quienes compartir sus experiencias.",
	"Solo puede ser intrépido quien conoce el miedo pero lo supera, quien ve el abismo con orgullo. Quien ve el abismo con ojos de águila, quien con garras de águila se aferra al abismo, ese tiene valor.",
	"A veces precisamos escalar de todas las montañas la más alta para ver de todos los amaneceres, el más hermoso. Eso lo leí una vez en el cojín decorativo de una señora.",
	"Cuando niño siempre pensé que las arenas movedizas iban a ser un problema mucho mayor de lo que resultaron ser. Porque si miras dibujos animados, las arenas movedizas son la tercera cosa más importante de la que te tienes que preocupar en la vida adulta detrás de la dinamita y yunques gigantes que caen desde el cielo. Me sentaba a pensar qué hacer con las arenas movedizas. Nunca pensé en cómo manejar problemas reales en la vida adulta, como qué hacer cuando tus familiares te piden dinero.",
	"Fue un gran fracaso como demonio, hizo todo lo posible para que sus cortas vidas fueran miserables, porque ese era su trabajo, pero nada de lo que se le ocurría era la mitad de malo que las cosas que ellos mismos inventaban. Parecían tener talento para ello. Estaba incorporado en su diseño, de alguna manera. Nacieron en un mundo que estaba en contra de ellos de mil maneras, y luego dedicaron la mayor parte de sus energías a empeorarlo.",
	"Nos sentamos para tener lo que podría haber sido la última cena de Navidad de la madre de Jordana. Francamente espero que no, porque el pavo estaba un poco seco, la col muy húmeda y todo en general un poco desenfocado.",
	"El ataque estaba lleno de ira. Entiendo esa rabia. La he visto en mi padre, y la he visto en mí, porque estoy enfadado con él por haberse ido. Pero cuando tomo esa ira y la pongo a un lado, todo lo que queda es dolor.",
	"Quiero irme adonde el sol continúe brillando a través de la lluvia. Donde el tiempo se adapte a mi ropa. Aprovechando los vientos del nordeste y zarpando en la brisa del verano, saltando como una piedra sobre el mar.",
	"Si confías en ti mismo y crees en tus sueños y sigues el camino que te muestran las estrellas, aún acabarás perdiendo contra personas que pasaron su tiempo trabajando duro y aprendiendo cosas y que no eran tan vagos.",
	"Conserva lo que tienes, olvida lo que te duele, lucha por lo que quieres, valora lo que posees, perdona a los que te hieren y disfruta de los que te aman.",
	"Crowley siempre había sabido que aún estaría vivo cuando el mundo terminara, porque era inmortal y no tenía alternativa. Pero esperaba que el fin estuviera muy lejos en el futuro. Porque le gustaba la gente. Fue un gran fracaso como demonio, hizo todo lo posible para que sus cortas vidas fueran miserables, porque ese era su trabajo, pero nada de lo que se le ocurría era la mitad de malo que las cosas que ellos mismos inventaban.",
	"Sin embargo, poder devolverlo también es la recompensa. Te pasarías la vida mirando una flor para disfrutar de su belleza en vez de ordeñar las vacas.",
	"Y ahora pretendo embarcarme en una melopea emocional de 72 horas. Lo más probable es que acabe bebiendo solo, pero no tengo problema con eso. Por tanto, la clase está cancelada.",
	"El purgatorio es un poco como lo que está en medio. No has sido malo del todo, pero tampoco has sido particularmente bueno. Como el equipo del Tottenham.",
	"El diablo cita la biblia en su provecho. El alma perversa que alega santo testimonio es como un canalla de cara sonriente o hermosa manzana podrida por dentro. Qué buena presencia tiene la impostura.",
	"Las personas mayores solo dicen que la vida sucede rápidamente para sentirse mejor. La verdad es que todo sucede en pequeños incrementos consecutivos. Y solo toma de veinte a treinta incrementos para darte cuenta de que estás destinado directamente a acabar solo en un banco del Parque de los Solteros.",
	"Escribo y miro la pared, vuelvo a escribir y vuelvo a mirar la pared. El sonido de las teclas retumban en mi cabeza y de pronto silencio total mientras miro la pared. Me gusta mirar la pared, es blanca y eso me gusta. Es como poner mi mente en blanco. El silencio es interrumpido por los fuertes tecleos y vuelvo a poner el punto para otro silencio total. Tecleo, punto y silencio.",
	"Estás desconectado de la realidad. Te has olvidado de las leyes de la jungla, menospreciándome. Cuando el espalda plateada tiene más plata que espalda, es mejor que se retire, antes de que lo retiren.",
	"Oyendo hablar a un hombre, fácil es acertar dónde vio la luz del sol; si os alaba Inglaterra, será inglés, si os habla mal de Prusia, es un francés, y si habla mal de España, es español.",
	"Miré al jardín. El sol se despedía de los árboles más altos. En un rincón oscuro una mujer apretaba a un niño contra su pecho como si quisiera apartarlo de la vida.",
	"Su imprecisión y pereza enloquecieron mis instintos compulsivos: su parche, la forma en que incluso su discurso estaba plagado de abandonos y problemas técnicos como un casete gastado, la forma en que sus sentidos cargados rechazaban el mundo.",
	"Entender cómo funciona algo facilita la convivencia con ello, mientras que la falta de comprensión y la ignorancia inducen al temor y a la superstición y nos convierten en críticos de otros seres humanos.",
	"El caos generalmente se encuentra en mayor abundancia en aquellos lugares donde se busca el orden. Siempre derrota a este último porque está mejor organizado.",
	"Personalmente, puedo digerir hierro y, si bien puede parecer que me duermo ocasionalmente, verán que me despierto con facilidad. Especialmente si soy agitado gentilmente por un buen abogado con un interesante punto de ley.",
	"Si hubieras mantenido mi amistad, los que maltrataron a tu hija lo hubieran pagado con creces. Porque cuando uno de mis amigos se crea enemigos, yo los convierto en mis enemigos.",
	"Advertir la vida mientras se vive, alcanzar a vislumbrar su implacable grandeza, disfrutar del tiempo y de las personas que lo habitan, celebrar la vida y el sueño de vivir, ese es su arte.",
	"Los regalos nos permiten demostrar exactamente lo poco que conocemos a una persona, y nada cabrea más a una persona que la tengan por algo que no es.",
	"Suele ser el caso que los enanos le tienen temor a las alturas, dado que por desgracia no disponen de muchas oportunidades para acostumbrarse a ellas.",
	"Con diez cañones por banda, viento en popa a toda vela, no corta el mar, sino vuela un velero bergantín; bajel pirata que llaman, por su bravura, el Temido, en todo el mar conocido de uno al otro confín. La luna en el mar riela, y en la lona gime el viento y alza en blando movimiento olas de plata y azul; y va el capitán pirata, cantando alegre en la popa, Asia a un lado, al otro Europa, y allá a su frente Estambul.",
	"Supe que mentía, como sin duda había hecho infinidad de veces en los últimos doce años. Pero no quise parecer inoportuno, de modo que cambié de tema. Su largo camino de ida y vuelta contenía episodios que me interesaban mucho más que las lecturas de la mujer que al fin tenía frente a mí, tras haber seguido sus huellas por tres continentes durante los últimos ocho meses. Decir que estaba decepcionado sería inexacto.",
	"Ella ya estaba lentamente aprendiendo que si ignoras las reglas, la gente, la mitad de las veces, las reescribirá en silencio para que no se apliquen a ti.",
	"Paré el coche y, con muchas dudas, descendí. Vi a lo lejos una luz y emprendí el camino hacia ella. Según me acercaba pude comprobar que se trataba de una ventana. Por fin llegué. Me asomé y pude ver mi cuerpo tumbado en una cama de hospital, mi mujer junto a mí, llorando a la vez que acariciándome la mano.",
	"La angustia la rodeaba... Él estaba allí otra vez. Lo sabía, no era posible, ella solo quería olvidarlo, olvidar todo aquello, pero él lo había conseguido. De nuevo dominaba su espacio.",
	"La cosa es que, para mí, una mujer tiene que tener una cierta cantidad de amortiguación. Algo que quede entre los dos, a modo de aislamiento. De lo contrario, estás justo contra su alma desnuda.",
	"¡Hola, soy Yogurt, el creador del universo! La longitud de estos textos guarda una relación de longitud, nunca deben ser demasiado largos. La calidad de estos también debe ser buena, que hablen de algo interesante y coherente. Esto no es por nada en concreto, pero lo hago por estilo.",
	"todo el como piensa correr las que son porque papas mis viste presidente mismo exactamente millones lindo quien palabra hiciste placer voy señor podría adelante busca diablo sueños habrá juntos",
	"contra pero que siento web foto poner buena permiso luna espera cuerpo contra toma tiene hablando siquiera algo ataque coche temo auto entonces horas venir mujeres más hola estamos",
	"hija hablar da decir pase seguridad existe tiene ella tengo nuestro será sabes ni todos hay padre también sí mis solo hacer hecho tiempo ver las pero son sus todo tú fuera a algo le hay decir gente nunca",
	"seguro banco estos recuerdo bueno visto van habrá empezar sistema maldita cuántos pagar pasar oí algunas luna cosa pobre gustan mí nunca vamos sabe quieres mismo donde los tú dije huevos vida mismo",
	"irnos ocurrió tenías señora puerta resto antes pasando libre camino culpa muchacho hará luna para odio hacemos cállate pobre ellas tendrás como verdad no aquí nuevo un bueno solo que hola sobre las años ella nos nombre hacer sabes siento ser ahí entonces les puedo una antes ese otro decir todos dije crees qué tenemos cuando ver haciendo una hacer las con de las por papas con padre",
	"Quisiera ser convexo para tu mano cóncava, y como un tronco hueco para acogerte en mi regazo y darte sombra y sueño. Suave y horizontal e interminable para la huella alterna y presurosa de tu pie izquierdo y de tu pie derecho. La de todas las formas como agua, siempre a gusto, en cualquier vaso siempre abrazándote por dentro. Y también como vaso para abrazarte por fuera al mismo tiempo. Como el agua hecha vaso tu confín -dentro y fuera- siempre exacto.",
	"Resulta evidente que los materiales digitales hipertextuales tienen ventajas sobre el material no digital y no hipertextual. Su capacidad para enlazar información, la posibilidad de proporcionar diferentes rutas de navegación y lectura, y la posibilidad de incorporar diferentes medios multimedia posibilitan un entorno rico en información que puede redundar en una mayor calidad en el aprendizaje del estudiante. Además, la capacidad del hipermedia de proporcionar entornos digitales capaces de incorporar representaciones visuales de unidades de conocimiento en forma estática o dinámica hace que este tipo de documentos sean especialmente indicados para poder explicar a los estudiantes representa.",
	"Es una verdad mundialmente reconocida que un hombre soltero, poseedor de una gran fortuna, necesita una esposa. Sin embargo, poco se sabe de los sentimientos u opiniones de un hombre de tales condiciones cuando entra a formar parte de un vecindario. Esta verdad está tan arraigada en las mentes de algunas de las familias que lo rodean, que algunas le consideran de su legítima propiedad y otras de la de sus hijas.",
	"Tenía que soltar al chico. Merece una oportunidad. Y si para conseguirlo, teníamos que distanciarnos, mandar todo a la mierda, y olvidar lo que ha habido entre nosotros, comprenderás que debía hacerlo.",
	"Ha matado a más hombres que la sequía, Alice. Crees que no es como tú pensabas que era, pero sigue siendo muy peligroso. Es un asesino, alguien debe tener la dignidad de llevarlo ante la justicia. No estoy solo en esto, confía un poco en mí.",
	"Al hombre le ocurre lo mismo que al árbol. Cuanto más quiere elevarse hacia la altura y hacia la luz, tanto más fuertemente tienden sus raíces hacia la tierra, hacia abajo, hacia lo oscuro, lo profundo - hacia el mal.",
	"Entre el discorde estruendo de la orgía acarició mi oído, como nota de música lejana, el eco de un suspiro. El eco de un suspiro que conozco, formado de un aliento que he bebido, perfume de una flor, que oculta crece en un claustro sombrío.",
	"Cuando vine a vivir a Roma con 26 años caí rápidamente y sin darme cuenta en lo que se podría llamar el vórtice de la mundanidad. Pero yo no quería ser simplemente un mundano, yo quería ser el rey de lo mundano. No quería solo participar en las fiestas, también quería tener el poder de hacerlas fallar. Y triunfé.",
	"Admiro tu coraje por venir solo, Roy. Viajando todo este camino siguiéndome hasta aquí. Me hace preguntarme qué podríamos haber logrado juntos. Pero supongo que el destino me ha privado del compañero que debería haber tenido.",
	"Wikipedia es un wiki. Esto significa que cualquier persona puede editar fácilmente cualquier página (excepto las protegidas) y sus cambios serán visibles inmediatamente.",
	"A veces se quedaba en silencio, y yo me preguntaba qué pasaba por su mente. Muchas tardes, antes de la noche, salimos a dar un paseo por su barrio. Rebeca era capaz de maravillarse mirando la enredadera de una fachada o de acercarse a sentir el aroma de unos jazmines al atardecer.",
	"Una pieza de puerco salvaje o jabalí tomarás; y córtalo a tajadas como dos dedos; y hacer tajadas de tocino gordo delgadas, y atravesarlas * por las piezas del puerco jabalí tantas cuantas quisieres.",
	"Aunque opino que si usan una botella, es un arma letal, así que deben asumir las consecuencias. Si te hubiera atacado solo con las manos sería diferente, y no habría sido justo.",
	"Los héroes no son siempre los que ganan. A veces, son los que pierden. Pero siguen luchando, y siguen aguantando. No se rinden. Eso es lo que los convierte en héroes.",
	"Y entonces les presento a Harvey. Y él es más grande e importante que todo lo que ellos me ofrecen. Y cuando se van, se van impresionados. La misma gente raramente vuelve, pero eso es... envidia, amigo mío.",
	"¿Quién fue Carl Sagan? A pesar de ser famoso por sus apariciones televisivas, este científico hizo mucho más de lo que mucha gente piensa. Como científico, Sagan marcó un cambio en la ciencia planetaria en la década de 1970 como el joven profesor de Harvard que era, \"en un momento en el cual la ciencia planetaria estaba en un remanso,\" dice Poundstone.",
	"Ella me daba la mano y no hacía falta más. Me alcanzaba para sentir que era bien acogido. Más que besarla, más que acostarnos juntos, más que ninguna otra cosa, ella me daba la mano y eso era amor.",
	"Un libro no es nada más que un lienzo de palabras, que son a su vez las notas en los márgenes del guión de la película más colosal que jamás se haya esculpido.",
	"La felicidad real siempre aparece escuálida por comparación con las compensaciones que ofrece la desdicha. Y, naturalmente, la estabilidad no es, ni con mucho, tan espectacular como la inestabilidad. Estar satisfecho de todo no posee el encanto que supone mantener una lucha justa contra la infelicidad, ni el pintoresquismo del combate contra la tentación o contra una pasión fatal o una duda. La felicidad nunca tiene grandeza.",
	"Es necesario esperar, aunque la esperanza haya de verse siempre frustrada, pues la esperanza misma constituye una dicha, y sus fracasos, por frecuentes que sean, son menos horribles que su extinción.",
	"Me di cuenta de que estábamos parados en el medio de la calle, el semáforo había cambiado y había un muro de tráfico que avanzaba hacia nosotros. Y entonces me congelé, no me pude mover.",
	"Se situó ante mí repentinamente. No era posible. Apenas un segundo antes la aleta del escualo distaba a más de veinte metros. Ahora lo tenía delante, sin poder esquivar su voracidad. Las risas de mis compañeras se esfumaron con el atardecer, y como si hubiesen robado dos horas al tiempo, se hizo noche cerrada.",
	"La clave de la felicidad no estaba en congelar todo placer momentáneo y aferrarse a ellos, sino en asegurarse de que la vida produjera muchos momentos futuros que esperar con ganas.",
	"Había notado que el sexo se parecía mucho a la cocina: fascinaba a la gente, a veces compraban libros llenos de recetas complicadas e imágenes interesantes, y a veces cuando tenían mucha hambre creaban grandes banquetes en su imaginación, pero al final del día se conformaban muy felizmente con huevo y patatas fritas, si estaban bien hechos y tal vez tenían una rodaja de tomate.",
	"En la batalla, en el bosque, en el precipicio de las montañas, en el vasto y obscuro océano, entre lanzas y flechas, durante el sueño, en la confusión, en la profunda humillación, los buenos actos que un hombre ha realizado con anterioridad lo defienden.",
	"Somos todos los trozos de lo que recordamos. Tenemos en nuestro interior las esperanzas y los temores de aquellos que nos aman. Mientras haya amor y memoria, no existe la auténtica pérdida.",
	"Viento del Sur, moreno, ardiente, llegas sobre mi carne, trayéndome semilla de brillantes miradas, empapado de azahares. Pones roja la luna y sollozantes los álamos cautivos, pero vienes ¡demasiado tarde! ¡ya he enrollado la noche de mi cuento en el estante!",
	"En los términos más simples, la frustración es una emoción que proviene de sentirse bloqueada para lograr un objetivo deseado. Hay fuentes internas de frustración, así como fuentes externas.",
	"Hemos ido pasando de los hogares de aire a los hogares de lluvia. Hemos ido pasando de los hogares ajenos, lacustres de un pasado que no era nuestro, a los hogares abiertos en dos por la proa del frío.",
	"En cuestión de segundos, se les estremeció el alma. No solo por lo que acababan de hacer, sino por la enorme serpiente dorada que estaba frente a ellos.",
	"Tenemos que luchar. No importa lo terrible que sea este mundo, no importa que tan cruel llegue a ser. Si estamos frente a una muerte segura, debemos seguir el camino que al menos nos dé una oportunidad de victoria. ¡Tenemos que luchar!",
	"Súbitamente comprendí mi peligro: me había dejado soterrar por un loco, luego de tomar un veneno. Las bravatas de Carlos transparentaban el íntimo terror de que yo no viera el prodigio; Carlos, para defender su delirio, para no saber que estaba loco, tenía que matarme. Sentí un confuso malestar, que traté de atribuir a la rigidez, y no a la operación de un narcótico. Cerré los ojos, los abrí. Entonces vi el Aleph.",
	"Había notado que el sexo se parecía mucho a la cocina: fascinaba a la gente, a veces compraban libros llenos de recetas complicadas e imágenes interesantes, y a veces cuando tenían mucha hambre creaban grandes banquetes en su imaginación, pero al final del día se conformaban muy felizmente con huevo y patatas fritas, si estaban bien hechos y tal vez tenían una rodaja de tomate.",
	"Sintió en su boca el suave olor de la fiebre y lo aspiró como si quisiera llenarse de las intimidades de su cuerpo. Y en ese momento se imaginó que ya llevaba muchos años en su casa y que se estaba muriendo. De pronto tuvo la clara sensación que no podría sobrevivir a la muerte de ella. Se acostaría a su lado y querría morir con ella. Conmovido por esa imagen hundió en ese momento la cara en la almohada junto a la cabeza de ella y permaneció así durante mucho tiempo...",
	"Tenemos que luchar. No importa lo terrible que sea este mundo, no importa que tan cruel llegue a ser. Si estamos frente a una muerte segura, debemos seguir el camino que al menos nos dé una oportunidad de victoria. ¡Tenemos que luchar!",
	"Para Lacan, el psicoanálisis no es en principio una teoría y una técnica de tratamiento de perturbaciones psíquicas, sino de una teoría y una práctica que confronta a los individuos con la dimensión más radical de la existencia humana.",
	"Mire, señora, llevo conduciendo esta ruta por quince años. Los he traído a que se traten y me los he llevado a casa después del tratamiento. Los cambia.",
	"En cuestión de segundos, se les estremeció el alma. No solo por lo que acababan de hacer, sino por la enorme serpiente dorada que estaba frente a ellos.",
	"En los términos más simples, la frustración es una emoción que proviene de sentirse bloqueada para lograr un objetivo deseado. Hay fuentes internas de frustración, así como fuentes externas.",
	"En mi vida he hecho ejercicio, me saltaba las clases de educación física en el instituto con la excusa de que estaba mal de las rodillas, en mi casa comemos albóndigas gigantes rellenas de queso cada vez que podemos, y tenemos para cada día de la semana un tarro enorme de dulces. El punto de todo esto es que no puedo ser fitness cuando toda mi vida he sido fatness.",
	"El viaje no termina jamás. Solo los viajeros terminan. Y también ellos pueden subsistir en memoria, en recuerdo, en narración... el objetivo de un viaje es solo el inicio de otro viaje.",
	"Descansar, cambiar de ocupación, hacer otras cosas, es muchas veces una manera de afilar nuestras herramientas. Seguir haciendo algo a la fuerza, en cambio, es un vano intento de reemplazar con voluntad la incapacidad de un individuo en un momento determinado.",
	"Primero vinieron los de capa gris a comerse todos mis cerdos. Después los de capa azul forzaron a mis hijos a luchar. Después los de capa verde hicieron a mi mujer prostituta. Después los de capa marrón quemaron mi casa. No tengo nada salvo mi vida; ahora vienen los de capa negra a robármela.",
	"Después de puesta la vida tantas veces por su ley al tablero; después de tan bien servida la corona de su rey verdadero; después de tanta hazaña a que no puede bastar cuenta cierta, en la su villa de Ocaña vino la Muerte a llamar a su puerta.",
	"Jamás pegaría a una mujer, Chloe. Me defendería si quisiera darme con una botella. Eso es diferente, es autodefensa. Y lucharía si ella supiera karate. Pero en general jamás pegaría a una mujer, Chloe, te lo aseguro.",
	"Para comprender el estado de la humanidad puede que baste con saber que la mayoría de los grandes triunfos y grandes catástrofes de la historia no se deben a que las personas son buenas en esencia o malas en esencia, sino a que las personas son en esencia personas.",
	"La jubilación no suena tan mal. Largas caminatas por el campo, podando rosas en un jardín con mi media naranja, criando algunos cachorros. Me lo he ganado.",
	"Nos dicen que recordemos a los ideales, no al hombre, porque con un hombre se puede acabar. Pueden detenerle, pueden matarle, pueden olvidarle, pero cuatrocientos años más tarde los ideales aún pueden seguir cambiando el mundo.",
	"Hace un siglo, la humanidad se enfrentó a un nuevo enemigo. La diferencia de poder entre el hombre y su enemigo era abrumadora, hasta un punto donde el hombre casi se extinguió. Los humanos que sobrevivieron construyeron tres muros: María, Rose y Sina. Y gracias a eso, pudieron vivir en paz durante un siglo.",
	"Sin embargo, su padre el rey sabía que el alma de la princesa regresaría quizá en otro cuerpo, en otro tiempo y en otro lugar y él la esperaría hasta su último aliento, hasta que el mundo dejara de girar.",
	"No te rodees de personas que lo único que hacen es pensar en su aspecto físico. Eres mucho más que piel y huesos. Eres el conjunto vivo y armonizado de infinidad de atributos que pueden enloquecer de placer a cualquiera, si te dispones a ello.",
	"Era muy tarde y estaba lejos de casa. No sabía cómo había llegado hasta esta parte de la ciudad y es que llevaba una temporada abusando en exceso del alcohol. Se trataba de una zona marginal con todos los males que nuestra sociedad nos brinda.",
	"Volverán las oscuras golondrinas en tu balcón sus nidos a colgar, y otra vez con el ala a sus cristales jugando llamarán. Pero aquellas que el vuelo refrenaban tu hermosura y mi dicha a contemplar, aquellas que aprendieron nuestros nombres... esas... no volverán.",
	"Los hombres en la habitación de repente se dieron cuenta de que no querían conocerla mejor. Era hermosa, pero en la forma en que un incendio forestal es hermoso, algo para admirar desde la distancia, no de cerca.",
	"Raras veces nos topamos con un individuo capaz de revisar la idea que tiene de su propia inteligencia y sus propios límites bajo la presión de los acontecimientos, hasta el punto de abrirse a nuevas perspectivas sobre lo que aún puede aprender.",
	"La mejor venganza es no parecerse a tu enemigo, para así nunca vivir en una envidia y sufrimiento que tendrá su fin en la agonía interminable luchando con tus propios pensamientos.",
	"Ahí estaba, enorme. Solos él y yo, él como experto depredador, en su terreno, yo como usurpadora, violando sus dominios. Me rodeaba, tan pronto lo tenía delante como detrás, inspeccionaba a su víctima con serenidad maldita.",
	"Sin embargo, su padre el rey sabía que el alma de la princesa regresaría quizá en otro cuerpo, en otro tiempo y en otro lugar y él la esperaría hasta su último aliento, hasta que el mundo dejara de girar.",
	"Pero en ese mismo instante hubo un resplandor, como si un rayo hubiese salido de las entrañas mismas de la tierra, bajo la ciudad. Durante un segundo vieron la forma incandescente, enceguecedora y lejana en blanco y negro, y la torre más alta resplandeció como una aguja rutilante; y un momento después, cuando volvió a cerrarse la oscuridad, un trueno ensordecedor y prolongado llegó desde los campos.",
	"Después de eso tuvimos una breve conversación sobre cómo tu cuerpo a veces puede parecer totalmente ajeno. Dijo que su cuerpo puede sentirse como una burocracia distante controlada por telegramas desde su cerebro, y yo le dije que a veces mi cuerpo es como el de Mario Mario, que se controla con un mando de videojuegos de Nintendo. El apellido de Mario es Mario.",
	"La libertad, Sancho, es uno de los más preciosos dones que a los hombres dieron los cielos; con ella no pueden igualarse los tesoros que encierran la tierra y el mar; por la libertad, así como por la honra, se puede y debe aventurar la vida.",
	"No hay más que tres acontecimientos importantes en la vida: nacer, vivir y morir. No sentimos lo primero, sufrimos al morir, y nos olvidamos de vivir.",
	"Arenas movedizas, cuchillos del revés, alambradas con espinas, serpientes de cascabel, un río de gasolina y una cerilla en la mano por si decide volver. Voy a cambiar de rostro, de nombre y de ciudad, voy a volverme invisible, ocultar mi identidad, construir un muro alto para encerrarme dentro y tirar la llave al mar; para que ella no me encuentre, para no volverla a ver, olvidar la vida que soñé.",
	"Hay cien mil calles en esta ciudad. Tú no necesitas saber la ruta. Me das un lugar y una hora y yo te doy un intervalo de 5 minutos. Si pasa algo durante ese período, soy tuyo... pase lo que pase. Si algo pasa un minuto antes o después... te las arreglas solo. ¿Entiendes?",
	"Si deseas ser el rey de la jungla, no es suficiente actuar como un rey. Debes ser el rey. Y no puede haber ninguna duda. Porque la duda provoca el caos y su propia desaparición. Mi reina me dijo eso.",
	"Mira, ambos sabemos que el crecimiento es solo el 50% del negocio. Necesito tus conexiones europeas. He visto cómo se hace la salchicha. Ahora háblame de las carnicerías.",
	"El día de Navidad de 1560, alcanzamos el último paso sobre los Andes y por primera vez miramos hacia abajo a la jungla legendaria. Por la mañana leí la misa y luego descendimos a través de las nubes.",
	"El primer viaje de Cristóbal Colón por el Atlántico rumbo al oeste, zarpó el 3 de agosto de 1492 desde el puerto de Palos. Los preparativos habían sido arduos y habían tomado mucho tiempo. Conseguir las embarcaciones y la tripulación resultó muy difícil, pues Colón era un desconocido para la gente de mar de la zona. De las tres naves que harían el viaje, dos fueron entregadas a Colón en Palos, en virtud de un castigo que pesaba sobre el puerto y que, por Real Provisión, obligaba a sus autoridades a cederlas. La tercera, en tanto, fue arrendada a Juan de la Cosa con fondos adquiridos de prestamistas.",
	"El segundo viaje colombino se hizo en prosecución de tres objetivos: socorrer a los españoles del fuerte de la Navidad; continuar los descubrimientos, tratando de alcanzar las tierras del Gran Khan; y colonizar las islas halladas anteriormente.",
	"Casi tres años tuvo que aguardar Colón para poder salir en su viaje siguiente. Durante ellos restableció su mermado prestigio. La nueva expedición, que costó 4.150.800 maravedises, estaba formada por ocho embarcaciones. Hubo mucha dificultad para buscar nuevos colonos, ya que los informes venidos de indias habían apagado el entusiasmo popular.",
	"El último viaje de Colón es notable sobre todo por los nuevos descubrimientos, principalmente a lo largo de la costa de América Central. También es de interés para los historiadores, que valoran las descripciones de las culturas nativas encontradas por la pequeña flota de Colón, especialmente las secciones relativas a los comerciantes mayas.",
	"Sin embargo, en su mayor parte, el cuarto viaje fue un fracaso desde cualquier punto de vista. Muchos de los hombres de Colón murieron, sus barcos se perdieron y nunca se encontró el paso hacia el oeste. Colón no volvió a navegar y, cuando murió en 1506, estaba convencido de haber encontrado Asia, aunque la mayor parte de Europa ya aceptaba que las Américas eran un \"Nuevo Mundo\" desconocido.",
	"Maté a aquel pequeño. Sé que no quería matarlo. Pero yo tuve la culpa de aquello. Y si no hubiera estado allí aquel niño aún seguiría aquí, y eso no puedo cambiarlo.",
	"No te obstines, pues, en mantener como única opinión la tuya creyéndola la única razonable. Todos los que creen que ellos solos poseen una inteligencia, una elocuencia o un genio superior a los de los demás, cuando se penetra dentro muestran solo la desnudez de su alma.",
	"Algún día en cualquier parte, en cualquier lugar indefectiblemente te encontrarás a ti mismo, y esa, solo esa, puede ser la más feliz o la más amarga de tus horas.",
	"Por todo esto, me parece que por ahora lo más conveniente es que ingrese en un centro adecuado y que allí se someta a una terapia. Es triste, pero no hay más remedio. Tal como te dije antes, hay que tener paciencia. Ir desenredando la madeja, hilo a hilo, sin perder la esperanza. Por más negra que esté la situación, el hilo principal existe, sin duda. Cuando uno está rodeado de tinieblas, la única alternativa es permanecer inmóvil hasta que sus ojos se acostumbren a la oscuridad.",
	"Le hemos dado la espalda al deber más importante que tenemos, a vivir una vida que es rica en experiencias elegidas por nosotros mismos. Carpe diem, amigos.",
	"Deseaba ver con mis propios ojos todos aquellos secretos guardados. La casualidad y el tiempo me habían concedido la oportunidad esperada. Por mi mente ya solo pasaba la idea de entrar a comprobar lo comentado en sordina por todos los otros.",
	"Esos cantos tétricos, cada vez más insoportables en medio del bosque, en medio de la lluvia, en medio de la espesa vegetación por la que sin embargo se filtraba la tenebrosa luz, hicieron que con las fuerzas que no tenía soltara un último grito de desesperación, de angustia.",
	"Se fue como un enfermo se va de la vida: despacio, con dolor, pero sabiendo que, al final, es lo mejor. Que después, quizás, vendrá el cielo y si no, al menos, el descanso.",
	"En un rincón del jardín, detrás de todos los frambuesos, había una maleza tupida donde no crecían ni flores ni frutales. En realidad, era un viejo seto que servía de frontera con el gran bosque, pero nadie lo había cuidado en los últimos veinte años, y se había convertido en una maleza impenetrable. La abuela había contado que el seto había dificultado el paso a las zorras que durante la guerra venían a la caza de las gallinas que andaban sueltas por el jardín.",
	"Se comprometió a avisar a los numerosos refugiados del Caribe que vivían en la ciudad, por si querían rendir los últimos honores a quien se había comportado como el más respetable de todos, el más activo y radical, aún después de que fue demasiado evidente que había sucumbido a la rémora del desencanto.",
	"Toda la vida es una gran cadena, cuya naturaleza se nos muestra en cada uno de los eslabones. Como todas las otras artes, la ciencia de la deducción y el análisis solo puede adquirirse mediante un estudio paciente y prolongado, y no hay vida lo bastante larga para permitir a un mortal alcanzar su grado máximo de perfección.",
	"Hay diferentes clases de oscuridad. Está la oscuridad que asusta, la oscuridad que tranquiliza, la oscuridad inquieta. Está la oscuridad de los amantes y la oscuridad de los asesinos. Se convierte en lo que cada uno desea que sea, necesita que sea. No es totalmente mala ni totalmente buena.",
	"La gente viene y me habla, pero soy incapaz de oírles por encima de las reverberaciones de mi mente. La gente viene y me mira, pero no puedo ver más que la sombra de sus ojos.",
	"Me gustas cuando callas porque estás como ausente, y me oyes desde lejos, y mi voz no te toca. Parece que los ojos se te hubieran volado, y parece que un beso te cerrara la boca. Me gustas cuando callas y estás como distante, y estás como quejándote, mariposa en arrullo. Y me oyes desde lejos, y mi voz no te alcanza, déjame que me calle con el silencio tuyo.",
	"No hay dioses en el universo, no hay naciones, no hay dinero, ni derechos humanos, ni leyes, ni justicia fuera de la imaginación común de los seres humanos.",
	"Las obligaciones forman parte de la brujería. Si de verdad quieres poner nerviosa a una bruja, hazle un favor que no tenga manera de devolverte. La obligación incumplida la roerá por dentro como un sordo dolor de tripas.",
	"Lo esencial es no perder la orientación. Siempre pendiente de la brújula, siguió guiando a sus hombres hacia el norte invisible, hasta que lograron salir de la región encantada.",
	"La verdadera inquietud del ser humano, el inexorable camino que abre paso en su torpe naturaleza es el inquebrantable, anodino, inverosímil y superficial deseo de poder. El sentimiento de una vida que se antoja breve pero intensa y que solo hace sentir satisfacción bajo los brebajes del amor y del placer no es más que una gota de lluvia que se interpone en el monstruo de la ambición.",
	"Nos preocupamos por ti y te conocemos. Pero también conocemos nuestras falsedades. Y por eso, a diferencia de ti, siempre acabamos hablando de tonterías y trivialidades, porque no nos queremos rebelar en nuestra ruindad.",
	"Lo siguiente que dice Jordana me hace darme cuenta de que es demasiado tarde para salvarla: \"Me di cuenta de que cuando enciendes un fósforo, la llama tiene la misma forma que una lágrima que cae.\"",
	"Nunca nadie nace odiando a otra persona a causa del color de su piel, su origen o su religión. Las personas tienen que aprender a odiar y, si pueden aprender a odiar, es posible enseñarles a amar, porque el amor llega de forma más natural al corazón que su opuesto.",
	"Algunos humanos harían cualquier cosa para ver si es posible hacerlo. Si colocas un interruptor grande en una cueva remota, con un letrero que dice \"Interruptor del fin del mundo. Por favor no tocar\", alguien lo presionaría antes siquiera de que se secara la pintura.",
	"Ella es... como el número pi. Me gustan tanto las matemáticas, porque son pura lógica. Los números son racionales, predecibles. Pero de repente, en medio de tanta armonía, aparece el número pi. Un número... misterioso, infinito, es un número que está vivo, crea su propio camino sin seguir patrones establecidos. Y eso hace que las matemáticas, además de lógica, también sean magia. Eso era Margarita para mí.",
	"El movimiento del vagón y los chirridos de las vías son gritos ensordecedores que me impiden dormir. Levanto la cabeza despacio resignado a no pegar ojo el resto del camino. Miro el reloj y me inquieta la soledad del vagón.",
	"Me enamoré de su coraje, su sinceridad y su llameante respeto propio. Y esas son las cosas en las que creo, incluso si el mundo entero se entrega a las salvajes sospechas de que ella no era todo lo que ella debe ser. La amo y eso es el principio de todo.",
	"Las escaleras se suben de frente, pues hacia atrás o de costado resultan particularmente incómodas. La actitud natural consiste en mantenerse de pie, los brazos colgando sin esfuerzo, la cabeza erguida aunque no tanto que los ojos dejen de ver los peldaños inmediatamente superiores al que se pisa, y respirando lenta y regularmente.",
	"Libre, y para mí sagrado, es el derecho de pensar... La educación es fundamental para la felicidad social; es el principio en el que descansan la libertad y el engrandecimiento de los pueblos.",
	"Hoy es siempre todavía. Toda la vida es ahora. Y ahora es momento de cumplir las promesas que nos hicimos. Porque ayer no lo hicimos. Porque mañana es tarde. Ahora.",
	"A través de todas las generaciones de la raza humana, ha habido una constante guerra: una guerra contra el miedo. Los que tienen el valor de vencerlo, son hechos libres y los que son conquistados por él, sufren hasta tener el valor para derrotarlo, o se los lleva la muerte.",
	"Conocí a un viajero de una tierra antigua que dijo: dos enormes piernas pétreas, sin su tronco se yerguen en el desierto. A su lado, en la arena, semihundido, yace un rostro hecho pedazos, cuyo ceño y mueca en la boca, y desdén de frío dominio, cuentan que su escultor comprendió bien esas pasiones las cuales aún sobreviven, grabadas en estos inertes objetos, a las manos que las tallaron y al corazón que las alimentó.",
	"No lo sé, Ray. No sé en qué creo. Las cosas que te enseñan de niño nunca te abandonan del todo. Siempre he intentado comportarme en la vida. Si veo una anciana cargada con la compra..., bueno, no le llevo las bolsas de la compra, no voy tan lejos, pero, no sé... le abriré la puerta y todo eso, la dejaré pasar...",
	"Te podría decir que he intentado destapar el abismo bajo mi ilusoria conexión con el mundo, que está escrito en las estrellas, o que soy un paranoico del revés: estoy bastante seguro de que el mundo conspira para hacerme feliz. Lo cual es todo cierto, pero es mucho más simple que eso: simplemente me gusta divertirme.",
	"Ambos hombres habían sido entrenados para este momento, sus vidas habían sido una larga preparación para ello, habían sido seleccionados al nacer como los que presenciarían la respuesta, pero aún así se encontraron jadeando y retorciéndose como niños emocionados.",
	"En el momento en que entiendo verdaderamente a mi enemigo, en el momento en que le entiendo lo suficientemente bien como para derrotarlo, entonces, en ese preciso instante, también le quiero. Creo que es imposible entender realmente a alguien, saber lo que quiere, saber lo que cree, y no amarle como se ama a sí mismo. Y entonces en ese preciso momento, cuando le quiero... le destruyo. Hago que le resulte imposible volver a hacerme daño, lo trituro hasta que no existe.",
	"Entre tantas rarezas venidas de todas partes, Florentino Ariza estaba de todos modos entre los más raros, pero no tanto como para llamar demasiado la atención. Lo más duro que oyó fue que alguien le gritara en la calle: Al pobre y al feo, todo se les va en deseo.",
	"He pasado toda mi vida asustado, asustado de las cosas que podrían suceder. Pero todo cambió desde que me dijeron que tenía cáncer. Me levanto para darle una patada en los dientes al miedo.",
	"El individuo ha luchado siempre para no ser absorbido por la tribu. Si lo intentas, a menudo estarás solo, y a veces asustado. Pero ningún precio es demasiado alto por el privilegio de ser uno mismo.",
	"Cualquier mago de mala muerte puede invocar a un demonio, pero nunca había visto que alguien invocara a un ángel. Mi ritual salió de maravilla... hasta que me rogó que no lo regresara al \"Cielo\".",
	"Es de verdad admirable la lucha que lleva la humanidad desde tiempos inmemoriales, lucha incesante con la que se esfuerza por arrancar y desgarrar todas las ataduras que intenta imponerle el ansia de dominio de uno solo, de una clase o también de un pueblo entero. Es esta una epopeya que ha tenido innumerables héroes y ha sido escrita por los historiadores de todo el mundo.",
	"La mayoría de nosotros no somos lo suficientemente ingenuos para creer en el viejo mito de que los científicos son dechados de objetividad desprovista de prejuicios, igualmente abiertos a todas las posibilidades, que sólo llegan a conclusiones mediante el peso de la evidencia y la lógica del argumento. Comprendemos que los prejuicios, las preferencias, los valores sociales y las actitudes psicológicas desempeñan un importante papel en el proceso del descubrimiento.",
	"La confianza en uno mismo no equivale a la vanidad. Todo lo contrario: solo las naciones o los hombres seguros de sí mismos, en el mejor sentido de la palabra, son capaces de escuchar la voz de los demás, aceptarlos como iguales, perdonar a sus enemigos y expiar sus propias culpas.",
	"Cuando bajé del avión, el hombre me esperaba con un pedazo de cartón en el que estaba escrito mi nombre. Yo iba a una conferencia de científicos y comentaristas de televisión dedicada a la aparentemente imposible tarea de mejorar la presentación de la ciencia en la televisión comercial. Amablemente, los organizadores me habían enviado un chófer.",
	"No todo lo que es oro reluce, ni toda la gente errante anda perdida; a las raíces profundas no llega la escarcha, el viejo vigoroso no se marchita. De las cenizas subirá un fuego, y una luz asomará en las sombras; el descoronado será de nuevo rey, forjarán otra vez la espada rota.",
	"Todas estas expresiones revelan que el mexicano considera la vida como lucha, concepción que no lo distingue del resto de los hombres modernos. El ideal de hombría para otros pueblos consiste en una abierta y agresiva disposición al combate; nosotros acentuamos el carácter defensivo, listos a repeler el ataque.",
	"Un huérfano es una persona a quien la muerte de sus progenitores ha privado de la posibilidad de la ingratitud filial, carencia que toca con singular elocuencia todas las cuerdas de la simpatía humana. Cuando es joven, el huérfano es enviado a un asilo, donde cultivando cuidadosamente su rudimentario sentido de la ubicación, se le enseña a conservar su lugar. Luego se lo instruye en las artes de la dependencia y el servilismo y finalmente se lo suelta para que vaya a vengarse del mundo...",
	"Ahora, tiempo después, sabía que todo estaba a punto de terminar. Él seguía siendo la sombra que la acechaba en el silencio sin dejarle un instante, ese sentimiento impregnado a sí misma tanto tiempo. Hubo un silencio largo al que le siguió otro todavía más prolongado. Entonces, alguien comenzó a forzar la cerradura de la casa.",
	"Una mujer hermosa e ingenua que roba con destreza a altos ejecutivos, un pintor habitado por fuerzas misteriosas y un sacerdote que se enfrenta a un caso de posesión demoníaca en La Candelaria, historias que se tejen en torno a la de Campo Elías, héroe de la guerra de Vietnam, quien inicia su particular descenso a los infiernos obsesionado por la dualidad entre el bien y el mal, entre Jekyll y Hyde, y se convertirá en un ángel exterminador.",
	"Me han dado los resultados de los tests y el análisis de sangre y el resultado es horrible. No encontraron nada, como si fuera un asesino silencioso. Voy a morir algún día, probablemente, sin haber dejado ninguna marca en este pequeño e insignificante planeta. Voy a morir como una insignificante mota de polvo anónima.",
	"Tyler me consigue un trabajo de camarero, después me mete una pistola en la boca y me dice que para alcanzar la vida eterna primero tienes que morirte. Sin embargo, durante mucho tiempo Tyler y yo fuimos muy buenos amigos. La gente siempre me pregunta si conocía bien a Tyler Durden.",
	"La luz era tan mansa al mediodía, que cuando Pelayo regresaba a la casa después de haber tirado los cangrejos, le costó trabajo ver qué era lo que se movía y se quejaba en el fondo del patio. Tuvo que acercarse mucho para descubrir que era un hombre viejo, que estaba tumbado boca abajo en el lodazal, y a pesar de sus grandes esfuerzos no podía levantarse, porque se lo impedían sus enormes alas.",
	"La industriosidad es el camino a la inmortalidad; la pereza es el camino a la muerte. Los industriosos no mueren; con los perezosos es como si ya estuvieran muertos.",
	"El mundo no es todo alegría y color. El mundo es un lugar terrible y por muy duro que seas, es capaz de arrodillarte a golpes y tenerte sometido permanentemente si tú no se lo impides. Ni tú, ni yo, ni nadie golpea más fuerte que la vida, pero no importa lo fuerte que golpeas, sino lo fuerte que pueden golpearte. Y lo aguantas mientras avanzas. Hay que soportar sin dejar de avanzar, así es como se gana. Si tú sabes lo que vales ve y consigue lo que mereces, pero tendrás que soportar los golpes.",
	"Pero después se quejan, me gritan para que me dé prisa, no tienen fe en mí o en mi taxi. Sin embargo, es el mismo taxi, el mismo taxista, y volvemos por la misma ruta. No es agradable y no me dan propina. Después de esto será un ser humano perfectamente normal, y ya sabes qué malos son.",
	"Ese instante que no se olvida, tan vacío devuelto por las sombras, tan vacío, rechazado por los relojes... Ampáralo niña ciega de alma, ponle tus cabellos escarchados por el fuego, abrázalo pequeña estatua de terror, señálale el mundo convulsionado a tus pies... dile que los suspiros del mar humedecen las únicas palabras por las que vale vivir. Pero ese instante sudoroso de nada, acurrucado en la cueva del destino, sin manos para decir nunca, sin manos para regalar mariposas a los niños muertos.",
	"Entonces opté por enfermar, lo que me salvó de la desesperación, como siempre ocurre en estos casos. Por mucho que se aborrezca la vida, cuando el destino nos mantiene vivos a nuestro pesar, la debilidad humana hace que no podamos evitar agradecer en secreto esta fatalidad. La muerte es tan horrenda, que nadie entre nosotros es capaz de verla de cerca sin sentir pavor.",
	"Giovani Battista Della Porta fue el típico estudioso del Renacimiento. Conocedor de varias disciplinas, dejó para el mundo de la criptografía uno de  los primeros métodos de cifrado polialfabético. Consistía en una tabla con 22 filas, cada dos de ellas con un semialfabeto reversible.",
	"Un pangrama es un texto en el que se encuentran todas las letras de un alfabeto determinado. El veloz murciélago hindú comía feliz cardillo y kiwi. La cigüeña tocaba el saxofón detrás del palenque de paja. El cadáver de Wamba, rey godo de España, fue exhumado y trasladado en una caja de zinc que pesó un kilo.",
	"En criptografía, el cifrado por sustitución es un método de cifrado por el que unidades de texto plano son sustituidas con texto cifrado siguiendo un sistema regular; las \"unidades\" pueden ser una sola letra (el caso más común), pares de letras, tríos de letras, mezclas de lo anterior, entre otros. El receptor descifra el texto realizando la sustitución inversa.",
	"La función SHA2 utiliza la función hash criptográfica SHA2 para convertir una cadena de longitud variable en una cadena de caracteres. La cadena de caracteres es una representación de texto del valor hexadecimal de la suma de comprobación con el número especificado de bits.",
	"En la anterior manera de calcular el WPM de este bot, usábamos la variable \"time_elapsed\", pero esta se ha visto sustituida por ((sent_when_unixmilli_float64-test_started_when_float)-1000). Restándole 1000 milisegundos tanto por el lag, como el retraso de reacción del usuario.",
	"Christopher Latham Sholes es conocido como el padre de la máquina de escribir. Nació en una ciudad remota en Mooresburg, Pensilvania, el 14 de febrero de 1819 y murió el 17 de febrero de 1890 después de haber luchado contra la tuberculosis pulmonar durante nueve años. Es conocido por diseñar el teclado \"QWERTY\"",
	"En los últimos minutos mi cuerpo ha experimentado un diferencial térmico de más de doscientos grados y aunque yo todavía estaba detenido, mis tejidos lo recuerdan y se retuercen, vibran, se dilatan. Me siento como una nuez empapada, llena de mil millones de hormigas rojas frenéticas peleando entre ellas y mordiendo la cáscara buscando una salida.",
	"Nos gustaba la casa porque aparte de espaciosa y antigua (hoy que las casas antiguas sucumben a la más ventajosa liquidación de sus materiales), guardaba los recuerdos de nuestros bisabuelos, el abuelo paterno, nuestros padres y toda la infancia.",
	"Cuando llega el momento de marchar, muchos no saben que su enemigo marcha a la cabeza. La voz que los manda es la voz del enemigo. El que habla del enemigo es el propio enemigo.",
	"Había una casa abajo, junto al estruendo de las olas desbaratándose contra los cantiles, donde el amor era más intenso porque tenía algo de naufragio.",
	"Debía de tener unos diez u once años cuando desapareció su madre. Era una mujer alta, elegante, más bien callada, de movimientos lentos y precioso cabello rubio. A su padre lo recordaba de manera más vaga como un hombre moreno y delgado, vestido siempre pulcramente de negro.",
	"La forma en que actuó muestra que le falta la madurez que es necesaria para este trabajo. Necesito a alguien que tenga, pues primero, un mínimo de cuidado para este trabajo, pero también, alguien que quiera aprender y ampliar sus conocimientos. Y por lo que puedo decir, él no tiene ninguna de esas cualidades. Entiendo que él necesita dinero y trabajo, pero yo no soy caridad para los que viven en los sótanos de sus padres. Dile que necesita salir de su trasero y trabajar.",
	"Por la ventana de la habitación entran los cálidos rayos del sol, un sol que sale sobre buenos y sobre malos, que estuvo ahí millones de años antes de que llegásemos a esta vida, a veces caótica y frenética, pero otras, llena de luz y de alegría. Un sol que perdurará y que sobrevivirá a santos y a tiranos, a imperios que caerán y otros que vendrán. Un sol que, a su vez, es tan sólo un suspiro dentro del cosmos infinito.",
	"Cuentan que el diablo una vez se enamoró, y no supo qué demonios hacer con ese vacío en el estómago. Cuentan que Satanás conoció su propio infierno en la boca de una mujer cuyos besos lo hacían un simple mortal. Cuentan que el diablo no es tan diablo, y se le puede ver caminando solo, con ojos de melancolía.",
	"Observa la marcha de los acontecimientos: considera las migraciones sucesivas de los pueblos y llegarás a la misma conclusión que yo... Asia es la primera abastecedora del mundo, ¿no es verdad? Durante 4000 años trabaja, es fertilizada, produce, y después, cuando las piedras han reemplazado a las espigas doradas de Homero, sus hijos abandonan su seno agotado y marchito. Los ves entonces arrojarse sobre Europa joven y potente que los alimentaba desde hace 2000 años.",
	"Puede que yo sea la única persona sobre la faz de la Tierra que sepa que eres la mujer más fantástica de la Tierra. Puede que yo sea el único que aprecie lo asombrosa que eres en cada una de las cosas que haces, y en cada uno de los pensamientos que tienes. En cómo dices lo que quieres decir y en cómo casi siempre quieres decir algo que tiene que ver con ser sincero y bueno. Y creo que la mayoría de la gente se pierde eso de ti. Y el hecho de que yo sí lo capte me hace sentir bien conmigo mismo.",
	"No creo que el arte haya llegado a ningún callejón sin salida. El hombre siempre irá abriendo nuevas puertas; lo importante es saber adónde conducen esas puertas. Y luego tener fuerza para emprender el camino que se ve desde ellas.",
	"Se apartó para dejarlo salir, y a través de la puerta entreabierta vio los almendros de la plaza, nevados por el resplandor del amanecer, pero no tuvo valor para ver nada más.",
	"Tu decepción entiende a mi sarcasmo, es que no pueden ser lo que son sin haberse conocido alguna vez. Probablemente se rían del amor, pobre aquel, el amor no tiene la culpa, el amor es naturaleza y la naturaleza no muere sola, el ser humano la mata.",
	"Cuando me llegaron los últimos días, y las feas trivialidades de la vida me hundieron en la locura como esas gotas de agua que el torturador deja caer sin cesar sobre un punto del cuerpo de su víctima, dormir se convirtió para mí en un refugio luminoso. En mis sueños encontré un poco de la belleza que había buscado en vano durante la vida, y pude vagar por viejos jardines y bosques encantados.",
	"Estoy estable, tranquilo. Dormí bien, sin pesadillas. Estoy activo e involucrado. Soy consciente de mi entorno y de aquellos en mi esfera inmediata. Estoy atento, enfocado en lo esencial excluyendo todo lo demás. No estoy seguro del futuro, pero no estoy preocupado. Confiaré en los más cercanos a mí. Y compartiré sus cargas como ellos comparten la mía. Viviré, y amaré.",
	"Abrí los ojos y allí no estaba ella. Desesperado la busqué por toda la habitación. Lo único que hallé fue el vago rostro de su fragancia fúnebre en el pasillo. Se había ido y no me había llevado con ella.",
	"Cada vez le costaba más trabajo encontrar la inquietud y el dolor necesarios para componer una nueva canción, porque con los años había llegado a tener una gran paz interior y la rebeldía que lo movilizaba en la juventud se había transformado en la mansedumbre del hombre satisfecho consigo mismo.",
	"Durante todo un día de otoño, oscuro, triste, silencioso, en que las nubes se cernían bajas y plomizas en los cielos, crucé solo, a caballo, una región singularmente monótona del país, y al fin, cuando se extendían las sombras, me encontré a la vista de la melancólica casa de Usher.",
	"De vez en cuando, al regreso de una fiesta loca, la nostalgia agazapada detrás de la puerta los tumbaba de un zarpazo, y entonces ocurría una explosión maravillosa en la que todo era otra vez como antes, y por cinco minutos volvían a ser los amantes desbraguetados de la luna de miel.",
	"Pero el castillo se quedó rondando por las colinas y se supo que no pertenecía a la bruja, sino al mago Howl. El mago Howl tampoco era un santo. Aunque al parecer no quería abandonar las colinas, se rumoreaba que le divertía atrapar a jovencitas y quitarles el alma. Otros aseguraban que se comía sus corazones. Era un mago absolutamente frío y sin escrúpulos y ninguna joven estaría segura si él andaba cerca.",
	"En la cintura, y no en las extremidades, se halla el muelle principal de los movimientos del cuerpo. Los movimientos de las extremidades son lentos y cortos, mientras que los de la cintura son libres y largos. Un giro de un gran eje equivale a cientos de giros de pequeños ejes. Consigue una completa vacuidad y con perseverancia conserva un estado de reposo. Yo opto por la blandura cuando mi oponente opta por la firmeza y por perseguir cuando mi oponente opta por retirarse.",
	"En mi opinión, el terreno donde el desconocimiento del factor infancia llama poderosamente la atención es el del cumplimiento de condenas. Es cierto que los establecimientos penitenciarios actuales ya no se asemejan a las tétricas cárceles de siglos pasados, pero hay algo que ha cambiado muy poco: raras veces se plantea la pregunta acerca de los motivos que llevan al individuo a convertirse en criminal y de lo que puede hacer para no volver a tropezar con la misma piedra.",
	"Antes de llegar al verso final ya había comprendido que no saldría jamás de ese cuarto, pues estaba previsto que la ciudad de los espejos (o los espejismos) sería arrasada por el viento y desterrada de la memoria de los hombres en el instante en que Aureliano Babilonia acabara de descifrar los pergaminos, y que todo lo escrito en ellos era irrepetible desde siempre y para siempre, porque las estirpes condenadas a cien años de soledad no tenían una segunda oportunidad sobre la tierra.",
	"Las personas que dejan las cosas para más tarde suelen ser bastante optimistas acerca de su capacidad para completar una tarea. Esta convicción suele ir acompañada de expresiones de reafirmación que utilizan para autoconvencerse de que todo está bajo control. Mientras se dejan llevar por una falsa sensación de seguridad, el tiempo pasa. De repente, cruzan una línea imaginaria de partida y entonces se dan cuenta de que han perdido el control y de que les falta tiempo.",
	"El cerebro insomne también es una especie de teórico de la conspiración, creyendo demasiado en su propia importancia paranoica. Pensando que si parpadease y se quedara dormido, el mundo sería invadido por una calamidad abrumadora, que sus reflexiones obsesivas de alguna manera están manteniendo a raya.",
	"A mí me da igual que los dioses existan o no, yo soy el amo de mi vida; paso de los temas religiosos, pero jamás menospreciaré a alguien que crea en ellos.",
	"Ofelia, ten cuidado, hermana mía; mantente en retaguardia del cariño, no te expongas al peligro del deseo. La más recatada se prodiga si a la luna revela su belleza. Ni la virtud escapa a la calumnia.",
	"Solté una risa sin malicia: era el monstruo del lago. Pero estaba viejo. Ni fuego por la boca ni postura imponente con la que amedrentar. Sin apenas dientes y con una lengua flácida bífida que ya no asustaba. Tuve la sensación de que había pasado años sin comer en la superficie, así que le di la cesta de mi merienda.",
	"Tengo que darle la razón, padre. Ha venido aquí con las armas cargadas esta vez. Y está en lo correcto sobre una cosa, sobre hombres más fuertes que yo alcanzando su salvación. Pero se equivoca con respecto a otra: lo que más atormenta a un hombre es lo que no se le ordena hacer.",
	"Las aplicaciones que tienen una mayor probabilidad de causar dependencia son los chats, las redes sociales y los juegos de rol on-line. Estas aplicaciones permiten la comunicación mediante el juego de identidades y, en muchos casos, la disociación con las consecuencias de la vida real. La red permite que el sujeto se reinvente a sí mismo, adoptando otras identidades y creando otro mundo alternativo contando con la protección del anonimato.",
	"La mayoría de las personas no llegamos a ser lo que podríamos ser. Porque el desarrollo no es ser tanto o mejor que los otros, sino todo lo que uno pueda llegar a ser. Casi nadie, yo el primero, llega a todo lo que pueda ser. Todavía soy aprendiz de mí mismo.",
	"¿No es verdad, ángel de amor, que en esta apartada orilla más pura la luna brilla y se respira mejor? Esta aura que vaga llena de los sencillos olores de las campesinas flores que brota esa orilla amena; esa agua limpia y serena que atraviesa sin temor la barca del pescador que espera cantando al día, ¿no es cierto, paloma mía, que están respirando amor?",
	"Ella tiene la manía de leer el periódico en la bañera y el periódico, por supuesto, se moja, y la parte que no se moja, se humedece por el vapor, y, por si eso fuera poco, el periódico me impide verla desnuda.",
	"Yo también pensé que si podía recibir todo de brazos, no pasaba nada. O pensaba que podía. Pero... quiero llegar más alto. Saber que hay más opciones de juego y no usarlas es un peñazo.",
	"Las aplicaciones que tienen una mayor probabilidad de causar dependencia son los chats, las redes sociales y los juegos de rol on-line. Estas aplicaciones permiten la comunicación mediante el juego de identidades y, en muchos casos, la disociación con las consecuencias de la vida real. La red permite que el sujeto se reinvente a sí mismo, adoptando otras identidades y creando otro mundo alternativo contando con la protección del anonimato.",
	"Solté una risa sin malicia: era el monstruo del lago. Pero estaba viejo. Ni fuego por la boca ni postura imponente con la que amedrentar. Sin apenas dientes y con una lengua flácida bífida que ya no asustaba. Tuve la sensación de que había pasado años sin comer en la superficie, así que le di la cesta de mi merienda.",
	"Hay un proverbio chino que dicta: \"Si estás planificando para un año, siembra arroz. Si planificas para una década, planta árboles. Si tienes planes para una vida, educa.\"",
	"¿Qué pasaría si te dijera que loco es trabajar cincuenta horas a la semana en una oficina durante cincuenta años tras los cuales te echan y terminas en una residencia deseando morir antes de sufrir la indignidad de no poder llegar al baño a tiempo? ¿No lo considerarías una locura?",
	"Mi mejor arma está en el cerebro. Mi hermano tiene su espada, el rey Robert tiene su maza, y yo tengo mi mente... Pero una mente necesita de los libros igual que una espada de una piedra de amolar, para conservar el filo. Por eso leo tanto, Jon Snow.",
	"Como dos aves doradas posadas en el mismo árbol, el ego y el yo, íntimos amigos, viven en el mismo cuerpo. El primero come los frutos dulces y amargos del árbol de la vida, mientras que el segundo observa con indiferencia.",
	"No me parece que la luciérnaga extraiga mayor suficiencia del hecho incontrovertible de que es una de las maravillas más fenomenales de este circo, y sin embargo basta suponerle una conciencia para comprender que cada vez que se le encandila la barriguita, el bicho de luz debe sentir como una cosquilla de privilegio.",
	"Había gastado mucho dinero, mucho ingenio y mucha fuerza de voluntad para que no se le notaran los setenta y seis años que había cumplido el último marzo, y estaba convencido en la soledad de su alma de haber amado en silencio mucho más que nadie jamás en este mundo.",
	"Así, con tal entender, todos sentidos humanos conservados, cercado de su mujer, y de sus hijos y hermanos y criados, dio el alma a quien se la dio, el cual la ponga en el cielo y en su gloria, y aunque la vida perdió, dejónos harto consuelo su memoria.",
	"Cayó la noche y se hizo el silencio; no se oía otra cosa que un débil temblor de tierra bajo los pies de los ents y un roce, la sombra de un susurro, como de muchas hojas llevadas por el viento.",
	"Los científicos han calculado que la probabilidad de que exista algo tan evidentemente absurdo es una entre muchos millones. Pero los magos han calculado que dicha probabilidad tiene lugar nueve de cada diez veces.",
	"Pero tengo un secreto. Aunque construyan murallas que lleguen hasta el cielo, yo encontraré la forma de volar sobre ellas. Aunque intenten atraparme con cientos de armas, yo encontraré un modo de resistir. Y hay muchos como yo ahí afuera, más de los que crees. Gente que se niega a dejar de creer. Gente que se niega a resignarse. Gente que ama en un mundo sin murallas, gente que ama frente al odio, frente al rechazo, sin miedo y contra toda esperanza. Te amo. Recordarlo. Eso no pueden quitártelo.",
	"El tiempo. Pierdo la noción de su paso cuando estoy contigo, cuando te vas... El tiempo se acelera hasta el momento hasta que vuelvo a verte. Cuando me besas el tiempo ya no existe, él vuela, al igual que nosotros lo hacemos mientras nos hacemos uno.",
	"Todos mis instintos son una forma, y todos los hechos son otras, y me temo mucho que los jurados británicos todavía no han alcanzado ese tono de inteligencia cuando van a dar la preferencia a mis teorías.",
	"Esta fabulilla, salga bien o mal, se me ha ocurrido ahora, por casualidad. Cerca de unos prados que hay en mi lugar, pasaba un borrico por casualidad.",
	"Cuando Hildebranda murió, casi centenaria en su hacienda de Flores de María, encontraron su copia bajo llave en el armario del dormitorio, escondida entre los pliegues de las sábanas perfumadas, junto con el fósil de un pensamiento en una carta borrada por los años.",
	"El convoy entrará en el depósito de chatarra por aquí. Primero neutralizaremos el primer coche, luego el último creando una trampa... una jaula hermética llena de muchos muertos. Nathan, pon a todos en posición y tráeme esos tanques de propano.",
	"Yo siempre vi el mundo de una manera distinta, sentí siempre que, entre dos cosas que parecen perfectamente delimitadas y separadas, hay intersticios por los cuales, para mí, al menos, pasaba, se colaba, un elemento, que no podía explicarse con leyes, que no podía explicarse con lógica, que no podía explicarse con la inteligencia.",
	"De modo que palpó sus órganos internos con más intención que atención, y mientras tanto iba olvidándose de su propia sabiduría y descubriendo asombrado que aquella criatura de maravilla era tan bella por dentro como por fuera, y entonces se abandonó a las delicias del tacto, no ya como el médico mejor calificado del litoral caribe, sino como un pobre hombre de Dios atormentado por el desorden de los instintos.",
	"Allí estaban desnudos los malvados; una mitad iba dando la espalda, otra de frente, con pasos más grandes; tal como en Roma la gran muchedumbre, del año jubilar, allí en el puente precisa de cruzar en doble vía, que por un lado todos van de cara hacia el castillo y a San Pedro marchan; y de otro lado marchan hacia el monte.",
	"Muchas de las cosas que pasarán mañana, son la consecuencia de lo que hicimos ayer. Por eso, pienso que no hay actitud más egoísta que la de querer que una persona regrese cuando la ves completamente feliz y a esa persona, en un pasado, no la valoraste. Así que, si empiezas a echarme de menos, recuerda que yo no me alejé de ti, tú me dejaste ir.",
	"Allí estaban desnudos los malvados; una mitad iba dando la espalda, otra de frente, con pasos más grandes; tal como en Roma la gran muchedumbre, del año jubilar, allí en el puente precisa de cruzar en doble vía, que por un lado todos van de cara hacia el castillo y a San Pedro marchan; y de otro lado marchan hacia el monte.",
	"No podrías tumbarte a hacer versos ni a pensar en las musarañas; buscarías trabajos particulares para redondear el presupuesto y te acostarías a las tres de la mañana contento de ahorrar sueño y dinero. Porque tendrías que ahorrar, ahorrar como una urraca; quitándolo de la comida, del vestido, del tabaco... Y cuando llevases un montón de años haciendo eso, y ensayando negocios y buscando caminos, acabarías por verte solicitando cualquier miserable empleo para no morirte de hambre.",
	"Cuando llegué a la playa encontré sus ropas en la orilla. Me metí en el agua, parecía tranquila, nadé, buceé, pero todo fue en vano. Mientras me ahogaba vi su cuerpo flotando, me pareció que sonreía, como si hubiera expirado con una risa liberadora e histérica, ya no le perseguía nadie, ya no tenía miedo, el miedo ahora estaba conmigo.",
	"Su madre lo comprende a la perfección. Capta el significado de lo que le dice de un modo que los monosílabos hablados de la Tierra jamás serán capaces de replicar. Sabe que a su hijo le apetece estar un rato a solas, sin que le pregunten si está bien, sin que se preocupen por él, sin necesidad de ir dejándole folletos explicativos por toda la casa.",
	"Mi madre se preocupa de que pueda tener problemas mentales. Encontré un libro sobre delirios paranoicos adolescentes durante una búsqueda rutinaria de la habitación de mis padres.",
	"Una sonrisa, no siempre es la solución a todos los problemas que se nos presentan a lo largo de nuestra vida, pero quizá puede ser la mejor medicina para sobrellevar aquel problema y que no se convierta en una enorme carga para cada uno de nosotros.",
	"No era adorable, señor Bryton. Adorable es como nos referimos a los bebés chinos. Mi compañero pinnípedo era hermoso. Nunca he tenido una relación tan maravillosa con ningún otro ser, animal o humano. Por seis meses conocí la felicidad más absoluta y es la única dicha que esta vida abominable jamás me ha proporcionado.",
	"Yo imaginaba ver aquello a través de los recuerdos de mi madre; de su nostalgia, entre retazos de suspiros. Siempre vivió ella suspirando por Comala, por el retorno; pero jamás volvió. Ahora yo vengo en su lugar, traigo los ojos con que ella miró estas cosas, porque me dio sus ojos para ver.",
	"Había compartido años de profundas emociones junto a él, años de verdadera pasión enloquecida, años junto a un desconocido que no permitió verla abandonar la casa cuando ni su cuerpo ni su mente le permitían pasar ni un solo minuto más junto a ese hombre.",
	"Su dominio de un estilo de maniobra prudente dio como resultado la introducción del adjetivo \"karpoviano\" en el vocabulario del ajedrez. Define una estrategia de estrangulación del adversario metódica y silenciosa, como la de una pitón.",
	"Sí, mi laúd tenía defectos, pero ¿qué importa eso cuando se trata de asuntos del corazón? Amamos lo que amamos. La razón no entra en juego. En muchos aspectos, el amor más insensato es el amor verdadero. Cualquiera puede amar algo por algún motivo. Es tan fácil como meterse un penique en el bolsillo. Pero amar algo a pesar de algo es otra cosa. Conocer los defectos y amarlos también. Eso es inusual, puro y perfecto.",
	"Esto va por los locos, los inadaptados, los rebeldes, los problemáticos. Los que no encajan en ningún sitio. Los que ven las cosas de otra manera. No siguen las reglas y no respetan lo establecido. Puedes citarlos, estar en desacuerdo, glorificarlos o satanizarles. Pero lo único que no puedes hacer es ignorarlos.",
	"Cuando buscabas cariño, en mí pronto lo encontraste. Y cuando más te quería, te fuiste y me dejaste. De momento sentí ira. Pero pronto comprendí que tu amor no vale nada. Y la nada nada inspira.",
	"En todos los casos probables de la vida, así como en el aspecto laboral y como en la vida cotidiana siempre lo bueno y rápido es enemigo de lo perfecto.",
	"Nos dejamos llevar por el miedo a las habladurías, por ese estúpido qué dirán. El qué dirán de unos cuantos a los que nada debíamos, que nada nos dieron, ni nos darán.",
	"A las doce, como cada noche. Desde que alquilé una habitación en aquella oscura pensión, dormir era un sueño. Mi casero debió advertirme, antes de cobrarme por adelantado, de su costumbre de tocar el violín a medianoche.",

}