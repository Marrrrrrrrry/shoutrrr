# Generators

Generators are used to create service configurations via the command line.  
The main generator is the reflection based [Basic generator](/generators/basic) that aims to be able to generator configurations for all the core services via a set of simple questions.
Two service-specific generators are also available: the [OAuth2 generator](/generators/oauth2) for the email service and the [Telegram generator](/generators/telegram).

## Usage

```bash
$ shoutrrr generate [OPTIONS] -g <GENERATOR> <SERVICE>
```