ignite
=======

An easy way to initialize application and eliminates the need to use the `func init()`. But retains the flexibility of being able to change the default settings using a custom `func init()` or by config files.

Contains factories for the main libraries of different domains: log, cloud computing, event data, cache, cloud monitoring, search engine, database, http client/server, dependency injection, pub/sub, rpc, data query, ftp client, kubernetes client , web framework, messaging system client, managing goroutines, software bus framework, command-line interface, etc.

To use just follow the examples contained in each package.

Installation
------------

	go get -u github.com/americanas-go/ignite

Contributing
--------
Every help is always welcome. Feel free do throw us a pull request, we'll do our best to check it out as soon as possible. But before that, let us establish some guidelines:

1. This is an open source project so please do not add any proprietary code or infringe any copyright of any sort.
2. Avoid unnecessary dependencies or messing up go.mod file.
3. Be aware of golang coding style. Use a lint to help you out.
4. Add tests to cover your contribution.
5. Add [godoc](https://elliotchance.medium.com/godoc-tips-tricks-cda6571549b) to your code. 
6. Use meaningful [messages](https://medium.com/@menuka/writing-meaningful-git-commit-messages-a62756b65c81) to your commits.
7. Use [pull requests](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/about-pull-requests).
8. At last, but also important, be kind and polite with the community.

Any submitted issue which disrespect one or more guidelines above, will be discarded and closed.


<hr>

Released under the [MIT License](LICENSE).
