# Simple guide on how to hack socialwhales.

### 0. How to run

```
$ revel run github.com/benzel/socialwhales
```

### 1. Code

* To format Go code changes, we use shell-script `scripts/format_code.sh.`.
* We use this [manual](http://google-styleguide.googlecode.com/svn/trunk/javascriptguide.xml) for JavaScript.
* This [article](https://gocardless.com/blog/building-a-large-angular-application/) shows good examples of how to organize Angular.js application.
* Socialwhales should be placed in `$GOCODE/src/github.com/socialwhales`.
* Maximum line lenght in both Go and JS is 120 characters, HTML, CSS and other files may have any width.

### 2. Data definition

* There is no ORM in social whales, so we rely on schema.sql, its comments and HACKING.md. Nuff said.
