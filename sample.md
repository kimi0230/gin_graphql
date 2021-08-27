```GraphQL
# Write your query or mutation here
query meetups{
	meetups{
    id
    name
    description
    user{
      account
    }
  }
}

query user{
	user(id:1){
    id
    account
    email
    createdAt
    meetups{
      id
      name
      description
    }
  }
}

mutation CreateMeetup{
  createMeetup(input:{name:"kkk123",description:"kkakakka"}){
    id
    name
    description
  }
}
```
