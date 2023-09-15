// For more information see https://aka.ms/fsharp-console-apps  //Jennifer González Solís
// 2)	Haciendo uso de la función filter, implemente una función que, a partir de una lista de cadenas de parámetro, filtre aquellas que contengan una subcadena que el usuario indique en otro argumento. Ej
//
// sub_cadenas “la” [“la casa, “el perro”, “pintando la cerca”]
//
// [“la casa, “pintando la cerca”]


let filtrarCadenas (subcadena:string) (listaCadenas:string list) =
    listaCadenas |> List.filter (fun cadena -> cadena.Contains(subcadena))

let subcadenaBuscada = "la"
let listaCadenas = ["la casa"; "el perro"; "pintando la cerca"]

let resultado = filtrarCadenas subcadenaBuscada listaCadenas

printfn "%A" resultado
