<p align="center"><img src="https://user-images.githubusercontent.com/31022056/192372549-3f4b2e53-8b18-4a0d-ab30-49527dcd255c.png" /></p>

<p align="center">A simple forward auth provider to protect <a href="https://traefik.io/traefik/">Traefik</a> services with user authentication.</p>

<p align="center">
<a href="https://hub.docker.com/r/marvinjwendt/traefik-guardian"><img alt="Docker Cloud Build Status" src="https://img.shields.io/docker/cloud/build/marvinjwendt/traefik-guardian?style=flat-square"></a>
<a href="https://hub.docker.com/r/marvinjwendt/traefik-guardian"><img alt="Docker Image Size (tag)" src="https://img.shields.io/docker/image-size/marvinjwendt/traefik-guardian/latest?style=flat-square"></a>
<a href="https://hub.docker.com/r/marvinjwendt/traefik-guardian"><img alt="Docker Pulls" src="https://img.shields.io/docker/pulls/marvinjwendt/traefik-guardian?style=flat-square"></a>
<a href="https://github.com/MarvinJWendt/traefik-guardian/blob/main/LICENCE"><img alt="GitHub" src="https://img.shields.io/github/license/MarvinJWendt/traefik-guardian?style=flat-square"></a>
</p>

![Screenshot](https://user-images.githubusercontent.com/31022056/192390005-428ff759-8a11-4e54-ba97-1c390e4bd1ed.png)

## Features

| Feature                           | Description                                                                                               |
|-----------------------------------|-----------------------------------------------------------------------------------------------------------|
| 🧸 Easy to use                    | Dead simple to use! No config files, no external dependencies, no setup. One single command to deploy.    | 
| 🔒 Authentication                 | Authenticate users with a password.                                                                       |
| 📝 Authorization                  | Authorize users to use services behind the Traefik proxy.                                                 |
| 💙 Conforms to Traefik standards  | Traefik Guardian implements Traefik Forward Auth. It also uses the same logging format as Traefik itself. |
| 🚄 Super fast                     | Easily handles hunderts of thousands authorization checks per second.                                     |
| 🤖 Header Authorization           | Authorize requests by passing the token in a header, to make guarded API connections possible.            |

## Getting started

### Docker Compose

_Full example soon._

```yaml
  auth:
    image: marvinjwendt/traefik-guardian
    environment:
      - AUTH_HOST=auth.example.com
      - PASSWORDS=plaintext:test1234|test1337
      # - PASSWORDS=bcrypt:$$2a$$12$$/n4Bb2g0YsW6rL9d0f2VquHkhl.iSaV88FOGiu5FEYXCEPW2Sl9yy|$$2a$$12$$UoUJQcz5W5wm9A98N4GC7.X.7x398zMl6Y/T5Vjycc.gel/xBzSGm
    networks:
      - proxy # your traefik network
    labels:
      - traefik.enable=true
      - traefik.docker.network=proxy # your traefik proxy
      - traefik.http.routers.auth.entrypoints=web
      - traefik.http.routers.auth.rule=Host(`auth.example.com`) || Path(`/traefik-guardian-session-share`)
      - traefik.http.middlewares.traefik-guardian.forwardauth.address=http://auth/check # Make sure the domain is the service name
```

## Configuration

> Environment variables are used to configure Traefik Guardian.

### Environment Variables

| Variable Name      | Description                              | Default Value              | Accepted Values                                 |
|--------------------|------------------------------------------|----------------------------|-------------------------------------------------|
| `AUTH_HOST`        | The host to use.                         |                            | Any valid host (e.g.: `auth.example.com`)       |
| `PASSWORDS`        | The passwords that can be used to login. |                            | See [Password Management](#password-management) |
| `DEBUG`            | Enable debug mode.                       | `false`                    | `true`, `false`                                 |
| `LOGIN_PAGE_TITLE` | Title of the login page.                 | `Traefik Guardian - Login` | Any string.                                     |

### Password Management

> Passwords are stored in the `PASSWORDS` environment variable.

The `PASSWORDS` environment variable is a separated list of passwords, prepended with the used algorithm. The seperator is a pipe (`|`).  
Example: `plaintext:pass1|pass2|pass3`

#### Supported Algorithms

| Algorithm   | Tool to generate hash                                                                                                                                             |
|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `plaintext` | No tool needed - just plain text passwords.                                                                                                                       |
| `bcrypt`    | You can use [Cyber Chef](https://gchq.github.io/CyberChef/#recipe=Bcrypt(12)) to generate your bcrypt hash. You need to escape every `$` with another one (`$$`). |
| `md5`       | You can use [Cyber Chef](https://gchq.github.io/CyberChef/#recipe=MD5()) to generate your md5 hash.                                                               |
| `sha512`    | You can use [Cyber Chef](https://gchq.github.io/CyberChef/#recipe=SHA2('512',64,1)) to generate your md5 hash.                                                    |

more to come...

## Authorization via Header

> You can authorize requests by passing a password in a header, to make guarded API connections possible.

To authorize requests to an API, you can pass the password in the header.  
The header name is `Guardian-Password` and the value should be one of your configured passwords.

---

> [MarvinJWendt.com](https://marvinjwendt.com) &nbsp;&middot;&nbsp;
> Twitter [@MarvinJWendt](https://twitter.com/MarvinJWendt)
