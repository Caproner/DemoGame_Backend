package proto

var (
    MSGC2SLogin = 100011
    MSGS2CLogin = 100012
    MSGC2SPlayerInfo = 110011
    MSGS2CPlayerInfo = 110012

)


type C2SLogin struct { 
    Code  interface{} 
}

type S2CLogin struct { 
    Token  interface{} 
    OpenID  interface{} 
    SessionKey  interface{} 
}

type C2SPlayerInfo struct { 
}

type S2CPlayerInfo struct { 
    Lv  interface{} 
    Exp  interface{} 
    Builds  interface{} 
    Npcs  interface{} 
    Maps  interface{} 
    UpdateTime  interface{} 
}

