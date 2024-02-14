package storage

import "github.com/strozz1/pinkbikers-web/internal/types"


type MockDB[T any] struct{
}

func (s *MockDB[T]) Get(id string) types.Ruta{
   return types.New(id,"Hola q ase")
}
func (s *MockDB[T]) Save(ruta types.Ruta){

}
func (s *MockDB[T]) Update(ruta T) {
}    

func (s *MockDB[T]) Delete(ruta T){

}
