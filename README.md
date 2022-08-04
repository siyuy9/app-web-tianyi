# `天意`

## logo

![black logo](./frontend/public/images/logo-dark.svg)
![white logo](./frontend/public/images/logo-white.svg)

the logo is made with [`google-font-to-svg-path`](https://danmarshall.github.io/google-font-to-svg-path/)
using font
[`zen-maru-gothic`](https://fonts.adobe.com/fonts/zen-maru-gothic#licensing-section)

## frontend

[`vuejs`](https://vuejs.org/) SPA using
[`primevue`](https://www.primefaces.org/primevue/)

### theme

modified [`sakai-vue` theme](https://github.com/primefaces/sakai-vue) theme
with pieces from
[primevue official website](https://github.com/primefaces/primevue)

[`sakai-vue` demo](https://www.primefaces.org/sakai-vue/)

## backend

[`golang`](https://go.dev/) with
[`fiber`](https://docs.gofiber.io/)

### database

[`gorm`](https://gorm.io/) is used as a driver, so any of
[these databases](https://gorm.io/docs/connecting_to_the_database.html)
can be used, theoretically

only `postgresql` is configured

#### passwords

passwords are hashed with [Argon2id](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id),
as suggested in
[Password Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
by [OWASP](https://owasp.org/)

## resources

- https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html
- https://github.com/gofiber/recipes
- https://github.com/gothinkster/golang-gin-realworld-example-app
- https://github.com/ansible-semaphore/semaphore
- https://github.com/alpody/golang-fiber-realworld-example-app
- https://github.com/primefaces/sakai-vue
- https://github.com/primefaces/primevue
- https://dev.to/aryaprakasa/serving-single-page-application-in-a-single-binary-file-with-go-12ij
- https://github.com/thomasvvugt/fiber-boilerplate/tree/f3afe188347523dc693017ac4450220ef792026c
- https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go
- https://github.com/hashicorp/nomad
- https://gorm.io/
- https://docs.gofiber.io/
