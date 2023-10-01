%5)Implemente un predicado que, a partir de una lista de cadenas string, filtre
%aquellas que contengan una subcadena que el usuario indique en otro
%argumento.
%Ej sub_cadenas(“la”, [“la casa, “el perro”, “pintando la
%cerca”],Filtradas).
%True
%Filtradas = [“la casa, “pintando la cerca”]


filtrar(_,[],[]).
filtrar(Subcadena,[Cadena|Resto],Filtradas):-
    sub_atom(Cadena, _, _, _, Subcadena),
    filtrar(Subcadena, Resto, FiltradasResto),
    Filtradas = [Cadena|FiltradasResto].

filtrar(Subcadena, [_|Resto], Filtradas) :-
    filtrar(Subcadena, Resto, Filtradas).
