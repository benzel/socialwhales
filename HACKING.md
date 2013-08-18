# Simple guide on how to hack socialwhales.

### 0. How to run

```
$ revel run github.com/benzel/socialwhales
```

### 1. Code

* To format Go code, we use shell-script `scripts/format_code.sh.`.
* We use this [manual](http://google-styleguide.googlecode.com/svn/trunk/javascriptguide.xml) for JavaScript.
* This [article](https://gocardless.com/blog/building-a-large-angular-application/) shows good examples of how to organize Angular.js application.
* Socialwhales should be placed in `$GOCODE/src/github.com/socialwhales`.
* Maximum line lenght in both Go and JS is 120 characters, HTML, CSS and other files may have any width.

### 2. Data definition

* There is no ORM in social whales, so we rely on schema.sql, its comments and HACKING.md. Nuff said.


### 3. Backend application structure

* Directory structure:

```
├── app                       # Root of Go backend application.
│   ├── controllers           # REST API implementations.
│   ├── init.go               # This module contains entry point and other modules initializers.
│   ├── models                # Application-level data definitions.
│   ├── routes                # This folder contains generated code which is responsible for url routing.
│   ├── services              # Difiirent implementations of database services.
│   │   ├── storage           # Definition of abstract database service.
│   │   ├── mem               # Simple service base on go maps.
│   │   └── pq                # Service based on gorp and pq libaries.
│   ├── tmp                   # 
│   ├── utils                 #
│   └── views                 #
│       └── App               #
│           └── Index.html    #
├── conf                      #
├── docs                      #
├── goPackages.txt            #
├── messages                  # We dont use this folder.
├── public                    #
│   ├── app                   #
│   ├── css                   #
│   ├── errors                #
│   ├── img                   #
│   └── js                    #
├── schema.sql                #
├── scripts                   #
└── tests                     #
```


### 4. Frontend application structure

