root
|_____github.com
         |_________Your-account-name-in-github
         |                |__________Your-project-name
         |                                |________main.go
         |                                |________handlers
         |                                |________models
         |                                |__________Your-project-submodule
         |                                                  |________main.go
         |                                                  |________handlers
         |                                                  |________models
         |               
         |_________gorilla
                         |__________sessions


import "module/submodule" => import "github.com/Your-account-name-in-github/Your-project-name/Your-project-submodule"