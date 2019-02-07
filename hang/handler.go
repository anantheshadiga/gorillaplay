package hang

import (
	"io"
	"net/http"
)

type HangHandler struct {
}

func (s *HangHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	jump := `
	_______AAAA____j_o_a_n____AAAA________
       VVVV               VVVV        
       (__)               (__)
        \ \               / /
         \ \   \\|||//   / /
          > \   _   _   / <
 hang      > \ / \ / \ / <
  in        > \\_o_o_// <
  there...   > ( (_) ) <
              >|     |<
             / |\___/| \
             / (_____) \
             /         \
              /   o   \
               ) ___ (   
              / /   \ \  
             ( /     \ )
             ><       ><
            ///\     /\\\
            '''       '''
	`
	io.WriteString(w, jump)
}

func NewHandler() http.Handler {
	return &HangHandler{}
}
