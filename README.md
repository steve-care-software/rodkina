# Rodkina is a network based data focused programming language
Rodkina is a programming language that focus on manipulating data.  It contains a built-in AST ([Abstract-Syntax-Tree](https://en.wikipedia.org/wiki/Abstract_syntax_tree)) engine that enables complex data [lexing](https://en.wikipedia.org/wiki/Lexical_analysis) using a visual syntax, a [parsing](https://en.wikipedia.org/wiki/Parsing) engine that helps convert an AST into an [arbitrary](https://en.wikipedia.org/wiki/Arbitrariness) [logical structure](https://en.wikipedia.org/wiki/Logic_in_computer_science), an instruction engine that makes it easy to convert data into a list of commands, a [compiler](https://en.wikipedia.org/wiki/Compiler) that can easily verify instruction's logic and form an executable [program](https://en.wikipedia.org/wiki/Computer_program) and an [interpreter](https://en.wikipedia.org/wiki/Interpreter_(computing)) that can execute a program.

It also contains a built-in [database](https://en.wikipedia.org/wiki/Database) system that makes it easy to manage data on disk and then share it with other people using a [digital identity](https://en.wikipedia.org/wiki/Digital_identity) system that enables a [decentralized management](https://en.wikipedia.org/wiki/Decentralization) of [role-based access controls](https://en.wikipedia.org/wiki/Role-based_access_control) on these databases.  Therefore, the Rodkina users forms a decentralized [computer network](https://en.wikipedia.org/wiki/Computer_network) of databases that can be queried using the built-in [query engine](https://en.wikipedia.org/wiki/Query_language) of the language.

The databases are state-based: each database contains a [blockchain](https://en.wikipedia.org/wiki/Blockchain) that represents the previous modification of its internal data.  Every time data is added, deleted or modified on a database, a list of [atomic transactions](https://en.wikipedia.org/wiki/Atomicity_(database_systems)) are send to the database and a block of transactions is formed.  Each data modification is made using a transaction and each block of transactions represents an atomic modification of data, which represents a state of the database that can be easily navigated using built-in Rodkina's programming language syntax.

At last, the evolution of the language, the path of links between databases and the database interfaces is managed and voted on by the community, using a software written in Rodkina that use decentralized databases and a human-usable interface which form our DAO ([decentralized autonomous organization](https://en.wikipedia.org/wiki/Decentralized_autonomous_organization)).

## Naming convention
The name of the network of databases and the web of links it forms is named "WebX".  The DAO organization that we are building is named "Legit DAO".

## Rodkina

### Syntax
#### Comments

#### Error Management

#### Types
##### Numbers
###### Byte
###### Signed Integer
###### Unsigned Integer
###### Floating-point
##### Boolean
##### List
##### Set
##### Sorted Set
##### Map
##### Key Store
##### Grammar
###### Variable
The first letter of a variable must be a lower-case letter, then can be followed by any letter, lower-case or upper-case.  A variable can contain only one (1) letter.
```
// valid one-letter variable:
a

// valid variable where the second letter is upper-case:
aVariable

// valid variable where the second letter is lower-case:
variable

```
###### Value
The "value" is represented by an unsigned integer and therefore must be a number between 0 and 255 (range: [0,255]).  It must always be assigned to a variable in order to be used in the composition of tokens.

There is no test suites for values.
```
// valid assignment of the value '122' to a variable:
myVariable: 122;

```

###### Compose
The "compose" is a series of "value" and other "compose" elements exclusively.  A "compose" always contains a specific series of bytes.  Each element of its composition is repeated only once if there is no occurence following the element.  If there is an "occurence" value, the element is repeated by this amount of time.  The "occurence" is an unsigned integer.

Please note each element of a "compose" is created using a variable name, a pipe (|) and an occurence.  The pipe (|) only acts as a syntaxic separator.  If the element is only repeated once, the occurence one (1) and its pipe (|) is optional.

There is no test suites for values.
```
// compose the world "actually"
wordActually: letterA letterC letterT letterU letterA letterL|2 letterY|1;

// values, where each byte is a ASCII representation of a letter:
letterA: 97;
letterC: 99;
letterT: 116;
letterU: 117;
letterA: 97;
letterL: 108;
letterY: 108;

```
###### Test suites
###### Cardinality
###### Token
###### Everything
###### External Grammar
###### Channels
###### Root
###### Assignment
###### Execution
###### Database
###### Tests

##### Query
###### Execution
###### Database
###### Tests
###### Type Casting

##### Program
###### Execution
###### Database
###### Tests
###### Type Casting

#### Operator
##### Arithmetic
##### Relational
##### Logical
##### Bitwise Shift

#### Scope
#### Assignment
#### Condition
#### Loop
#### Function

### Architecture
#### Applications
##### Abstract-Syntax-Tree
##### Query
##### Interpreter
##### Database
##### Cryptography
#### Instances
##### Grammar
##### Query
#### Modules
#### Language

### Compiling libraries
#### Embedding
##### Web Assembly
##### C/C++
##### Python

### Management
#### Alpha
##### Source Code
##### Social Networks
#### Beta
#### Stable
##### Versioning

### Contributors

## WebX
### Social networks
### Contributors

## Legit DAO
### Phases
#### Pre-infrastructure
##### Group funds handling
##### Requesting a refund
#### Infrastructure voting
#### Infrastructure usage
#### Moving to another infrastructure
### Social networks
### Contributors
