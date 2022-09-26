<p align="center"><img src="https://user-images.githubusercontent.com/31022056/192372549-3f4b2e53-8b18-4a0d-ab30-49527dcd255c.png" /></p>

<p align="center">A simple forward auth provider to protect <a href="https://traefik.io/traefik/">Traefik</a> services with user authentication.</p>

<p align="center">
<a href="https://hub.docker.com/r/marvinjwendt/traefik-guardian"><img alt="Docker Cloud Build Status" src="https://img.shields.io/docker/cloud/build/marvinjwendt/traefik-guardian?style=flat-square"></a>
<a href="https://hub.docker.com/r/marvinjwendt/traefik-guardian"><img alt="Docker Image Size (tag)" src="https://img.shields.io/docker/image-size/marvinjwendt/traefik-guardian/latest?style=flat-square"></a>
<a href="https://hub.docker.com/r/marvinjwendt/traefik-guardian"><img alt="Docker Pulls" src="https://img.shields.io/docker/pulls/marvinjwendt/traefik-guardian?style=flat-square"></a>
<a href="https://github.com/MarvinJWendt/traefik-guardian/blob/main/LICENCE"><img alt="GitHub" src="https://img.shields.io/github/license/MarvinJWendt/traefik-guardian?style=flat-square"></a>
</p>

![Screenshot](https://user-images.githubusercontent.com/31022056/188281754-ebe3e6f8-b92c-4155-8683-896ad8e86c65.png)

<p align="center">
<table>
<tbody>
<td align="center">
<img width="2000" height="0" /><br>
<h1>‼️ WORK IN PROGRESS ‼️</h1><br>
<h4>This project is work in progress and <b>unstable</b>. It's adviced to NOT use it in production in its current state.</h4>
<img width="2000" height="0" />
</td>
</tbody>
</table>
</p>

## Features

| Feature                          | Description                                                             |
|----------------------------------|-------------------------------------------------------------------------|
| 🧸 Easy to use                   | Easy to use and configure.                                              | 
| 🔒 Authentication                | Authenticate users with username and password.                          |
| 📝 Authorization                 | Authorize users with user groups and permissions for separate routes.   |
| 💙 Conforms to Traefik standards | Traefik Auth Provider uses the same logging format as Traefik is using. |
| 🚄 Super fast                    | Easily handles hunderts of thousands authorization checks per second.   |

## Getting started

WORK IN PROGRESS

## Configuration

> Environment variables are used to configure the application itself

### Environment Variables

| Variable Name | Description              | Default Value | Accepted Values   |
|---------------|--------------------------|---------------|-------------------|
| `Debug`       | Enable debug mode.       | `false`       | `true`, `false`   |
