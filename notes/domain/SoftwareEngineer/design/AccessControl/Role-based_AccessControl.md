# Role-based Access Control

Good for separation of duties (SoD) requirements

Allows it to implement DAC or MAC.
DAC with groups (e.g., as implemented in POSIX file systems) can emulate RBAC. MAC can simulate RBAC if the role graph is restricted to a tree rather than a partially ordered set.

The Bell-LaPadula (BLP) model was synonymous with MAC and file system permissions were synonymous with DAC.

## Term

* S = Subject = A person or automated agent
* R = Role = Job function or title which defines an authority level
* P = Permissions = An approval of a mode of access to a resource
* SE = Session = A mapping involving S, R and/or P
* SA = Subject Assignment
* PA = Permission Assignment
* RH = Partially ordered Role Hierarchy. RH can also be written: ≥ (The notation: x ≥ y means that x inherits the permissions of y.)

Description:

* A subject can have multiple roles.
* A role can have multiple subjects.
* A role can have many permissions.
* A permission can be assigned to many roles.
* An operation can be assigned to many permissions.
* A permission can be assigned to many operations.
