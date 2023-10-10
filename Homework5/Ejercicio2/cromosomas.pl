%Jennifer González Solís
% Hechos de las personas con sus nombres, apellidos y cromosomas
persona("Persona1", "Apellido1", [1, 0, 0, 0, 1, 0, 0, 0, 1, 1]).
persona("Persona2", "Apellido2", [1, 1, 1, 1, 0, 0, 0, 0, 0, 0]).
persona("Persona3", "Apellido3", [0, 0, 0, 0, 1, 1, 1, 1, 1, 1]).
persona("Persona4", "Apellido4", [1, 0, 1, 0, 1, 0, 1, 0, 1, 0]).
persona("Persona5", "Apellido5", [0, 1, 0, 1, 0, 1, 0, 1, 0, 1]).
persona("Persona6", "Apellido6", [1, 1, 0, 0, 1, 1, 0, 0, 1, 1]).
persona("Persona7", "Apellido7", [0, 0, 1, 1, 0, 0, 1, 1, 0, 0]).
persona("Persona8", "Apellido8", [1, 1, 1, 0, 0, 0, 1, 1, 1, 0]).
persona("Persona9", "Apellido9", [0, 0, 0, 1, 1, 1, 0, 0, 0, 1]).
persona("Persona10", "Apellido10", [1, 0, 0, 1, 1, 0, 0, 1, 1, 0]).

% Predicado para calcular el porcentaje de exactitud entre dos cromosomas
porcentaje_exactitud(CromosomaSujeto, CromosomaCandidato, Porcentaje) :-
    length(CromosomaSujeto, Longitud),
    exactitud(CromosomaSujeto, CromosomaCandidato, 0, Longitud, Exactitud),
    Porcentaje is (Exactitud / Longitud) * 100.

exactitud(_, _, Acumulador, 0, Acumulador).  % Caso base
exactitud([X|RestoSujeto], [X|RestoCandidato], Acumulador, N, Exactitud) :-
    N > 0,
    NuevoAcumulador is Acumulador + 1,
    N1 is N - 1,
    exactitud(RestoSujeto, RestoCandidato, NuevoAcumulador, N1, Exactitud).
exactitud([_|RestoSujeto], [_|RestoCandidato], Acumulador, N, Exactitud) :-
    N > 0,
    N1 is N - 1,
    exactitud(RestoSujeto, RestoCandidato, Acumulador, N1, Exactitud).

% Encontrar el sujeto más parecido a una muestra
sujeto_mas_parecido(Muestra, Personas, SujetoParecido, PorcentajeParecido) :-
    sujeto_parecido_aux(Muestra, Personas, _, -1, SujetoParecido, PorcentajeParecido).

sujeto_parecido_aux(_, [], SujetoParecido, PorcentajeParecido, SujetoParecido, PorcentajeParecido).  % Caso base
sujeto_parecido_aux(Muestra, [Persona|RestoPersonas], _, PorcentajeActual, SujetoParecido, PorcentajeParecido) :-
    porcentaje_exactitud(Muestra, Persona, Porcentaje),
    Porcentaje >= PorcentajeActual,
    sujeto_parecido_aux(Muestra, RestoPersonas, Persona, Porcentaje, SujetoParecido, PorcentajeParecido).
sujeto_parecido_aux(Muestra, [Persona|RestoPersonas], SujetoActual, PorcentajeActual, SujetoParecido, PorcentajeParecido) :-
    porcentaje_exactitud(Muestra, Persona, Porcentaje),
    Porcentaje < PorcentajeActual,
    sujeto_parecido_aux(Muestra, RestoPersonas, SujetoActual, PorcentajeActual, SujetoParecido, PorcentajeParecido).
% Muestra de cromosoma
muestra([0, 0, 0, 1, 1, 1, 0, 0, 0, 1]).

% Encontrar el sujeto más parecido a la muestra
encontrar_sujeto_parecido(SujetoParecido, PorcentajeParecido) :-
    muestra(Muestra),
    findall((Nombre, Apellido, Porcentaje), (persona(Nombre, Apellido, Cromosoma), porcentaje_exactitud(Muestra, Cromosoma, Porcentaje)), ListaParecidos),
    maximo_parecido(ListaParecidos, SujetoParecido, PorcentajeParecido).

% Encontrar el máximo porcentaje de parecido en la lista
maximo_parecido([(Nombre, _, Porcentaje)], Nombre, Porcentaje).
maximo_parecido([(Nombre, _, Porcentaje)|Resto], SujetoParecido, PorcentajeParecido) :-
    maximo_parecido(Resto, OtroNombre, OtroPorcentaje),
    (Porcentaje >= OtroPorcentaje ->
        SujetoParecido = Nombre,
        PorcentajeParecido = Porcentaje
    ;
        SujetoParecido = OtroNombre,
        PorcentajeParecido = OtroPorcentaje
    ).

% Consultar
%?- encontrar_sujeto_parecido(Sujeto, Porcentaje).
