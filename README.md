# xmail
xmail is a Go tool that uses [haccer/available](https://github.com/haccer/available) to detect which email addresses have domains which are able to be registered.

## Install:

```
go install github.com/haccer/xmail@latest
```

or

```
git clone https://github.com/haccer/xmail
cd xmail
go install
```

## Usage:

```
cat emails.txt | xmail
```

or

```
xmail -w emails.txt
```

With JSON:

```
% cat emails.txt| xmail --json
[
  {
    "domain": "doalkdjfaklsjdfk.com"
  },
  {
    "domain": "doesntexisthahshhsh.com"
  }
]
```

![xmail](https://raw.githubusercontent.com/haccer/xmail/main/image.png)

