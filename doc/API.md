# Routes for event API

## Get event by ID

### Request

```http request
GET https://foo.bar/events/{id}
```

### Response

#### Success

- **Code**: 200 OK
- **Content**: JSON with event information

```json
{
  "id": 123,
  "name": "Funny Event",
  "host": "Organisation",
  "location": "North Pole",
  "start": "",
  "end": "",
  "dress-code": "Formal",
  "theme": "Christmas",
  "price": 9.99,
  "signup-link": "https://foo.bar/ilmomasiina..."
}
```

#### Errors

- **Code**: 404 Not found
  - Description: No event found with given `id`.

## Get multiple events

### Request

```http request
GET https://foo.bar/events?count=<number>
```

#### Parameters

`count` (optional): Number of events to return. If left empty, the API returns all upcoming events. 

### Response

#### Success

- **Code**: 200 OK
- **Content**: JSON a list of events

```json
[
  {
  "id": 123,
  "name": "Funny Event",
  "host": "Organisation",
  "location": "North Pole",
  "start": "",
  "end": "",
  "dress-code": "Formal",
  "theme": "Christmas",
  "price": 9.99,
  "signup-link": "https://foo.bar/ilmomasiina..."
  },  
  {
    "id": 456,
    "name": "Boring Event",
    "host": "Individual",
    "location": "Living Room",
    "start": "",
    "end": "",
    "dress-code": "Pants",
    "theme": "KPK",
    "price": null,
    "signup-link": null
  }
]
```
