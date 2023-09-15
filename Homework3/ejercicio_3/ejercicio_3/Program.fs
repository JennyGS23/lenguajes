// For more information see https://aka.ms/fsharp-console-apps //Jennifer Gonzalez Solis

// 3)	Realizar una función que obtenga el n-esimo elemento de una lista pero utilizando solamente map (sin recursión). Tip: Utilizar lista de índices como segunda lista de argumento para el map.
//
// n_esimo 2 [1,2,3,4,5]
// 3
// n_esimo 3 [1,2,3,4,5]
// 4
// n_esimo 6 [1,2,3,4,5]
// False

let n_esimo n lista =
    let indices = [0..(List.length lista - 1)]
    let elementos = List.map2 (fun i x -> (i, x)) indices lista
    match List.tryFind (fun (i, _) -> i = n) elementos with
    | Some (_, elemento) -> Some elemento
    | None -> None


let resultado1 = n_esimo 2 [1; 2; 3; 4; 5]
let resultado2 = n_esimo 3 [1; 2; 3; 4; 5]
let resultado3 = n_esimo 6 [1; 2; 3; 4; 5]

match resultado1 with
| Some valor -> printfn "%A" valor
| None -> printfn "False"

match resultado2 with
| Some valor -> printfn "%A" valor
| None -> printfn "False"

match resultado3 with
| Some valor -> printfn "%A" valor
| None -> printfn "False"

