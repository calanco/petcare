# Petcare

Simple RESTful API server to store and retrieve information of your pets.

## Usage

## Pet structure
  ```yaml
  Name    Name    `json:"name"`
  Species Species `json:"species"`
  Breed   Breed   `json:"breed,omitempty"`
  Size    Size    `json:"size,omitempty"`
  Date    Date    `json:"date,omitempty"`
  Weight  Weight  `json:"weight,omitempty"`
   ```
 Every attribute of the __Pet__ struct is converted to lower case.


## Pet endpoints
- _/api/pet/{name}_ GET: retrieve the information of _name_
- _/api/pet/_ POST: store the information of your new pet
- _/api/pet/{name}_ PUT: add a new information for your existing pet
- _/api/pet/{name}_ DELETE: delete _name_ from your stored pets

## Pets endpoints
- _/api/pets GET: retrieve the information of all your stored pets

## Food endpoints
- _/api/food/{name}_ GET: get the suggested food to feed _name_
