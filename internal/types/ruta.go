package types



type Ruta struct{
    id string
    title string
}

func New(id string, title string) Ruta{
    return Ruta{
        id,
        title,
    }
}
