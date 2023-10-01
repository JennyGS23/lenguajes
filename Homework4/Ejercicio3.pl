%3)Defina la función aplanar(L,L2). Esta función recibe una lista con
% múltiples listas anidadas como elementos otra lista que contendría los
% mismos elementos de manera lineal (sin listas).

aplanar([],[]).
aplanar([Lista|Resto],Aplanada):-
    is_list(Lista),
    aplanar(Lista, AplanadaCabeza),
    aplanar(Resto, AplanadaResto),
    append(AplanadaCabeza, AplanadaResto, Aplanada).
aplanar([Elemento|Resto],[Elemento| AplanadaResto]):-
    not(is_list(Elemento)),
    aplanar(Resto,AplanadaResto).

