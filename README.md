# ficsit-switcher

A simple redirect webserver for [smm.ficsit.app](https://smm.ficsit.app) and [cli.ficsit.app](https://cli.ficsit.app)

## Special Routing

Both domains support route matching to download a specific version of the binaries rather than auto-detecting from your
user-agent.

### SMM

```
https://smm.ficsit.app
https://smm.ficsit.app/:platform
https://smm.ficsit.app/:platform/:arch
```

|           | Windows                                      | Linux                                      |
|-----------|----------------------------------------------|--------------------------------------------|
| **amd64** | [link](https://smm.ficsit.app/windows/amd64) | [link](https://smm.ficsit.app/linux/amd64) |

### CLI

```
https://smm.ficsit.app
https://smm.ficsit.app/:platform
https://smm.ficsit.app/:platform/:arch
https://smm.ficsit.app/:platform/:arch/:packaging
```

<table>
  <tr>
    <th></th>
    <th>Windows</th>
    <th>macOS</th>
    <th>Linux</th>
  </tr>

  <tr>
    <td rowspan="4"><strong>amd64</strong></td>
    <td rowspan="4"><a href="https://smm.ficsit.app/windows/amd64/binary">binary</a></td>
    <td rowspan="8"><a href="https://smm.ficsit.app/darwin/amd64/binary">binary</a></td>
    <td><a href="https://smm.ficsit.app/linux/amd64/binary">binary</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/amd64/deb">deb</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/amd64/rpm">rpm</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/amd64/apk">apk</a></td>
  </tr>

  <tr>
    <td rowspan="4"><strong>386</strong></td>
    <td rowspan="4"><a href="https://smm.ficsit.app/windows/386/binary">binary</a></td>
    <td><a href="https://smm.ficsit.app/linux/386/binary">binary</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/386/deb">deb</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/386/rpm">rpm</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/386/apk">apk</a></td>
  </tr>

  <tr>
    <td rowspan="4"><strong>arm64</strong></td>
    <td rowspan="4"><a href="https://smm.ficsit.app/windows/arm64/binary">binary</a></td>
    <td rowspan="12"><i>n/a</i></td>
    <td><a href="https://smm.ficsit.app/linux/arm64/binary">binary</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/arm64/deb">deb</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/arm64/rpm">rpm</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/arm64/apk">apk</a></td>
  </tr>

  <tr>
    <td rowspan="4"><strong>armv7</strong></td>
    <td rowspan="4"><a href="https://smm.ficsit.app/windows/armv7/binary">binary</a></td>
    <td><a href="https://smm.ficsit.app/linux/armv7/binary">binary</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/armv7/deb">deb</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/armv7/rpm">rpm</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/armv7/apk">apk</a></td>
  </tr>

  <tr>
    <td rowspan="4"><strong>ppc64le</strong></td>
    <td rowspan="4"><i>n/a</i></td>
    <td><a href="https://smm.ficsit.app/linux/ppc64le/binary">binary</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/ppc64le/deb">deb</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/ppc64le/rpm">rpm</a></td>
  </tr>
  <tr>
    <td><a href="https://smm.ficsit.app/linux/ppc64le/apk">apk</a></td>
  </tr>
</table>