# Introduction

This repo is used to practice creating clean code using golang. Created as hands on for my presentation at [COMPFEST](https://compfest.id) 2021.

# Requirement

- docker > 20.10.7

# How to Run

```bash
docker-compose run app
```

# How to Use

Clone or fork this repo. Then you can try to refactor on your own. Easy to follow walkthough is available. And the solution is there too if you want to cheat ;)
Follow git version to walk through refactoring process and eliminate code smell. Happy refactoring ;)

# Walkthrough

In general, this walkthough is implemented using TDD fashion. Your job is only 2: make test pass and make sure the app is running fine. There are several branch for each code smell. You should implement in in correct order.

1. Code smell: **duplicate code**.\
Checkout using `git checkout code-smell/duplicate-code`. You will find TODO. Your job is to implement those TODO.
2. (optional) Solution is at `git checkout code-smell-refactored/duplicate-code`.
3. Code smell: **long method**.\
Checkout using `git checkout code-smell/long-method`. You will find TODO. Your job is to implement those TODO.
4. (optional) Solution is at `git checkout code-smell-refactored/long-method`.
