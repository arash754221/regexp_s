package detector

import "regexp"



type Detector struct{
    Input string
}


func (d Detector) variable() bool{
    
    // varables  ^[a-zA-Z_][a-zA-Z_0-9]*$

    pattern := "^[a-zA-Z_][a-zA-Z_0-9]*$"


	ok, err := regexp.MatchString(pattern, d.Input)
    if err != nil{
        return false
    }
    return  ok

}

func (d Detector) ZNumber(){
    
    // z-number  ^[0-9]+$


    pattern := "^[0-9]+$"


	ok, err := regexp.MatchString(pattern, d.input)
    if err != nil{
        return false
    }
    return  ok

}

func (d Detector) QNumber(){
    // q-number  ^[0-9]+\.[0-9]+$
    
    pattern := "^[0-9]+\.[0-9]+$"


	ok, err := regexp.MatchString(pattern, d.input)
    if err != nil{
        return false
    }
    return  ok
}

func (d Detector) Strings(){
    // strings   ^".*"$



    pattern := "^".*"$"


	ok, err := regexp.MatchString(pattern, d.input)
    if err != nil{
        return false
    }
    return  ok
    
}

func (d Detector) Operators(){
    // Operators:  ^[><=!][=]{0,1}$


    pattern := "^[><=!][=]{0,1}$"


	ok, err := regexp.MatchString(pattern, d.input)
    if err != nil{
        return false
    }
    return  ok
    
}

func (d Detector) Commands(){
    // Commands:    (^if$|^else$|^for$)
    


    pattern := "(^if$|^else$|^for$)"


	ok, err := regexp.MatchString(pattern, d.input)
    if err != nil{
        return false
    }
    return  ok
}

func (d Detector) Reserved(){
    // Reserved :(^int$|^float$|^string$|^bool$|^char$)
    
    pattern := "(^int$|^float$|^string$|^bool$|^char$)"


	ok, err := regexp.MatchString(pattern, d.input)
    if err != nil{
        return false
    }
    return  ok
}


func (d Detector) Bracketes(){
    // Bracketes:^[(){}]{1}$
    
    pattern := "^[(){}]{1}$"


	ok, err := regexp.MatchString(pattern, d.input)
    if err != nil{
        return false
    }
    return  ok
}


func (d Detector) Reserve2(){
    // Reserve2: (^return$|^break$)
    

    pattern := "(^return$|^break$)"


	ok, err := regexp.MatchString(pattern, d.input)
    if err != nil{
        return false
    }
    return  ok
}

func (d Detector) Comma(){
    
    // Comma: ^[;,]{1}$


    pattern := "^[;,]{1}$"


	ok, err := regexp.MatchString(pattern, d.input)
    if err != nil{
        return false
    }
    return  ok
}




