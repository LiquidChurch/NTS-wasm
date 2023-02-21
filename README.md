# Project Title

2023 Night to Shine Check-in System

## Description

Volunteer and guest check-in system for Liquid Church's 2023 Night to Shine. This is a frontend interface that utilizes Liquid Church's Rock backend api to load up volunteer and guest information as well as provide the ability to check-in volunteer and guests.

## Getting Started

### Dependencies

* [GoLang 1.19](https://go.dev/)
* [Go-App](https://go-app.dev/)
* [Tailwind CLI](https://tailwindcss.com/blog/standalone-cli)
* [html5-qrcode](https://github.com/mebjas/html5-qrcode)

### Building application

* This repo creates a static WASM app. A separate static server will be required to serve the static WASM app.
* Makefile is used to build the static WASM app.
    * make web //This builds the static web files, including the CSS through tailwind CLI.
    * make app //This builds the WASM app files, needed to run the applicaion as a WASM.

## Authors

[George Tuan](mailto:george@tuan.pro)

## Version History

* v0.2.15 - Inital release for 2023 NTS

## License

This project is licensed under the MIT License - see the LICENSE.md file for details
