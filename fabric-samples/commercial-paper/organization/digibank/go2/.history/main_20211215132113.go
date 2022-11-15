Marco Unisa, [15/12/21 13:21]
func main() { 
 
 idUser := os.Args[1] 
 
 //f.Create(idUser, "ERICSSON", "A1234", "COSTUMID", "DESCR", "K", "TYPEP", "SWLT", "1-1-2020", "1-1-2021", "3", "1", "ERICSSON") 
 f.QueryAll(idUser) 
 f.GetItemFromId(idUser, "A1234") 
 f.Transfer(idUser, "A1234", "TIM") 
 f.QueryAll(idUser) 
 
}