# go-oxtel

go-oxtel is a library implementation of the Oxtel protocol in Golang, with no other dependencies. Currently this library supports version 10.1.x of the Harmonic Spectrum System. This library was created off v1.23 of the Harmonic Oxtel Protocol Specification and implements all documented commands and unsolicited tallies.

Please refer to the Harmonic Oxtel Protocol Specification documentation provided by Harmonic for more information.

> ⚠️ **Note:** Some commands, specifically relating to Audio Profiles, Audio Loudness, and External I/O were not tested against an Oxtel engine. They are implemented according to the Specification, but untested.

# Install

Install the package with:
`go get github.com/ryansavara/go-oxtel`

Import it with:
`import "github.com/ryansavara/go-oxtel"`

and use Oxtel as the package name inside the code.
