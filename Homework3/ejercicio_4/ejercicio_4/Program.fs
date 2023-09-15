//Jennifer González Solís
//4) Modifique el ejercicio de búsqueda en profundidad visto en clase para ingresar pesos a
//los vértices (una recomendación es representarlos como [["i"],["a","b"],["3","6","5"]]
//para evitar tener que cambiar algo en el algoritmo actual).
//Implemente una función que encuentre el camino más corto en la búsqueda en profundidad basándose en la
//sumatoria de pesos. Implica cambiar la búsqueda en profundidad para que se lleven los pesos además de los vértices.

type Node = (string * string list * int list)
let private grafo1 = [
    ("i", ["a"; "b"], [7; 70]);
    ("a", ["i"; "c"; "d"], [2; 4; 1]);
    ("b", ["i"; "c"; "d"], [23; 5; 9]);
    ("c", ["a"; "b"; "x"], [2; 7; 4]);
    ("d", ["a"; "b"; "f"], [19; 30; 64]);
    ("f", ["d"], [2]);
    ("x", ["c"], [4])
]
let private grafo2 = [
    ("i", ["a"; "b"], [40; 2]);
    ("a", ["i"; "c"; "d"], [2; 4; 5]);
    ("b", ["i"; "c"; "d"], [90; 5; 9]);
    ("c", ["a"; "b"; "x"], [2; 7; 4]);
    ("d", ["a"; "b"; "f"], [19;2; 64]);
    ("f", ["d"], [5]);
    ("x", ["c"], [7])
]
type grafo = Node list

let rec vecinos (nodo: string) =
    function
    | [] -> []
    | (head, neigh, _) :: rest ->
        if head = nodo then neigh
        else vecinos nodo rest

let miembro (elem) (lista): bool =
    List.exists (fun x -> x = elem) lista

let extender ruta grafo = 
    List.filter
        (fun x -> x <> [])
        (List.map (fun x -> if (miembro x ruta) then [] else x::ruta) (vecinos (List.head ruta) grafo)) 

let rec obtenerPeso desde ``hasta`` = function
    | (inicio, camino, peso)::rest when inicio = desde ->
        List.tryFindIndex (fun x -> x = hasta) camino
        |> Option.bind (fun index -> if List.item index peso <> 0 then Some (List.item index peso) else None)
        |> Option.defaultValue 0
    | _::rest -> obtenerPeso desde hasta rest
    | [] -> 0


let calcularPeso (grafo: grafo) =
    fun ruta ->
        let rec acumularPeso acum = function
            | [] | [_] -> acum
            | x::y::rest -> acumularPeso (acum + obtenerPeso x y grafo) (y::rest)
        (ruta, acumularPeso 0 ruta)

let rutaCorta (ini: string) (fin: string) (grafo: grafo): string list =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif (List.head (List.head ruta)) = fin then
            List.rev (List.head ruta) :: prof_aux fin ((extender (List.head ruta) grafo) @ (List.tail ruta)) grafo       
        else
            prof_aux fin ((List.tail ruta) @ (extender (List.head ruta) grafo)) grafo
    prof_aux fin [[ini]] grafo |> List.map (calcularPeso grafo) |> List.minBy (fun n -> snd n) |> fst
    
    

[<EntryPoint>]
let main (args: string array): int =
    printf"Ruta más corta en el grafo 1: \n"
    rutaCorta "i" "x" grafo1 |> printfn "%A"
    printf"Ruta más corta en el grafo 2: \n"
    rutaCorta "i" "x" grafo2 |> printfn "%A"
    0

