package server

type Server interface {
	Add(int, any) error
	Get(int) (any, error)
}

type Server1 struct{}

func NewServer() Server {
	return &Server1{}
}

func (s *Server1) Add(key int, val any) (err error) {

}

func (s *Server1) Get(key int) (val any, err error) {
	return val, err
}
