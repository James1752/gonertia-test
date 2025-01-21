# gonertia-test

https://docs.gofiber.io/recipes/hello-world/

https://github.com/romsar/gonertia?tab=readme-ov-file

[Makefile](https://gist.github.com/alexedwards/3b40775846535d0014ab1ff477e4a568)

https://github.com/mehdihadeli/Go-MediatR


## Thoughts

1) Proper logging - maybe some request decorator logging
2) Config - how to get config through to the application like server port address
3) https - use a self signed cert?
4) Database and some kind of unit of work pattern for transactional
5) Go through and add documentation to all aspects of the application
6) Take a look at composability - how can I pull parts of module implementation and have it reused in other modules etc... Can we do this for any other files?
Maybe some kind of application abstraction also


Traceability - assign and id or guid or something that is traeable
Figure out some way to start modules independly - for instance only start users module
    - Currently main starts all moduls and main api - modules started in go routines



### Route Thoughts

- Admin
  - User
    - Register
    - Details

- Company
  - Session
    - Login
    - Logout
  - Protected
    - Test
