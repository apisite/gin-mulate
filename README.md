# gin-mulate
> [gin](https://github.com/gin-gonic/gin) bindings for [mulate](https://github.com/apisite/mulate) library

[![GoCard][gc1]][gc2]
 [![GitHub Release][gr1]][gr2]
 [![GitHub code size in bytes][sz]]()
 [![GitHub license][gl1]][gl2]

[gc1]: https://goreportcard.com/badge/apisite/gin-mulate
[gc2]: https://goreportcard.com/report/github.com/apisite/gin-mulate
[gr1]: https://img.shields.io/github/release/apisite/gin-mulate.svg
[gr2]: https://github.com/apisite/gin-mulate/releases
[sz]: https://img.shields.io/github/languages/code-size/apisite/gin-mulate.svg
[gl1]: https://img.shields.io/github/license/apisite/gin-mulate.svg
[gl2]: LICENSE


## Usage

```go
    mlt, _ := mulate.New(cfg.Template, log)
    mlt.DisableCache(true)

    allFuncs := make(template.FuncMap, 0)
    err = mlt.LoadTemplates(allFuncs)
    if err != nil {
        log.Fatal(err)
    }

    router := gin.Default()
    templates := ginmulate.New(mlt, log)
    templates.Route("", router)
```

See also: [sample](sample/)

## TODO

* [ ] rpc
* [ ] test coverage
* [ ] docs

## See also

* https://github.com/gin-contrib/multitemplate

## License

The MIT License (MIT), see [LICENSE](LICENSE).

Copyright (c) 2018 Aleksei Kovrizhkin <lekovr+apisite@gmail.com>

