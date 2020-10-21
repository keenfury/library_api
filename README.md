# Library API

Library API give you the abilty to interface with our catalog of books via restFUL requests.

All of our books are found in this format:

### Book
- id integer [required]
- author string [required; max len: 100]
- imageLink string [optional; max len: 250]
- language string [optional; max len: 50]
- link string [optional; max len: 250]
- pages int [required]
- title string [required; max len: 250]
- year int [required]

### Endpoints

#### List
GET: v1/book/list\
Response: []Book\
Code: 200

#### Get
GET: v1/book/{id}\
Response: Book\
Code: 200

#### Add
POST: v1/book\
Body: json\
{\
    "author": "",\
    "pages": 100,\
    "title": "",\
    "year": 1990\
}\
Response: json of book\
Code: 200

#### Update
PUT: v1/book\
Body: json\
{\
    "id": 1,\
    ...\
} (id is required; add only key/value pairs when needed)\
Code: 204

#### Delete
DELETE: v1/book/{id}\
Code: 204

### Running locally

- Create director: github.com/keenfury (for imports to work correctly) and clone.
- Build and run from cmd/library_api, it using go mod so it should get all the required modules.
- Local port is set by default to 12572 or can be set with env. var of: APP_PORT or set via command line argument flag of: restport.
- A sample set of books can be found in data/books.json, these are loaded in memory to provide a data set to interact with, this is done with file internal/v1/book/data_file.go.
- A sample of accessing via Postgresql can be done, see internal/v1/book/data.go, to do this, these steps will need to be done:
    - comment/uncomment lines 16 & 17 in internal/v1/common.go to use the correct struct (interface)
    - schema is found in data/schema.sql
    - provide env. vars for PG_USER, PG_PASS, PG_DB to your local postgres instance