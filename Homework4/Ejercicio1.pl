%1)Defina sumlist(L, S) que es verdadero si S es la suma de los
% elementos de L.
sumlist([],0).
sumlist([Cabeza|Cola],S):-
    sumlist(Cola, SRest),
    S is Cabeza + SRest.
