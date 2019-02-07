package jump

import (
	"io"
	"net/http"
)

type SiteHandler struct {
}

func (s *SiteHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
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
	io.WriteString(w, jump)
}

func NewHandler() http.Handler {
	return &SiteHandler{}
}
