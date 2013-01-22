# safekilla

Testing tool for network devices such as SafeConnect

## Use Case

This can be used to test networking applications that rely on clients using specific user agent strings.
For example, in a contrived scenario, a network router may limit access to the WAN to users that are
authenticated. Said router may allow exceptions for clients based on their user agent string. This
application would assist in testing that.

### Building

* [Install Go](http://golang.org/doc/install)

* `go get github.com/colemickens/safekilla`

* The resulting binary is in the `bin` directory in one of your GOPATHs. (If you have $GOPATH/bin on your $PATH, you can simply run `safekilla` anywhere. Otherwise, you'll need to run `$GOPATH/bin/safekilla` or something else if you have multiple $GOPATH directories.)

## Additional Options

* `--agent=USER_AGENT_TO_FAKE`

* `--url=URL_TO_PING`

* `--c=N` (causes it to run continuous, trying to ping out every N minutes)

For example, running it without any options is equivalent to running:

`./safekilla --agent=wii --url=http://www.google.com --c=1`