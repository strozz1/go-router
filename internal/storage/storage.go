package storage

type Storage[T any] interface{
    Get(string) T
    Save(T)
    Update(T)
    Delete(T)
}
