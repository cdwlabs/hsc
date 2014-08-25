package main

// The commit that was compiled.  Populated by the compiler.
var GitCommit string

// The main version number.
const Version = "0.0.1"

// Borrowed this idea from HashiCorp guys.  A pre-release marker for the version.  If this is "" (empty string)
// then it means that it is a final release.  Otherwise, this is a pre-release such as "dev", "beta", "rc1", etc.
const VersionPrerelease = "dev"

// Use this to determine if an upgrade breaks backwards compatibility.  If this is "" (empty string)
// then it means that this version is backwards compatible.  Otherwise, this is the Version that breaks.
const BreaksCompatibilityWithVersion = ""
