%2) Defina la relaci�n subconj(S, S1), donde S y S1 son listas
% representando conjuntos, que es verdadera si S1 es subconjunto de S.

subconj([],_).
subconj([X|S1],S):-
    member(X, S),
    subconj(S1,S).
