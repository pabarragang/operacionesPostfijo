package main

import ("fmt"
  "bufio"
  "os"
  "strings"
  "strconv")

type Arbol struct{
  izq *Arbol
  valor string
  der *Arbol
}

type Pila struct{
  dato *Arbol
  sig *Pila
}

func push(p *Pila, a *Arbol)*Pila{
  if p!=nil{
    nuevo :=&Pila{a, p}
    p=nuevo;
  }else{
    p=&Pila{a,nil}
  }
  return p
}

func pop(p *Pila) (*Arbol, *Pila){
  salida:=p
  p=p.sig
  return salida.dato, p
}


func expToArbol(p *Pila,s []string)*Pila{
  cant:=0
  var arbAux *Arbol
  for _,dato:=range s{
     var a=&Arbol{valor:dato}
     if(dato!="+" && dato!="-" && dato!="*" && dato!="/"){
       p=push(p,a)
       cant++
     }else{
       if cant<2{
         fmt.Println("Expresion no valida")
         return nil
       }else{
         arbAux,p =pop(p)
         a.der=arbAux
         arbAux,p =pop(p)
         a.izq=arbAux
         p=push(p,a)
         cant--
       }
     }
  }
  if(cant!=1){fmt.Println("Expresion no valida")
    return nil
  }
  return p
}



func recorrerInorden(t *Arbol){
  if t==nil{
    return
  }
  recorrerInorden(t.izq)
  fmt.Print(t.valor, " ")
  recorrerInorden(t.der)
}

 func calcular(t *Arbol) int{
   if t!=nil{
     switch  t.valor{
     case "*":
       return calcular(t.izq)*calcular(t.der)
     case "/":
       return calcular(t.izq)/calcular(t.der)
     case "+":
       return calcular(t.izq)+calcular(t.der)
     case "-":
       return calcular(t.izq)-calcular(t.der)
     default:
       i,_:=strconv.Atoi(t.valor)
       return i
     }
   }
   return -1
 }

func evaluar(t *Arbol){
  recorrerInorden(t)
  fmt.Println("=", calcular(t))
}

func comprobar(t *Arbol, er *string) int{
  if t!=nil{
    _,e:=strconv.Atoi(t.valor)
    if t.der==nil && t.izq==nil{
      if e!=nil{
        *er+=t.valor+" debe ser un valor numerico\n"
        return 1
      }
      return 0
    }else if t.der!=nil && t.izq!=nil{
      if e!=nil{
        if t.valor =="*" ||t.valor =="/"||t.valor =="+"||t.valor =="-"{
          return comprobar(t.der,er)+comprobar(t.izq,er)
        }
          *er+=t.valor+" operacion no valida\n"
          return comprobar(t.der,er)+comprobar(t.izq,er)+1
      }
      *er+=t.valor+" debe ser una operacion, no un numero\n"
      return comprobar(t.der,er)+comprobar(t.izq,er)+1
    }else{
      return 1
    }
  }
  return 0
}

func main(){
  var  p *Pila
  var a *Arbol
  var e string
  sc := bufio.NewScanner(os.Stdin)
  fmt.Println("Escriba la expresion postfija(cada termino separado por espacio)")
  sc.Scan()
  op:=sc.Text()
  p=expToArbol(p,strings.Split(op," "))
  if(p!=nil){
    a,p=pop(p)
    b:=comprobar(a, &e)
    if(b==0){
      evaluar(a)
    }else{
      fmt.Print("Hubo ", -b," errores:\n",e)
    }
  }
}
