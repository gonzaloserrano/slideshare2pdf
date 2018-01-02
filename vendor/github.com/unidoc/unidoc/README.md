# UniDoc

[UniDoc](http://unidoc.io) is a fast and powerful open source PDF library for Go (golang). The library is written and supported by the owners of the [FoxyUtils.com](https://foxyutils.com) website, where the library is used to power many of the PDF services offered. 

[![wercker status](https://app.wercker.com/status/22b50db125a6d376080f3f0c80d085fa/s/master "wercker status")](https://app.wercker.com/project/bykey/22b50db125a6d376080f3f0c80d085fa)
[![GoDoc](https://godoc.org/github.com/unidoc/unidoc?status.svg)](https://godoc.org/github.com/unidoc/unidoc)

# Version 2

Version 2.0.0 has been released. Version 2 represents a major improvement over version 1 with capabilities for modifying
and generating PDF contents. The library has been split up into three major packages and a
few smaller ones. The **core** package contains core PDF file parsing functionality and
primitive objects, whereas the **model** subpackage provides a higher level interface to the PDF.
The **creator** package provides a convenient interface for creating image and text based PDF files
and reports.

See the release announcement: [https://unidoc.io/news/unidoc-v2-released](https://unidoc.io/news/unidoc-v2-released)


## Installation
~~~
go get github.com/unidoc/unidoc/...
~~~

## Examples

Multiple examples are provided in our example repository.
Many features for processing PDF files with [documented examples](https://unidoc.io/examples) on our website.

Contact us if you need any specific examples.

## Vendoring
For reliability, we recommend using specific versions and the vendoring capability of golang.
Check out the Releases section to see the tagged releases.

## Copying/License

UniDoc is licensed as [AGPL][agpl] software (with extra terms as specified in our license).

AGPL is a free / open source software license.

This doesn't mean the software is gratis!

Buying a license is mandatory as soon as you develop activities
distributing the UniDoc software inside your product or deploying it on a network
without disclosing the source code of your own applications under the AGPL license.
These activities include:

 * offering services as an application service provider or over-network application programming interface (API)
 * creating/manipulating documents for users in a web/server/cloud application
 * shipping UniDoc with a closed source product

Please see [pricing](http://unidoc.io/pricing) to purchase a commercial license or contact sales at sales@unidoc.io for more info.

## Contributing

Contributors need to approve the [Contributor License Agreement](https://docs.google.com/a/owlglobal.io/forms/d/1PfTjEAi67-x0JOTU45SDonJnWy1fWB_J1aopGss34bY/viewform) before any code will be reviewed. Preferably add a test case to make sure there is no regression and that the new behaviour is as expected.

## Support and consulting

Please email us at support@unidoc.io for any queries.

Technical support is included with a purchase of a license, as listed on our [pricing](http://unidoc.io/pricing) page.

If you have any specific tasks that need to be done, we offer consulting in certain cases.
Please contact us with a brief summary of what you need and we will get back to you with a quote, if appropriate.

## Stay up to date

* Follow us on [twitter](https://twitter.com/unidoclib)
* Sign-up for our [newsletter](http://eepurl.com/b9Idt9)

[agpl]: LICENSE.md
[contributing]: CONTRIBUTING.md
