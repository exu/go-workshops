## Mangos SP message system

> Package mangos is an implementation in pure Go of the SP ("Scalability Protocols") messaging system. This makes heavy use of go channels, internally, but it can operate on systems that lack support for cgo.

> The reference implementation of the SP protocols is available as nanomsg™; there is also an effort to implement an improved and more capable version of nanomsg called NNG™.

> The design is intended to make it easy to add new transports with almost trivial effort, as well as new topologies ("protocols" in SP terminology.)

> At present, all of the Req/Rep, Pub/Sub, Pair, Bus, Push/Pull, and Surveyor/Respondent patterns are supported.

> Additionally, there is an experimental new pattern called STAR available. This pattern is like Bus, except that the messages are delivered not just to immediate peers, but to all members of the topology. Developers must be careful not to create cycles in their network when using this pattern, otherwise infinite loops can occur.

> Supported transports include TCP, inproc, IPC, Websocket, Websocket/TLS and TLS. Use addresses of the form "tls+tcp://:" to access TLS. Note that ipc:// is not supported on Windows (by either this or the reference implementation.) Forcing the local TCP port in Dial is not supported yet (this is rarely useful).