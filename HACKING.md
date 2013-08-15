# Simple guide on how to hack socialwhales.

### 0. How to run

```
$ revel run github.com/benzel/socialwhales
```

### 1. Code style

* To format Go code changes, just run `scripts/format_code.sh.`.
* Use this [manual](http://google-styleguide.googlecode.com/svn/trunk/javascriptguide.xml) for JavaScript.
* Socialwhales should be placed in `$GOCODE/src/github.com/socialwhales`.
* Maximum line lenght is 120 characters.

### 2. Data definition

* We rely on schema.sql, its comments and HACKING.md. Nuff said.
