detect :
    varables  ^[a-zA-Z_][a-zA-Z_0-9]*$
    z-number  ^[0-9]+$
    q-number  ^[0-9]+\.[0-9]+$
    strings   ^".*"$

    Operators:
        <   \<$
        >   \>$
        =   \=$
        <=  \<=$
        >=  \>=$
        ==  \==$
        !=  \!=$

        All: ^[><=!][=]{0,1}$


    Commands: 
        if  if$
        else  else$
        for for$

        All:    (^if$|^else$|^for$)

    word: 
        int 
        fload 
        string 
        bool 
        char

        All : (^int$|^float$|^string$|^bool$|^char$)

    terms: 
    (  \($

    ) :\)$  

    } :\}$ 

    {  :\{$

        All:^[(){}]{1}$

    Reserve2: 
        return 
        break

        All : (^return$|^break$)

    term: 
        ; 
        ,

        All:    ^[;,]{1}$

    