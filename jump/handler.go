package jump

import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

type JumpHandler struct {
	logger *zap.Logger
}

func (s *JumpHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	jump := `
	   ||
         ||
        _;|
       /__3
      / /||
     / / // .--.
     \ \// / (OO)
      \//  |( _ )
      // \__/'-'\__
     // \__      _ \
 _.-'/    | ._._.|\ \
(_.-'     |      \ \ \
   .-._   /    o ) / /
  /_ \ \ /   \__/ / /
    \ \_/   / /  E_/
     \     / /
      '-._/-' 
   `
	s.logger.Info("Request to jump")
	io.WriteString(w, jump)
}

func NewHandler(logger *zap.Logger) http.Handler {
	return &JumpHandler{logger}
}
