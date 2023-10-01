%3)Defina la funci�n aplanar(L,L2). Esta funci�n recibe una lista con
% m�ltiples listas anidadas como elementos otra lista que contendr�a los
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

