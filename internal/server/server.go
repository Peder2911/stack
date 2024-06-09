package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/peder2911/stack/pkg/stack"
)

type Server struct {
	*http.ServeMux
	stack *stack.Stack
}

func (s *Server) push(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodPost:
			content := make([]byte, stack.ChunkSize)
			_,err := r.Body.Read(content)
			if err != nil && err != io.EOF {
				log.Printf("Error while writing: %s\n",err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			s.stack.Push(string(content))
			w.WriteHeader(http.StatusCreated)
		default: 
			w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) pop(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type","text/plain")
			content,err := s.stack.Pop()
			if err != nil {
				log.Printf("Error while reading: %s\n",err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			if len(content) == 0 {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			_,err = w.Write([]byte(content))
			if err != nil {
				log.Printf("Error while writing to connection: %s\n",err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		default: 
			w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) root(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type","text/plain")
			size,err := s.stack.Size()
			if err != nil {
				log.Printf("Error while getting size: %s\n",err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write([]byte(fmt.Sprintf("%v",size)))
		case http.MethodDelete:
			s.stack.Truncate()
			w.WriteHeader(http.StatusNoContent)
		default: 
			w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func NewServer(stack *stack.Stack) (*Server, error){
	mux := http.NewServeMux()
	server := Server{
		mux,
		stack,
	}
	mux.HandleFunc("/",server.root)
	mux.HandleFunc("/push",server.push)
	mux.HandleFunc("/pop",server.pop)
	return &server, nil
}
