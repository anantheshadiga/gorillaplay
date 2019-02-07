package hang

import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

type HangHandler struct {
	logger *zap.Logger
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
	s.logger.Info("Request to hang")
	io.WriteString(w, jump)
}

func NewHandler(logger *zap.Logger) http.Handler {
	return &HangHandler{logger}
}
