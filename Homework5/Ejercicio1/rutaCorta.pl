% Definición de conexiones con pesos
conectado(i, a, 2).
conectado(i, b, 3).
conectado(a, b, 1).
conectado(a, c, 4).
conectado(b, f, 2).
conectado(b, c, 3).
conectado(c, f, 5).

% Regla para conexiones considerando el peso
conectados(X, Y, Peso) :- conectado(X, Y, Peso).
conectados(X, Y, Peso) :- conectado(Y, X, Peso).

% Regla para encontrar la ruta más corta y su peso
ruta(Ini, Fin, Ruta, Peso) :-
    findall((P, R), (ruta2(Ini, Fin, [], R, P)), Rutas),
    sort(Rutas, [(Peso, Ruta) | _]).

% Regla para encontrar una ruta
ruta2(Fin, Fin, _, [Fin], 0).
ruta2(Ini, Fin, Visitados, [Ini | Resto], PesoTotal) :-
    conectados(Ini, Alguien, Peso),
    not(member(Alguien, Visitados)),
    ruta2(Alguien, Fin, [Ini | Visitados], Resto, PesoRestante),
    PesoTotal is Peso + PesoRestante.



%Forma de consultar:
%?-ruta(i,f,Ruta,Peso).

