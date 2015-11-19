restless
========

**THIS IS A WIP, USE UNDER YOUR OWN RISK!**

restless is a magic mocked API generator from json schemas. Let me explain myself:

1. you have some url routes that you want to mock.
2. you know the schema of what's going to be returned.

For the point 1, if you want your handlerhere: `/api/v1/people/1`, then you just need to do this:

    $ mkdir -p api/v1/people/1

Now, if in that folder you have an `schema.json`, restless will serve a mocked object for you. Let's see it with an example:

Example
-------

I am pretty sure you know [this example](http://json-schema.org/examples.html). We are going to be serving a url `/people/1` with it:

    $ mkdir -p people/1
    $ cat << EOF > people/1/schema.json
    {
      "title": "Example Schema",
      "type": "object",
      "properties": {
        "firstName": {
          "type": "string"
        },
        "lastName": {
          "type": "string"
        },
        "age": {
          "description": "Age in years",
          "type": "integer",
          "minimum": 0
        }
      },
      "required": ["firstName", "lastName"]
    }
    EOF

Now you can run restless:

    $ restless
    2015/11/20 00:14:58 GET endpoint registered: /people/1
    2015/11/20 00:14:58 Listening :8080

And query your shinny API!

    $ curl -i localhost:8080/people/1
    HTTP/1.1 200 OK
    Content-Type: application/json; charset=UTF-8
    Date: Thu, 19 Nov 2015 23:15:56 GMT
    Content-Length: 69

    {"age":545,"firstName":"Mella et","lastName":"Hi calidum ipsa coram"}%

Developing
----------

### Dependencies

We are using [glide](https://github.com/Masterminds/glide) for the dependency management, please do so.

### Running

If you are running this in development mode I will recommend you to use [gin](https://github.com/codegangsta/gin):

    $ gin run

TODO
----

- Organize the code properly and split the huge method.
- Support more verbs than GET.
- Support more types, a difficult one will be `array`.
- Support more properties from json schema, ex: `minimum`.
