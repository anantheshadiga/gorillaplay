### Gorilla Play

Test application to play around with the gorilla web server

## Prerequisite

1. Install dep
2. Fetch the dependency
```
dep ensure -v
```
3. Build 

```
cd gorilla
go build .
```

## Running the app

```
cd gorilla
./gorilla
```

## Play

```bash
➜  ~ curl  'http://localhost:4242/hang'

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

```


```bash
➜  ~ curl  'http://localhost:4242/jump'

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

```
