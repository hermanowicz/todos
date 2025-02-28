Simple todo app, build with:

- go chi

- sql/sqlite (for now)

---

1.
Ok lessons learned from this few lines.
Start with data model and add DB handlers for crud.

aka:
- create record
- update record
- read all records
- read a record
- delete record

2.

Then go with implementation of buisnes logic.
go is cool but this is c on steroids and i have to remember it
*This is not a python or ruby*

3.

Don't use std lib when you dont need it fiber is much more productive.
yes typs are cool and safe but f... me amount of boiler plate required for
data manipulation is painfull.
