welcome to REGEX Practice
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





test's:
if ( 3 > 2 ) {

println ( " 3 is greater than 2 " )

}

for ( int i = 0 ; i < 10 ; i++) { println(i); }

while (true) { println("infinite loop"); }

if ( true ) { println("true"); } else { println("false"); }

switch (day) { case 1: println("Monday"); break; case 2: println("Tuesday"); break; default: println("Invalid day"); }

if (true) { println("true"); } else if (false) { println("false"); } else { println("neither"); }

function main() { return 0; }