# Ory/Keto Testing

> ORY Keto is a permission server that implements best practice access control mechanisms. https://www.ory.sh/keto/docs/

## Commands
### Check
- Check allowance

### Clean
- Deletes policies & roles

### List
- List policies & roles

### Seed
- Create policies & roles 

## Examples
### List
#### Policies
```json
[
  {
    "id": "p:manager:project:1951823",
    "description": "a description string",
    "subjects": [
      "s:admin",
      "s:manager:project:1951823"
    ],
    "resources": [
      "r:project:1951823"
    ],
    "actions": [
      "a:read",
      "a:edit",
      "a:delete"
    ],
    "effect": "allow",
    "conditions": null
  },
  {
    "id": "p:editor:project:1951823",
    "description": "a description string",
    "subjects": [
      "s:editor:project:1951823"
    ],
    "resources": [
      "r:project:1951823"
    ],
    "actions": [
      "a:read",
      "a:edit"
    ],
    "effect": "allow",
    "conditions": null
  },
  {
    "id": "p:reader:project:1951823",
    "description": "a description string",
    "subjects": [
      "s:reader:project:1951823"
    ],
    "resources": [
      "r:project:1951823"
    ],
    "actions": [
      "a:read"
    ],
    "effect": "allow",
    "conditions": null
  },
  {
    "id": "p:manager:project:5725100",
    "description": "a description string",
    "subjects": [
      "s:admin",
      "s:manager:project:5725100"
    ],
    "resources": [
      "r:project:5725100"
    ],
    "actions": [
      "a:read",
      "a:edit",
      "a:delete"
    ],
    "effect": "allow",
    "conditions": null
  },
  {
    "id": "p:editor:project:5725100",
    "description": "a description string",
    "subjects": [
      "s:editor:project:5725100"
    ],
    "resources": [
      "r:project:5725100"
    ],
    "actions": [
      "a:read",
      "a:edit"
    ],
    "effect": "allow",
    "conditions": null
  },
  {
    "id": "p:reader:project:5725100",
    "description": "a description string",
    "subjects": [
      "s:reader:project:5725100"
    ],
    "resources": [
      "r:project:5725100"
    ],
    "actions": [
      "a:read"
    ],
    "effect": "allow",
    "conditions": null
  }
]
```
#### Roles
```json
[
  {
    "id": "s:admin",
    "description": "a description string",
    "members": [
      "u:0"
    ]
  },
  {
    "id": "s:manager:project:1951823",
    "description": "a description string",
    "members": [
      "u:1"
    ]
  },
  {
    "id": "s:editor:project:1951823",
    "description": "a description string",
    "members": [
      "u:2"
    ]
  },
  {
    "id": "s:reader:project:1951823",
    "description": "a description string",
    "members": [
      "u:3",
      "u:5"
    ]
  },
  {
    "id": "s:manager:project:5725100",
    "description": "a description string",
    "members": [
      "u:1"
    ]
  },
  {
    "id": "s:editor:project:5725100",
    "description": "a description string",
    "members": [
      "u:4"
    ]
  },
  {
    "id": "s:reader:project:5725100",
    "description": "a description string",
    "members": [
      "u:5"
    ]
  }
]
```

### Query
```text
//****************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:read u:0} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:read u:1} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:read u:2} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:read u:3} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:read u:4} //
//****************************************************************//
403 Forbidden
//****************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:read u:5} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:edit u:0} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:edit u:1} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:edit u:2} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:edit u:3} //
//****************************************************************//
403 Forbidden
//****************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:edit u:4} //
//****************************************************************//
403 Forbidden
//****************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:edit u:5} //
//****************************************************************//
403 Forbidden
//******************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:delete u:0} //
//******************************************************************//
200 OK
//******************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:delete u:1} //
//******************************************************************//
200 OK
//******************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:delete u:2} //
//******************************************************************//
403 Forbidden
//******************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:delete u:3} //
//******************************************************************//
403 Forbidden
//******************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:delete u:4} //
//******************************************************************//
403 Forbidden
//******************************************************************//
// Check If a Request is Allowed - {r:project:1951823 a:delete u:5} //
//******************************************************************//
403 Forbidden
//****************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:read u:0} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:read u:1} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:read u:2} //
//****************************************************************//
403 Forbidden
//****************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:read u:3} //
//****************************************************************//
403 Forbidden
//****************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:read u:4} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:read u:5} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:edit u:0} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:edit u:1} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:edit u:2} //
//****************************************************************//
403 Forbidden
//****************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:edit u:3} //
//****************************************************************//
403 Forbidden
//****************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:edit u:4} //
//****************************************************************//
200 OK
//****************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:edit u:5} //
//****************************************************************//
403 Forbidden
//******************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:delete u:0} //
//******************************************************************//
200 OK
//******************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:delete u:1} //
//******************************************************************//
200 OK
//******************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:delete u:2} //
//******************************************************************//
403 Forbidden
//******************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:delete u:3} //
//******************************************************************//
403 Forbidden
//******************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:delete u:4} //
//******************************************************************//
403 Forbidden
//******************************************************************//
// Check If a Request is Allowed - {r:project:5725100 a:delete u:5} //
//******************************************************************//
403 Forbidden
```

## Setup
```
./start-psql.sh
./start-keto.sh
./scripts/migrate.sh # Only the first time
```
