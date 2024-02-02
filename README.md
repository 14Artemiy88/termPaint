# termPaint
Drawing in the terminal

## Features
- Drawing by cursor
- Choosing any symbol from the keyboard
- Choosing the color of RGB
- Save image
- Load image

## Examples

![term_paint.gif](screenshots/term_paint.gif)

![2023-07-09_12-02.png](screenshots/2023-07-09_12-02.png)

## Menus

![menu.png](screenshots/menu.png)   ![helpMenu.png](screenshots/helpMenu.png)   ![file.png](screenshots/file.png)


## Requirements
```agsl
go 1.20
```

## Installation

### Docker
```bash
docker run -ti artemiy88/termpaint
```

### Go

```bash
go install github.com/14Artemiy88/termPaint@latest
```
Make sure the Go executables directory ($GOPATH/bin) is added to your PATH environment variable. You can achieve this using the following command:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```
#### Usage
```bash
termPaint
```

### Snap
```bash
sudo snap install --beta termpaint
```
#### Usage
```bash
termpaint
```