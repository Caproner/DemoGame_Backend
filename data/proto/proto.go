package proto

var (
    MSGS2CErr = 100012
    MSGS2COk = 100022
    MSGS2CLogin = 110012
    MSGC2SPlayerInfo = 120011
    MSGS2CPlayerInfo = 120012
    MSGC2SLoopSync = 130001
    MSGC2SMoneySync = 130011
    MSGC2SLvSync = 130021
    MSGC2SBuildsSync = 130031
    MSGC2SNpcsSync = 130041
    MSGC2SMapsSync = 130051
    MSGC2SMoneyOpt = 140021

)


type S2CErr struct { 
    Error  interface{} `json:"Error"` 
}

type S2COk struct { 
}

type S2CLogin struct { 
    Token  interface{} `json:"Token"` 
    OpenID  interface{} `json:"OpenID"` 
    SessionKey  interface{} `json:"SessionKey"` 
}

type C2SPlayerInfo struct { 
}

type S2CPlayerInfo struct { 
    Lv  interface{} `json:"Lv"` 
    Exp  interface{} `json:"Exp"` 
    Money  interface{} `json:"Money"` 
    Builds  interface{} `json:"Builds"` 
    Npcs  interface{} `json:"Npcs"` 
    Maps  interface{} `json:"Maps"` 
    Times  interface{} `json:"Times"` 
}

type C2SLoopSync struct { 
}

type C2SMoneySync struct { 
    Money  interface{} `json:"Money"` 
}

type C2SLvSync struct { 
    Lv  interface{} `json:"Lv"` 
    Exp  interface{} `json:"Exp"` 
}

type C2SBuildsSync struct { 
    Builds  interface{} `json:"Builds"` 
}

type C2SNpcsSync struct { 
    Npcs  interface{} `json:"Npcs"` 
}

type C2SMapsSync struct { 
    Maps  interface{} `json:"Maps"` 
}

type C2SMoneyOpt struct { 
    Opt  interface{} `json:"Opt"` 
    Type  interface{} `json:"Type"` 
    Num  interface{} `json:"Num"` 
}

