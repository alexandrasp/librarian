# Librarian

Librarian is a book recommendation web api implemented in Golang in order to improve and practice Go technical skills.
# Instructions

  - Make sure you have the latest version of Docker Desktop installed
  - In this repo's root dir, run this command to start the back-end app (on port 5001) and PostgreSQL database (on port 5432):

```bash

$ docker-compose up --build

```

- To test and see results from back-end you can run cURL commands as the examples below:

  Request example 1:
```bash
curl -X GET "localhost:5001/api/v1/genres"
```

```json
[{"id":1,"title":"Young Adult"},{"id":2,"title":"SciFi/Fantasy"},{"id":3,"title":"Romance"},{"id":4,"title":"Nonfiction"},{"id":5,"title":"Mystery"},{"id":6,"title":"Memoir"},{"id":7,"title":"Fiction"},{"id":8,"title":"Childrens"}]
```
Request example 2:
```bash
curl -X GET "localhost:5001/api/v1/books?limit=10"
```
```json
[{"ID":1,"Title":"Alanna Saves the Day","YearPublished":1972,"Rating":1.62,"Pages":169,"GenderId":8,"AuthorId":6},{"ID":2,"Title":"Adventures of Kaya","YearPublished":1999,"Rating":2.13,"Pages":619,"GenderId":1,"AuthorId":40},{"ID":3,"Title":"A Horrible Human with the Habits of a Monster","YearPublished":1976,"Rating":1.14,"Pages":258,"GenderId":7,"AuthorId":25},{"ID":4,"Title":"And I Said Yes","YearPublished":1954,"Rating":3.3,"Pages":183,"GenderId":7,"AuthorId":16},{"ID":5,"Title":"Ballinby Boys","YearPublished":1960,"Rating":1.88,"Pages":205,"GenderId":2,"AuthorId":4},{"ID":6,"Title":"Banana Slug and the Lost Cow","YearPublished":1983,"Rating":2.53,"Pages":527,"GenderId":8,"AuthorId":20},{"ID":7,"Title":"Banana Slug and Xyr Friends","YearPublished":1989,"Rating":3.64,"Pages":558,"GenderId":8,"AuthorId":20},{"ID":8,"Title":"Banana Slug and the Glass Half Full","YearPublished":1952,"Rating":4.51,"Pages":796,"GenderId":8,"AuthorId":17},{"ID":9,"Title":"Banana Slug and the Mossy Rock","YearPublished":2006,"Rating":4.43,"Pages":70,"GenderId":8,"AuthorId":31},{"ID":10,"Title":"Burnished Silver","YearPublished":1932,"Rating":1.2,"Pages":202,"GenderId":3,"AuthorId":30}]
```
Request example 3:
```bash
curl -X GET "localhost:5001/api/v1/books?authors=1,3&limit=10"
```
```json
[{"id":29,"title":"Nothing But Capers","yearPublished":2004,"rating":3.87,"pages":347,"genre":{"id":4,"title":"Nonfiction"},"author":{"id":1,"firstName":"Wendell","lastName":"Stackhouse"}}]
```
