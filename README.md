<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/ALTA-BE7-HAFIDH-IRSYAD-K/event-planner">
    <img src="https://github.com/ALTA-BE7-HAFIDH-IRSYAD-K/event-planner/blob/main/bg-go.png" alt="Logo" height="80">
  </a>

<h3 align="center">EVENT PLANNING</h3>
  <p align="center">
    a Golang REST API with Echo and GORM
    <br />
    <br />
    <a href="https://github.com/ALTA-BE7-HAFIDH-IRSYAD-K/event-planner/issues">Report Bug</a>
    ·
    <a href="https://github.com/ALTA-BE7-HAFIDH-IRSYAD-K/event-planner/issues">Request Feature</a>
    .
    <a href=""><b>Documentation</b></a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
    * [Built With](#built-with)
* [Getting Started](#getting-started)
    * [Prerequisites](#prerequisites)
    * [Installation](#installation)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [License](#license)
* [Contact](#contact)
* [Acknowledgements](#acknowledgements)



<!-- ABOUT THE PROJECT -->
## About The Project

### Folder
Standard foldering that we use
* configs := Initial configs
* Repository := crud to database
* Service := Business logic
* Delivery := this folder contains handler, routes, middleware and response
* Driver := connection to database
* Entity := Model struct

### Feature
* Register & Login
* CRUD User
* CRUD Event
* Create and Read Event Participant
* Create and Read Comment on Event

### Built With
List of technology that we use
* [Golang as Language](https://golang.org/)
* [Echo as Framework](https://echo.labstack.com/)
* [Golang Validation as Validator Field](https://github.com/go-playground/validator)
* [JWT Auth as Autentication](https://github.com/dgrijalva/jwt-go)
* [Gorm as ORM](https://gorm.io/index.html)
* [MySql as Database](https://www.mysql.com/)


<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.

### Installation

1. Clone the repo (in Folder htdocs)
```sh
git clone https://github.com/ALTA-BE7-HAFIDH-IRSYAD-K/event-planner
```
2. Install module with get
```sh
go get ./...
go mod tidy
```
3. Run
```sh
go run main.go
```
4. Create .env file
```sh
source .env
```
5. Access via url
```JS
localhost:8081
```


<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/ALTA-BE7-HAFIDH-IRSYAD-K/event-planner/issues) for a list of proposed features (and known issues).


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


<!-- CONTACT -->    
## Contact

- Hafidh Irsyad K - [@hafidhIrsyad](https://www.linkedin.com/in/hafidhirsyad/)

- Abbi Yudha - [@abbiyudha](https://www.linkedin.com/in/abbi-yudha-370397165/)


---

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-blue)](https://opensource.org/licenses/MIT)

- **[MIT license](https://opensource.org/licenses/MIT)**
- Created by <a href="https://www.linkedin.com/in/hafidhirsyad/" target="_blank">Hafidh Irsyad Khairuddin</a> and <a href="https://www.linkedin.com/in/abbi-yudha-370397165/" target="_blank">Abbi Yudha</a>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=flat-square
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=flat-square
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=flat-square
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/ubaidillahhf/alterra-store
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/badge/License-MIT-blue
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/othneildrew
[product-screenshot]: images/screenshot.png
