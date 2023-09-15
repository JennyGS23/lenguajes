// For more information see https://aka.ms/fsharp-console-apps       //Jennifer González Solís
// 1)	Realizar ejercicio para desplazar (SHIFT) una lista de elementos n posiciones a la izquierda o la derecha según se indique por argumento. Ej:
//
// desplazar “izq” 3  [1,2,3,4,5]
//
// [4,5,0,0,0]
//
// desplazar “der” 2  [1,2,3,4,5]
//
// [0,0,1,2,3]
//
// desplazar “izq” 6  [1,2,3,4,5]
//
// [0,0,0,0,0]

let rec desplazarIzquierda n lista =
    match n with
    | 0 -> lista
    | _ ->
        match lista with
        | [] -> []
        | hd::tl -> desplazarIzquierda (n-1) (tl @ [0])
        

let rec desplazarDerecha n lista =
    match n with
    | 0 -> lista
    | _ ->
        match lista with
        | [] -> []
        | lst ->
            let resto = List.init (List.length lst - 1) (fun i -> List.item i lst)
            desplazarDerecha (n-1) (0::resto)

let desplazarElementos direccion n lista =
    
    match direccion with
    | "izq" ->
         desplazarIzquierda n lista
    | "der" ->
         desplazarDerecha n lista
    | _ -> failwith "Dirección no válida"



let listaOriginal = [1; 2; 3; 4; 5]

let resultadoIzquierda = desplazarElementos "izq" 3 listaOriginal
printfn "Desplazar a la izquierda 3 posiciones: %A" resultadoIzquierda

let resultadoDerecha = desplazarElementos "der" 2 listaOriginal
printfn "Desplazar a la derecha 2 posiciones: %A" resultadoDerecha

let resultadoIzquierda2 = desplazarElementos "izq" 6 listaOriginal
printfn "Desplazar a la izquierda 6 posiciones: %A" resultadoIzquierda2
