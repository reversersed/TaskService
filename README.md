## Table of Contents
- [Introduction](#introduction)
- [Startup](#startup)
  - [Template](#startup-template)
  
<a name="introduction"></a>
## Introduction

This is REST API service with tasks CRUD. Made as test task for verba group.
Project was made in DDD architecture (Domain Driven Design) as pet project.
Made by © reversersed

<a name="startup"></a>
## Startup

Project has a makefile to ease startup.
Available makefile commands
- `run - run the service`
- `install (alias - i) - install project dependencies`
- `gen - generate documentation and mock files`
- `upgrade - upgrade dependencies`
- `clean - clean mod files`
- `start - start docker compose with rebuild`
- `up - start docker compose without rebuild`
- `stop - stop docker container`
  
You can display it again with typing `make` in console in project root folder.

<a name="startup-template"></a>
### Template

If you have make installed, you can:
- run `make i` to install all project dependencies, and then run `make run` to regenerate documentation and rebuild the project
- or just run `make start` to build project without installing dependencies and build existing pre-generated files

Neither, if you dont have make installed, you can build and run the project with `docker compose --env-file ./config/.env up --build -d` command

<h4>Make sure you have Docker started before building and starting project</h4>

After project started you can open `localhost:9000/swagger/index.html` (with default confing) in browser to see OpenAPI swagger documentation.<br>
Here's the screenshot of currently swagger:<br><br>
![swagger screenshot](https://github.com/user-attachments/assets/f72bb627-b985-4d86-8d8e-6060f761d8ef)
