# SmallPoint
[![Build Status](https://api.travis-ci.org/Cloud-Foundations/SmallPoint.svg?branch=master)](https://travis-ci.org/Cloud-Foundations/SmallPoint)


Smallpoint is a LDAP group management system, it is developed to help with different operations based on LDAP database. First, super administrators can create and delete LDAP groups, then all users can ask for access to any LDAP group, and managers of the requested group can approve or decline such request. What's more, users can list all groups they belong to as well as all groups in the LDAP database. And this system is easy to use, configure and administer.


## Getting Start
You can build it from source. The rpm package contain the binary.

### Prerequisites
* go >= 1.10.1

### Building
1. make deps
2. make

This will leave you with a binary: smallpoint.

### Running
You will need to create a new valid config file. And run the binary file yourself.


## Contributions
All contributions must be unencumbered. It is the responsibility of
the contributor to ensure compliance with all laws, copyrights,
patents and contracts.


## LICENSE

Copyright 2016 Symantec Corporation.
Copyright 2019 cloud-foundations.org

Licensed under the Apache License, Version 2.0 (the “License”); you
may not use this file except in compliance with the License.

You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0 Unless required by
applicable law or agreed to in writing, software distributed under the
License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for
the specific language governing permissions and limitations under the
License.

